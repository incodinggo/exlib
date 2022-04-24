package slice

import "reflect"

type in struct {
	v int
}

func (i *in) index() int {
	return i.v
}

func (i *in) Bool() bool {
	return i.v != -1
}

//万能方法，牺牲性能
func contains(slice interface{}, v interface{}) (index int) {
	index = -1
	if slice == nil {
		return
	}
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(v, s.Index(i).Interface()) {
				index = i
				return
			}
		}
	}
	return
}

//强类型，保证性能
func containsString(slice []string, v string) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInt(slice []int, v int) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsInt64(slice []int64, v int64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsBool(slice []bool, v bool) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsUint64(slice []uint64, v uint64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsFloat(slice []float64, v float64) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}

func containsComplex(slice []complex128, v complex128) (index int) {
	index = -1
	if slice == nil {
		return
	}
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return
}
