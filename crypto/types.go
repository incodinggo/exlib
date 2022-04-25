package crypto

import "crypto"

const (
	NONE = iota
	PaddingZero
	PaddingPKCS5
	PaddingPKCS7
	NoPadding
)

const (
	ECB = iota
	CBC
	CTR
	OFB
	CFB
)

const (
	Bit64   = 64
	Bit128  = 128
	Bit256  = 256
	Bit512  = 512
	Bit1024 = 1024
	Bit2048 = 2048
	Bit4096 = 4096
)

const (
	PirTypEC = iota
	PirTypPKCS1
	PirTypPKCS8
)

const (
	SignTypSHA1   = crypto.SHA1
	SignTypSHA256 = crypto.SHA256
	SignTypSHA512 = crypto.SHA512
	SignTypMD5    = crypto.MD5
)
