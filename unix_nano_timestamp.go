package util

import (
	"strconv"
	"time"
)

func UnixNanoTimestamp() string {
	return strconv.Itoa(int(time.Now().UnixNano()))
}
