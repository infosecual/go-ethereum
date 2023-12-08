package trienode

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_MergedNodeSet_Flatten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var set *NodeSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if set == nil {
			return
		}

		s1 := NewWithNodeSet(set)
		s1.Flatten()
	})
}

func Fuzz_Nosy_MergedNodeSet_Merge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var set *NodeSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		var other *NodeSet
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if set == nil || other == nil {
			return
		}

		s1 := NewWithNodeSet(set)
		s1.Merge(other)
	})
}

func Fuzz_Nosy_NodeSet_AddLeaf__(f *testing.F) {
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
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}

		set := NewNodeSet(owner)
		set.AddLeaf(parent, blob)
	})
}

func Fuzz_Nosy_NodeSet_AddNode__(f *testing.F) {
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
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		set := NewNodeSet(owner)
		set.AddNode(path, n)
	})
}

// skipping Fuzz_Nosy_NodeSet_ForEachWithOrder__ because parameters include func, chan, or unsupported interface: func(path string, n *github.com/ethereum/go-ethereum/trie/trienode.Node)

func Fuzz_Nosy_NodeSet_Hashes__(f *testing.F) {
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

		set := NewNodeSet(owner)
		set.Hashes()
	})
}

func Fuzz_Nosy_NodeSet_Merge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o1 common.Hash
		fill_err = tp.Fill(&o1)
		if fill_err != nil {
			return
		}
		var o2 common.Hash
		fill_err = tp.Fill(&o2)
		if fill_err != nil {
			return
		}
		var nodes map[string]*Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		set := NewNodeSet(o1)
		set.Merge(o2, nodes)
	})
}

func Fuzz_Nosy_NodeSet_Size__(f *testing.F) {
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

		set := NewNodeSet(owner)
		set.Size()
	})
}

func Fuzz_Nosy_NodeSet_Summary__(f *testing.F) {
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

		set := NewNodeSet(owner)
		set.Summary()
	})
}

func Fuzz_Nosy_ProofList_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ProofList
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Delete(key)
	})
}

func Fuzz_Nosy_ProofList_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ProofList
		fill_err = tp.Fill(&n)
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
		if n == nil {
			return
		}

		n.Put(key, value)
	})
}

func Fuzz_Nosy_ProofSet_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := NewProofSet()
		db.Delete(key)
	})
}

func Fuzz_Nosy_ProofSet_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := NewProofSet()
		db.Get(key)
	})
}

func Fuzz_Nosy_ProofSet_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		db := NewProofSet()
		db.Has(key)
	})
}

func Fuzz_Nosy_ProofSet_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, value []byte) {
		db := NewProofSet()
		db.Put(key, value)
	})
}

// skipping Fuzz_Nosy_ProofSet_Store__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_ProofList_DataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n ProofList
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.DataSize()
	})
}

func Fuzz_Nosy_ProofList_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n ProofList
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.Set()
	})
}

// skipping Fuzz_Nosy_ProofList_Store__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter
