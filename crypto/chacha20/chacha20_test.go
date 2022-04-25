package chacha20_test

import (
	"exlib/crypto"
	"exlib/crypto/chacha20"
	"fmt"
	"testing"
)

func TestChacha20_XorKeyStream(t *testing.T) {
	key := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	iv := []byte("12345678")
	text := []byte("欢迎使用exlib")
	c := chacha20.New(key, iv, 20)
	fmt.Println(c.XorKeyStream(text).Base64())
}

func TestChacha20_XorKeyStream2(t *testing.T) {
	key, _ := crypto.FromHex("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	iv, _ := crypto.FromHex("1234567898765412")
	b, _ := crypto.FromBase64("JTyQHN3UNt1B5PzKVvPPxk82dKA=")
	c := chacha20.New(key, iv, 8)
	fmt.Println(c.XorKeyStream(b).String())
}
