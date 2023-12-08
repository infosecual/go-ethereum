package rawdb

import (
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
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

func Fuzz_Nosy_Freezer_Ancient__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Ancient(kind, number)
	})
}

func Fuzz_Nosy_Freezer_AncientRange__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var maxBytes uint64
		fill_err = tp.Fill(&maxBytes)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.AncientRange(kind, start, count, maxBytes)
	})
}

func Fuzz_Nosy_Freezer_AncientSize__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.AncientSize(kind)
	})
}

func Fuzz_Nosy_Freezer_Ancients__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Ancients()
	})
}

func Fuzz_Nosy_Freezer_Close__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Close()
	})
}

func Fuzz_Nosy_Freezer_HasAncient__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.HasAncient(kind, number)
	})
}

// skipping Fuzz_Nosy_Freezer_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_Freezer_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

// skipping Fuzz_Nosy_Freezer_ReadAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_Freezer_Sync__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Sync()
	})
}

func Fuzz_Nosy_Freezer_Tail__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Tail()
	})
}

func Fuzz_Nosy_Freezer_TruncateHead__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.TruncateHead(items)
	})
}

func Fuzz_Nosy_Freezer_TruncateTail__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var tail uint64
		fill_err = tp.Fill(&tail)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.TruncateTail(tail)
	})
}

func Fuzz_Nosy_Freezer_repair__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.repair()
	})
}

func Fuzz_Nosy_Freezer_validate__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.validate()
	})
}

func Fuzz_Nosy_KeyLengthIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *KeyLengthIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_ResettableFreezer_Ancient__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Ancient(kind, number)
	})
}

func Fuzz_Nosy_ResettableFreezer_AncientRange__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var maxBytes uint64
		fill_err = tp.Fill(&maxBytes)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.AncientRange(kind, start, count, maxBytes)
	})
}

func Fuzz_Nosy_ResettableFreezer_AncientSize__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.AncientSize(kind)
	})
}

func Fuzz_Nosy_ResettableFreezer_Ancients__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Ancients()
	})
}

func Fuzz_Nosy_ResettableFreezer_Close__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Close()
	})
}

func Fuzz_Nosy_ResettableFreezer_HasAncient__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.HasAncient(kind, number)
	})
}

// skipping Fuzz_Nosy_ResettableFreezer_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_ResettableFreezer_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

// skipping Fuzz_Nosy_ResettableFreezer_ReadAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_ResettableFreezer_Reset__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Reset()
	})
}

func Fuzz_Nosy_ResettableFreezer_Sync__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Sync()
	})
}

func Fuzz_Nosy_ResettableFreezer_Tail__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.Tail()
	})
}

func Fuzz_Nosy_ResettableFreezer_TruncateHead__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.TruncateHead(items)
	})
}

func Fuzz_Nosy_ResettableFreezer_TruncateTail__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var maxTableSize uint32
		fill_err = tp.Fill(&maxTableSize)
		if fill_err != nil {
			return
		}
		var tables map[string]bool
		fill_err = tp.Fill(&tables)
		if fill_err != nil {
			return
		}
		var tail uint64
		fill_err = tp.Fill(&tail)
		if fill_err != nil {
			return
		}

		f1, err := NewResettableFreezer(datadir, namespace, readonly, maxTableSize, tables)
		if err != nil {
			return
		}
		f1.TruncateTail(tail)
	})
}

func Fuzz_Nosy_chainFreezer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string, namespace string, readonly bool) {
		f1, err := newChainFreezer(datadir, namespace, readonly)
		if err != nil {
			return
		}
		f1.Close()
	})
}

// skipping Fuzz_Nosy_chainFreezer_freeze__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

func Fuzz_Nosy_chainFreezer_freezeRange__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var nfdb *nofreezedb
		fill_err = tp.Fill(&nfdb)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if nfdb == nil {
			return
		}

		f1, err := newChainFreezer(datadir, namespace, readonly)
		if err != nil {
			return
		}
		f1.freezeRange(nfdb, number, limit)
	})
}

// skipping Fuzz_Nosy_freezerBatch_Append__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_freezerBatch_AppendRaw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Freezer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var item []byte
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		batch := newFreezerBatch(f1)
		batch.AppendRaw(kind, num, item)
	})
}

