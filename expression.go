package util

type Expression struct {
	IsLiteral bool
	Value     string

	IsCall bool
	Fn     *Ref
	Args   []*Expression

	IsRef bool
	Ref   *Ref
}
