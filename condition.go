package util

type Condition struct {
	LHS *Expression
	Op  string
	RHS *Expression
}
