package snap

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
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

func Fuzz_Nosy_AccountRangePacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *AccountRangePacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_AccountRangePacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *AccountRangePacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_AccountRangePacket_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *AccountRangePacket
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Unpack()
	})
}

func Fuzz_Nosy_ByteCodesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ByteCodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_ByteCodesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ByteCodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetAccountRangePacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetAccountRangePacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetAccountRangePacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetAccountRangePacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetByteCodesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetByteCodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetByteCodesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetByteCodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetStorageRangesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetStorageRangesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetStorageRangesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetStorageRangesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetTrieNodesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetTrieNodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetTrieNodesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetTrieNodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Peer_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ID()
	})
}

func Fuzz_Nosy_Peer_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Log()
	})
}

func Fuzz_Nosy_Peer_RequestAccountRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var origin common.Hash
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		var limit common.Hash
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var bytes uint64
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestAccountRange(id, root, origin, limit, bytes)
	})
}

func Fuzz_Nosy_Peer_RequestByteCodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var bytes uint64
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestByteCodes(id, hashes, bytes)
	})
}

func Fuzz_Nosy_Peer_RequestStorageRanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var accounts []common.Hash
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		var origin []byte
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		var limit []byte
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		var bytes uint64
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestStorageRanges(id, root, accounts, origin, limit, bytes)
	})
}

func Fuzz_Nosy_Peer_RequestTrieNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var paths []TrieNodePathSet
		fill_err = tp.Fill(&paths)
		if fill_err != nil {
			return
		}
		var bytes uint64
		fill_err = tp.Fill(&bytes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestTrieNodes(id, root, paths, bytes)
	})
}

func Fuzz_Nosy_Peer_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Version()
	})
}

func Fuzz_Nosy_StorageRangesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StorageRangesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_StorageRangesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StorageRangesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_StorageRangesPacket_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *StorageRangesPacket
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Unpack()
	})
}

// skipping Fuzz_Nosy_Syncer_OnAccounts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_Syncer_OnByteCodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_Syncer_OnStorage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_Syncer_OnTrieNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

func Fuzz_Nosy_Syncer_Progress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Progress()
	})
}

// skipping Fuzz_Nosy_Syncer_Register__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_Syncer_Sync__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_Syncer_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Unregister(id)
	})
}

// skipping Fuzz_Nosy_Syncer_assignAccountTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/snap.accountResponse

// skipping Fuzz_Nosy_Syncer_assignBytecodeHealTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/snap.bytecodeHealResponse

// skipping Fuzz_Nosy_Syncer_assignBytecodeTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/snap.bytecodeResponse

// skipping Fuzz_Nosy_Syncer_assignStorageTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/snap.storageResponse

// skipping Fuzz_Nosy_Syncer_assignTrienodeHealTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/snap.trienodeHealResponse

func Fuzz_Nosy_Syncer_cleanAccountTasks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.cleanAccountTasks()
	})
}

// skipping Fuzz_Nosy_Syncer_cleanPath__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Batch

func Fuzz_Nosy_Syncer_cleanStorageTasks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.cleanStorageTasks()
	})
}

func Fuzz_Nosy_Syncer_commitHealer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.commitHealer(force)
	})
}

func Fuzz_Nosy_Syncer_forwardAccountTask__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var task *accountTask
		fill_err = tp.Fill(&task)
		if fill_err != nil {
			return
		}
		if s == nil || task == nil {
			return
		}

		s.forwardAccountTask(task)
	})
}

func Fuzz_Nosy_Syncer_loadSyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.loadSyncStatus()
	})
}

// skipping Fuzz_Nosy_Syncer_onByteCodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_Syncer_onHealByteCodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

func Fuzz_Nosy_Syncer_onHealState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var paths [][]byte
		fill_err = tp.Fill(&paths)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.onHealState(paths, value)
	})
}

func Fuzz_Nosy_Syncer_processAccountResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *accountResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processAccountResponse(res)
	})
}

func Fuzz_Nosy_Syncer_processBytecodeHealResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *bytecodeHealResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processBytecodeHealResponse(res)
	})
}

func Fuzz_Nosy_Syncer_processBytecodeResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *bytecodeResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processBytecodeResponse(res)
	})
}

func Fuzz_Nosy_Syncer_processStorageResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *storageResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processStorageResponse(res)
	})
}

func Fuzz_Nosy_Syncer_processTrienodeHealResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *trienodeHealResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processTrienodeHealResponse(res)
	})
}

func Fuzz_Nosy_Syncer_report__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.report(force)
	})
}

func Fuzz_Nosy_Syncer_reportHealProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.reportHealProgress(force)
	})
}

func Fuzz_Nosy_Syncer_reportSyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.reportSyncProgress(force)
	})
}

func Fuzz_Nosy_Syncer_revertAccountRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *accountRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertAccountRequest(req)
	})
}

