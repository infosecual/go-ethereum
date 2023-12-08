package secp256k1

import (
	"math/big"
	"testing"
	"unsafe"

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

func Fuzz_Nosy_BitCurve_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x1 *big.Int
		fill_err = tp.Fill(&x1)
		if fill_err != nil {
			return
		}
		var y1 *big.Int
		fill_err = tp.Fill(&y1)
		if fill_err != nil {
			return
		}
		var x2 *big.Int
		fill_err = tp.Fill(&x2)
		if fill_err != nil {
			return
		}
		var y2 *big.Int
		fill_err = tp.Fill(&y2)
		if fill_err != nil {
			return
		}
		if x1 == nil || y1 == nil || x2 == nil || y2 == nil {
			return
		}

		BitCurve := S256()
		BitCurve.Add(x1, y1, x2, y2)
	})
}

func Fuzz_Nosy_BitCurve_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x1 *big.Int
		fill_err = tp.Fill(&x1)
		if fill_err != nil {
			return
		}
		var y1 *big.Int
		fill_err = tp.Fill(&y1)
		if fill_err != nil {
			return
		}
		if x1 == nil || y1 == nil {
			return
		}

		BitCurve := S256()
		BitCurve.Double(x1, y1)
	})
}

func Fuzz_Nosy_BitCurve_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		BitCurve := S256()
		BitCurve.IsOnCurve(x, y)
	})
}

func Fuzz_Nosy_BitCurve_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		BitCurve := S256()
		BitCurve.Marshal(x, y)
	})
}

func Fuzz_Nosy_BitCurve_ScalarBaseMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, k []byte) {
		BitCurve := S256()
		BitCurve.ScalarBaseMult(k)
	})
}

func Fuzz_Nosy_BitCurve_ScalarMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var Bx *big.Int
		fill_err = tp.Fill(&Bx)
		if fill_err != nil {
			return
		}
		var By *big.Int
		fill_err = tp.Fill(&By)
		if fill_err != nil {
			return
		}
		var scalar []byte
		fill_err = tp.Fill(&scalar)
		if fill_err != nil {
			return
		}
		if Bx == nil || By == nil {
			return
		}

		BitCurve := S256()
		BitCurve.ScalarMult(Bx, By, scalar)
	})
}

func Fuzz_Nosy_BitCurve_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		BitCurve := S256()
		BitCurve.Unmarshal(d1)
	})
}

func Fuzz_Nosy_BitCurve_addJacobian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x1 *big.Int
		fill_err = tp.Fill(&x1)
		if fill_err != nil {
			return
		}
		var y1 *big.Int
		fill_err = tp.Fill(&y1)
		if fill_err != nil {
			return
		}
		var z1 *big.Int
		fill_err = tp.Fill(&z1)
		if fill_err != nil {
			return
		}
		var x2 *big.Int
		fill_err = tp.Fill(&x2)
		if fill_err != nil {
			return
		}
		var y2 *big.Int
		fill_err = tp.Fill(&y2)
		if fill_err != nil {
			return
		}
		var z2 *big.Int
		fill_err = tp.Fill(&z2)
		if fill_err != nil {
			return
		}
		if x1 == nil || y1 == nil || z1 == nil || x2 == nil || y2 == nil || z2 == nil {
			return
		}

		BitCurve := S256()
		BitCurve.addJacobian(x1, y1, z1, x2, y2, z2)
	})
}

func Fuzz_Nosy_BitCurve_affineFromJacobian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var z *big.Int
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil || z == nil {
			return
		}

		BitCurve := S256()
		BitCurve.affineFromJacobian(x, y, z)
	})
}

func Fuzz_Nosy_BitCurve_doubleJacobian__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		var z *big.Int
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil || z == nil {
			return
		}

		BitCurve := S256()
		BitCurve.doubleJacobian(x, y, z)
	})
}

func Fuzz_Nosy_CompressPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		CompressPubkey(x, y)
	})
}

func Fuzz_Nosy_DecompressPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pubkey []byte) {
		DecompressPubkey(pubkey)
	})
}

func Fuzz_Nosy_RecoverPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, msg []byte, sig []byte) {
		RecoverPubkey(msg, sig)
	})
}

func Fuzz_Nosy_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, msg []byte, seckey []byte) {
		Sign(msg, seckey)
	})
}

func Fuzz_Nosy_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pubkey []byte, msg []byte, signature []byte) {
		VerifySignature(pubkey, msg, signature)
	})
}

func Fuzz_Nosy__Cfunc_GoString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *_Ctype_char
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		_Cfunc_GoString(p)
	})
}

func Fuzz_Nosy__Cgo_ptr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ptr unsafe.Pointer
		fill_err = tp.Fill(&ptr)
		if fill_err != nil {
			return
		}

		_Cgo_ptr(ptr)
	})
}

// skipping Fuzz_Nosy__Cgo_use__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__cgoCheckPointer__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__cgoCheckResult__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy__cgo_runtime_cgocall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 unsafe.Pointer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 uintptr
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_cgo_runtime_cgocall(_x1, _x2)
	})
}

func Fuzz_Nosy__cgo_runtime_gostring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *_Ctype_char
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_cgo_runtime_gostring(_x1)
	})
}

func Fuzz_Nosy__cgoexp_e3fed97ebf03_secp256k1GoPanicError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *struct {
			p0 *_Ctype_char
			p1 unsafe.Pointer
		}
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		_cgoexp_e3fed97ebf03_secp256k1GoPanicError(a)
	})
}

func Fuzz_Nosy__cgoexp_e3fed97ebf03_secp256k1GoPanicIllegal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *struct {
			p0 *_Ctype_char
			p1 unsafe.Pointer
		}
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		_cgoexp_e3fed97ebf03_secp256k1GoPanicIllegal(a)
	})
}

func Fuzz_Nosy_readBits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		readBits(bigint, buf)
	})
}

func Fuzz_Nosy_secp256k1GoPanicError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *_Ctype_char
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var d2 unsafe.Pointer
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		secp256k1GoPanicError(msg, d2)
	})
}

func Fuzz_Nosy_secp256k1GoPanicIllegal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *_Ctype_char
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var d2 unsafe.Pointer
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		secp256k1GoPanicIllegal(msg, d2)
	})
}