func Fuzz_Nosy_freezerBatch_commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Freezer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		batch := newFreezerBatch(f1)
		batch.commit()
	})
}

func Fuzz_Nosy_freezerBatch_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Freezer
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		batch := newFreezerBatch(f1)
		batch.reset()
	})
}

func Fuzz_Nosy_freezerInfo_count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *freezerInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if info == nil {
			return
		}

		info.count()
	})
}

func Fuzz_Nosy_freezerInfo_size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *freezerInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if info == nil {
			return
		}

		info.size()
	})
}

func Fuzz_Nosy_freezerTable_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.Close()
	})
}

func Fuzz_Nosy_freezerTable_Retrieve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, item uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.Retrieve(item)
	})
}

func Fuzz_Nosy_freezerTable_RetrieveItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, start uint64, count uint64, maxBytes uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.RetrieveItems(start, count, maxBytes)
	})
}

func Fuzz_Nosy_freezerTable_Sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.Sync()
	})
}

func Fuzz_Nosy_freezerTable_advanceHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.advanceHead()
	})
}

func Fuzz_Nosy_freezerTable_dumpIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var disableSnappy bool
		fill_err = tp.Fill(&disableSnappy)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var start int64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var stop int64
		fill_err = tp.Fill(&stop)
		if fill_err != nil {
			return
		}

		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.dumpIndex(w, start, stop)
	})
}

func Fuzz_Nosy_freezerTable_dumpIndexStdout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, start int64, stop int64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.dumpIndexStdout(start, stop)
	})
}

func Fuzz_Nosy_freezerTable_dumpIndexString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, start int64, stop int64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.dumpIndexString(start, stop)
	})
}

func Fuzz_Nosy_freezerTable_getIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, from uint64, count uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.getIndices(from, count)
	})
}

func Fuzz_Nosy_freezerTable_has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, number uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.has(number)
	})
}

func Fuzz_Nosy_freezerTable_newBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.newBatch()
	})
}

// skipping Fuzz_Nosy_freezerTable_openFile__ because parameters include func, chan, or unsupported interface: func(string) (*os.File, error)

func Fuzz_Nosy_freezerTable_preopen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.preopen()
	})
}

func Fuzz_Nosy_freezerTable_releaseFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, num uint32) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.releaseFile(num)
	})
}

func Fuzz_Nosy_freezerTable_releaseFilesAfter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, num uint32, remove bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.releaseFilesAfter(num, remove)
	})
}

func Fuzz_Nosy_freezerTable_releaseFilesBefore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, num uint32, remove bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.releaseFilesBefore(num, remove)
	})
}

func Fuzz_Nosy_freezerTable_repair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.repair()
	})
}

func Fuzz_Nosy_freezerTable_retrieveItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, start uint64, count uint64, maxBytes uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.retrieveItems(start, count, maxBytes)
	})
}

func Fuzz_Nosy_freezerTable_size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.size()
	})
}

func Fuzz_Nosy_freezerTable_sizeNolock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.sizeNolock()
	})
}

func Fuzz_Nosy_freezerTable_truncateHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, items uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.truncateHead(items)
	})
}

func Fuzz_Nosy_freezerTable_truncateTail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, name string, disableSnappy bool, readonly bool, items uint64) {
		t1, err := newFreezerTable(path, name, disableSnappy, readonly)
		if err != nil {
			return
		}
		t1.truncateTail(items)
	})
}

// skipping Fuzz_Nosy_freezerTableBatch_Append__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_freezerTableBatch_AppendRaw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batch *freezerTableBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var item uint64
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if batch == nil {
			return
		}

		batch.AppendRaw(item, blob)
	})
}

func Fuzz_Nosy_freezerTableBatch_appendItem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batch *freezerTableBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if batch == nil {
			return
		}

		batch.appendItem(d2)
	})
}

func Fuzz_Nosy_freezerTableBatch_commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batch *freezerTableBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if batch == nil {
			return
		}

		batch.commit()
	})
}

func Fuzz_Nosy_freezerTableBatch_maybeCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batch *freezerTableBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if batch == nil {
			return
		}

		batch.maybeCommit()
	})
}

func Fuzz_Nosy_freezerTableBatch_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var batch *freezerTableBatch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if batch == nil {
			return
		}

		batch.reset()
	})
}

