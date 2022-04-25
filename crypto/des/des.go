package des

import (
	"crypto/cipher"
	cptDes "crypto/des"
	"exlib/crypto"
)

type des struct {
	key     []byte
	iv      []byte
	padding int
	mode    int
}

func New(key, iv string, padding, mode int) *des {
	d := &des{
		key:     []byte(key),
		iv:      []byte(iv),
		padding: padding,
		mode:    mode,
	}
	kl := len(d.key)
	ivl := len(d.iv)
	if kl != 8 && kl != 16 && kl != 24 && kl != 32 {
		panic("des key length must 8/16/24/32.")
	}
	if ivl != 8 && ivl != 16 && ivl != 24 && ivl != 32 {
		panic("des iv length must 8/16/24/32.")
	}
	return d
}

func (d *des) cipher(block cipher.Block, de bool) (cipher.BlockMode, cipher.Stream) {
	if de {
		switch d.mode {
		case crypto.ECB:
			return crypto.NewECBDecrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCDecrypter(block, d.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, d.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, d.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBDecrypter(block, d.iv)
		default:
			return nil, nil
		}
	} else {
		switch d.mode {
		case crypto.ECB:
			return crypto.NewECBEncrypter(block), nil
		case crypto.CBC:
			return cipher.NewCBCEncrypter(block, d.iv), nil
		case crypto.CTR:
			return nil, cipher.NewCTR(block, d.iv)
		case crypto.OFB:
			return nil, cipher.NewOFB(block, d.iv)
		case crypto.CFB:
			return nil, cipher.NewCFBEncrypter(block, d.iv)
		default:
			return nil, nil
		}
	}
}

func (d *des) Encrypt(origin string) *Ret {
	b := []byte(origin)
	block, err := cptDes.NewCipher(d.key)
	if err != nil {
		return &Ret{err: err}
	}
	padded := crypto.Padding(b, d.padding, block.BlockSize())
	en := make([]byte, len(padded))
	bm, s := d.cipher(block, false)
	if s != nil {
		s.XORKeyStream(en, padded)
		return &Ret{v: en}
	}
	bm.CryptBlocks(en, padded)
	return &Ret{v: en}
}

func (d *des) Decrypt(origin []byte) *Ret {
	block, err := cptDes.NewCipher(d.key)
	if err != nil {
		return &Ret{err: err}
	}
	de := make([]byte, len(origin))
	bm, s := d.cipher(block, true)
	if s != nil {
		s.XORKeyStream(de, origin)
	} else {
		bm.CryptBlocks(de, origin)
	}
	unPadded := crypto.UnPadding(de, d.padding)
	return &Ret{v: unPadded}
}
