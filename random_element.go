package util

import (
	"fmt"
	"math/rand"
)

func RandomElement[T any](arr []T) (T, error) {
	if len(arr) == 0 {
		var zero T
		return zero, fmt.Errorf("cannot select from an empty slice")
	}
	index := rand.Intn(len(arr))
	return arr[index], nil
}
