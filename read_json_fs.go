package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func ReadJSONFS(w io.Writer, t *Type, path string) error {
	if t.IsScalar {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	}
	if t.IsArray {
		_, err := w.Write([]byte("["))
		if err != nil {
			return err
		}
		i := 0
		for {
			elPath := filepath.Join(path, strconv.Itoa(i))
			if !Exists(elPath) {
				break
			}
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			err = ReadJSONFS(w, t.ElemType, elPath)
			if err != nil {
				return err
			}
			i++
		}
		_, err = w.Write([]byte("]"))
		return err
	}
	if t.IsMap {
		_, err := w.Write([]byte("{"))
		if err != nil {
			return err
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		for i, e := range entries {
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			_, err = fmt.Fprintf(w, "\"%s\":", e.Name())
			if err != nil {
				return err
			}
			err = ReadJSONFS(w, t.ElemType, filepath.Join(path, e.Name()))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("}"))
		return err
	}
	if t.IsStruct {
		_, err := w.Write([]byte("{"))
		if err != nil {
			return err
		}
		for i, f := range t.Fields {
			if i > 0 {
				fmt.Fprintf(w, ",")
			}
			_, err = fmt.Fprintf(w, "\"%s\":", f.ID())
			if err != nil {
				return err
			}
			err = ReadJSONFS(w, f.Type, filepath.Join(path, f.ID()))
			if err != nil {
				return err
			}
		}
		_, err = w.Write([]byte("}"))
		return err
	}
	return fmt.Errorf("bad type")
}
