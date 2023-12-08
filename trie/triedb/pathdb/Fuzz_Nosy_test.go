package pathdb

import (
	"io"
	"testing"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/rlp"
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

func Fuzz_Nosy_Config_sanitize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.sanitize()
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

// skipping Fuzz_Nosy_Database_Recover__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.TrieLoader

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
		var parentRoot common.Hash
		fill_err = tp.Fill(&parentRoot)
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

		db.Update(root, parentRoot, block, nodes, states)
	})
}

// skipping Fuzz_Nosy_Database_loadDiffLayer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triedb/pathdb.layer

func Fuzz_Nosy_Database_loadDiskLayer__(f *testing.F) {
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
		var r *rlp.Stream
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if db == nil || r == nil {
			return
		}

		db.loadDiskLayer(r)
	})
}

func Fuzz_Nosy_Database_loadJournal__(f *testing.F) {
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
		var diskRoot common.Hash
		fill_err = tp.Fill(&diskRoot)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.loadJournal(diskRoot)
	})
}

func Fuzz_Nosy_Database_loadLayers__(f *testing.F) {
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

		db.loadLayers()
	})
}

func Fuzz_Nosy_Database_modifyAllowed__(f *testing.F) {
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

		db.modifyAllowed()
	})
}

func Fuzz_Nosy_accountIndex_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *accountIndex
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.decode(blob)
	})
}

func Fuzz_Nosy_accountIndex_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *accountIndex
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.encode()
	})
}

func Fuzz_Nosy_decoder_readAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *decoder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var pos int
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.readAccount(pos)
	})
}

func Fuzz_Nosy_decoder_readStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *decoder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var accIndex accountIndex
		fill_err = tp.Fill(&accIndex)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.readStorage(accIndex)
	})
}

func Fuzz_Nosy_decoder_verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *decoder
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.verify()
	})
}

func Fuzz_Nosy_diffLayer_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
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
		if dl == nil {
			return
		}

		dl.Node(owner, path, hash)
	})
}

func Fuzz_Nosy_diffLayer_journal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.journal(w)
	})
}

func Fuzz_Nosy_diffLayer_node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
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
		var depth int
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.node(owner, path, hash, depth)
	})
}

func Fuzz_Nosy_diffLayer_parentLayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.parentLayer()
	})
}

func Fuzz_Nosy_diffLayer_persist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.persist(force)
	})
}

func Fuzz_Nosy_diffLayer_rootHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.rootHash()
	})
}

func Fuzz_Nosy_diffLayer_stateID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.stateID()
	})
}

func Fuzz_Nosy_diffLayer_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dl *diffLayer
		fill_err = tp.Fill(&dl)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var states *triestate.Set
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		if dl == nil || states == nil {
			return
		}

		dl.update(root, id, block, nodes, states)
	})
}

func Fuzz_Nosy_diskLayer_Node__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
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
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.Node(owner, path, hash)
	})
}

func Fuzz_Nosy_diskLayer_commit__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil || bottom == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.commit(bottom, force)
	})
}

func Fuzz_Nosy_diskLayer_isStale__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.isStale()
	})
}

func Fuzz_Nosy_diskLayer_journal__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.journal(w)
	})
}

func Fuzz_Nosy_diskLayer_markStale__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.markStale()
	})
}

func Fuzz_Nosy_diskLayer_parentLayer__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.parentLayer()
	})
}

func Fuzz_Nosy_diskLayer_resetCache__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.resetCache()
	})
}

// skipping Fuzz_Nosy_diskLayer_revert__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.TrieLoader

func Fuzz_Nosy_diskLayer_rootHash__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.rootHash()
	})
}

func Fuzz_Nosy_diskLayer_setBufferSize__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.setBufferSize(size)
	})
}

func Fuzz_Nosy_diskLayer_size__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.size()
	})
}

func Fuzz_Nosy_diskLayer_stateID__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil {
			return
		}

		dl := newDiskLayer(root, id, db, cleans, buffer)
		dl.stateID()
	})
}

func Fuzz_Nosy_diskLayer_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 common.Hash
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		var i2 uint64
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var db *Database
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var cleans *fastcache.Cache
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var buffer *nodebuffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		var r6 common.Hash
		fill_err = tp.Fill(&r6)
		if fill_err != nil {
			return
		}
		var i7 uint64
		fill_err = tp.Fill(&i7)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var states *triestate.Set
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		if db == nil || cleans == nil || buffer == nil || states == nil {
			return
		}

		dl := newDiskLayer(r1, i2, db, cleans, buffer)
		dl.update(r6, i7, block, nodes, states)
	})
}

func Fuzz_Nosy_hashLoader_OpenStorageTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accounts map[common.Hash][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storages map[common.Hash]map[common.Hash][]byte
		fill_err = tp.Fill(&storages)
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

		l := newHashLoader(accounts, storages)
		l.OpenStorageTrie(stateRoot, addrHash, root)
	})
}

func Fuzz_Nosy_hashLoader_OpenTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accounts map[common.Hash][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storages map[common.Hash]map[common.Hash][]byte
		fill_err = tp.Fill(&storages)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}

		l := newHashLoader(accounts, storages)
		l.OpenTrie(root)
	})
}

func Fuzz_Nosy_hasher_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		h := newHasher()
		h.hash(d1)
	})
}

func Fuzz_Nosy_history_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var freezer *rawdb.ResettableFreezer
		fill_err = tp.Fill(&freezer)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var accountData []byte
		fill_err = tp.Fill(&accountData)
		if fill_err != nil {
			return
		}
		var storageData []byte
		fill_err = tp.Fill(&storageData)
		if fill_err != nil {
			return
		}
		var accountIndexes []byte
		fill_err = tp.Fill(&accountIndexes)
		if fill_err != nil {
			return
		}
		var storageIndexes []byte
		fill_err = tp.Fill(&storageIndexes)
		if fill_err != nil {
			return
		}
		if freezer == nil {
			return
		}

		h, err := readHistory(freezer, id)
		if err != nil {
			return
		}
		h.decode(accountData, storageData, accountIndexes, storageIndexes)
	})
}

