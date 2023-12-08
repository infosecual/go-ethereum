package memorydb

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

func Fuzz_Nosy_Database_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, start []byte, limit []byte) {
		db := New()
		db.Compact(start, limit)
	})
}

func Fuzz_Nosy_Database_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := New()
		db.Delete(key)
	})
}

func Fuzz_Nosy_Database_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := New()
		db.Get(key)
	})
}

func Fuzz_Nosy_Database_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := New()
		db.Has(key)
	})
}

func Fuzz_Nosy_Database_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
		db := New()
		db.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_Database_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix []byte, start []byte) {
		db := New()
		db.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_Database_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, value []byte) {
		db := New()
		db.Put(key, value)
	})
}

func Fuzz_Nosy_Database_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, property string) {
		db := New()
		db.Stat(property)
	})
}

func Fuzz_Nosy_batch_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batch
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

func Fuzz_Nosy_batch_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batch
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

// skipping Fuzz_Nosy_batch_Replay__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_batch_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batch
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

func Fuzz_Nosy_batch_ValueSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batch
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

func Fuzz_Nosy_batch_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batch
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

func Fuzz_Nosy_iterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *iterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Error()
	})
}

func Fuzz_Nosy_iterator_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *iterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Key()
	})
}

func Fuzz_Nosy_iterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *iterator
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

func Fuzz_Nosy_iterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *iterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Release()
	})
}

func Fuzz_Nosy_iterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *iterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Value()
	})
}

func Fuzz_Nosy_snapshot_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		snap := newSnapshot(db)
		snap.Get(key)
	})
}

func Fuzz_Nosy_snapshot_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		snap := newSnapshot(db)
		snap.Has(key)
	})
}

func Fuzz_Nosy_snapshot_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		snap := newSnapshot(db)
		snap.Release()
	})
}
