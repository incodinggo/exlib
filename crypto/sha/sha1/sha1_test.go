package sha1_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto/sha/sha1"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(sha1.Sum([]byte("a")).Hex())
}
