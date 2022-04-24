package convert

import "github.com/gogf/gf/v2/util/gconv"

func Ints(in interface{}) []int {
	return gconv.SliceInt(in)
}
func Int32s(in interface{}) []int32 {
	return gconv.SliceInt32(in)
}
func Int64s(in interface{}) []int64 {
	return gconv.SliceInt64(in)
}
