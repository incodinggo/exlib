package convert

import (
	"strconv"
)

func Int(in string) int {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic("not a number")
	}
	return int(i)
}
func Int32(in string) int32 {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic("not a number")
	}
	return int32(i)
}
func Int64(in string) int64 {
	i, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic("not a number")
	}
	return i
}
