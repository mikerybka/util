package util

import (
	"io"
	"net/http"
)

type HTMLHead struct {
	Title  string
	Desc   string
	Author string
	CSS    string
}

func (h *HTMLHead) ServeHTTP(w http.ResponseWriter, r http.Request) {
	h.Write(w)
}

func (h *HTMLHead) Write(w io.Writer) (int64, error) {
	return h.XML().Write(w)
}

func (h *HTMLHead) XML() *XML {
	children := []*XML{
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					K: "name",
					V: "viewport",
				},
				{
					K: "content",
					V: "width=device-width, initial-scale=1.0",
				},
			},
		},
		{
			El: "title",
			Children: []*XML{
				XMLString(h.Title),
			},
		},
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					K: "name",
					V: "description",
				},
				{
					K: "content",
					V: h.Desc,
				},
			},
		},
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					K: "name",
					V: "author",
				},
				{
					K: "content",
					V: h.Author,
				},
			},
		},
		{
			El: "link",
			Attrs: []Pair[string, string]{
				{
					K: "rel",
					V: "stylesheet",
				},
				{
					K: "href",
					V: h.CSS,
				},
				{
					K: "type",
					V: "text/css",
				},
			},
		},
	}

	xml := &XML{
		El:       "head",
		Children: children,
	}

	return xml
}
