package pebble

import (
	"testing"
	"time"

	"github.com/cockroachdb/pebble"
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
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Close()
	})
}

func Fuzz_Nosy_Database_Compact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, start []byte, limit []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Compact(start, limit)
	})
}

func Fuzz_Nosy_Database_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, key []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Delete(key)
	})
}

func Fuzz_Nosy_Database_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, key []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Get(key)
	})
}

func Fuzz_Nosy_Database_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, key []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Has(key)
	})
}

func Fuzz_Nosy_Database_NewBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.NewBatch()
	})
}

func Fuzz_Nosy_Database_NewBatchWithSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, size int) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_Database_NewIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, prefix []byte, start []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_Database_NewSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.NewSnapshot()
	})
}

func Fuzz_Nosy_Database_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Path()
	})
}

func Fuzz_Nosy_Database_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, key []byte, value []byte) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Put(key, value)
	})
}

func Fuzz_Nosy_Database_Stat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool, property string) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.Stat(property)
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
		var ephemeral bool
		fill_err = tp.Fill(&ephemeral)
		if fill_err != nil {
			return
		}
		var refresh time.Duration
		fill_err = tp.Fill(&refresh)
		if fill_err != nil {
			return
		}
		var n8 string
		fill_err = tp.Fill(&n8)
		if fill_err != nil {
			return
		}

		d, err := New(file, cache, handles, n4, readonly, ephemeral)
		if err != nil {
			return
		}
		d.meter(refresh, n8)
	})
}

func Fuzz_Nosy_Database_onCompactionBegin__(f *testing.F) {
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
		var ephemeral bool
		fill_err = tp.Fill(&ephemeral)
		if fill_err != nil {
			return
		}
		var info pebble.CompactionInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.onCompactionBegin(info)
	})
}

func Fuzz_Nosy_Database_onCompactionEnd__(f *testing.F) {
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
		var ephemeral bool
		fill_err = tp.Fill(&ephemeral)
		if fill_err != nil {
			return
		}
		var info pebble.CompactionInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.onCompactionEnd(info)
	})
}

func Fuzz_Nosy_Database_onWriteStallBegin__(f *testing.F) {
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
		var ephemeral bool
		fill_err = tp.Fill(&ephemeral)
		if fill_err != nil {
			return
		}
		var b pebble.WriteStallBeginInfo
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.onWriteStallBegin(b)
	})
}

func Fuzz_Nosy_Database_onWriteStallEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, cache int, handles int, namespace string, readonly bool, ephemeral bool) {
		d, err := New(file, cache, handles, namespace, readonly, ephemeral)
		if err != nil {
			return
		}
		d.onWriteStallEnd()
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

func Fuzz_Nosy_pebbleIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *pebbleIterator
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

func Fuzz_Nosy_pebbleIterator_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *pebbleIterator
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

func Fuzz_Nosy_pebbleIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *pebbleIterator
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

func Fuzz_Nosy_pebbleIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *pebbleIterator
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

func Fuzz_Nosy_pebbleIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *pebbleIterator
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

// skipping Fuzz_Nosy_panicLogger_Errorf__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_panicLogger_Fatalf__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_panicLogger_Infof__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_upperBound__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix []byte) {
		upperBound(prefix)
	})
}
