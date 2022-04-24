package sha256

import (
	"crypto/sha256"
)

func Sum(r []byte) *Ret {
	h := sha256.New()
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}
