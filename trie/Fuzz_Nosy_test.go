package trie

import (
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/triestate"
	"github.com/ethereum/go-ethereum/trie/utils"
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

func Fuzz_Nosy_Database_Cap__(f *testing.F) {
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
		var limit common.StorageSize
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Cap(limit)
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

func Fuzz_Nosy_Database_Commit__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var report bool
		fill_err = tp.Fill(&report)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Commit(root, report)
	})
}

func Fuzz_Nosy_Database_Dereference__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Dereference(root)
	})
}

func Fuzz_Nosy_Database_Disable__(f *testing.F) {
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

		db.Disable()
	})
}

func Fuzz_Nosy_Database_Enable__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Enable(root)
	})
}

func Fuzz_Nosy_Database_Initialized__(f *testing.F) {
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
		var genesisRoot common.Hash
		fill_err = tp.Fill(&genesisRoot)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Initialized(genesisRoot)
	})
}

func Fuzz_Nosy_Database_IsVerkle__(f *testing.F) {
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

		db.IsVerkle()
	})
}

func Fuzz_Nosy_Database_Journal__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Journal(root)
	})
}

func Fuzz_Nosy_Database_Preimage__(f *testing.F) {
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
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Preimage(hash)
	})
}

func Fuzz_Nosy_Database_Reader__(f *testing.F) {
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
		var blockRoot common.Hash
		fill_err = tp.Fill(&blockRoot)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Reader(blockRoot)
	})
}

func Fuzz_Nosy_Database_Recover__(f *testing.F) {
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
		var t2 common.Hash
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Recover(t2)
	})
}

func Fuzz_Nosy_Database_Recoverable__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Recoverable(root)
	})
}

func Fuzz_Nosy_Database_Reference__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Reference(root, parent)
	})
}

func Fuzz_Nosy_Database_Scheme__(f *testing.F) {
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

		db.Scheme()
	})
}

func Fuzz_Nosy_Database_SetBufferSize__(f *testing.F) {
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

		db.SetBufferSize(size)
	})
}

func Fuzz_Nosy_Database_Size__(f *testing.F) {
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

		db.Size()
	})
}

func Fuzz_Nosy_Database_Update__(f *testing.F) {
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var nodes *trienode.MergedNodeSet
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var states *triestate.Set
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		if db == nil || nodes == nil || states == nil {
			return
		}

		db.Update(root, parent, block, nodes, states)
	})
}

func Fuzz_Nosy_Database_WritePreimages__(f *testing.F) {
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

		db.WritePreimages()
	})
}

func Fuzz_Nosy_Iterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *Iterator
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

func Fuzz_Nosy_Iterator_Prove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *Iterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Prove()
	})
}

func Fuzz_Nosy_MissingNodeError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *MissingNodeError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_MissingNodeError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *MissingNodeError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Unwrap()
	})
}

func Fuzz_Nosy_StackTrie_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
		if fill_err != nil {
			return
		}
		if options == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.Commit()
	})
}

func Fuzz_Nosy_StackTrie_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
		if fill_err != nil {
			return
		}
		if options == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.Hash()
	})
}

func Fuzz_Nosy_StackTrie_MustUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
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
		if options == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.MustUpdate(key, value)
	})
}

func Fuzz_Nosy_StackTrie_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
		if fill_err != nil {
			return
		}
		if options == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.Reset()
	})
}

func Fuzz_Nosy_StackTrie_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
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
		if options == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.Update(key, value)
	})
}

func Fuzz_Nosy_StackTrie_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
		if fill_err != nil {
			return
		}
		var st *stNode
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if options == nil || st == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.hash(st, path)
	})
}

func Fuzz_Nosy_StackTrie_insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var options *StackTrieOptions
		fill_err = tp.Fill(&options)
		if fill_err != nil {
			return
		}
		var st *stNode
		fill_err = tp.Fill(&st)
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
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if options == nil || st == nil {
			return
		}

		t1 := NewStackTrie(options)
		t1.insert(st, key, value, path)
	})
}

// skipping Fuzz_Nosy_StackTrieOptions_WithCleaner__ because parameters include func, chan, or unsupported interface: func(path []byte)

// skipping Fuzz_Nosy_StackTrieOptions_WithSkipBoundary__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Gauge

