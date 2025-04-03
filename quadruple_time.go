package util

import (
	"fmt"
	"time"
)

// Run each fn 250 ms apart forever.
// Returns an error if a func takes longer than 249 ms to run.
func QuadrupleTime(fn1, fn2, fn3, fn4 func()) error {
	// Wait until the top of the second
	d := 1_000_000_000 - time.Now().Nanosecond()
	time.Sleep(time.Duration(d))

	for {
		// Run first func
		fn1()

		// Wait until 250 ms
		now := time.Now()
		if now.Nanosecond() > 249_000_000 {
			return fmt.Errorf("fn1 took too long")
		}
		d = 250_000_000 - time.Now().Nanosecond()
		time.Sleep(time.Duration(d))

		// Run second func
		fn2()

		// Wait until 500 ms
		now = time.Now()
		if now.Nanosecond() > 499_000_000 || now.Nanosecond() < 250_000_000 {
			return fmt.Errorf("fn2 took too long")
		}
		d = 500_000_000 - time.Now().Nanosecond()
		time.Sleep(time.Duration(d))

		// Run third func
		fn3()

		// Wait until 750 ms
		now = time.Now()
		if now.Nanosecond() > 749_000_000 || now.Nanosecond() < 500_000_000 {
			return fmt.Errorf("fn3 took too long")
		}
		d = 750_000_000 - time.Now().Nanosecond()
		time.Sleep(time.Duration(d))

		// Run final func
		fn4()

		// Sleep until the top of the second again
		now = time.Now()
		if now.Nanosecond() > 999_000_000 || now.Nanosecond() < 750_000_000 {
			return fmt.Errorf("fn4 took too long")
		}
		d = 1_000_000_000 - time.Now().Nanosecond()
		time.Sleep(time.Duration(d))
	}
}
