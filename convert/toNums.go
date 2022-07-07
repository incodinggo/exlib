package convert

import (
	"strconv"
)

func Ints(in []string) []int {
	var out []int
	for _, e := range in {
		i, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil
		}
		out = append(out, int(i))
	}
	return out
}
func Int32s(in []string) []int32 {
	var out []int32
	for _, e := range in {
		i, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil
		}
		out = append(out, int32(i))
	}
	return out
}
func Int64s(in []string) []int64 {
	var out []int64
	for _, e := range in {
		i, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil
		}
		out = append(out, i)
	}
	return out
}
