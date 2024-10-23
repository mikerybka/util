package util

type Condition struct {
	LHS *Expression `json:"lhs"`
	Op  string      `json:"op"`
	RHS *Expression `json:"rhs"`
}
