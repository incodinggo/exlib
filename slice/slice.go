package slice

import (
	"fmt"
	"reflect"
)

// RmDup remove duplicate & empty element
// 是否需要删除空字符 rmEmp
func RmDup(i []string, rmEmp bool) []string {
	set := make(map[interface{}]struct{}, len(i))
	j := 0
	for _, v := range i {
		if rmEmp && v == "" {
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

// RmDupI remove duplicate & zero element
// 是否需要删除0值 rmEmp
func RmDupI(i []int64, rmEmp bool) []int64 {
	set := make(map[interface{}]struct{}, len(i))
	j := 0
	for _, v := range i {
		if rmEmp && v == 0 {
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
func In(slice interface{}, v interface{}) *in {
	switch reflect.ValueOf(v).Kind() {
	case reflect.String:
		return &in{containsString(slice.([]string), v.(string))}
	case reflect.Int:
		return &in{containsInt(slice.([]int), v.(int))}
	case reflect.Int64:
		return &in{containsInt64(slice.([]int64), v.(int64))}
	case reflect.Bool:
		return &in{containsBool(slice.([]bool), v.(bool))}
	case reflect.Uint64:
		return &in{containsUint64(slice.([]uint64), v.(uint64))}
	case reflect.Float64:
		return &in{containsFloat(slice.([]float64), v.(float64))}
	case reflect.Complex128:
		return &in{containsComplex(slice.([]complex128), v.(complex128))}
	default:
		return &in{contains(slice, v)}
	}
}

// Insert Add element(s) at any index of the slice
// orig	Need a slice ptr
// ins	Need a slice
// Func will change orig slice
func Insert(orig interface{}, ins interface{}, idx int, t reflect.Kind) {
	v := reflect.ValueOf(orig)
	if v.Len()-1 < idx {
		panic("insert slice index out of range")
	}
	typ := v.Type().Elem().(reflect.Type)
	switch typ.Kind() {
	case reflect.String:
		insertString(orig.(*[]string), ins.([]string), idx)
	case reflect.Int:
		insertInt(orig.(*[]int), ins.([]int), idx)
	case reflect.Int64:
		insertInt64(orig.(*[]int64), ins.([]int64), idx)
	case reflect.Bool:
		insertBool(orig.(*[]bool), ins.([]bool), idx)
	case reflect.Uint64:
		insertUint64(orig.(*[]uint64), ins.([]uint64), idx)
	case reflect.Float64:
		insertFloat(orig.(*[]float64), ins.([]float64), idx)
	case reflect.Complex128:
		insertComplex(orig.(*[]complex128), ins.([]complex128), idx)
	default:
		panic(fmt.Sprintf("type %T not supported", t))
	}
}