// skipping Fuzz_Nosy_StackTrieOptions_WithWriter__ because parameters include func, chan, or unsupported interface: func(path []byte, hash github.com/ethereum/go-ethereum/common.Hash, blob []byte)

func Fuzz_Nosy_StateTrie_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var collectLeaf bool
		fill_err = tp.Fill(&collectLeaf)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.Commit(collectLeaf)
	})
}

func Fuzz_Nosy_StateTrie_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.Copy()
	})
}

func Fuzz_Nosy_StateTrie_DeleteAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.DeleteAccount(address)
	})
}

func Fuzz_Nosy_StateTrie_DeleteStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var _x3 common.Address
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.DeleteStorage(_x3, key)
	})
}

func Fuzz_Nosy_StateTrie_GetAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.GetAccount(address)
	})
}

func Fuzz_Nosy_StateTrie_GetAccountByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var addrHash common.Hash
		fill_err = tp.Fill(&addrHash)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.GetAccountByHash(addrHash)
	})
}

func Fuzz_Nosy_StateTrie_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var shaKey []byte
		fill_err = tp.Fill(&shaKey)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.GetKey(shaKey)
	})
}

func Fuzz_Nosy_StateTrie_GetNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.GetNode(path)
	})
}

func Fuzz_Nosy_StateTrie_GetStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var _x3 common.Address
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.GetStorage(_x3, key)
	})
}

func Fuzz_Nosy_StateTrie_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.Hash()
	})
}

func Fuzz_Nosy_StateTrie_MustDelete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.MustDelete(key)
	})
}

func Fuzz_Nosy_StateTrie_MustGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.MustGet(key)
	})
}

func Fuzz_Nosy_StateTrie_MustNodeIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.MustNodeIterator(start)
	})
}

func Fuzz_Nosy_StateTrie_MustUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.MustUpdate(key, value)
	})
}

func Fuzz_Nosy_StateTrie_NodeIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.NodeIterator(start)
	})
}

// skipping Fuzz_Nosy_StateTrie_Prove__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_StateTrie_UpdateAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acc *types.StateAccount
		fill_err = tp.Fill(&acc)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil || acc == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.UpdateAccount(address, acc)
	})
}

func Fuzz_Nosy_StateTrie_UpdateContractCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var _x3 common.Address
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 common.Hash
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var _x5 []byte
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.UpdateContractCode(_x3, _x4, _x5)
	})
}

func Fuzz_Nosy_StateTrie_UpdateStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var _x3 common.Address
		fill_err = tp.Fill(&_x3)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.UpdateStorage(_x3, key, value)
	})
}

func Fuzz_Nosy_StateTrie_getSecKeyCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.getSecKeyCache()
	})
}

func Fuzz_Nosy_StateTrie_hashKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id *ID
		fill_err = tp.Fill(&id)
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
		if id == nil || db == nil {
			return
		}

		t1, err := NewStateTrie(id, db)
		if err != nil {
			return
		}
		t1.hashKey(key)
	})
}

func Fuzz_Nosy_Sync_AddCodeEntry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var parentPath []byte
		fill_err = tp.Fill(&parentPath)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddCodeEntry(hash, path, parent, parentPath)
	})
}

// skipping Fuzz_Nosy_Sync_AddSubTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.LeafCallback

// skipping Fuzz_Nosy_Sync_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

func Fuzz_Nosy_Sync_MemSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MemSize()
	})
}

func Fuzz_Nosy_Sync_Missing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Missing(max)
	})
}

func Fuzz_Nosy_Sync_Pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Pending()
	})
}

func Fuzz_Nosy_Sync_ProcessCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var result CodeSyncResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ProcessCode(result)
	})
}

func Fuzz_Nosy_Sync_ProcessNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var result NodeSyncResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ProcessNode(result)
	})
}

// skipping Fuzz_Nosy_Sync_children__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_Sync_commitCodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *codeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.commitCodeRequest(req)
	})
}

func Fuzz_Nosy_Sync_commitNodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *nodeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.commitNodeRequest(req)
	})
}

func Fuzz_Nosy_Sync_scheduleCodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *codeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleCodeRequest(req)
	})
}

func Fuzz_Nosy_Sync_scheduleNodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Sync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *nodeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleNodeRequest(req)
	})
}