func Fuzz_Nosy_history_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var freezer *rawdb.ResettableFreezer
		fill_err = tp.Fill(&freezer)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if freezer == nil {
			return
		}

		h, err := readHistory(freezer, id)
		if err != nil {
			return
		}
		h.encode()
	})
}

func Fuzz_Nosy_layerTree_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *layerTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var parentRoot common.Hash
		fill_err = tp.Fill(&parentRoot)
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
		if tree == nil || nodes == nil || states == nil {
			return
		}

		tree.add(root, parentRoot, block, nodes, states)
	})
}

func Fuzz_Nosy_layerTree_bottom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *layerTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		tree.bottom()
	})
}

func Fuzz_Nosy_layerTree_cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *layerTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var layers int
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		tree.cap(root, layers)
	})
}

// skipping Fuzz_Nosy_layerTree_forEach__ because parameters include func, chan, or unsupported interface: func(github.com/ethereum/go-ethereum/trie/triedb/pathdb.layer)

func Fuzz_Nosy_layerTree_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *layerTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		tree.get(root)
	})
}

func Fuzz_Nosy_layerTree_len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *layerTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		tree.len()
	})
}

// skipping Fuzz_Nosy_layerTree_reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triedb/pathdb.layer

func Fuzz_Nosy_meta_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.decode(blob)
	})
}

func Fuzz_Nosy_meta_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.encode()
	})
}

func Fuzz_Nosy_nodebuffer_commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var n2 map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&n2)
		if fill_err != nil {
			return
		}
		var layers uint64
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}
		var n4 map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&n4)
		if fill_err != nil {
			return
		}

		b := newNodeBuffer(limit, n2, layers)
		b.commit(n4)
	})
}

func Fuzz_Nosy_nodebuffer_empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var layers uint64
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}

		b := newNodeBuffer(limit, nodes, layers)
		b.empty()
	})
}

// skipping Fuzz_Nosy_nodebuffer_flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

func Fuzz_Nosy_nodebuffer_node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var layers uint64
		fill_err = tp.Fill(&layers)
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

		b := newNodeBuffer(limit, nodes, layers)
		b.node(owner, path, hash)
	})
}

func Fuzz_Nosy_nodebuffer_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var layers uint64
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}

		b := newNodeBuffer(limit, nodes, layers)
		b.reset()
	})
}

// skipping Fuzz_Nosy_nodebuffer_revert__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueReader

// skipping Fuzz_Nosy_nodebuffer_setSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

func Fuzz_Nosy_nodebuffer_updateSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var layers uint64
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}
		var delta int64
		fill_err = tp.Fill(&delta)
		if fill_err != nil {
			return
		}

		b := newNodeBuffer(limit, nodes, layers)
		b.updateSize(delta)
	})
}

func Fuzz_Nosy_slotIndex_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *slotIndex
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.decode(blob)
	})
}

func Fuzz_Nosy_slotIndex_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *slotIndex
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		i.encode()
	})
}

func Fuzz_Nosy_testHasher_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var owner common.Hash
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var cleans map[common.Hash][]byte
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var collectLeaf bool
		fill_err = tp.Fill(&collectLeaf)
		if fill_err != nil {
			return
		}

		h, err := newTestHasher(owner, root, cleans)
		if err != nil {
			return
		}
		h.Commit(collectLeaf)
	})
}

func Fuzz_Nosy_testHasher_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var owner common.Hash
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var cleans map[common.Hash][]byte
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		h, err := newTestHasher(owner, root, cleans)
		if err != nil {
			return
		}
		h.Delete(key)
	})
}

func Fuzz_Nosy_testHasher_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var owner common.Hash
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var cleans map[common.Hash][]byte
		fill_err = tp.Fill(&cleans)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		h, err := newTestHasher(owner, root, cleans)
		if err != nil {
			return
		}
		h.Get(key)
	})
}

func Fuzz_Nosy_testHasher_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var owner common.Hash
		fill_err = tp.Fill(&owner)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var cleans map[common.Hash][]byte
		fill_err = tp.Fill(&cleans)
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

		h, err := newTestHasher(owner, root, cleans)
		if err != nil {
			return
		}
		h.Update(key, value)
	})
}

func Fuzz_Nosy_layer_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
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
		if layer == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.Node(owner, path, hash)
	})
}

func Fuzz_Nosy_layer_journal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if layer == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.journal(w)
	})
}

func Fuzz_Nosy_layer_parentLayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if layer == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.parentLayer()
	})
}

func Fuzz_Nosy_layer_rootHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if layer == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.rootHash()
	})
}

func Fuzz_Nosy_layer_stateID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if layer == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.stateID()
	})
}

func Fuzz_Nosy_layer_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var layer *diffLayer
		fill_err = tp.Fill(&layer)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var nodes map[common.Hash]map[string]*trienode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var states *triestate.Set
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		if layer == nil || states == nil {
			return
		}

		_x1, err := diffToDisk(layer, force)
		if err != nil {
			return
		}
		_x1.update(root, id, block, nodes, states)
	})
}

func Fuzz_Nosy_cacheKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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

		cacheKey(owner, path)
	})
}

func Fuzz_Nosy_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var states map[common.Hash][]byte
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}

		hash(states)
	})
}

// skipping Fuzz_Nosy_truncateFromHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batcher

// skipping Fuzz_Nosy_truncateFromTail__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batcher

// skipping Fuzz_Nosy_writeNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch
