package util

import "time"

//MakeTimestamp will create a timestamp since epoch in ms
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
