package rsa

import (
	"crypto/rand"
	cptRsa "crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

//获取密钥对
func KeyPair(bits int) (pubPem, priPem *Ret, err error) {
	priKey, err := cptRsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	priKeyDer := x509.MarshalPKCS1PrivateKey(priKey)
	priKeyBlock := pem.Block{
		Type:    "rsa PRIVATE KEY",
		Headers: nil,
		Bytes:   priKeyDer,
	}
	priKeyPem := pem.EncodeToMemory(&priKeyBlock)
	pubKey := priKey.PublicKey
	pubKeyDer, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubKeyDer,
	}
	pubKeyPem := pem.EncodeToMemory(&pubKeyBlock)
	return &Ret{v: pubKeyPem}, &Ret{v: priKeyPem}, nil
}
