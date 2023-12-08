package hexutil

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

func Fuzz_Nosy_Big_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Big_ToInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ToInt()
	})
}

// skipping Fuzz_Nosy_Big_UnmarshalGraphQL__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Big_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Big_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

// skipping Fuzz_Nosy_Bytes_UnmarshalGraphQL__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Bytes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Bytes_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

func Fuzz_Nosy_U256_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *U256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_U256_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *U256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_U256_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *U256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Uint_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Uint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Uint_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Uint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

// skipping Fuzz_Nosy_Uint64_UnmarshalGraphQL__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Uint64_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Uint64
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Uint64_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Uint64
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Big_ImplementsGraphQLType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		b.ImplementsGraphQLType(name)
	})
}

func Fuzz_Nosy_Big_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Big
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes_ImplementsGraphQLType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		b.ImplementsGraphQLType(name)
	})
}

func Fuzz_Nosy_Bytes_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_U256_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b U256
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Uint_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Uint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Uint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Uint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_Uint64_ImplementsGraphQLType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Uint64
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		b.ImplementsGraphQLType(name)
	})
}

func Fuzz_Nosy_Uint64_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Uint64
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Uint64_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Uint64
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_decError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err decError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		Decode(input)
	})
}

func Fuzz_Nosy_DecodeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		DecodeUint64(input)
	})
}

func Fuzz_Nosy_Encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		Encode(b)
	})
}

func Fuzz_Nosy_EncodeBig__(f *testing.F) {
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
		if bigint == nil {
			return
		}

		EncodeBig(bigint)
	})
}

func Fuzz_Nosy_EncodeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		EncodeUint64(i)
	})
}

func Fuzz_Nosy_MustDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		MustDecode(input)
	})
}

func Fuzz_Nosy_MustDecodeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		MustDecodeUint64(input)
	})
}

func Fuzz_Nosy_bytesHave0xPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		bytesHave0xPrefix(input)
	})
}

func Fuzz_Nosy_checkNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		checkNumber(input)
	})
}

func Fuzz_Nosy_checkNumberText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		checkNumberText(input)
	})
}

func Fuzz_Nosy_checkText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, wantPrefix bool) {
		checkText(input, wantPrefix)
	})
}

func Fuzz_Nosy_decodeNibble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in byte) {
		decodeNibble(in)
	})
}

func Fuzz_Nosy_has0xPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		has0xPrefix(input)
	})
}

func Fuzz_Nosy_isString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		isString(input)
	})
}
