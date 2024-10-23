package util

var FieldType = &Type{
	IsStruct: true,
	Fields: []Field{
		{
			Name:        NewName("Name"),
			Description: "",
			Type:        "Name",
		},
		{
			Name:        NewName("Description"),
			Description: "",
			Type:        "string",
		},
		{
			Name:        NewName("Type"),
			Description: "",
			Type:        "string",
		},
	},
}
