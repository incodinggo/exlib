package crc

import "fmt"

type Ret struct {
	V uint32
}

func (r *Ret) Uint32() uint32 {
	return r.V
}

func (r *Ret) Uint64() uint64 {
	return uint64(r.V)
}

func (r *Ret) Uint() uint {
	return uint(r.V)
}

func (r *Ret) Int() int {
	return int(r.V)
}

func (r *Ret) Int64() int64 {
	return int64(r.V)
}

func (r *Ret) Hex() string {
	return fmt.Sprintf("%X", r.V)
}
