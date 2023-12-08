package hashdb

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/trie/triestate"
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
		var node common.Hash
		fill_err = tp.Fill(&node)
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

		db.Commit(node, report)
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
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Reader(root)
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
		var child common.Hash
		fill_err = tp.Fill(&child)
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

		db.Reference(child, parent)
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

// skipping Fuzz_Nosy_Database_commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

func Fuzz_Nosy_Database_dereference__(f *testing.F) {
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

		db.dereference(hash)
	})
}

func Fuzz_Nosy_Database_insert__(f *testing.F) {
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
		var node []byte
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.insert(hash, node)
	})
}

func Fuzz_Nosy_Database_node__(f *testing.F) {
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

		db.node(hash)
	})
}

func Fuzz_Nosy_Database_reference__(f *testing.F) {
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
		var child common.Hash
		fill_err = tp.Fill(&child)
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

		db.reference(child, parent)
	})
}

// skipping Fuzz_Nosy_cachedNode_forChildren__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triedb/hashdb.ChildResolver

func Fuzz_Nosy_cleaner_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cleaner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Delete(key)
	})
}

func Fuzz_Nosy_cleaner_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cleaner
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var rlp []byte
		fill_err = tp.Fill(&rlp)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Put(key, rlp)
	})
}

func Fuzz_Nosy_reader_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader *reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var owner common.Hash
		fill_err = tp.Fill(&owner)
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
		if reader == nil {
			return
		}

		reader.Node(owner, path, hash)
	})
}

// skipping Fuzz_Nosy_ChildResolver_ForEach__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triedb/hashdb.ChildResolver
