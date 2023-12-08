package bls

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

func Fuzz_Nosy_checkInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id byte, inputLen int) {
		checkInput(id, inputLen)
	})
}

func Fuzz_Nosy_fuzz__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id byte, d2 []byte) {
		fuzz(id, d2)
	})
}

func Fuzz_Nosy_fuzzCrossG1Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzCrossG1Add(d1)
	})
}

func Fuzz_Nosy_fuzzCrossG1MultiExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzCrossG1MultiExp(d1)
	})
}

func Fuzz_Nosy_fuzzCrossG2Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzCrossG2Add(d1)
	})
}

func Fuzz_Nosy_fuzzCrossPairing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzzCrossPairing(d1)
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

func Fuzz_Nosy_massageBLST__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		massageBLST(in)
	})
}
