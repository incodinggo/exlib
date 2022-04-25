package crypto

import (
	"encoding/base64"
	"encoding/hex"
)

//Base64
func Base64(raw []byte) string {
	return base64.StdEncoding.EncodeToString(raw)
}

func FromBase64(raw string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(raw)
}

//HEX
func Hex(raw []byte) string {
	return hex.EncodeToString(raw)
}

func FromHex(raw string) ([]byte, error) {
	return hex.DecodeString(raw)
}
