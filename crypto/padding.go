package crypto

import (
	"bytes"
)

//Zero padding
func ZeroPadding(r []byte, blockSize int) []byte {
	padding := blockSize - len(r)%blockSize
	if padding == blockSize {
		return r
	}
	return append(r, bytes.Repeat([]byte{0}, padding)...)
}

func ZeroUnPadding(r []byte) []byte {
	return bytes.TrimFunc(r, func(ru rune) bool {
		return ru == rune(0)
	})
}

//PKCS5
func PKCS5Padding(r []byte, blockSize int) []byte {
	padding := blockSize - len(r)%blockSize
	p := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(r, p...)
}

func PKCS5UnPadding(r []byte) []byte {
	length := len(r)
	// 去掉最后一个字节 p 次
	p := int(r[length-1])
	return r[:(length - p)]
}

//PKCS7
func PKCS7Padding(r []byte, blockSize int) []byte {
	n := blockSize - len(r)%blockSize
	p := bytes.Repeat([]byte{byte(n)}, n)
	return append(r, p...)
}

func PKCS7UnPadding(r []byte) []byte {
	l := len(r)
	p := int(r[l-1])
	return r[:(l - p)]
}

func Padding(r []byte, typ, bz int) []byte {
	switch typ {
	case PaddingZero:
		return ZeroPadding(r, bz)
	case PaddingPKCS5:
		return PKCS5Padding(r, bz)
	case PaddingPKCS7:
		return PKCS7Padding(r, bz)
	default:
		return r
	}
}

func UnPadding(r []byte, typ int) []byte {
	switch typ {
	case PaddingZero:
		return ZeroUnPadding(r)
	case PaddingPKCS5:
		return PKCS5UnPadding(r)
	case PaddingPKCS7:
		return PKCS7UnPadding(r)
	default:
		return r
	}
}
