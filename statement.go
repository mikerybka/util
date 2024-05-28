package util

// Statement represents a line of code in a function or method body.
// There are 4 types of statements: returns, assignments, ifs and loops.
// Ifs and loops have substatements.
type Statement struct {
	IsReturn bool
	Return   *Expression

	IsAssign bool
	Name     string
	Value    *Expression

	IsIf      bool
	Condition *Condition
	Body      []*Statement
}
