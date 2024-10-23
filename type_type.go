package util

var TypeType = &Type{
	IsStruct: true,
	Fields: []Field{
		{
			Name:        NewName("Is scalar"),
			Description: "",
			Type:        "bool",
		},
		{
			Name:        NewName("Kind"),
			Description: "",
			Type:        "string",
		},
		{
			Name:        NewName("Is pointer"),
			Description: "",
			Type:        "bool",
		},
		{
			Name:        NewName("Is array"),
			Description: "",
			Type:        "bool",
		},
		{
			Name:        NewName("Is map"),
			Description: "",
			Type:        "bool",
		},
		{
			Name:        NewName("Elem type"),
			Description: "",
			Type:        "string",
		},
		{
			Name:        NewName("Is struct"),
			Description: "",
			Type:        "bool",
		},
		{
			Name:        NewName("Fields"),
			Description: "",
			Type:        "[]util.Field",
		},
	},
}
