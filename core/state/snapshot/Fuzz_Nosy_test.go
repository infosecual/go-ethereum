package snapshot

import (
	"bytes"
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

func Fuzz_Nosy_Tree_AccountIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.AccountIterator(root, seek)
	})
}

func Fuzz_Nosy_Tree_Cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
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
		if t1 == nil {
			return
		}

		t1.Cap(root, layers)
	})
}

func Fuzz_Nosy_Tree_Disable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Disable()
	})
}

func Fuzz_Nosy_Tree_DiskRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.DiskRoot()
	})
}

func Fuzz_Nosy_Tree_Journal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Journal(root)
	})
}

func Fuzz_Nosy_Tree_Rebuild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Rebuild(root)
	})
}

func Fuzz_Nosy_Tree_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Release()
	})
}

func Fuzz_Nosy_Tree_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Size()
	})
}

func Fuzz_Nosy_Tree_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var blockRoot common.Hash
		fill_err = tp.Fill(&blockRoot)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Snapshot(blockRoot)
	})
}

func Fuzz_Nosy_Tree_Snapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var limits int
		fill_err = tp.Fill(&limits)
		if fill_err != nil {
			return
		}
		var nodisk bool
		fill_err = tp.Fill(&nodisk)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Snapshots(root, limits, nodisk)
	})
}

func Fuzz_Nosy_Tree_StorageIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.StorageIterator(root, account, seek)
	})
}

func Fuzz_Nosy_Tree_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var blockRoot common.Hash
		fill_err = tp.Fill(&blockRoot)
		if fill_err != nil {
			return
		}
		var parentRoot common.Hash
		fill_err = tp.Fill(&parentRoot)
		if fill_err != nil {
			return
		}
		var destructs map[common.Hash]struct{}
		fill_err = tp.Fill(&destructs)
		if fill_err != nil {
			return
		}
		var accounts map[common.Hash][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storage map[common.Hash]map[common.Hash][]byte
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Update(blockRoot, parentRoot, destructs, accounts, storage)
	})
}

func Fuzz_Nosy_Tree_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Verify(root)
	})
}

func Fuzz_Nosy_Tree_cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var diff *diffLayer
		fill_err = tp.Fill(&diff)
		if fill_err != nil {
			return
		}
		var layers int
		fill_err = tp.Fill(&layers)
		if fill_err != nil {
			return
		}
		if t1 == nil || diff == nil {
			return
		}

		t1.cap(diff, layers)
	})
}

func Fuzz_Nosy_Tree_diskRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.diskRoot()
	})
}

func Fuzz_Nosy_Tree_disklayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.disklayer()
	})
}

func Fuzz_Nosy_Tree_generating__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.generating()
	})
}

func Fuzz_Nosy_Tree_waitBuild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Tree
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.waitBuild()
	})
}

func Fuzz_Nosy_abortErr_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *abortErr
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

func Fuzz_Nosy_binaryIterator_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Account()
	})
}

func Fuzz_Nosy_binaryIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
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

func Fuzz_Nosy_binaryIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
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

func Fuzz_Nosy_binaryIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
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

func Fuzz_Nosy_binaryIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
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

func Fuzz_Nosy_binaryIterator_Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *binaryIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Slot()
	})
}

func Fuzz_Nosy_diffAccountIterator_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffAccountIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Account()
	})
}

func Fuzz_Nosy_diffAccountIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffAccountIterator
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

func Fuzz_Nosy_diffAccountIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffAccountIterator
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

func Fuzz_Nosy_diffAccountIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffAccountIterator
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

func Fuzz_Nosy_diffAccountIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffAccountIterator
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

func Fuzz_Nosy_diffLayer_Account__(f *testing.F) {
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
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.Account(hash)
	})
}

func Fuzz_Nosy_diffLayer_AccountIterator__(f *testing.F) {
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
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.AccountIterator(seek)
	})
}

func Fuzz_Nosy_diffLayer_AccountList__(f *testing.F) {
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

		dl.AccountList()
	})
}

func Fuzz_Nosy_diffLayer_AccountRLP__(f *testing.F) {
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
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.AccountRLP(hash)
	})
}

func Fuzz_Nosy_diffLayer_Journal__(f *testing.F) {
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
		var buffer *bytes.Buffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if dl == nil || buffer == nil {
			return
		}

		dl.Journal(buffer)
	})
}

func Fuzz_Nosy_diffLayer_Parent__(f *testing.F) {
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

		dl.Parent()
	})
}

func Fuzz_Nosy_diffLayer_Root__(f *testing.F) {
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

		dl.Root()
	})
}

