package leveldb

import (
	"testing"
	"time"

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

func Fuzz_Nosy_Database_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Close()
	})
}

func Fuzz_Nosy_Database_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, start []byte, limit []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Compact(start, limit)
	})
}

func Fuzz_Nosy_Database_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, key []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Delete(key)
	})
}

func Fuzz_Nosy_Database_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, key []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Get(key)
	})
}

func Fuzz_Nosy_Database_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, key []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Has(key)
	})
}

func Fuzz_Nosy_Database_NewBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.NewBatch()
	})
}

func Fuzz_Nosy_Database_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, size int) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_Database_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, prefix []byte, start []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_Database_NewSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.NewSnapshot()
	})
}

func Fuzz_Nosy_Database_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Path()
	})
}

func Fuzz_Nosy_Database_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, key []byte, value []byte) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Put(key, value)
	})
}

func Fuzz_Nosy_Database_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, property string) {
		db, err := New(file, cache, handles, namespace, readonly)
		if err != nil {
			return
		}
		db.Stat(property)
	})
}

func Fuzz_Nosy_Database_meter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var cache int
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var handles int
		fill_err = tp.Fill(&handles)
		if fill_err != nil {
			return
		}
		var n4 string
		fill_err = tp.Fill(&n4)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		var refresh time.Duration
		fill_err = tp.Fill(&refresh)
		if fill_err != nil {
			return
		}
		var n7 string
		fill_err = tp.Fill(&n7)
		if fill_err != nil {
			return
		}

		db, err := New(file, cache, handles, n4, readonly)
		if err != nil {
			return
		}
		db.meter(refresh, n7)
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

func Fuzz_Nosy_replayer_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *replayer
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

func Fuzz_Nosy_replayer_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *replayer
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

func Fuzz_Nosy_snapshot_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var snap *snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if snap == nil {
			return
		}

		snap.Get(key)
	})
}

func Fuzz_Nosy_snapshot_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var snap *snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if snap == nil {
			return
		}

		snap.Has(key)
	})
}

func Fuzz_Nosy_snapshot_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var snap *snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		if snap == nil {
			return
		}

		snap.Release()
	})
}
