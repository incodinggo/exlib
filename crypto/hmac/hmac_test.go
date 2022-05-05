package hmac_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto/hmac"
	"testing"
)

func TestSum1(t *testing.T) {
	fmt.Println(hmac.Sum1([]byte("a"), []byte("a")).Hex())
}
func TestSum256(t *testing.T) {
	fmt.Println(hmac.Sum256([]byte("a"), []byte("a")).Hex())
}
func TestSum512(t *testing.T) {
	fmt.Println(hmac.Sum512([]byte("a"), []byte("a")).Hex())
}
