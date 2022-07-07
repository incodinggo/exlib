package slice

//NI Slice A element not in Slice B [T]A数组是否有B不包含的元素,并返回B中不包含的A的元素数组
//n   不在B的元素
//idx index
func NI[T comparable](sliceA []T, sliceB []T) (n []T, idx []int) {
	for i, e := range sliceA {
		if !In(sliceB, e).Bool() {
			n = append(n, e)
			idx = append(idx, i)
		}
	}
	return n, idx
}
