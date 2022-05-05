package tripledes

import (
	"crypto/cipher"
	"crypto/des"
	"github.com/incodinggo/exlib/crypto"
)

type Tripledes struct {
	key     []byte
	iv      []byte
	padding int
	mode    int
}

func New(key, iv string, padding, mode int) *Tripledes {
	d := &Tripledes{
		key:     []byte(key),
		iv:      []byte(iv),
		padding: padding,
		mode:    mode,
	}
	kl := len(d.key)
	ivl := len(d.iv)
	if kl != 24 && kl != 32 {
		panic("Tripledes key length must 24/32.")
	}
	if ivl != 24 && ivl != 32 {
		panic("Tripledes iv length must 24/32.")
	}
	return d
}

func (t *Tripledes) cipher(block cipher.Block, de bool) (cipher.BlockMode, cipher.Stream) {
	if de {
		switch t.mode {
		case crypto.ECB:
			return crypto.NewECBDecrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCDecrypter(block, t.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, t.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, t.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBDecrypter(block, t.iv)
		default:
			return nil, nil
		}
	} else {
		switch t.mode {
		case crypto.ECB:
			return crypto.NewECBEncrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCEncrypter(block, t.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, t.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, t.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBEncrypter(block, t.iv)
		default:
			return nil, nil
		}
	}
}
func (t *Tripledes) Encrypt(origin string) *Ret {
	b := []byte(origin)
	block, err := des.NewTripleDESCipher(t.key)
	if err != nil {
		return &Ret{err: err}
	}
	padded := crypto.Padding(b, t.padding, block.BlockSize())
	en := make([]byte, len(padded))
	bm, s := t.cipher(block, false)
	if s != nil {
		s.XORKeyStream(en, padded)
		return &Ret{v: en}
	}
	bm.CryptBlocks(en, padded)
	return &Ret{v: en}
}

func (t *Tripledes) Decrypt(origin []byte) *Ret {
	block, err := des.NewTripleDESCipher(t.key)
	if err != nil {
		return &Ret{err: err}
	}
	de := make([]byte, len(origin))
	bm, s := t.cipher(block, true)
	if s != nil {
		s.XORKeyStream(de, origin)
	} else {
		bm.CryptBlocks(de, origin)
	}
	unPadded := crypto.UnPadding(de, t.padding)
	return &Ret{v: unPadded}
}
