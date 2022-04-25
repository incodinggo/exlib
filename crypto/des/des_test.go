package des_test

import (
	"exlib/crypto"
	"exlib/crypto/des"
	"fmt"
	"testing"
)

func TestDes_Encrypt(t *testing.T) {
	o := des.New("aaaaaaaa", "bbbbbbbb", crypto.PaddingPKCS7, crypto.CFB)
	fmt.Println(o.Encrypt("欢迎使用exlib ").Base64())
}

func TestDes_Decrypt(t *testing.T) {
	o := des.New("aaaaaaaa", "bbbbbbbb", crypto.PaddingPKCS7, crypto.CFB)
	//d, _ := crypto.FromHex("9c21ec39e3a5a3febd3686c31180fac03388c6f76a52d2f1") //ECB PKCS7
	//d, _ := crypto.FromHex("a5460f99c3863798dd0ae9ca78596e5c8edbb166c7c86d06") //CBC PKCS7
	//d, _ := crypto.FromHex("048924980e310d6a1bb0d52393e99d4c65c028617502c824") //CTR PKCS7
	//d, _ := crypto.FromHex("048924980e310d6a2a51f4a40d7f33ed942dfe900ed9f562") //OFB PKCS7
	d, _ := crypto.FromHex("048924980e310d6a017c9396e5ee02e8e0949cd240729cde") //CFB PKCS7
	fmt.Println(o.Decrypt(d).String())
}
