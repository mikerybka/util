package main

import (
	"fmt"
	"os"

	"github.com/mikerybka/util"
)

func main() {
	root := &util.Type{
		IsStruct: true,
		Fields: []util.Field{
			{
				Name: util.NewName("One"),
				Type: &util.Type{
					IsScalar: true,
					Scalar: &util.Ref{
						Name: "string",
					},
				},
			},
			{
				Name: util.NewName("Two"),
				Type: &util.Type{
					IsArray: true,
					ElemType: &util.Type{
						IsScalar: true,
						Scalar: &util.Ref{
							Name: "int",
						},
					},
				},
			},
			{
				Name: util.NewName("Three"),
				Type: &util.Type{
					IsMap: true,
					ElemType: &util.Type{
						IsScalar: true,
						Scalar: &util.Ref{
							Name: "bool",
						},
					},
				},
			},
		},
	}
	fmt.Println(util.NewJSONReader(map[string]*util.Type{}).Read(os.Stdout, root, "cmd/jsonreadertest/root"))
}
