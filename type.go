package util

import (
	"fmt"
)

type Type struct {
	Name        Name                 `json:"name"`
	Description string               `json:"description"`
	IsScalar    bool                 `json:"isScalar"`
	IsArray     bool                 `json:"isArray"`
	IsMap       bool                 `json:"isMap"`
	ElemType    string               `json:"elemType"`
	IsStruct    bool                 `json:"isStruct"`
	Fields      []Field              `json:"fields"`
	Methods     map[string]*Function `json:"methods"`
	DefaultJSON string               `json:"defaultJSON"`
}

func (t *Type) WriteTypeScriptFile(path string) error {
	if t.IsScalar {
		st := &ScalarType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return st.WriteTypeScriptFile(path)
	}

	if t.IsArray {
		at := &ArrayType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return at.WriteTypeScriptFile(path)
	}

	if t.IsMap {
		mt := &MapType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return mt.WriteTypeScriptFile(path)
	}

	if t.IsStruct {
		st := &StructType{
			Name:        t.Name,
			Description: t.Description,
			Fields:      t.Fields,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return st.WriteTypeScriptFile(path)
	}

	return fmt.Errorf("invalid type")
}

func (t *Type) WriteGoFile(path string) error {
	if t.IsScalar {
		st := &ScalarType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return st.WriteGoFile(path)
	}

	if t.IsArray {
		at := &ArrayType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return at.WriteGoFile(path)
	}

	if t.IsMap {
		mt := &MapType{
			Name:        t.Name,
			Description: t.Description,
			ElemType:    t.ElemType,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return mt.WriteGoFile(path)
	}

	if t.IsStruct {
		st := &StructType{
			Name:        t.Name,
			Description: t.Description,
			Fields:      t.Fields,
			Methods:     t.Methods,
			DefaultJSON: t.DefaultJSON,
		}
		return st.WriteGoFile(path)
	}

	return fmt.Errorf("invalid type")
}
