package aes

import (
	cptAes "crypto/aes"
	"crypto/cipher"
	"github.com/incodinggo/exlib/crypto"
)

type aes struct {
	key     []byte
	iv      []byte
	padding int
	mode    int
}

//填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func New(key, iv string, padding, mode int) *aes {
	a := &aes{
		key:     []byte(key),
		iv:      []byte(iv),
		padding: padding,
		mode:    mode,
	}
	l := len(a.key)
	if l != 16 && l != 24 && l != 32 {
		panic("AES key length must 16/24/32 Bits")
	}
	return a
}

func (a *aes) cipher(block cipher.Block, de bool) (cipher.BlockMode, cipher.Stream) {
	if de {
		switch a.mode {
		case crypto.ECB:
			return crypto.NewECBDecrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCDecrypter(block, a.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, a.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, a.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBDecrypter(block, a.iv)
		default:
			return nil, nil
		}
	} else {
		switch a.mode {
		case crypto.ECB:
			return crypto.NewECBEncrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCEncrypter(block, a.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, a.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, a.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBEncrypter(block, a.iv)
		default:
			return nil, nil
		}
	}
}

func (a *aes) Encrypt(origin string) *Ret {
	b := []byte(origin)
	block, err := cptAes.NewCipher(a.key)
	if err != nil {
		return &Ret{err: err}
	}
	var padded []byte
	bs := block.BlockSize()
	padded = crypto.Padding(b, a.padding, bs)
	en := make([]byte, len(padded))
	bm, s := a.cipher(block, false)
	if s != nil {
		s.XORKeyStream(en, padded)
		return &Ret{v: en}
	}
	bm.CryptBlocks(en, padded)
	return &Ret{v: en}
}

func (a *aes) Decrypt(origin []byte) *Ret {
	block, err := cptAes.NewCipher(a.key)
	if err != nil {
		return &Ret{err: err}
	}
	de := make([]byte, len(origin))
	bm, s := a.cipher(block, true)
	if s != nil {
		s.XORKeyStream(de, origin)
	} else {
		bm.CryptBlocks(de, origin)
	}
	unPadded := crypto.UnPadding(de, a.padding)
	return &Ret{v: unPadded}
}