func Fuzz_Nosy_diffLayer_Stale__(f *testing.F) {
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

		dl.Stale()
	})
}

func Fuzz_Nosy_diffLayer_Storage__(f *testing.F) {
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
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		var storageHash common.Hash
		fill_err = tp.Fill(&storageHash)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.Storage(accountHash, storageHash)
	})
}

func Fuzz_Nosy_diffLayer_StorageIterator__(f *testing.F) {
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
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.StorageIterator(account, seek)
	})
}

func Fuzz_Nosy_diffLayer_StorageList__(f *testing.F) {
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
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.StorageList(accountHash)
	})
}

func Fuzz_Nosy_diffLayer_Update__(f *testing.F) {
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
		var blockRoot common.Hash
		fill_err = tp.Fill(&blockRoot)
		if fill_err != nil {
			return
		}
		var destructs map[common.Hash]struct{}
		fill_err = tp.Fill(&destructs)
		if fill_err != nil {
			return
		}
		var accounts map[common.Hash][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storage map[common.Hash]map[common.Hash][]byte
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.Update(blockRoot, destructs, accounts, storage)
	})
}

func Fuzz_Nosy_diffLayer_accountRLP__(f *testing.F) {
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

		dl.accountRLP(hash, depth)
	})
}

func Fuzz_Nosy_diffLayer_flatten__(f *testing.F) {
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

		dl.flatten()
	})
}

func Fuzz_Nosy_diffLayer_initBinaryAccountIterator__(f *testing.F) {
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

		dl.initBinaryAccountIterator()
	})
}

func Fuzz_Nosy_diffLayer_initBinaryStorageIterator__(f *testing.F) {
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
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.initBinaryStorageIterator(account)
	})
}

func Fuzz_Nosy_diffLayer_newBinaryAccountIterator__(f *testing.F) {
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

		dl.newBinaryAccountIterator()
	})
}

func Fuzz_Nosy_diffLayer_newBinaryStorageIterator__(f *testing.F) {
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
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if dl == nil {
			return
		}

		dl.newBinaryStorageIterator(account)
	})
}

func Fuzz_Nosy_diffLayer_rebloom__(f *testing.F) {
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
		var origin *diskLayer
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if dl == nil || origin == nil {
			return
		}

		dl.rebloom(origin)
	})
}

func Fuzz_Nosy_diffLayer_storage__(f *testing.F) {
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
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		var storageHash common.Hash
		fill_err = tp.Fill(&storageHash)
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

		dl.storage(accountHash, storageHash, depth)
	})
}

func Fuzz_Nosy_diffStorageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffStorageIterator
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

func Fuzz_Nosy_diffStorageIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffStorageIterator
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

func Fuzz_Nosy_diffStorageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffStorageIterator
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

func Fuzz_Nosy_diffStorageIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffStorageIterator
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

func Fuzz_Nosy_diffStorageIterator_Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diffStorageIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Slot()
	})
}

func Fuzz_Nosy_diskAccountIterator_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskAccountIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Account()
	})
}

func Fuzz_Nosy_diskAccountIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskAccountIterator
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

func Fuzz_Nosy_diskAccountIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskAccountIterator
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

func Fuzz_Nosy_diskAccountIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskAccountIterator
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

func Fuzz_Nosy_diskAccountIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskAccountIterator
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

func Fuzz_Nosy_diskLayer_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Account(hash)
	})
}

func Fuzz_Nosy_diskLayer_AccountIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.AccountIterator(seek)
	})
}

func Fuzz_Nosy_diskLayer_AccountRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.AccountRLP(hash)
	})
}

func Fuzz_Nosy_diskLayer_Journal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var buffer *bytes.Buffer
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		if bottom == nil || buffer == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Journal(buffer)
	})
}

func Fuzz_Nosy_diskLayer_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Parent()
	})
}

func Fuzz_Nosy_diskLayer_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Release()
	})
}

func Fuzz_Nosy_diskLayer_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Root()
	})
}

func Fuzz_Nosy_diskLayer_Stale__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Stale()
	})
}

func Fuzz_Nosy_diskLayer_Storage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var accountHash common.Hash
		fill_err = tp.Fill(&accountHash)
		if fill_err != nil {
			return
		}
		var storageHash common.Hash
		fill_err = tp.Fill(&storageHash)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Storage(accountHash, storageHash)
	})
}

func Fuzz_Nosy_diskLayer_StorageIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.StorageIterator(account, seek)
	})
}

