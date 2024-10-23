package util

// Statement represents a line of code in a function or method body.
// There are 4 types of statements: returns, assignments, ifs and loops.
// Ifs and loops have substatements.
type Statement struct {
	IsReturn bool        `json:"isReturn"`
	Return   *Expression `json:"return"`

	IsAssign bool        `json:"isAssign"`
	Name     string      `json:"name"`
	Value    *Expression `json:"value"`

	IsIf      bool         `json:"isIf"`
	Condition *Condition   `json:"condition"`
	Body      []*Statement `json:"body"`

	// TODO: loops
}
