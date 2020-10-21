package base

import (
	"bytes"
	"runtime"
	"strconv"
	"time"
)

func GoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func GetMilliTimeStamp(time time.Time) int64 {
	return time.UnixNano() / 1e6
}
