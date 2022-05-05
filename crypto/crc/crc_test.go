package crc_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto/crc"
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
