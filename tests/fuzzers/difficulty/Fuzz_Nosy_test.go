package difficulty

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

func Fuzz_Nosy_fuzzer_fuzz__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fuzzer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.fuzz()
	})
}

func Fuzz_Nosy_fuzzer_read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fuzzer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.read(size)
	})
}

func Fuzz_Nosy_fuzzer_readBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fuzzer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.readBool()
	})
}

func Fuzz_Nosy_fuzzer_readSlice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fuzzer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var min int
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.readSlice(min, max)
	})
}

func Fuzz_Nosy_fuzzer_readUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fuzzer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var min uint64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max uint64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.readUint64(min, max)
	})
}

func Fuzz_Nosy_fuzz__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		fuzz(d1)
	})
}
