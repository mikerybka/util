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
	children := []XML{
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					Key:   "name",
					Value: "viewport",
				},
				{
					Key:   "content",
					Value: "width=device-width, initial-scale=1.0",
				},
			},
		},
		{
			El: "title",
			Children: []XML{
				XMLString(h.Title),
			},
		},
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					Key:   "name",
					Value: "description",
				},
				{
					Key:   "content",
					Value: h.Desc,
				},
			},
		},
		{
			El: "meta",
			Attrs: []Pair[string, string]{
				{
					Key:   "name",
					Value: "author",
				},
				{
					Key:   "content",
					Value: h.Author,
				},
			},
		},
		{
			El: "link",
			Attrs: []Pair[string, string]{
				{
					Key:   "rel",
					Value: "stylesheet",
				},
				{
					Key:   "href",
					Value: h.CSS,
				},
				{
					Key:   "type",
					Value: "text/css",
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