func Fuzz_Nosy_Syncer_revertBytecodeHealRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *bytecodeHealRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertBytecodeHealRequest(req)
	})
}

func Fuzz_Nosy_Syncer_revertBytecodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *bytecodeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertBytecodeRequest(req)
	})
}

func Fuzz_Nosy_Syncer_revertRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.revertRequests(peer)
	})
}

func Fuzz_Nosy_Syncer_revertStorageRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *storageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertStorageRequest(req)
	})
}

func Fuzz_Nosy_Syncer_revertTrienodeHealRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *trienodeHealRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertTrienodeHealRequest(req)
	})
}

func Fuzz_Nosy_Syncer_saveSyncStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.saveSyncStatus()
	})
}

func Fuzz_Nosy_Syncer_scheduleRevertAccountRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *accountRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertAccountRequest(req)
	})
}

func Fuzz_Nosy_Syncer_scheduleRevertBytecodeHealRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *bytecodeHealRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertBytecodeHealRequest(req)
	})
}

func Fuzz_Nosy_Syncer_scheduleRevertBytecodeRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *bytecodeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertBytecodeRequest(req)
	})
}

func Fuzz_Nosy_Syncer_scheduleRevertStorageRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *storageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertStorageRequest(req)
	})
}

func Fuzz_Nosy_Syncer_scheduleRevertTrienodeHealRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Syncer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *trienodeHealRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertTrienodeHealRequest(req)
	})
}

func Fuzz_Nosy_TrieNodesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TrieNodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_TrieNodesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TrieNodesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_capacitySort_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *capacitySort
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_capacitySort_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *capacitySort
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Less(i, j)
	})
}

func Fuzz_Nosy_capacitySort_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *capacitySort
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Swap(i, j)
	})
}

func Fuzz_Nosy_hashRange_End__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var start common.Hash
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}

		r := newHashRange(start, num)
		r.End()
	})
}

func Fuzz_Nosy_hashRange_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var start common.Hash
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}

		r := newHashRange(start, num)
		r.Next()
	})
}

func Fuzz_Nosy_hashRange_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var start common.Hash
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var num uint64
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}

		r := newHashRange(start, num)
		r.Start()
	})
}

func Fuzz_Nosy_healRequestSort_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *healRequestSort
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Len()
	})
}

func Fuzz_Nosy_healRequestSort_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *healRequestSort
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Less(i, j)
	})
}

func Fuzz_Nosy_healRequestSort_Merge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *healRequestSort
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Merge()
	})
}

func Fuzz_Nosy_healRequestSort_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *healRequestSort
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Swap(i, j)
	})
}

// skipping Fuzz_Nosy_Backend_Chain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Backend

// skipping Fuzz_Nosy_Backend_Handle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Backend

// skipping Fuzz_Nosy_Backend_PeerInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Backend

// skipping Fuzz_Nosy_Backend_RunPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Backend

// skipping Fuzz_Nosy_Packet_Kind__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Packet

// skipping Fuzz_Nosy_Packet_Name__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Packet

// skipping Fuzz_Nosy_SyncPeer_ID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_SyncPeer_Log__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_SyncPeer_RequestAccountRange__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_SyncPeer_RequestByteCodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_SyncPeer_RequestStorageRanges__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

// skipping Fuzz_Nosy_SyncPeer_RequestTrieNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.SyncPeer

func Fuzz_Nosy_enrEntry_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e enrEntry
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.ENRKey()
	})
}

// skipping Fuzz_Nosy_MakeProtocols__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Backend

func Fuzz_Nosy_ServiceGetAccountRangeQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var req *GetAccountRangePacket
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if chain == nil || req == nil {
			return
		}

		ServiceGetAccountRangeQuery(chain, req)
	})
}

func Fuzz_Nosy_ServiceGetByteCodesQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var req *GetByteCodesPacket
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if chain == nil || req == nil {
			return
		}

		ServiceGetByteCodesQuery(chain, req)
	})
}

func Fuzz_Nosy_ServiceGetStorageRangesQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var req *GetStorageRangesPacket
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if chain == nil || req == nil {
			return
		}

		ServiceGetStorageRangesQuery(chain, req)
	})
}

func Fuzz_Nosy_ServiceGetTrieNodesQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var req *GetTrieNodesPacket
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if chain == nil || req == nil {
			return
		}

		ServiceGetTrieNodesQuery(chain, req, start)
	})
}

func Fuzz_Nosy_estimateRemainingSlots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hashes int
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var last common.Hash
		fill_err = tp.Fill(&last)
		if fill_err != nil {
			return
		}

		estimateRemainingSlots(hashes, last)
	})
}

func Fuzz_Nosy_sortByAccountPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var paths []string
		fill_err = tp.Fill(&paths)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

		sortByAccountPath(paths, hashes)
	})
}
