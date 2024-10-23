package util

type MethodCall struct {
	Method string   `json:"m"`
	Args   []string `json:"args"`
}