func Fuzz_Nosy_Trie_Commit__(f *testing.F) {
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
		var collectLeaf bool
		fill_err = tp.Fill(&collectLeaf)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.Commit(collectLeaf)
	})
}

func Fuzz_Nosy_Trie_Copy__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Copy()
	})
}

func Fuzz_Nosy_Trie_Delete__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Delete(key)
	})
}

func Fuzz_Nosy_Trie_Get__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Get(key)
	})
}

func Fuzz_Nosy_Trie_GetNode__(f *testing.F) {
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
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.GetNode(path)
	})
}

func Fuzz_Nosy_Trie_Hash__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Hash()
	})
}

func Fuzz_Nosy_Trie_MustDelete__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.MustDelete(key)
	})
}

func Fuzz_Nosy_Trie_MustGet__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.MustGet(key)
	})
}

func Fuzz_Nosy_Trie_MustGetNode__(f *testing.F) {
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
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.MustGetNode(path)
	})
}

func Fuzz_Nosy_Trie_MustNodeIterator__(f *testing.F) {
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
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.MustNodeIterator(start)
	})
}

func Fuzz_Nosy_Trie_MustUpdate__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.MustUpdate(key, value)
	})
}

func Fuzz_Nosy_Trie_NodeIterator__(f *testing.F) {
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
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.NodeIterator(start)
	})
}

// skipping Fuzz_Nosy_Trie_Prove__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_Trie_Reset__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Reset()
	})
}

func Fuzz_Nosy_Trie_Update__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.Update(key, value)
	})
}

// skipping Fuzz_Nosy_Trie_delete__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

// skipping Fuzz_Nosy_Trie_get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

// skipping Fuzz_Nosy_Trie_getNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_Trie_hashRoot__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.hashRoot()
	})
}

// skipping Fuzz_Nosy_Trie_insert__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_Trie_newFlag__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.newFlag()
	})
}

// skipping Fuzz_Nosy_Trie_resolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_Trie_resolveAndTrack__(f *testing.F) {
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
		var n hashNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		t1 := NewEmpty(db)
		t1.resolveAndTrack(n, prefix)
	})
}

func Fuzz_Nosy_Trie_update__(f *testing.F) {
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

		t1 := NewEmpty(db)
		t1.update(key, value)
	})
}

func Fuzz_Nosy_VerkleTrie_Commit__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var _x4 bool
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.Commit(_x4)
	})
}

func Fuzz_Nosy_VerkleTrie_Copy__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.Copy()
	})
}

func Fuzz_Nosy_VerkleTrie_DeleteAccount__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.DeleteAccount(addr)
	})
}

func Fuzz_Nosy_VerkleTrie_DeleteStorage__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.DeleteStorage(addr, key)
	})
}

func Fuzz_Nosy_VerkleTrie_GetAccount__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.GetAccount(addr)
	})
}

func Fuzz_Nosy_VerkleTrie_GetKey__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.GetKey(key)
	})
}

func Fuzz_Nosy_VerkleTrie_GetStorage__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.GetStorage(addr, key)
	})
}

func Fuzz_Nosy_VerkleTrie_Hash__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.Hash()
	})
}

func Fuzz_Nosy_VerkleTrie_IsVerkle__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.IsVerkle()
	})
}

func Fuzz_Nosy_VerkleTrie_NodeIterator__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var startKey []byte
		fill_err = tp.Fill(&startKey)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.NodeIterator(startKey)
	})
}

// skipping Fuzz_Nosy_VerkleTrie_Prove__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_VerkleTrie_ToDot__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.ToDot()
	})
}

func Fuzz_Nosy_VerkleTrie_UpdateAccount__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var acc *types.StateAccount
		fill_err = tp.Fill(&acc)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil || acc == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.UpdateAccount(addr, acc)
	})
}

func Fuzz_Nosy_VerkleTrie_UpdateContractCode__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		var code []byte
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.UpdateContractCode(addr, codeHash, code)
	})
}

func Fuzz_Nosy_VerkleTrie_UpdateStorage__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
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
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.UpdateStorage(address, key, value)
	})
}

func Fuzz_Nosy_VerkleTrie_nodeResolver__(f *testing.F) {
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
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cache *utils.PointCache
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if db == nil || cache == nil {
			return
		}

		t1, err := NewVerkleTrie(root, db, cache)
		if err != nil {
			return
		}
		t1.nodeResolver(path)
	})
}

