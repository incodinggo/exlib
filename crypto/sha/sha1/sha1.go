package sha1

import (
	"crypto/sha1"
)

func Sum(r []byte) *Ret {
	h := sha1.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}
