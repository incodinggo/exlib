package crc

import (
	"hash/crc32"
)

func C32Q(raw []byte) *Ret {
	// 在此包中，crc多项式以反转符号表示，
	// 或LSB优先表示。
	//
	// LSB优先表示是一个带有n位的十六进制数，其中
	// 最高有效位表示x⁰和最低有效系数
	// bit表示xⁿ-1的系数（xⁿ的系数是隐含的）。
	//
	// 例如，crc32-Q，由以下多项式定义，
	//	x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
	// 具有反转符号0b11010101100000101000001010000001，所以该值
	// 应该传递给MakeTable的是0xD5828281。
	crc32q := crc32.MakeTable(0xD5828281)
	return &Ret{crc32.Checksum(raw, crc32q)}
}

func CWithTab(tab *crc32.Table, raw []byte) *Ret {
	return &Ret{crc32.Checksum(raw, tab)}
}

func IEEE(raw []byte) *Ret {
	return &Ret{crc32.ChecksumIEEE(raw)}
}