// skipping Fuzz_Nosy_committer_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

// skipping Fuzz_Nosy_committer_commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_committer_commitChildren__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nodeset *trienode.NodeSet
		fill_err = tp.Fill(&nodeset)
		if fill_err != nil {
			return
		}
		var tracer *tracer
		fill_err = tp.Fill(&tracer)
		if fill_err != nil {
			return
		}
		var collectLeaf bool
		fill_err = tp.Fill(&collectLeaf)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var n *fullNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if nodeset == nil || tracer == nil || n == nil {
			return
		}

		c := newCommitter(nodeset, tracer, collectLeaf)
		c.commitChildren(path, n)
	})
}

// skipping Fuzz_Nosy_committer_store__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_decodeError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *decodeError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Error()
	})
}

// skipping Fuzz_Nosy_differenceIterator_AddResolver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeResolver

func Fuzz_Nosy_differenceIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
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

func Fuzz_Nosy_differenceIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Hash()
	})
}

func Fuzz_Nosy_differenceIterator_Leaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Leaf()
	})
}

func Fuzz_Nosy_differenceIterator_LeafBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafBlob()
	})
}

func Fuzz_Nosy_differenceIterator_LeafKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafKey()
	})
}

func Fuzz_Nosy_differenceIterator_LeafProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafProof()
	})
}

func Fuzz_Nosy_differenceIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var _x2 bool
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next(_x2)
	})
}

func Fuzz_Nosy_differenceIterator_NodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.NodeBlob()
	})
}

func Fuzz_Nosy_differenceIterator_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Parent()
	})
}

func Fuzz_Nosy_differenceIterator_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *differenceIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Path()
	})
}

func Fuzz_Nosy_fullNode_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var elems []byte
		fill_err = tp.Fill(&elems)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.EncodeRLP(w)
	})
}

func Fuzz_Nosy_fullNode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, elems []byte) {
		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.String()
	})
}

func Fuzz_Nosy_fullNode_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, elems []byte) {
		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.cache()
	})
}

func Fuzz_Nosy_fullNode_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, elems []byte) {
		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.copy()
	})
}

func Fuzz_Nosy_fullNode_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var elems []byte
		fill_err = tp.Fill(&elems)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.encode(w)
	})
}

func Fuzz_Nosy_fullNode_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, elems []byte, ind string) {
		n, err := decodeFull(hash, elems)
		if err != nil {
			return
		}
		n.fstring(ind)
	})
}

func Fuzz_Nosy_hasher_encodedBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, parallel bool) {
		h := newHasher(parallel)
		h.encodedBytes()
	})
}

func Fuzz_Nosy_hasher_fullnodeToHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parallel bool
		fill_err = tp.Fill(&parallel)
		if fill_err != nil {
			return
		}
		var n *fullNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		h := newHasher(parallel)
		h.fullnodeToHash(n, force)
	})
}

// skipping Fuzz_Nosy_hasher_hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_hasher_hashData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, parallel bool, d2 []byte) {
		h := newHasher(parallel)
		h.hashData(d2)
	})
}

func Fuzz_Nosy_hasher_hashFullNodeChildren__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parallel bool
		fill_err = tp.Fill(&parallel)
		if fill_err != nil {
			return
		}
		var n *fullNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		h := newHasher(parallel)
		h.hashFullNodeChildren(n)
	})
}

func Fuzz_Nosy_hasher_hashShortNodeChildren__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parallel bool
		fill_err = tp.Fill(&parallel)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		h := newHasher(parallel)
		h.hashShortNodeChildren(n)
	})
}

// skipping Fuzz_Nosy_hasher_proofHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_hasher_shortnodeToHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var parallel bool
		fill_err = tp.Fill(&parallel)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		h := newHasher(parallel)
		h.shortnodeToHash(n, force)
	})
}

// skipping Fuzz_Nosy_nodeIterator_AddResolver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeResolver

func Fuzz_Nosy_nodeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
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

func Fuzz_Nosy_nodeIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Hash()
	})
}

func Fuzz_Nosy_nodeIterator_Leaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Leaf()
	})
}

func Fuzz_Nosy_nodeIterator_LeafBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafBlob()
	})
}

func Fuzz_Nosy_nodeIterator_LeafKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafKey()
	})
}

