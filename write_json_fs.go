package util

import (
	"bytes"
	"encoding/json"
	"encoding/json/jsontext"
	"io"
	"log"
	"os"
)

// WriteJSONFS takes a JSON value and writes it recursively to the filesystem.
func WriteJSONFS(path string, b []byte) error {
	dec := jsontext.NewDecoder(bytes.NewReader(b))
	for {
		tok, err := dec.ReadToken()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		switch tok.Kind() {
		case '{':

		case '}':
		case '[':
		case ']':
		case '"':
			f, err := os.Create(path)
			if err != nil {
				return err
			}
			err = json.NewEncoder(f).Encode(tok.String())
			if err != nil {
				return err
			}
			err = f.Close()
			if err != nil {
				return err
			}
		default:
			return os.WriteFile(path, []byte(tok.String()), os.ModePerm)
		}
	}
	return nil
}
