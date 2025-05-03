package util

import (
	"sync"
)

// RunInParallel runs each function in fns in its own goroutine,
// waits for all to finish, and returns a slice of errors (same order, same length).
func RunInParallel(fns []func() error) []error {
	var wg sync.WaitGroup
	errs := make([]error, len(fns))

	wg.Add(len(fns))
	for i, fn := range fns {
		i, fn := i, fn // capture loop variables
		go func() {
			defer wg.Done()
			errs[i] = fn()
		}()
	}

	wg.Wait()
	return errs
}
