package slice

import (
	"reflect"
	"sort"
	"unsafe"
)

// FastSortF64 The way for float64 sort faster cause type float64 follow IEEE754 features of floating standard
func FastSortF64(a []float64) {
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
	*cHdr = *aHdr
	sort.Ints(c)
}
