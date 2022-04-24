package stringex

import "strconv"

type T struct {
	v []string
}

func (t *T) Strings() []string {
	return t.v
}

func (t *T) Int64s() []int64 {
	if len(t.v) == 0 {
		return []int64{}
	}
	var elems []int64
	for _, v := range t.v {
		elem, err := strconv.ParseInt(v, 0, 0)
		if err != nil {
			return []int64{}
		}
		elems = append(elems, elem)
	}
	return elems
}

func (t *T) Ints() []int {
	if len(t.v) == 0 {
		return []int{}
	}
	var elems []int
	for _, v := range t.v {
		elem, err := strconv.Atoi(v)
		if err != nil {
			return []int{}
		}
		elems = append(elems, elem)
	}
	return elems
}

func (t *T) Float64s() []float64 {
	if len(t.v) == 0 {
		return []float64{}
	}
	var elems []float64
	for _, v := range t.v {
		elem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return []float64{}
		}
		elems = append(elems, elem)
	}
	return elems
}

func (t *T) Float32s() []float32 {
	if len(t.v) == 0 {
		return []float32{}
	}
	var elems []float32
	for _, v := range t.v {
		elem, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return []float32{}
		}
		elems = append(elems, float32(elem))
	}
	return elems
}
