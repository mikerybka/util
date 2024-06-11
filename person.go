package util

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}