func Fuzz_Nosy_freezerdb_AncientDatadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frdb *freezerdb
		fill_err = tp.Fill(&frdb)
		if fill_err != nil {
			return
		}
		if frdb == nil {
			return
		}

		frdb.AncientDatadir()
	})
}

func Fuzz_Nosy_freezerdb_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frdb *freezerdb
		fill_err = tp.Fill(&frdb)
		if fill_err != nil {
			return
		}
		if frdb == nil {
			return
		}

		frdb.Close()
	})
}

func Fuzz_Nosy_freezerdb_Freeze__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var frdb *freezerdb
		fill_err = tp.Fill(&frdb)
		if fill_err != nil {
			return
		}
		var threshold uint64
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		if frdb == nil {
			return
		}

		frdb.Freeze(threshold)
	})
}

func Fuzz_Nosy_hasher_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		h := newHasher()
		h.hash(d1)
	})
}

func Fuzz_Nosy_indexEntry_append__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexEntry
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.append(b)
	})
}

func Fuzz_Nosy_indexEntry_bounds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexEntry
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var end *indexEntry
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if i == nil || end == nil {
			return
		}

		i.bounds(end)
	})
}

func Fuzz_Nosy_indexEntry_unmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *indexEntry
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.unmarshalBinary(b)
	})
}

func Fuzz_Nosy_nofreezedb_Ancient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Ancient(kind, number)
	})
}

func Fuzz_Nosy_nofreezedb_AncientDatadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.AncientDatadir()
	})
}

func Fuzz_Nosy_nofreezedb_AncientRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var max uint64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var maxByteSize uint64
		fill_err = tp.Fill(&maxByteSize)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.AncientRange(kind, start, max, maxByteSize)
	})
}

func Fuzz_Nosy_nofreezedb_AncientSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.AncientSize(kind)
	})
}

func Fuzz_Nosy_nofreezedb_Ancients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Ancients()
	})
}

func Fuzz_Nosy_nofreezedb_HasAncient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.HasAncient(kind, number)
	})
}

// skipping Fuzz_Nosy_nofreezedb_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_nofreezedb_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

// skipping Fuzz_Nosy_nofreezedb_ReadAncients__ because parameters include func, chan, or unsupported interface: func(reader github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_nofreezedb_Sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Sync()
	})
}

func Fuzz_Nosy_nofreezedb_Tail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Tail()
	})
}

func Fuzz_Nosy_nofreezedb_TruncateHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.TruncateHead(items)
	})
}

func Fuzz_Nosy_nofreezedb_TruncateTail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *nofreezedb
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.TruncateTail(items)
	})
}

func Fuzz_Nosy_receiptLogs_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *receiptLogs
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if r == nil || s == nil {
			return
		}

		r.DecodeRLP(s)
	})
}

func Fuzz_Nosy_snappyBuffer_compress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *snappyBuffer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.compress(d2)
	})
}

func Fuzz_Nosy_stat_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stat
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var size common.StorageSize
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Add(size)
	})
}

func Fuzz_Nosy_stat_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stat
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Count()
	})
}

func Fuzz_Nosy_stat_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stat
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Size()
	})
}

func Fuzz_Nosy_table_Ancient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ancient(kind, number)
	})
}

func Fuzz_Nosy_table_AncientDatadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.AncientDatadir()
	})
}

func Fuzz_Nosy_table_AncientRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var maxBytes uint64
		fill_err = tp.Fill(&maxBytes)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.AncientRange(kind, start, count, maxBytes)
	})
}

func Fuzz_Nosy_table_AncientSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.AncientSize(kind)
	})
}

func Fuzz_Nosy_table_Ancients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Ancients()
	})
}

func Fuzz_Nosy_table_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_table_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var limit []byte
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Compact(start, limit)
	})
}

func Fuzz_Nosy_table_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Delete(key)
	})
}

func Fuzz_Nosy_table_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Get(key)
	})
}

func Fuzz_Nosy_table_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Has(key)
	})
}

func Fuzz_Nosy_table_HasAncient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.HasAncient(kind, number)
	})
}

// skipping Fuzz_Nosy_table_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_table_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

func Fuzz_Nosy_table_NewBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.NewBatch()
	})
}

func Fuzz_Nosy_table_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_table_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_table_NewSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.NewSnapshot()
	})
}

func Fuzz_Nosy_table_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Put(key, value)
	})
}

