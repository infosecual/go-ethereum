package rlpstruct

import (
	"reflect"
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

func Fuzz_Nosy_TagError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e TagError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_Type_DefaultNilValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.DefaultNilValue()
	})
}

func Fuzz_Nosy_ProcessFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allFields []Field
		fill_err = tp.Fill(&allFields)
		if fill_err != nil {
			return
		}

		ProcessFields(allFields)
	})
}

func Fuzz_Nosy_isByte__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		isByte(typ)
	})
}

func Fuzz_Nosy_isByteArray__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}

		isByteArray(typ)
	})
}

func Fuzz_Nosy_isUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k reflect.Kind
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		isUint(k)
	})
}

func Fuzz_Nosy_lastPublicField__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fields []Field
		fill_err = tp.Fill(&fields)
		if fill_err != nil {
			return
		}

		lastPublicField(fields)
	})
}
