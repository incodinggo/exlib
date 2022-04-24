

package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func Sum1(k, r []byte) *Ret {
	h := hmac.New(sha1.New, k)
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}

func Sum256(k, r []byte) *Ret {
	h := hmac.New(sha256.New, k)
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}

func Sum512(k, r []byte) *Ret {
	h := hmac.New(sha512.New, k)
	h.Write(r)
	return &Ret{v: h.Sum(nil)}
}
