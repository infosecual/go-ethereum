package dbtest

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

// skipping Fuzz_Nosy_BenchDatabaseSuite__ because parameters include func, chan, or unsupported interface: func() github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_TestDatabaseSuite__ because parameters include func, chan, or unsupported interface: func() github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_iterateKeys__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

func Fuzz_Nosy_makeDataset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, ksize int, vsize int, order bool) {
		makeDataset(size, ksize, vsize, order)
	})
}

func Fuzz_Nosy_randBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, len int) {
		randBytes(len)
	})
}
