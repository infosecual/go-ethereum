package pruner

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_Pruner_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Pruner
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Prune(root)
	})
}

func Fuzz_Nosy_stateBloom_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64, filename string, tempname string) {
		bloom, err := newStateBloomWithSize(size)
		if err != nil {
			return
		}
		bloom.Commit(filename, tempname)
	})
}

func Fuzz_Nosy_stateBloom_Contain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64, key []byte) {
		bloom, err := newStateBloomWithSize(size)
		if err != nil {
			return
		}
		bloom.Contain(key)
	})
}

func Fuzz_Nosy_stateBloom_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64, key []byte) {
		bloom, err := newStateBloomWithSize(size)
		if err != nil {
			return
		}
		bloom.Delete(key)
	})
}

func Fuzz_Nosy_stateBloom_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64, key []byte, value []byte) {
		bloom, err := newStateBloomWithSize(size)
		if err != nil {
			return
		}
		bloom.Put(key, value)
	})
}

func Fuzz_Nosy_stateBloomHasher_BlockSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.BlockSize()
	})
}

func Fuzz_Nosy_stateBloomHasher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.Reset()
	})
}

func Fuzz_Nosy_stateBloomHasher_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.Size()
	})
}

func Fuzz_Nosy_stateBloomHasher_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		f1.Sum(b)
	})
}

func Fuzz_Nosy_stateBloomHasher_Sum64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.Sum64()
	})
}

func Fuzz_Nosy_stateBloomHasher_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 stateBloomHasher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		f1.Write(p)
	})
}

func Fuzz_Nosy_bloomFilterName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		bloomFilterName(datadir, hash)
	})
}

func Fuzz_Nosy_findBloomFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string) {
		findBloomFilter(datadir)
	})
}

func Fuzz_Nosy_isBloomFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filename string) {
		isBloomFilter(filename)
	})
}
