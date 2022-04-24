package md5

import "crypto/md5"

func Sum(r []byte) *Ret {
	h := md5.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}
