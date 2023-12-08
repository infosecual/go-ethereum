package bn256

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

func Fuzz_Nosy_fuzzAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzAdd(d1)
	})
}

func Fuzz_Nosy_fuzzMul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzMul(d1)
	})
}

func Fuzz_Nosy_fuzzPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzPair(d1)
	})
}

func Fuzz_Nosy_getG1Points__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input io.Reader
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		getG1Points(input)
	})
}

func Fuzz_Nosy_getG2Points__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input io.Reader
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		getG2Points(input)
	})
}
