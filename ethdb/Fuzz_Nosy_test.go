package ethdb

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

// skipping Fuzz_Nosy_AncientReader_ReadAncients__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReader

// skipping Fuzz_Nosy_AncientReaderOp_Ancient__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientReaderOp_AncientRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientReaderOp_AncientSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientReaderOp_Ancients__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientReaderOp_HasAncient__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientReaderOp_Tail__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientReaderOp

// skipping Fuzz_Nosy_AncientStater_AncientDatadir__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientStater

// skipping Fuzz_Nosy_AncientWriteOp_Append__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriteOp

// skipping Fuzz_Nosy_AncientWriteOp_AppendRaw__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriteOp

// skipping Fuzz_Nosy_AncientWriter_MigrateTable__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_AncientWriter_ModifyAncients__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_AncientWriter_Sync__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_AncientWriter_TruncateHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_AncientWriter_TruncateTail__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.AncientWriter

// skipping Fuzz_Nosy_Batch_Replay__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

// skipping Fuzz_Nosy_Batch_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

// skipping Fuzz_Nosy_Batch_ValueSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

// skipping Fuzz_Nosy_Batch_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

// skipping Fuzz_Nosy_Batcher_NewBatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batcher

// skipping Fuzz_Nosy_Batcher_NewBatchWithSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batcher

// skipping Fuzz_Nosy_Compacter_Compact__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Compacter

func Fuzz_Nosy_HookedBatch_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b HookedBatch
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		b.Delete(key)
	})
}

func Fuzz_Nosy_HookedBatch_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b HookedBatch
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

		b.Put(key, value)
	})
}

// skipping Fuzz_Nosy_Iteratee_NewIterator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iteratee

// skipping Fuzz_Nosy_Iterator_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

// skipping Fuzz_Nosy_Iterator_Key__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

// skipping Fuzz_Nosy_Iterator_Next__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

// skipping Fuzz_Nosy_Iterator_Release__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

// skipping Fuzz_Nosy_Iterator_Value__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Iterator

// skipping Fuzz_Nosy_KeyValueReader_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_KeyValueReader_Has__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_KeyValueStater_Stat__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStater

// skipping Fuzz_Nosy_KeyValueWriter_Delete__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_KeyValueWriter_Put__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_Snapshot_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Snapshot

// skipping Fuzz_Nosy_Snapshot_Has__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Snapshot

// skipping Fuzz_Nosy_Snapshot_Release__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Snapshot

// skipping Fuzz_Nosy_Snapshotter_NewSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Snapshotter
