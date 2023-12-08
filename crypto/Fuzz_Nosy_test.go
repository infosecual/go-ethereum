package crypto

import (
	"bufio"
	"crypto/ecdsa"
	"math/big"
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

func Fuzz_Nosy_KeccakState_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 []byte) {
		_x1 := NewKeccakState()
		_x1.Read(_x1)
	})
}

func Fuzz_Nosy_CompressPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if pubkey == nil {
			return
		}

		CompressPubkey(pubkey)
	})
}

func Fuzz_Nosy_Ecrecover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, sig []byte) {
		Ecrecover(hash, sig)
	})
}

func Fuzz_Nosy_FromECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var priv *ecdsa.PrivateKey
		fill_err = tp.Fill(&priv)
		if fill_err != nil {
			return
		}
		if priv == nil {
			return
		}

		FromECDSA(priv)
	})
}

func Fuzz_Nosy_FromECDSAPub__(f *testing.F) {
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

		FromECDSAPub(pub)
	})
}

func Fuzz_Nosy_Keccak256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 [][]byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		Keccak256(d1...)
	})
}

func Fuzz_Nosy_Keccak512__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 [][]byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		Keccak512(d1...)
	})
}

func Fuzz_Nosy_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var digestHash []byte
		fill_err = tp.Fill(&digestHash)
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

		Sign(digestHash, prv)
	})
}

func Fuzz_Nosy_ValidateSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v byte
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r *big.Int
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var homestead bool
		fill_err = tp.Fill(&homestead)
		if fill_err != nil {
			return
		}
		if r == nil || s == nil {
			return
		}

		ValidateSignatureValues(v, r, s, homestead)
	})
}

func Fuzz_Nosy_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pubkey []byte, digestHash []byte, signature []byte) {
		VerifySignature(pubkey, digestHash, signature)
	})
}

func Fuzz_Nosy_readASCII__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		var r *bufio.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		readASCII(buf, r)
	})
}

func Fuzz_Nosy_zeroBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bytes []byte) {
		zeroBytes(bytes)
	})
}