func Fuzz_Nosy_nodeIterator_LeafProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafProof()
	})
}

func Fuzz_Nosy_nodeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var descend bool
		fill_err = tp.Fill(&descend)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next(descend)
	})
}

func Fuzz_Nosy_nodeIterator_NodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.NodeBlob()
	})
}

func Fuzz_Nosy_nodeIterator_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Parent()
	})
}

func Fuzz_Nosy_nodeIterator_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Path()
	})
}

func Fuzz_Nosy_nodeIterator_findChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var n *fullNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var ancestor common.Hash
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		if it == nil || n == nil {
			return
		}

		it.findChild(n, index, ancestor)
	})
}

func Fuzz_Nosy_nodeIterator_getFromPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.getFromPool()
	})
}

func Fuzz_Nosy_nodeIterator_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.init()
	})
}

func Fuzz_Nosy_nodeIterator_nextChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var parent *nodeIteratorState
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var ancestor common.Hash
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		if it == nil || parent == nil {
			return
		}

		it.nextChild(parent, ancestor)
	})
}

func Fuzz_Nosy_nodeIterator_nextChildAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var parent *nodeIteratorState
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var ancestor common.Hash
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if it == nil || parent == nil {
			return
		}

		it.nextChildAt(parent, ancestor, key)
	})
}

func Fuzz_Nosy_nodeIterator_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var descend bool
		fill_err = tp.Fill(&descend)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.peek(descend)
	})
}

func Fuzz_Nosy_nodeIterator_peekSeek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var seekKey []byte
		fill_err = tp.Fill(&seekKey)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.peekSeek(seekKey)
	})
}

func Fuzz_Nosy_nodeIterator_pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.pop()
	})
}

func Fuzz_Nosy_nodeIterator_push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var state *nodeIteratorState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var parentIndex *int
		fill_err = tp.Fill(&parentIndex)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if it == nil || state == nil || parentIndex == nil {
			return
		}

		it.push(state, parentIndex, path)
	})
}

func Fuzz_Nosy_nodeIterator_putInPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var item *nodeIteratorState
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}
		if it == nil || item == nil {
			return
		}

		it.putInPool(item)
	})
}

func Fuzz_Nosy_nodeIterator_resolveBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var hash hashNode
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.resolveBlob(hash, path)
	})
}

func Fuzz_Nosy_nodeIterator_resolveHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var hash hashNode
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.resolveHash(hash, path)
	})
}

func Fuzz_Nosy_nodeIterator_seek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.seek(prefix)
	})
}

// skipping Fuzz_Nosy_nodeIteratorHeap_Pop__ because parameters include func, chan, or unsupported interface: *github.com/ethereum/go-ethereum/trie.nodeIteratorHeap

// skipping Fuzz_Nosy_nodeIteratorHeap_Push__ because parameters include func, chan, or unsupported interface: *github.com/ethereum/go-ethereum/trie.nodeIteratorHeap

func Fuzz_Nosy_nodeIteratorState_resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *nodeIteratorState
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var it *nodeIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if st == nil || it == nil {
			return
		}

		st.resolve(it, path)
	})
}

func Fuzz_Nosy_preimageStore_commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store *preimageStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if store == nil {
			return
		}

		store.commit(force)
	})
}

func Fuzz_Nosy_preimageStore_insertPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store *preimageStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var preimages map[common.Hash][]byte
		fill_err = tp.Fill(&preimages)
		if fill_err != nil {
			return
		}
		if store == nil {
			return
		}

		store.insertPreimage(preimages)
	})
}

func Fuzz_Nosy_preimageStore_preimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store *preimageStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if store == nil {
			return
		}

		store.preimage(hash)
	})
}

func Fuzz_Nosy_preimageStore_size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store *preimageStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		if store == nil {
			return
		}

		store.size()
	})
}

func Fuzz_Nosy_shortNode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_shortNode_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.cache()
	})
}

func Fuzz_Nosy_shortNode_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.copy()
	})
}

func Fuzz_Nosy_shortNode_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.encode(w)
	})
}

func Fuzz_Nosy_shortNode_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *shortNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ind string
		fill_err = tp.Fill(&ind)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.fstring(ind)
	})
}

