package util

func Pop[T any](items []T) (T, []T) {
	return items[0], items[1:]
}
