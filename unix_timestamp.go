package util

import (
	"strconv"
	"time"
)

func UnixTimestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