func Fuzz_Nosy_stNode_getDiffIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k1 []byte
		fill_err = tp.Fill(&k1)
		if fill_err != nil {
			return
		}
		var child *stNode
		fill_err = tp.Fill(&child)
		if fill_err != nil {
			return
		}
		var k3 []byte
		fill_err = tp.Fill(&k3)
		if fill_err != nil {
			return
		}
		if child == nil {
			return
		}

		n := newExt(k1, child)
		n.getDiffIndex(k3)
	})
}

func Fuzz_Nosy_stNode_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var child *stNode
		fill_err = tp.Fill(&child)
		if fill_err != nil {
			return
		}
		if child == nil {
			return
		}

		n := newExt(key, child)
		n.reset()
	})
}

func Fuzz_Nosy_syncMemBatch_hasCode__(f *testing.F) {
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

		batch := newSyncMemBatch()
		batch.hasCode(hash)
	})
}

func Fuzz_Nosy_syncMemBatch_hasNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte) {
		batch := newSyncMemBatch()
		batch.hasNode(path)
	})
}

func Fuzz_Nosy_tracer_onDelete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte) {
		t1 := newTracer()
		t1.onDelete(path)
	})
}

func Fuzz_Nosy_tracer_onInsert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte) {
		t1 := newTracer()
		t1.onInsert(path)
	})
}

func Fuzz_Nosy_tracer_onRead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte, val []byte) {
		t1 := newTracer()
		t1.onRead(path, val)
	})
}

func Fuzz_Nosy_trieLoader_OpenStorageTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *trieLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var stateRoot common.Hash
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var addrHash common.Hash
		fill_err = tp.Fill(&addrHash)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.OpenStorageTrie(stateRoot, addrHash, root)
	})
}

func Fuzz_Nosy_trieLoader_OpenTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *trieLoader
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.OpenTrie(root)
	})
}

func Fuzz_Nosy_trieReader_node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path []byte
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		r := newEmptyReader()
		r.node(path, hash)
	})
}

// skipping Fuzz_Nosy_unionIterator_AddResolver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeResolver

func Fuzz_Nosy_unionIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
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

func Fuzz_Nosy_unionIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Hash()
	})
}

func Fuzz_Nosy_unionIterator_Leaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Leaf()
	})
}

func Fuzz_Nosy_unionIterator_LeafBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafBlob()
	})
}

func Fuzz_Nosy_unionIterator_LeafKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafKey()
	})
}

func Fuzz_Nosy_unionIterator_LeafProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.LeafProof()
	})
}

func Fuzz_Nosy_unionIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var descend bool
		fill_err = tp.Fill(&descend)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next(descend)
	})
}

func Fuzz_Nosy_unionIterator_NodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.NodeBlob()
	})
}

func Fuzz_Nosy_unionIterator_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Parent()
	})
}

func Fuzz_Nosy_unionIterator_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *unionIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Path()
	})
}

// skipping Fuzz_Nosy_NodeIterator_AddResolver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeResolver

func Fuzz_Nosy_NodeIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.Error()
	})
}

func Fuzz_Nosy_NodeIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.Hash()
	})
}

func Fuzz_Nosy_NodeIterator_Leaf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.Leaf()
	})
}

func Fuzz_Nosy_NodeIterator_LeafBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.LeafBlob()
	})
}

func Fuzz_Nosy_NodeIterator_LeafKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.LeafKey()
	})
}

func Fuzz_Nosy_NodeIterator_LeafProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.LeafProof()
	})
}

func Fuzz_Nosy_NodeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var _x3 bool
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1 := newNodeIterator(t1, start)
		_x1.Next(_x3)
	})
}

func Fuzz_Nosy_NodeIterator_NodeBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.NodeBlob()
	})
}

func Fuzz_Nosy_NodeIterator_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.Parent()
	})
}

func Fuzz_Nosy_NodeIterator_Path__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trie
		fill_err = tp.Fill(&t1)
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

		_x1 := newNodeIterator(t1, start)
		_x1.Path()
	})
}

// skipping Fuzz_Nosy_Reader_Node__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.Reader

// skipping Fuzz_Nosy_backend_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

// skipping Fuzz_Nosy_backend_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

// skipping Fuzz_Nosy_backend_Initialized__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

// skipping Fuzz_Nosy_backend_Scheme__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

// skipping Fuzz_Nosy_backend_Size__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

