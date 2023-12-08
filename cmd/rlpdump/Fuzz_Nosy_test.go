package main

import (
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

// skipping Fuzz_Nosy_die__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_isASCII__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		isASCII(b)
	})
}

func Fuzz_Nosy_textToRlp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		textToRlp(r)
	})
}

func Fuzz_Nosy_ws__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		ws(n)
	})
}
