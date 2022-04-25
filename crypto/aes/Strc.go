package aes

import (
	"encoding/base64"
	"encoding/hex"
)

type Ret struct {
	v   []byte
	err error
}

func (r *Ret) Byte() []byte {
	return r.v
}

func (r *Ret) String() string {
	return string(r.v)
}

func (r *Ret) Hex() string {
	return hex.EncodeToString(r.v)
}

func (r *Ret) Base64() string {
	return base64.StdEncoding.EncodeToString(r.v)
}

func (r *Ret) Error() error {
	return r.err
}

func (r *Ret) Result() ([]byte, error) {
	return r.v, r.err
}
