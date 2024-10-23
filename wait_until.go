package util

import "time"

func WaitUntil(h, m, s int) {
	now := time.Now().UTC()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), h, m, s, 0, now.Location())
	if startTime.Before(now) {
		startTime = startTime.Add(24 * time.Hour)
	}
	time.Sleep(startTime.Sub(now))
}
