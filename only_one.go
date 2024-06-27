package util

func OnlyOne[T any](m map[string]T) (id string, value T, ok bool) {
	if len(m) != 1 {
		return "", value, false
	}
	for k, v := range m {
		id = k
		value = v
	}
	return id, value, true
}
