package rsa_test

import (
	"exlib/crypto"
	"exlib/crypto/rsa"
	"fmt"
	"testing"
)

func TestKeyPair(t *testing.T) {
	fmt.Println(rsa.KeyPair(crypto.Bit2048))
}

//-----------pkcs8
var pub = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCoh1TkpfKVn0tECx7Xvr8ALysl
UFyECGiQbEv+wFcuOsgiW1Hpk1Su1fuDh/DlCcJbNGuYF4R61gh4OKcnMQW3t4hW
NojZx9e022J4ODP44O5W6O293mpjxkMe9s0+HkNnq6FxJ7pk5a7FdhaTTi9eZOWe
b33QF2vOOZEdpu6zBQIDAQAB
-----END PUBLIC KEY-----`

var pri = `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAKiHVOSl8pWfS0QL
Hte+vwAvKyVQXIQIaJBsS/7AVy46yCJbUemTVK7V+4OH8OUJwls0a5gXhHrWCHg4
pycxBbe3iFY2iNnH17TbYng4M/jg7lbo7b3eamPGQx72zT4eQ2eroXEnumTlrsV2
FpNOL15k5Z5vfdAXa845kR2m7rMFAgMBAAECgYAtIpGJ6yfdCpyNzKyl+2AMHJXh
mHQuqFHY3Cg+QiUXLIcNLIfzlpHHgWerAm1x6fszkgZ+5U2F2GzMRd/+jxKIggA0
HBawQO3QyKY1n3o4cXRn0AIfv7qPYUaD/kgZzssCVtWGIScDcmPOjLl8siGuBpH2
V0guApCEHYshjFHMAQJBAN+Bi5MzRonnQxu7jXDVdHn81ZD4MAGgkNbb6JsAK+5r
gDlj93wgPJ5rLUqG6to0cEYSiaWGGyheTgnnM0+A+vECQQDBB6LJcsyJz4PbUf93
GCjA1ZmMjfG3EQNNXN4Tv6a9MQhS2s6EcFFyVJ3DXR0y5TmcJBp8+qnPGni/MYNR
x3FVAkBiwR2Hq8dGFW28ajFDorz1TXabuF2VynlUxhPPsNK4ZL1iHj5ylNfT820u
mdBZG4RkIbq57ThArPJ0Q7iTL10xAkEAg9cNTn5ESZQcwT2+OL+dhifeGmwQvjK2
iK53WBG8dtf4kW61Qyvb1TUKFBeTC+IYeRcHqHkkZjn0YtIUmVSZiQJBAIFNhugu
SgXHB90v4TjQywakA22VB5L5dFaaiJHTdgOR/d3T7ybAePKSO36TGI/BCdC/rzD7
q0NtEA27TL+FrtY=
-----END PRIVATE KEY-----`

func TestNew(t *testing.T) {
	_ = rsa.New(pub, pri)
}

func TestEncrypt(t *testing.T) {
	r := rsa.New(pub, pri)
	fmt.Println(r.Encrypt([]byte("欢迎使用exlib ")).Base64())
}

func TestDecrypt(t *testing.T) {
	r := rsa.New(pub, pri)
	b, _ := crypto.FromBase64("E1WrQecs/k4FlnvoGdRaNwRKln8EaAmDPmi33uMCfzVRBUAJ9ZwENOHh4VF2OWcccsSmkUS8GQ2h8Fn6XcankUnEqH6dhPTUWmV4L56tbd/1fkondT2WTwhBAeGJ7pVYt1Q5TV5ZF6LhTVdkizG1PIORYpWgwBtynIQw+1zMkOI=")
	fmt.Println(r.Decrypt(b).String())
}

func TestSign(t *testing.T) {
	r := rsa.New(pub, pri, crypto.SignTypSHA1)
	fmt.Println(r.Sign([]byte("欢迎使用exlib ")).Base64())
}

func TestVerify(t *testing.T) {
	r := rsa.New(pub, pri, crypto.SignTypSHA1)
	b, _ := crypto.FromBase64("OpMFkN+aJOWPdQp0elx7AH2e6C5gJlQEcalOiQOEsooeRMLbwQGuqUcYwTZz3ccU0P2OUAliWPKWjarJ792sYcUHV6vvZFN3p4iQqv0wTcRXPibWKE6+9Ua+5dVpmplE2MRA85blUD7hq9QAzxd0HHlOklzzHU8QysPMMAxM+/Y=")
	fmt.Println(r.Verify([]byte("欢迎使用exlib "), b).Error())
}
