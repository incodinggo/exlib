package slice

import "reflect"

// RmDup RmDupI remove duplicate & zero element
// rmEmp 是否需要删除zero值
func RmDup[T comparable](i []T, rmEmp bool) []T {
	set := make(map[T]struct{}, len(i))
	j := 0
	for _, v := range i {
		if rmEmp && reflect.ValueOf(v).IsZero() {
			continue
		}
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		i[j] = v
		j++
	}
	return i[:j]
}

// In 通配
func In[T comparable](sl []T, v T) *in {
	for i, t := range sl {
		if t == v {
			return &in{i}
		}
	}
	return &in{-1}
}

// Insert Add element(s) at any index of the slice
// orig	Need a slice
// ins	Need a slice
// Func will change orig slice
func Insert[T comparable](sl *[]T, ins []T, idx int) {
	if idx >= len(*sl) {
		panic("insert slice index out of range")
	}
	*sl = append(*sl, ins...)
	copy((*sl)[idx+len(ins):], (*sl)[idx:])
	copy((*sl)[idx:], ins)
}