// skipping Fuzz_Nosy_table_ReadAncients__ because parameters include func, chan, or unsupported interface: func(reader github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_table_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Stat(property)
	})
}

func Fuzz_Nosy_table_Sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Sync()
	})
}

func Fuzz_Nosy_table_Tail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Tail()
	})
}

func Fuzz_Nosy_table_TruncateHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.TruncateHead(items)
	})
}

func Fuzz_Nosy_table_TruncateTail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *table
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var items uint64
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.TruncateTail(items)
	})
}

func Fuzz_Nosy_tableBatch_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *tableBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Delete(key)
	})
}

func Fuzz_Nosy_tableBatch_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *tableBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Put(key, value)
	})
}

// skipping Fuzz_Nosy_tableBatch_Replay__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_tableBatch_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *tableBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Reset()
	})
}

func Fuzz_Nosy_tableBatch_ValueSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *tableBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ValueSize()
	})
}

func Fuzz_Nosy_tableBatch_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *tableBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Write()
	})
}

func Fuzz_Nosy_tableIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *tableIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Error()
	})
}

func Fuzz_Nosy_tableIterator_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *tableIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Key()
	})
}

func Fuzz_Nosy_tableIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *tableIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Next()
	})
}

func Fuzz_Nosy_tableIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *tableIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Release()
	})
}

func Fuzz_Nosy_tableIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *tableIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Value()
	})
}

func Fuzz_Nosy_tableReplayer_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *tableReplayer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Delete(key)
	})
}

func Fuzz_Nosy_tableReplayer_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *tableReplayer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Put(key, value)
	})
}

func Fuzz_Nosy_writeBuffer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wb *writeBuffer
		fill_err = tp.Fill(&wb)
		if fill_err != nil {
			return
		}
		if wb == nil {
			return
		}

		wb.Reset()
	})
}

func Fuzz_Nosy_writeBuffer_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wb *writeBuffer
		fill_err = tp.Fill(&wb)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if wb == nil {
			return
		}

		wb.Write(d2)
	})
}

func Fuzz_Nosy_counter_Percentage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c counter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var current uint64
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}

		c.Percentage(current)
	})
}

func Fuzz_Nosy_counter_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c counter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.String()
	})
}

// skipping Fuzz_Nosy_DeleteAccountSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteAccountTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteBadBlocks__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteBlockWithoutNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteBloombits__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_DeleteBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteCanonicalHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteHeaderNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteLegacyTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSkeletonHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSkeletonSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSnapshotDisabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSnapshotGenerator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSnapshotJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSnapshotRecoveryNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteSnapshotRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteStateID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteStorageSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteStorageTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteTrieJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteTxLookupEntries__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_DeleteTxLookupEntry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_ExistsAccountTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ExistsStorageTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasAccountTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_HasCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasCodeWithPrefix__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_HasLegacyTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_HasStorageTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_HasTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_IndexTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_InitDatabaseFromFreezer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_IsAccountTrieNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		IsAccountTrieNode(key)
	})
}

func Fuzz_Nosy_IsCodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		IsCodeKey(key)
	})
}

func Fuzz_Nosy_IsLegacyTrieNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, val []byte) {
		IsLegacyTrieNode(key, val)
	})
}

func Fuzz_Nosy_IsStorageTrieNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		IsStorageTrieNode(key)
	})
}

// skipping Fuzz_Nosy_ParseStateScheme__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_PopUncleanShutdownMarker__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

func Fuzz_Nosy_PreexistingDatabase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		PreexistingDatabase(path)
	})
}

// skipping Fuzz_Nosy_PushUncleanShutdownMarker__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_ReadAccountSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadAccountTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadAllBadBlocks__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadAllCanonicalHashes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iteratee

// skipping Fuzz_Nosy_ReadAllHashes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iteratee

// skipping Fuzz_Nosy_ReadAllHashesInRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iteratee

// skipping Fuzz_Nosy_ReadBloomBits__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadChainMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_ReadCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadCodeWithPrefix__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadDatabaseVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadFastTxLookupLimit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadGenesisStateSpec__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadHeaderNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadHeaderRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadLastPivotNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadLegacyTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadPersistentStateID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadSkeletonSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapSyncStatusFlag__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapshotDisabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapshotGenerator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapshotJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapshotRecoveryNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadSnapshotSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadStateAccountHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateAccountIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateHistoryMeta__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateHistoryMetaList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadStateScheme__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadStateStorageHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStateStorageIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_ReadStorageSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadStorageTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