// skipping Fuzz_Nosy_backend_Update__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.backend

func Fuzz_Nosy_hashNode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n hashNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_hashNode_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n hashNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.cache()
	})
}

func Fuzz_Nosy_hashNode_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n hashNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n.encode(w)
	})
}

func Fuzz_Nosy_hashNode_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n hashNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ind string
		fill_err = tp.Fill(&ind)
		if fill_err != nil {
			return
		}

		n.fstring(ind)
	})
}

// skipping Fuzz_Nosy_mptResolver_ForEach__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/common.Hash)

func Fuzz_Nosy_node_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, buf []byte) {
		_x1, err := decodeNode(hash, buf)
		if err != nil {
			return
		}
		_x1.cache()
	})
}

func Fuzz_Nosy_node_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		_x1, err := decodeNode(hash, buf)
		if err != nil {
			return
		}
		_x1.encode(w)
	})
}

func Fuzz_Nosy_node_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, buf []byte, _x3 string) {
		_x1, err := decodeNode(hash, buf)
		if err != nil {
			return
		}
		_x1.fstring(_x3)
	})
}

// skipping Fuzz_Nosy_nodeIteratorHeap_Len__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.nodeIteratorHeap

// skipping Fuzz_Nosy_nodeIteratorHeap_Less__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.nodeIteratorHeap

// skipping Fuzz_Nosy_nodeIteratorHeap_Swap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.nodeIteratorHeap

func Fuzz_Nosy_rawNode_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n rawNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n.EncodeRLP(w)
	})
}

func Fuzz_Nosy_rawNode_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n rawNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.cache()
	})
}

func Fuzz_Nosy_rawNode_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n rawNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n.encode(w)
	})
}

func Fuzz_Nosy_rawNode_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n rawNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ind string
		fill_err = tp.Fill(&ind)
		if fill_err != nil {
			return
		}

		n.fstring(ind)
	})
}

func Fuzz_Nosy_seekError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e seekError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_valueNode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n valueNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_valueNode_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n valueNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.cache()
	})
}

func Fuzz_Nosy_valueNode_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n valueNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var w rlp.EncoderBuffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		n.encode(w)
	})
}

func Fuzz_Nosy_valueNode_fstring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n valueNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ind string
		fill_err = tp.Fill(&ind)
		if fill_err != nil {
			return
		}

		n.fstring(ind)
	})
}

// skipping Fuzz_Nosy_NewDifferenceIterator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeIterator

// skipping Fuzz_Nosy_NewUnionIterator__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/trie.NodeIterator

func Fuzz_Nosy_ResolvePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path []byte) {
		ResolvePath(path)
	})
}

// skipping Fuzz_Nosy_VerifyProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_VerifyRangeProof__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

func Fuzz_Nosy_compactToHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, compact []byte) {
		compactToHex(compact)
	})
}

// skipping Fuzz_Nosy_compareNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.NodeIterator

func Fuzz_Nosy_concat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s1 []byte, s2 []byte) {
		concat(s1, s2...)
	})
}

func Fuzz_Nosy_decodeNibbles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nibbles []byte, bytes []byte) {
		decodeNibbles(nibbles, bytes)
	})
}

func Fuzz_Nosy_decodeRef__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		decodeRef(buf)
	})
}

// skipping Fuzz_Nosy_forGatherChildren__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

// skipping Fuzz_Nosy_get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

// skipping Fuzz_Nosy_hasRightElement__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_hasTerm__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		hasTerm(s)
	})
}

func Fuzz_Nosy_hexToCompact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hex []byte) {
		hexToCompact(hex)
	})
}

func Fuzz_Nosy_hexToCompactInPlace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hex []byte) {
		hexToCompactInPlace(hex)
	})
}

func Fuzz_Nosy_hexToKeybytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hex []byte) {
		hexToKeybytes(hex)
	})
}

func Fuzz_Nosy_keybytesToHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str []byte) {
		keybytesToHex(str)
	})
}

// skipping Fuzz_Nosy_nodeToBytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_prefixLen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		prefixLen(a, b)
	})
}

// skipping Fuzz_Nosy_proofToPath__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node

func Fuzz_Nosy_returnHasherToPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *hasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		returnHasherToPool(h)
	})
}

// skipping Fuzz_Nosy_unsetInternal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie.node
