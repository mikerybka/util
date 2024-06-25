package util

func XMLString(s string) XML {
	return XML{
		El: "string",
		Attrs: []Pair[string, string]{
			{
				K: "value",
				V: s,
			},
		},
	}
}
