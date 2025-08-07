package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Backend struct {
	DataURL   string `json:"dataURL"`
	GoWorkDir string `json:"goWorkDir"`
	BinDir    string `json:"binDir"`
}

func (b *Backend) Get(path string) (*Response, error) {
	u := b.DataURL + "/data" + path
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	resp := &Response{}
	err = json.NewDecoder(res.Body).Decode(resp)
	if err != nil {
		panic(err)
	}
	return resp, nil
}

func (b *Backend) listTypes() ([]string, error) {
	u := b.DataURL + "/types"
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	res := &Response{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		panic(err)
	}
	if res.Type != "dir" {
		panic(b.DataURL + "/types not dir")
	}
	typeIDs := []string{}
	err = json.Unmarshal([]byte(res.Value), &typeIDs)
	if err != nil {
		panic("bad response")
	}
	return typeIDs, nil
}

func (b *Backend) getType(t string) (*Type, error) {
	u := b.DataURL + "/types/" + t
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	resp := &Type{}
	err = json.NewDecoder(res.Body).Decode(resp)
	if err != nil {
		panic(err)
	}
	return resp, nil
}

func (b *Backend) getTypes() (map[string]*Type, error) {
	typeIDs, err := b.listTypes()
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	types := map[string]*Type{}
	for _, id := range typeIDs {
		t, err := b.getType(id)
		if err != nil {
			return nil, err
		}
		types[id] = t
	}
	return types, nil
}

// buildHandler builds a Go HTTP server to handle modifications to custom types.
// It returns the path to the built executable or an error.
func (backend *Backend) buildHandler(typeID string) (string, error) {
	// Create workdir
	cmd := exec.Command("go", "work", "init")
	cmd.Dir = backend.GoWorkDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Generate Go code
	path := filepath.Join(backend.GoWorkDir, "types")
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	cmd = exec.Command("go", "mod", "init", "types")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	types, err := backend.getTypes()
	if err != nil {
		return "", err
	}
	// for id, t := range types {
	// 	err = t.WriteGoFile(filepath.Join(path, id+".go"), "types")
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }

	// Generate cmd package
	path = filepath.Join(backend.GoWorkDir, "cmd")
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}
	cmd = exec.Command("go", "mod", "init", "cmd")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	for id, t := range types {
		f, err := os.Create(filepath.Join(path, id, "main.go"))
		if err != nil {
			return "", err
		}
		tmpl := `package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"types"
)

func main() {
	method := os.Args[1]
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	data := &types.{{ .Name }}{}
	args := []string{}
	json.Unmarshal([]byte(lines[0]), data)
	json.Unmarshal([]byte(lines[1]), &args)
	switch method {
{{ range .Methods }}
	case "{{ .Name }}":
{{ range .Inputs }}
		arg{{ i+1 }} := {{ zeroValue(.Type) }}
		json.Unmarshal([]byte(args[{{ i }}]), &arg{{ i+1 }})
{{ end }}
		{{ ranage .Outputs }}{{ if i > 0 }}, {{ end }}out{{ i+1 }}{{ end }} := data.{{ .Name }}({{ ranage .Inputs }}{{ if i > 0 }}, {{ end }}arg{{ i+1 }}{{ end }})
		json.NewEncoder(os.Stdout).Encode(data)
		fmt.Println()
		res := []string{}
{{ i := range .Outputs }}
		b, err := json.Marshal(out{{ i+1 }})
		if err != nil {
			panic(err)
		}
		res = append(res, string(b))
{{ end }}
		json.NewEncoder(os.Stdout).Encode(res)
		fmt.Println()
{{ end }}
	default:
		panic("no such method")
	}
}
		
`
		err = template.Must(template.New("cmd").Parse(tmpl)).Execute(f, t)
		if err != nil {
			return "", err
		}
	}

	// build handler
	outfile := filepath.Join(backend.BinDir, typeID)
	cmd = exec.Command("go", "build", "-o", outfile, filepath.Join(path, typeID))
	cmd.Dir = backend.GoWorkDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return outfile, nil
}

func (backend *Backend) Call(path string, method string, args []string) ([]string, error) {
	// Get data
	data, err := backend.Get(path)
	if err != nil {
		return nil, err
	}

	// Build handler
	c, err := backend.buildHandler(data.Type)
	if err != nil {
		panic(err)
	}

	// Call handler
	cmd := exec.Command(c, method)
	input, err := json.Marshal(struct {
		Data   string   `json:"data"`
		Method string   `json:"method"`
		Args   []string `json:"args"`
	}{
		Data:   data.Value,
		Method: method,
		Args:   args,
	})
	if err != nil {
		panic(err)
	}
	input = append(input, '\n')
	cmd.Stdin = bytes.NewReader(input)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(out), "\n")
	newData := lines[0]
	result := lines[1]

	// Save data
	req, err := http.NewRequest("PUT", backend.DataURL+"/data"+path, strings.NewReader(newData))
	if err != nil {
		panic(err)
	}
	saveResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if saveResp.StatusCode != 200 {
		return nil, fmt.Errorf("while saving: %s", saveResp.Status)
	}

	// Return result
	res := []string{}
	err = json.Unmarshal([]byte(result), &res)
	if err != nil {
		panic(err)
	}
	return res, nil
}
