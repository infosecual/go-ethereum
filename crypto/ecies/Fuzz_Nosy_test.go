package ecies

import (
	"crypto/ecdsa"
	"io"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_PrivateKey_Decrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		var c []byte
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var s1 []byte
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var s2 []byte
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		if prv == nil {
			return
		}

		p1 := ImportECDSA(prv)
		p1.Decrypt(c, s1, s2)
	})
}

func Fuzz_Nosy_PrivateKey_ExportECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if prv == nil {
			return
		}

		p1 := ImportECDSA(prv)
		p1.ExportECDSA()
	})
}

func Fuzz_Nosy_PrivateKey_GenerateShared__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		var pub *PublicKey
		fill_err = tp.Fill(&pub)
		if fill_err != nil {
			return
		}
		var skLen int
		fill_err = tp.Fill(&skLen)
		if fill_err != nil {
			return
		}
		var macLen int
		fill_err = tp.Fill(&macLen)
		if fill_err != nil {
			return
		}
		if prv == nil || pub == nil {
			return
		}

		p1 := ImportECDSA(prv)
		p1.GenerateShared(pub, skLen, macLen)
	})
}

func Fuzz_Nosy_PublicKey_ExportECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pub *ecdsa.PublicKey
		fill_err = tp.Fill(&pub)
		if fill_err != nil {
			return
		}
		if pub == nil {
			return
		}

		p1 := ImportECDSAPublic(pub)
		p1.ExportECDSA()
	})
}

// skipping Fuzz_Nosy_AddParamsForCurve__ because parameters include func, chan, or unsupported interface: crypto/elliptic.Curve

func Fuzz_Nosy_Encrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rand io.Reader
		fill_err = tp.Fill(&rand)
		if fill_err != nil {
			return
		}
		var pub *PublicKey
		fill_err = tp.Fill(&pub)
		if fill_err != nil {
			return
		}
		var m []byte
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var s1 []byte
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var s2 []byte
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		if pub == nil {
			return
		}

		Encrypt(rand, pub, m, s1, s2)
	})
}

func Fuzz_Nosy_MaxSharedKeyLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pub *PublicKey
		fill_err = tp.Fill(&pub)
		if fill_err != nil {
			return
		}
		if pub == nil {
			return
		}

		MaxSharedKeyLength(pub)
	})
}

// skipping Fuzz_Nosy_concatKDF__ because parameters include func, chan, or unsupported interface: hash.Hash

// skipping Fuzz_Nosy_deriveKeys__ because parameters include func, chan, or unsupported interface: hash.Hash

func Fuzz_Nosy_generateIV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *ECIESParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var rand io.Reader
		fill_err = tp.Fill(&rand)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		generateIV(params, rand)
	})
}

// skipping Fuzz_Nosy_messageTag__ because parameters include func, chan, or unsupported interface: func() hash.Hash

func Fuzz_Nosy_roundup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, blocksize int) {
		roundup(size, blocksize)
	})
}

func Fuzz_Nosy_symDecrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *ECIESParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var ct []byte
		fill_err = tp.Fill(&ct)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		symDecrypt(params, key, ct)
	})
}

func Fuzz_Nosy_symEncrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rand io.Reader
		fill_err = tp.Fill(&rand)
		if fill_err != nil {
			return
		}
		var params *ECIESParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var m []byte
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		symEncrypt(rand, params, key, m)
	})
}
