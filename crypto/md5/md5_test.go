package md5_test

import (
	"ExLib/crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(md5.Sum([]byte("a")).Hex())
	fmt.Println(md5.Sum([]byte("a")).Hex16())
}
