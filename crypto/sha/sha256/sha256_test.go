package sha256_test

import (
	"exlib/crypto/sha/sha256"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(sha256.Sum([]byte("a")).Hex())
}
