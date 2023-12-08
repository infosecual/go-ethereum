package math

import (
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

func Fuzz_Nosy_Decimal256_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64) {
		i := NewDecimal256(x)
		i.MarshalText()
	})
}

func Fuzz_Nosy_Decimal256_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64) {
		i := NewDecimal256(x)
		i.String()
	})
}

func Fuzz_Nosy_Decimal256_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64, input []byte) {
		i := NewDecimal256(x)
		i.UnmarshalText(input)
	})
}

func Fuzz_Nosy_HexOrDecimal256_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64) {
		i := NewHexOrDecimal256(x)
		i.MarshalText()
	})
}

func Fuzz_Nosy_HexOrDecimal256_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64, input []byte) {
		i := NewHexOrDecimal256(x)
		i.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_HexOrDecimal256_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64, input []byte) {
		i := NewHexOrDecimal256(x)
		i.UnmarshalText(input)
	})
}

func Fuzz_Nosy_HexOrDecimal64_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *HexOrDecimal64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_HexOrDecimal64_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *HexOrDecimal64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.UnmarshalText(input)
	})
}

func Fuzz_Nosy_HexOrDecimal64_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i HexOrDecimal64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.MarshalText()
	})
}

func Fuzz_Nosy_Byte__(f *testing.F) {
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
		var padlength int
		fill_err = tp.Fill(&padlength)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		Byte(bigint, padlength, n)
	})
}

func Fuzz_Nosy_FirstBitSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		FirstBitSet(v)
	})
}

func Fuzz_Nosy_MustParseUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		MustParseUint64(s)
	})
}

func Fuzz_Nosy_PaddedBigBytes__(f *testing.F) {
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
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		PaddedBigBytes(bigint, n)
	})
}

func Fuzz_Nosy_ParseBig256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		ParseBig256(s)
	})
}

func Fuzz_Nosy_ParseUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		ParseUint64(s)
	})
}

func Fuzz_Nosy_ReadBits__(f *testing.F) {
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

		ReadBits(bigint, buf)
	})
}

func Fuzz_Nosy_SafeAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64, y uint64) {
		SafeAdd(x, y)
	})
}

func Fuzz_Nosy_SafeMul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64, y uint64) {
		SafeMul(x, y)
	})
}

func Fuzz_Nosy_SafeSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64, y uint64) {
		SafeSub(x, y)
	})
}

func Fuzz_Nosy_U256Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *big.Int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		U256Bytes(n)
	})
}

func Fuzz_Nosy_bigEndianByteAt__(f *testing.F) {
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
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		bigEndianByteAt(bigint, n)
	})
}
