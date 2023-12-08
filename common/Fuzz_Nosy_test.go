package common

import (
	"math/rand"
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

// skipping Fuzz_Nosy_Address_Scan__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Address_SetBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b1 []byte, b2 []byte) {
		a := BytesToAddress(b1)
		a.SetBytes(b2)
	})
}

// skipping Fuzz_Nosy_Address_UnmarshalGraphQL__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Address_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, input []byte) {
		a := BytesToAddress(b)
		a.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Address_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, input []byte) {
		a := BytesToAddress(b)
		a.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Address_checksumHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.checksumHex()
	})
}

func Fuzz_Nosy_Decimal_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Decimal
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.UnmarshalJSON(input)
	})
}

// skipping Fuzz_Nosy_Hash_Scan__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Hash_SetBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b1 []byte, b2 []byte) {
		h := BytesToHash(b1)
		h.SetBytes(b2)
	})
}

// skipping Fuzz_Nosy_Hash_UnmarshalGraphQL__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Hash_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, input []byte) {
		h := BytesToHash(b)
		h.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Hash_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, input []byte) {
		h := BytesToHash(b)
		h.UnmarshalText(input)
	})
}

func Fuzz_Nosy_MixedcaseAddress_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.Address()
	})
}

func Fuzz_Nosy_MixedcaseAddress_Original__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.Original()
	})
}

func Fuzz_Nosy_MixedcaseAddress_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.String()
	})
}

func Fuzz_Nosy_MixedcaseAddress_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_MixedcaseAddress_ValidChecksum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.ValidChecksum()
	})
}

func Fuzz_Nosy_UnprefixedAddress_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *UnprefixedAddress
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.UnmarshalText(input)
	})
}

func Fuzz_Nosy_UnprefixedHash_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *UnprefixedHash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Address_Big__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.Big()
	})
}

func Fuzz_Nosy_Address_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.Bytes()
	})
}

func Fuzz_Nosy_Address_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var other Address
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		a := BytesToAddress(b)
		a.Cmp(other)
	})
}

// skipping Fuzz_Nosy_Address_Format__ because parameters include func, chan, or unsupported interface: fmt.State

func Fuzz_Nosy_Address_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.Hex()
	})
}

func Fuzz_Nosy_Address_ImplementsGraphQLType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, name string) {
		a := BytesToAddress(b)
		a.ImplementsGraphQLType(name)
	})
}

func Fuzz_Nosy_Address_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.MarshalText()
	})
}

func Fuzz_Nosy_Address_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.String()
	})
}

func Fuzz_Nosy_Address_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.Value()
	})
}

func Fuzz_Nosy_Address_hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		a := BytesToAddress(b)
		a.hex()
	})
}

func Fuzz_Nosy_AddressEIP55_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr AddressEIP55
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		addr.MarshalJSON()
	})
}

func Fuzz_Nosy_AddressEIP55_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr AddressEIP55
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		addr.String()
	})
}

func Fuzz_Nosy_Hash_Big__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.Big()
	})
}

func Fuzz_Nosy_Hash_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.Bytes()
	})
}

func Fuzz_Nosy_Hash_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var other Hash
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		h := BytesToHash(b)
		h.Cmp(other)
	})
}

// skipping Fuzz_Nosy_Hash_Format__ because parameters include func, chan, or unsupported interface: fmt.State

func Fuzz_Nosy_Hash_Generate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var rand *rand.Rand
		fill_err = tp.Fill(&rand)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if rand == nil {
			return
		}

		h := BytesToHash(b)
		h.Generate(rand, size)
	})
}

func Fuzz_Nosy_Hash_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.Hex()
	})
}

func Fuzz_Nosy_Hash_ImplementsGraphQLType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, name string) {
		_x1 := BytesToHash(b)
		_x1.ImplementsGraphQLType(name)
	})
}

func Fuzz_Nosy_Hash_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.MarshalText()
	})
}

func Fuzz_Nosy_Hash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.String()
	})
}

func Fuzz_Nosy_Hash_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.TerminalString()
	})
}

func Fuzz_Nosy_Hash_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		h := BytesToHash(b)
		h.Value()
	})
}

func Fuzz_Nosy_MixedcaseAddress_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ma := NewMixedcaseAddress(addr)
		ma.MarshalJSON()
	})
}

func Fuzz_Nosy_PrettyAge_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 PrettyAge
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.String()
	})
}

func Fuzz_Nosy_PrettyDuration_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d PrettyDuration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

func Fuzz_Nosy_StorageSize_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StorageSize
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_StorageSize_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StorageSize
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.TerminalString()
	})
}

func Fuzz_Nosy_UnprefixedAddress_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a UnprefixedAddress
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.MarshalText()
	})
}

func Fuzz_Nosy_UnprefixedHash_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h UnprefixedHash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.MarshalText()
	})
}

func Fuzz_Nosy_AbsolutePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string, filename string) {
		AbsolutePath(datadir, filename)
	})
}

func Fuzz_Nosy_Bytes2Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d []byte) {
		Bytes2Hex(d)
	})
}

func Fuzz_Nosy_CopyBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		CopyBytes(b)
	})
}

func Fuzz_Nosy_FileExist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		FileExist(filePath)
	})
}

func Fuzz_Nosy_FromHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		FromHex(s)
	})
}

func Fuzz_Nosy_Hex2Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		Hex2Bytes(str)
	})
}

func Fuzz_Nosy_Hex2BytesFixed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string, flen int) {
		Hex2BytesFixed(str, flen)
	})
}

func Fuzz_Nosy_IsHexAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		IsHexAddress(s)
	})
}

func Fuzz_Nosy_LeftPadBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, slice []byte, l int) {
		LeftPadBytes(slice, l)
	})
}

func Fuzz_Nosy_ParseHexOrString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		ParseHexOrString(str)
	})
}

func Fuzz_Nosy_PrintDeprecationWarning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		PrintDeprecationWarning(str)
	})
}

// skipping Fuzz_Nosy_Report__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_RightPadBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, slice []byte, l int) {
		RightPadBytes(slice, l)
	})
}

func Fuzz_Nosy_TrimLeftZeroes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		TrimLeftZeroes(s)
	})
}

func Fuzz_Nosy_TrimRightZeroes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		TrimRightZeroes(s)
	})
}

func Fuzz_Nosy_findLine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, offset int64) {
		findLine(d1, offset)
	})
}

func Fuzz_Nosy_has0xPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		has0xPrefix(str)
	})
}

func Fuzz_Nosy_isHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		isHex(str)
	})
}

func Fuzz_Nosy_isHexCharacter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, c byte) {
		isHexCharacter(c)
	})
}

func Fuzz_Nosy_isString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		isString(input)
	})
}
