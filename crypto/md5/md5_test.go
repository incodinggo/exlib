package md5_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto/md5"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(md5.Sum([]byte("a")).Hex())
	fmt.Println(md5.Sum([]byte("a")).Hex16())
}
