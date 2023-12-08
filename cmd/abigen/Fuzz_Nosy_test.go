package main

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

func Fuzz_Nosy_nameFilter_Matches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var patterns []string
		fill_err = tp.Fill(&patterns)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		f1, err := newNameFilter(patterns...)
		if err != nil {
			return
		}
		f1.Matches(name)
	})
}

func Fuzz_Nosy_nameFilter_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var patterns []string
		fill_err = tp.Fill(&patterns)
		if fill_err != nil {
			return
		}
		var pattern string
		fill_err = tp.Fill(&pattern)
		if fill_err != nil {
			return
		}

		f1, err := newNameFilter(patterns...)
		if err != nil {
			return
		}
		f1.add(pattern)
	})
}