// skipping Fuzz_Nosy_ReadTransitionStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadTrieJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadTxIndexTail__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_ReadTxLookupEntry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Reader

func Fuzz_Nosy_ResolveAccountTrieNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		ResolveAccountTrieNodeKey(key)
	})
}

func Fuzz_Nosy_ResolveStorageTrieNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		ResolveStorageTrieNode(key)
	})
}

// skipping Fuzz_Nosy_UnindexTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_UpdateUncleanShutdownMarker__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_WriteAccountSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteAccountTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteAncientBlocks__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_WriteBadBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_WriteBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteBloomBits__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteBodyRLP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteCanonicalHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteChainConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteDatabaseVersion__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteFastTxLookupLimit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteFinalizedBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteGenesisStateSpec__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteHeadBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteHeadFastBlockHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteHeadHeaderHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteHeaderNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteLastPivotNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteLegacyTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WritePersistentStateID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WritePreimages__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSkeletonHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSkeletonSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapSyncStatusFlag__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotDisabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotGenerator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotRecoveryNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteSnapshotSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteStateHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_WriteStateID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteStorageSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteStorageTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTransitionStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTrieJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTrieNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTxIndexTail__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTxLookupEntries__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_WriteTxLookupEntriesByBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_accountSnapshotKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		accountSnapshotKey(hash)
	})
}

func Fuzz_Nosy_accountTrieNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte) {
		accountTrieNodeKey(path)
	})
}

func Fuzz_Nosy_blockBodyKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		blockBodyKey(number, hash)
	})
}

func Fuzz_Nosy_blockReceiptsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		blockReceiptsKey(number, hash)
	})
}

func Fuzz_Nosy_bloomBitsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bit uint
		fill_err = tp.Fill(&bit)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		bloomBitsKey(bit, section, hash)
	})
}

func Fuzz_Nosy_codeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		codeKey(hash)
	})
}

func Fuzz_Nosy_configKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		configKey(hash)
	})
}

// skipping Fuzz_Nosy_deleteHeaderWithoutNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_encodeBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		encodeBlockNumber(number)
	})
}

func Fuzz_Nosy_genesisStateSpecKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		genesisStateSpecKey(hash)
	})
}

func Fuzz_Nosy_grow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte, n int) {
		grow(buf, n)
	})
}

func Fuzz_Nosy_headerHashKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		headerHashKey(number)
	})
}

func Fuzz_Nosy_headerKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		headerKey(number, hash)
	})
}

func Fuzz_Nosy_headerKeyPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		headerKeyPrefix(number)
	})
}

func Fuzz_Nosy_headerNumberKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		headerNumberKey(hash)
	})
}

func Fuzz_Nosy_headerTDKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		headerTDKey(number, hash)
	})
}

// skipping Fuzz_Nosy_indexTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_indexTransactionsForTesting__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_inspectFreezers__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_isCanon__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_iterateTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_preimageKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		preimageKey(hash)
	})
}

// skipping Fuzz_Nosy_printChainMetadata__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

func Fuzz_Nosy_resolveChainFreezerDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ancient string) {
		resolveChainFreezerDir(ancient)
	})
}

func Fuzz_Nosy_skeletonHeaderKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, number uint64) {
		skeletonHeaderKey(number)
	})
}

func Fuzz_Nosy_stateIDKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}

		stateIDKey(root)
	})
}

func Fuzz_Nosy_storageSnapshotKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		var storageHash common.Hash
		fill_err = tp.Fill(&storageHash)
		if fill_err != nil {
			return
		}

		storageSnapshotKey(accountHash, storageHash)
	})
}

func Fuzz_Nosy_storageSnapshotsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}

		storageSnapshotsKey(accountHash)
	})
}

func Fuzz_Nosy_storageTrieNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}

		storageTrieNodeKey(accountHash, path)
	})
}

func Fuzz_Nosy_tmpName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		tmpName(path)
	})
}

func Fuzz_Nosy_txLookupKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		txLookupKey(hash)
	})
}

// skipping Fuzz_Nosy_unindexTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_unindexTransactionsForTesting__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_writeTxLookupEntry__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter
