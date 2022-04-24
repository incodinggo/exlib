package slice

type ni struct {
	n   []interface{} //不在B的元素
	idx []int         //index
}

// Strings 类型转换函数请务必保证所转换类型与传入类型一致
func (i *ni) Strings() ([]string, []int) {
	var r []string
	for _, e := range i.n {
		c, ok := e.(string)
		if !ok {
			return []string{}, []int{}
		}
		r = append(r, c)
	}
	return r, i.idx
}

// Int64s 类型转换函数请务必保证所转换类型与传入类型一致
func (i *ni) Int64s() ([]int64, []int) {
	var r []int64
	for _, e := range i.n {
		c, ok := e.(int64)
		if !ok {
			return []int64{}, []int{}
		}
		r = append(r, c)
	}
	return r, i.idx
}

// Int32s 类型转换函数请务必保证所转换类型与传入类型一致
func (i *ni) Int32s() ([]int32, []int) {
	var r []int32
	for _, e := range i.n {
		c, ok := e.(int32)
		if !ok {
			return []int32{}, []int{}
		}
		r = append(r, c)
	}
	return r, i.idx
}

// Ints 类型转换函数请务必保证所转换类型与传入类型一致
func (i *ni) Ints() ([]int, []int) {
	var r []int
	for _, e := range i.n {
		c, ok := e.(int)
		if !ok {
			return []int{}, []int{}
		}
		r = append(r, c)
	}
	return r, i.idx
}

// Bool 仅返回是否有不包含信息,true表示有不包含元素，false表示无
func (i *ni) Bool() bool {
	return len(i.n) != 0
}

//NIStr Slice A element not in Slice B [string]A数组是否有B不包含的元素,并返回B中不包含的A的元素数组
func NIStr(sliceA []string, sliceB []string) *ni {
	var n []interface{}
	var idx []int
	for i, e := range sliceA {
		if !In(sliceB, e).Bool() {
			n = append(n, e)
			idx = append(idx, i)
		}
	}
	return &ni{n: n, idx: idx}
}

//NIInt Slice A element not in Slice B [int]A数组是否有B不包含的元素,并返回B中不包含的A的元素数组
func NIInt(sliceA []int, sliceB []int) *ni {
	var n []interface{}
	var idx []int
	for i, e := range sliceA {
		if !In(sliceB, e).Bool() {
			n = append(n, e)
			idx = append(idx, i)
		}
	}
	return &ni{n: n, idx: idx}
}

//NII64 Slice A element not in Slice B [int64]A数组是否有B不包含的元素,并返回B中不包含的A的元素数组
func NII64(sliceA []int64, sliceB []int64) *ni {
	var n []interface{}
	var idx []int
	for i, e := range sliceA {
		if !In(sliceB, e).Bool() {
			n = append(n, e)
			idx = append(idx, i)
		}
	}
	return &ni{n: n, idx: idx}
}