func Fuzz_Nosy_diskLayer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var destructs map[common.Hash]struct{}
		fill_err = tp.Fill(&destructs)
		if fill_err != nil {
			return
		}
		var accounts map[common.Hash][]byte
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var storage map[common.Hash]map[common.Hash][]byte
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		if bottom == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.Update(blockHash, destructs, accounts, storage)
	})
}

func Fuzz_Nosy_diskLayer_checkAndFlush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var current []byte
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}
		if bottom == nil || ctx == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.checkAndFlush(ctx, current)
	})
}

func Fuzz_Nosy_diskLayer_generate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bottom *diffLayer
		fill_err = tp.Fill(&bottom)
		if fill_err != nil {
			return
		}
		var stats *generatorStats
		fill_err = tp.Fill(&stats)
		if fill_err != nil {
			return
		}
		if bottom == nil || stats == nil {
			return
		}

		dl := diffToDisk(bottom)
		dl.generate(stats)
	})
}

// skipping Fuzz_Nosy_diskLayer_generateRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.onStateCallback

// skipping Fuzz_Nosy_diskLayer_proveRange__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

func Fuzz_Nosy_diskStorageIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskStorageIterator
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

func Fuzz_Nosy_diskStorageIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskStorageIterator
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

func Fuzz_Nosy_diskStorageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskStorageIterator
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

func Fuzz_Nosy_diskStorageIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskStorageIterator
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

func Fuzz_Nosy_diskStorageIterator_Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *diskStorageIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Slot()
	})
}

func Fuzz_Nosy_fastIterator_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Account()
	})
}

func Fuzz_Nosy_fastIterator_Debug__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Debug()
	})
}

func Fuzz_Nosy_fastIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Error()
	})
}

func Fuzz_Nosy_fastIterator_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Hash()
	})
}

func Fuzz_Nosy_fastIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Next()
	})
}

func Fuzz_Nosy_fastIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Release()
	})
}

func Fuzz_Nosy_fastIterator_Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.Slot()
	})
}

func Fuzz_Nosy_fastIterator_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.init()
	})
}

func Fuzz_Nosy_fastIterator_move__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var newpos int
		fill_err = tp.Fill(&newpos)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.move(index, newpos)
	})
}

func Fuzz_Nosy_fastIterator_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		var accountIterator bool
		fill_err = tp.Fill(&accountIterator)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		fi, err := newFastIterator(tree, root, account, seek, accountIterator)
		if err != nil {
			return
		}
		fi.next(idx)
	})
}

func Fuzz_Nosy_generateStats_finishAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, done uint64) {
		stat := newGenerateStats()
		stat.finishAccounts(done)
	})
}

func Fuzz_Nosy_generateStats_finishContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var done uint64
		fill_err = tp.Fill(&done)
		if fill_err != nil {
			return
		}

		stat := newGenerateStats()
		stat.finishContract(account, done)
	})
}

func Fuzz_Nosy_generateStats_progressAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var done uint64
		fill_err = tp.Fill(&done)
		if fill_err != nil {
			return
		}

		stat := newGenerateStats()
		stat.progressAccounts(account, done)
	})
}

func Fuzz_Nosy_generateStats_progressContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var done uint64
		fill_err = tp.Fill(&done)
		if fill_err != nil {
			return
		}

		stat := newGenerateStats()
		stat.progressContract(account, slot, done)
	})
}

func Fuzz_Nosy_generatorContext_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.close()
	})
}

func Fuzz_Nosy_generatorContext_iterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.iterator(kind)
	})
}

func Fuzz_Nosy_generatorContext_openIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.openIterator(kind, start)
	})
}

func Fuzz_Nosy_generatorContext_removeStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.removeStorageAt(account)
	})
}

func Fuzz_Nosy_generatorContext_removeStorageBefore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.removeStorageBefore(account)
	})
}

func Fuzz_Nosy_generatorContext_removeStorageLeft__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.removeStorageLeft()
	})
}

func Fuzz_Nosy_generatorContext_reopenIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *generatorContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		ctx.reopenIterator(kind)
	})
}

func Fuzz_Nosy_generatorStats_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gs *generatorStats
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var marker []byte
		fill_err = tp.Fill(&marker)
		if fill_err != nil {
			return
		}
		if gs == nil {
			return
		}

		gs.Log(msg, root, marker)
	})
}

func Fuzz_Nosy_holdableIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
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

func Fuzz_Nosy_holdableIterator_Hold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Hold()
	})
}

func Fuzz_Nosy_holdableIterator_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
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

