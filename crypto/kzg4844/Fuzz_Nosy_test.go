package kzg4844

import (
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

func Fuzz_Nosy_ComputeProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blob Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var point Point
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}

		ComputeProof(blob, point)
	})
}

func Fuzz_Nosy_ckzgComputeProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blob Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var point Point
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}

		ckzgComputeProof(blob, point)
	})
}

func Fuzz_Nosy_gokzgComputeProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blob Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var point Point
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}

		gokzgComputeProof(blob, point)
	})
}
