package remotedb

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

func Fuzz_Nosy_Database_Ancient__(f *testing.F) {
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

func Fuzz_Nosy_Database_AncientDatadir__(f *testing.F) {
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

		db.AncientDatadir()
	})
}

func Fuzz_Nosy_Database_AncientRange__(f *testing.F) {
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
		if db == nil {
			return
		}

		db.AncientRange(kind, start, count, maxBytes)
	})
}

func Fuzz_Nosy_Database_AncientSize__(f *testing.F) {
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

func Fuzz_Nosy_Database_Ancients__(f *testing.F) {
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

		db.Ancients()
	})
}

func Fuzz_Nosy_Database_Close__(f *testing.F) {
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

		db.Close()
	})
}

func Fuzz_Nosy_Database_Compact__(f *testing.F) {
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
		if db == nil {
			return
		}

		db.Compact(start, limit)
	})
}

func Fuzz_Nosy_Database_Delete__(f *testing.F) {
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

		db.Delete(key)
	})
}

func Fuzz_Nosy_Database_Get__(f *testing.F) {
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

		db.Get(key)
	})
}

func Fuzz_Nosy_Database_Has__(f *testing.F) {
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

		db.Has(key)
	})
}

func Fuzz_Nosy_Database_HasAncient__(f *testing.F) {
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

// skipping Fuzz_Nosy_Database_MigrateTable__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_Database_ModifyAncients__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/ethdb.AncientWriteOp) error

func Fuzz_Nosy_Database_NewBatch__(f *testing.F) {
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

		db.NewBatch()
	})
}

func Fuzz_Nosy_Database_NewBatchWithSize__(f *testing.F) {
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
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.NewBatchWithSize(size)
	})
}

func Fuzz_Nosy_Database_NewIterator__(f *testing.F) {
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
		if db == nil {
			return
		}

		db.NewIterator(prefix, start)
	})
}

func Fuzz_Nosy_Database_NewSnapshot__(f *testing.F) {
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

		db.NewSnapshot()
	})
}

func Fuzz_Nosy_Database_Put__(f *testing.F) {
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
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Put(key, value)
	})
}

// skipping Fuzz_Nosy_Database_ReadAncients__ because parameters include func, chan, or unsupported interface: func(op github.com/ethereum/go-ethereum/ethdb.AncientReaderOp) error

func Fuzz_Nosy_Database_Stat__(f *testing.F) {
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
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Stat(property)
	})
}

func Fuzz_Nosy_Database_Sync__(f *testing.F) {
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

		db.Sync()
	})
}

func Fuzz_Nosy_Database_Tail__(f *testing.F) {
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

		db.Tail()
	})
}

func Fuzz_Nosy_Database_TruncateHead__(f *testing.F) {
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
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.TruncateHead(n)
	})
}

func Fuzz_Nosy_Database_TruncateTail__(f *testing.F) {
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
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.TruncateTail(n)
	})
}
