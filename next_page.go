package util

type NextPage struct {
	Name  Name
	State []*Field
	Body  []*Statement
}

func (p *NextPage) Write(path string) error {
	c := &ReactComponent{
		Name:  p.Name,
		State: p.State,
		Body:  p.Body,
	}
	return c.Write(path)
}
