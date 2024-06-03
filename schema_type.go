package util

var SchemaType *Type = &Type{
	IsStruct: true,
	Fields: []Field{
		{
			ID:   "id",
			Name: "ID",
			Type: StringType,
		},
		{
			ID:   "name",
			Name: "Name",
			Type: StringType,
		},
		{
			ID:   "fields",
			Name: "Fields",
			Type: &Type{
				IsArray:  true,
				ElemType: FieldType,
			},
		},
	},
}

var FieldType *Type = &Type{
	IsStruct: true,
	Fields: []Field{
		{
			ID:   "id",
			Name: "ID",
			Type: StringType,
		},
		{
			ID:   "name",
			Name: "Name",
			Type: StringType,
		},
		{
			ID:   "type",
			Name: "Type",
			Type: TypeType,
		},
	},
}

var TypeType *Type = &Type{
	Name:       "type",
	PluralName: "types",
	IsStruct:   true,
	Fields: []Field{
		{
			ID:   "name",
			Name: "name",
			Type: StringType,
		},
		{
			ID:   "plural_name",
			Name: "plural name",
			Type: StringType,
		},
		{
			ID:   "is_scalar",
			Name: "is scalar",
			Type: BoolType,
		},
		{
			ID:   "kind",
			Name: "kind",
			Type: StringType,
		},
		{
			ID:   "is_pointer",
			Name: "is pointer",
			Type: BoolType,
		},
		{
			ID:   "is_array",
			Name: "is array",
			Type: BoolType,
		},
		{
			ID:   "is_map",
			Name: "is map",
			Type: BoolType,
		},
		{
			ID:   "elem_type",
			Name: "elem type",
			Type: TypeType,
		},
	},
}