func Fuzz_Nosy_holdableIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
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

func Fuzz_Nosy_holdableIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
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

func Fuzz_Nosy_holdableIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *holdableIterator
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

// skipping Fuzz_Nosy_proofResult_forEach__ because parameters include func, chan, or unsupported interface: func(key []byte, val []byte) error

func Fuzz_Nosy_proofResult_last__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var result *proofResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if result == nil {
			return
		}

		result.last()
	})
}

func Fuzz_Nosy_proofResult_valid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var result *proofResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if result == nil {
			return
		}

		result.valid()
	})
}

func Fuzz_Nosy_weightedIterator_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *weightedIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var other *weightedIterator
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if it == nil || other == nil {
			return
		}

		it.Cmp(other)
	})
}

func Fuzz_Nosy_AccountIterator_Account__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		_x1, err := newFastAccountIterator(tree, root, seek)
		if err != nil {
			return
		}
		_x1.Account()
	})
}

// skipping Fuzz_Nosy_Iterator_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Iterator

// skipping Fuzz_Nosy_Iterator_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Iterator

// skipping Fuzz_Nosy_Iterator_Next__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Iterator

// skipping Fuzz_Nosy_Iterator_Release__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Iterator

// skipping Fuzz_Nosy_Snapshot_Account__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Snapshot

// skipping Fuzz_Nosy_Snapshot_AccountRLP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Snapshot

// skipping Fuzz_Nosy_Snapshot_Root__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Snapshot

// skipping Fuzz_Nosy_Snapshot_Storage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.Snapshot

func Fuzz_Nosy_StorageIterator_Slot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tree *Tree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var account common.Hash
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var seek common.Hash
		fill_err = tp.Fill(&seek)
		if fill_err != nil {
			return
		}
		if tree == nil {
			return
		}

		_x1, err := newFastStorageIterator(tree, root, account, seek)
		if err != nil {
			return
		}
		_x1.Slot()
	})
}

func Fuzz_Nosy_accountBloomHasher_BlockSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.BlockSize()
	})
}

func Fuzz_Nosy_accountBloomHasher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Reset()
	})
}

func Fuzz_Nosy_accountBloomHasher_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Size()
	})
}

func Fuzz_Nosy_accountBloomHasher_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		h.Sum(b)
	})
}

func Fuzz_Nosy_accountBloomHasher_Sum64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Sum64()
	})
}

func Fuzz_Nosy_accountBloomHasher_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h accountBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		h.Write(p)
	})
}

func Fuzz_Nosy_destructBloomHasher_BlockSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.BlockSize()
	})
}

func Fuzz_Nosy_destructBloomHasher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Reset()
	})
}

func Fuzz_Nosy_destructBloomHasher_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Size()
	})
}

func Fuzz_Nosy_destructBloomHasher_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		h.Sum(b)
	})
}

func Fuzz_Nosy_destructBloomHasher_Sum64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Sum64()
	})
}

func Fuzz_Nosy_destructBloomHasher_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h destructBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		h.Write(p)
	})
}

// skipping Fuzz_Nosy_snapshot_AccountIterator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

// skipping Fuzz_Nosy_snapshot_Journal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

// skipping Fuzz_Nosy_snapshot_Parent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

// skipping Fuzz_Nosy_snapshot_Stale__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

// skipping Fuzz_Nosy_snapshot_StorageIterator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

// skipping Fuzz_Nosy_snapshot_Update__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state/snapshot.snapshot

func Fuzz_Nosy_storageBloomHasher_BlockSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.BlockSize()
	})
}

func Fuzz_Nosy_storageBloomHasher_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Reset()
	})
}

func Fuzz_Nosy_storageBloomHasher_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Size()
	})
}

func Fuzz_Nosy_storageBloomHasher_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		h.Sum(b)
	})
}

func Fuzz_Nosy_storageBloomHasher_Sum64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Sum64()
	})
}

func Fuzz_Nosy_storageBloomHasher_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageBloomHasher
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		h.Write(p)
	})
}

func Fuzz_Nosy_ParseGeneratorStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, generatorBlob []byte) {
		ParseGeneratorStatus(generatorBlob)
	})
}

func Fuzz_Nosy_increaseKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		increaseKey(key)
	})
}

// skipping Fuzz_Nosy_journalProgress__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

// skipping Fuzz_Nosy_loadAndParseJournal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_loadSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStore

// skipping Fuzz_Nosy_runReport__ because parameters include func, chan, or unsupported interface: chan bool

// skipping Fuzz_Nosy_stackTrieGenerate__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter
