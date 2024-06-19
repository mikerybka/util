package main

import "github.com/mikerybka/util"

func main() {
	pkg := &util.GoPackage{
		Types: map[string]util.Type{
			"Schema": util.Type{
				IsStruct: true,
				Fields: []util.Field{
					{
						Name: "Fields",
						Type: &util.Type{
							IsArray:  true,
							ElemType: "SchemaField",
						},
					},
				},
			},
			"Schema Field": util.Type{
				IsStruct: true,
				Fields: []util.Field{
					{
						Name: "Name",
						Type: &util.Type{
							IsScalar: true,
							Kind:     "string",
						},
					},
					{
						Name: "Type",
						Type: &util.Type{
							IsScalar: true,
							Kind:     "string",
						},
					},
				},
			},
		},
	}

	err := pkg.Write(".")
	if err != nil {
		panic(err)
	}
}
