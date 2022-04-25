package aes_test

import (
	"encoding/hex"
	"exlib/crypto"
	"exlib/crypto/aes"
	"fmt"
	"testing"
)

func TestAes_Encrypt(t *testing.T) {
	o := aes.New("aaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbb", crypto.PaddingPKCS7, crypto.OFB)
	fmt.Println(o.Encrypt("欢迎使用exlib ").Base64())
}

func TestAes_Decrypt(t *testing.T) {
	o := aes.New("aaaaaaaaaaaaaaaa", "bbbbbbbbbbbbbbbb", crypto.PaddingPKCS7, crypto.OFB)
	//d, _ := hex.DecodeString("78c5d8e5858f22be99c3e582f9e73077abc1a98b6cddaf7bf1685f042a3bc7e3") //CBC PKCS7
	//d, _ := hex.DecodeString("5e129409275ca7997a731ccfc1d65a0ced532c0da092e408efdb266c755d5fcb") //ECB PKCS7
	//d, _ := hex.DecodeString("ffe5174ccfd9a68a6e60109cd0dba379a71e220ed8b946cf72a2be23bfe695b6") //CTR PKCS7
	d, _ := hex.DecodeString("ffe5174ccfd9a68a6e60109cd0dba379c6ab9263ab51cab1622f4163ac334e22") //OFB NONE
	//d, _ := hex.DecodeString("ffe5174ccfd9a68a6e60109cd0dba379dec7d5db94") //CFB NONE
	fmt.Println(o.Decrypt(d).String())
}
