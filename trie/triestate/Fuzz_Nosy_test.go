package triestate

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

func Fuzz_Nosy_Set_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var accounts map[common.Address][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storages map[common.Address]map[common.Hash][]byte
		fill_err = tp.Fill(&storages)
		if fill_err != nil {
			return
		}
		var incomplete map[common.Address]struct{}
		fill_err = tp.Fill(&incomplete)
		if fill_err != nil {
			return
		}

		s := New(accounts, storages, incomplete)
		s.Size()
	})
}

func Fuzz_Nosy_hasher_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		h := newHasher()
		h.hash(d1)
	})
}

// skipping Fuzz_Nosy_Trie_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.Trie

// skipping Fuzz_Nosy_Trie_Delete__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.Trie

// skipping Fuzz_Nosy_Trie_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.Trie

// skipping Fuzz_Nosy_Trie_Update__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.Trie

// skipping Fuzz_Nosy_TrieLoader_OpenStorageTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.TrieLoader

// skipping Fuzz_Nosy_TrieLoader_OpenTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.TrieLoader

// skipping Fuzz_Nosy_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/trie/triestate.TrieLoader
