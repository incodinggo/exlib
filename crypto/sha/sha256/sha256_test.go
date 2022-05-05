package sha256_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto/sha/sha256"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(sha256.Sum([]byte("a")).Hex())
}
