package util

func XMLString(s string) XML {
	return XML{
		El: "string",
		Attrs: []Pair[string, string]{
			{
				Key:   "value",
				Value: s,
			},
		},
	}
}
