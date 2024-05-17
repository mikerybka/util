package util

type Collection struct {
	Schema *Schema
	Items  []string
}

// List returns a list of the first `limit` items after
func (c *Collection) List(limit, after int) {
	panic("not implemented")
}
