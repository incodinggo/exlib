package crc_test

import (
	"exlib/crypto/crc"
	"fmt"
	"testing"
)

func Testcrc32Q(t *testing.T) {
	str := "测试字符串"
	fmt.Println(crc.C32Q([]byte(str)).Hex())
}

func TestcrcIEEE(t *testing.T) {
	str := "测试字符串"
	fmt.Println(crc.IEEE([]byte(str)).Hex())
}
