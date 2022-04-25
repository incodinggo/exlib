package sha1_test

import (
	"exlib/crypto/sha/sha1"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(sha1.Sum([]byte("a")).Hex())
}
