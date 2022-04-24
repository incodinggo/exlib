package slice

func insertString(slice *[]string, v []string, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertInt(slice *[]int, v []int, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertInt64(slice *[]int64, v []int64, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertBool(slice *[]bool, v []bool, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertUint64(slice *[]uint64, v []uint64, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertFloat(slice *[]float64, v []float64, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}

func insertComplex(slice *[]complex128, v []complex128, idx int) {
	*slice = append(*slice, v...)
	copy((*slice)[idx+len(v):], (*slice)[idx:])
	copy((*slice)[idx:], v)
}
