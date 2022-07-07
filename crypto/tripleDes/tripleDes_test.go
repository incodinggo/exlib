package tripledes_test

import (
	"fmt"
	"github.com/incodinggo/exlib/crypto"
	"github.com/incodinggo/exlib/crypto/tripleDes"
	"testing"
)

func TestTripledes_Encrypt(t *testing.T) {
	td := tripledes.New("aaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb", crypto.PaddingPKCS7, crypto.ECB)
	fmt.Println(td.Encrypt("欢迎使用Tripledes ").Hex())
}

func TestTripledes_Decrypt(t *testing.T) {
	td := tripledes.New("aaaaaaaaaaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbbbbbbbbbb", crypto.PaddingPKCS7, crypto.ECB)
	b, _ := crypto.FromHex("9c21ec39e3a5a3fe8a6c841546af5e3810e608b209bd59dc")
	fmt.Println(td.Decrypt(b).String())
}
