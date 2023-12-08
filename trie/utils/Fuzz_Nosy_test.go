package utils

import (
	"testing"

	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/holiman/uint256"
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

func Fuzz_Nosy_PointCache_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxItems int, addr []byte) {
		c := NewPointCache(maxItems)
		c.Get(addr)
	})
}

func Fuzz_Nosy_PointCache_GetStem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxItems int, addr []byte) {
		c := NewPointCache(maxItems)
		c.GetStem(addr)
	})
}

func Fuzz_Nosy_BalanceKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte) {
		BalanceKey(address)
	})
}

func Fuzz_Nosy_BalanceKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		BalanceKeyWithEvaluatedAddress(evaluated)
	})
}

func Fuzz_Nosy_CodeChunkKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address []byte
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var chunk *uint256.Int
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if chunk == nil {
			return
		}

		CodeChunkKey(address, chunk)
	})
}

func Fuzz_Nosy_CodeChunkKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addressPoint *banderwagon.Element
		fill_err = tp.Fill(&addressPoint)
		if fill_err != nil {
			return
		}
		var chunk *uint256.Int
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if addressPoint == nil || chunk == nil {
			return
		}

		CodeChunkKeyWithEvaluatedAddress(addressPoint, chunk)
	})
}

func Fuzz_Nosy_CodeKeccakKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte) {
		CodeKeccakKey(address)
	})
}

func Fuzz_Nosy_CodeKeccakKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		CodeKeccakKeyWithEvaluatedAddress(evaluated)
	})
}

func Fuzz_Nosy_CodeSizeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte) {
		CodeSizeKey(address)
	})
}

func Fuzz_Nosy_CodeSizeKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		CodeSizeKeyWithEvaluatedAddress(evaluated)
	})
}

func Fuzz_Nosy_GetTreeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address []byte
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var treeIndex *uint256.Int
		fill_err = tp.Fill(&treeIndex)
		if fill_err != nil {
			return
		}
		var subIndex byte
		fill_err = tp.Fill(&subIndex)
		if fill_err != nil {
			return
		}
		if treeIndex == nil {
			return
		}

		GetTreeKey(address, treeIndex, subIndex)
	})
}

func Fuzz_Nosy_GetTreeKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		var treeIndex *uint256.Int
		fill_err = tp.Fill(&treeIndex)
		if fill_err != nil {
			return
		}
		var subIndex byte
		fill_err = tp.Fill(&subIndex)
		if fill_err != nil {
			return
		}
		if evaluated == nil || treeIndex == nil {
			return
		}

		GetTreeKeyWithEvaluatedAddress(evaluated, treeIndex, subIndex)
	})
}

func Fuzz_Nosy_NonceKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte) {
		NonceKey(address)
	})
}

func Fuzz_Nosy_NonceKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		NonceKeyWithEvaluatedAddress(evaluated)
	})
}

func Fuzz_Nosy_StorageSlotKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte, storageKey []byte) {
		StorageSlotKey(address, storageKey)
	})
}

func Fuzz_Nosy_StorageSlotKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		var storageKey []byte
		fill_err = tp.Fill(&storageKey)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		StorageSlotKeyWithEvaluatedAddress(evaluated, storageKey)
	})
}

func Fuzz_Nosy_VersionKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address []byte) {
		VersionKey(address)
	})
}

func Fuzz_Nosy_VersionKeyWithEvaluatedAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		VersionKeyWithEvaluatedAddress(evaluated)
	})
}

func Fuzz_Nosy_codeChunkIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chunk *uint256.Int
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if chunk == nil {
			return
		}

		codeChunkIndex(chunk)
	})
}

func Fuzz_Nosy_pointToHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evaluated *banderwagon.Element
		fill_err = tp.Fill(&evaluated)
		if fill_err != nil {
			return
		}
		var suffix byte
		fill_err = tp.Fill(&suffix)
		if fill_err != nil {
			return
		}
		if evaluated == nil {
			return
		}

		pointToHash(evaluated, suffix)
	})
}

func Fuzz_Nosy_storageIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bytes []byte) {
		storageIndex(bytes)
	})
}
