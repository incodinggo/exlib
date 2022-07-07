package slice

type in struct {
	v int
}

func (i *in) index() int {
	return i.v
}

func (i *in) Bool() bool {
	return i.v != -1
}
