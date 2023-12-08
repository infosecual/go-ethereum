package state

import (
	"encoding/json"
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/trienode"
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

func Fuzz_Nosy_Dump_OnAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Dump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var addr *common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var account DumpAccount
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if d == nil || addr == nil {
			return
		}

		d.OnAccount(addr, account)
	})
}

func Fuzz_Nosy_Dump_OnRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Dump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.OnRoot(root)
	})
}

func Fuzz_Nosy_StateDB_AddAddressToAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddAddressToAccessList(addr)
	})
}

func Fuzz_Nosy_StateDB_AddBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if s == nil || amount == nil {
			return
		}

		s.AddBalance(addr, amount)
	})
}

func Fuzz_Nosy_StateDB_AddLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var log *types.Log
		fill_err = tp.Fill(&log)
		if fill_err != nil {
			return
		}
		if s == nil || log == nil {
			return
		}

		s.AddLog(log)
	})
}

func Fuzz_Nosy_StateDB_AddPreimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var preimage []byte
		fill_err = tp.Fill(&preimage)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddPreimage(hash, preimage)
	})
}

func Fuzz_Nosy_StateDB_AddRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddRefund(gas)
	})
}

func Fuzz_Nosy_StateDB_AddSlotToAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddSlotToAccessList(addr, slot)
	})
}

func Fuzz_Nosy_StateDB_AddressInAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.AddressInAccessList(addr)
	})
}

func Fuzz_Nosy_StateDB_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var deleteEmptyObjects bool
		fill_err = tp.Fill(&deleteEmptyObjects)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Commit(block, deleteEmptyObjects)
	})
}

func Fuzz_Nosy_StateDB_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Copy()
	})
}

func Fuzz_Nosy_StateDB_CreateAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.CreateAccount(addr)
	})
}

func Fuzz_Nosy_StateDB_Database__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Database()
	})
}

func Fuzz_Nosy_StateDB_Dump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var opts *DumpConfig
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if s == nil || opts == nil {
			return
		}

		s.Dump(opts)
	})
}

// skipping Fuzz_Nosy_StateDB_DumpToCollector__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.DumpCollector

func Fuzz_Nosy_StateDB_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Empty(addr)
	})
}

func Fuzz_Nosy_StateDB_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Error()
	})
}

func Fuzz_Nosy_StateDB_Exist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Exist(addr)
	})
}

func Fuzz_Nosy_StateDB_Finalise__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var deleteEmptyObjects bool
		fill_err = tp.Fill(&deleteEmptyObjects)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Finalise(deleteEmptyObjects)
	})
}

func Fuzz_Nosy_StateDB_GetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBalance(addr)
	})
}

func Fuzz_Nosy_StateDB_GetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetCode(addr)
	})
}

func Fuzz_Nosy_StateDB_GetCodeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetCodeHash(addr)
	})
}

func Fuzz_Nosy_StateDB_GetCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetCodeSize(addr)
	})
}

func Fuzz_Nosy_StateDB_GetCommittedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetCommittedState(addr, hash)
	})
}

func Fuzz_Nosy_StateDB_GetLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var blockNumber uint64
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetLogs(hash, blockNumber, blockHash)
	})
}

func Fuzz_Nosy_StateDB_GetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetNonce(addr)
	})
}

func Fuzz_Nosy_StateDB_GetOrNewStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetOrNewStateObject(addr)
	})
}

func Fuzz_Nosy_StateDB_GetRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetRefund()
	})
}

func Fuzz_Nosy_StateDB_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetState(addr, hash)
	})
}

func Fuzz_Nosy_StateDB_GetStorageRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetStorageRoot(addr)
	})
}

func Fuzz_Nosy_StateDB_GetTransientState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetTransientState(addr, key)
	})
}

func Fuzz_Nosy_StateDB_HasSelfDestructed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.HasSelfDestructed(addr)
	})
}

func Fuzz_Nosy_StateDB_IntermediateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var deleteEmptyObjects bool
		fill_err = tp.Fill(&deleteEmptyObjects)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.IntermediateRoot(deleteEmptyObjects)
	})
}

func Fuzz_Nosy_StateDB_IterativeDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var opts *DumpConfig
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var output *json.Encoder
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if s == nil || opts == nil || output == nil {
			return
		}

		s.IterativeDump(opts, output)
	})
}

func Fuzz_Nosy_StateDB_Logs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Logs()
	})
}

func Fuzz_Nosy_StateDB_Preimages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Preimages()
	})
}

func Fuzz_Nosy_StateDB_Prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var rules params.Rules
		fill_err = tp.Fill(&rules)
		if fill_err != nil {
			return
		}
		var sender common.Address
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var coinbase common.Address
		fill_err = tp.Fill(&coinbase)
		if fill_err != nil {
			return
		}
		var dst *common.Address
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var list types.AccessList
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		if s == nil || dst == nil {
			return
		}

		s.Prepare(rules, sender, coinbase, dst, precompiles, list)
	})
}

func Fuzz_Nosy_StateDB_RawDump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var opts *DumpConfig
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if s == nil || opts == nil {
			return
		}

		s.RawDump(opts)
	})
}

func Fuzz_Nosy_StateDB_RevertToSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var revid int
		fill_err = tp.Fill(&revid)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RevertToSnapshot(revid)
	})
}

func Fuzz_Nosy_StateDB_SelfDestruct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SelfDestruct(addr)
	})
}

func Fuzz_Nosy_StateDB_Selfdestruct6780__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Selfdestruct6780(addr)
	})
}

func Fuzz_Nosy_StateDB_SetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if s == nil || amount == nil {
			return
		}

		s.SetBalance(addr, amount)
	})
}

func Fuzz_Nosy_StateDB_SetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var code []byte
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetCode(addr, code)
	})
}

func Fuzz_Nosy_StateDB_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetNonce(addr, nonce)
	})
}

func Fuzz_Nosy_StateDB_SetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetState(addr, key, value)
	})
}

func Fuzz_Nosy_StateDB_SetStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var storage map[common.Hash]common.Hash
		fill_err = tp.Fill(&storage)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetStorage(addr, storage)
	})
}

func Fuzz_Nosy_StateDB_SetTransientState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetTransientState(addr, key, value)
	})
}

func Fuzz_Nosy_StateDB_SetTxContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var thash common.Hash
		fill_err = tp.Fill(&thash)
		if fill_err != nil {
			return
		}
		var ti int
		fill_err = tp.Fill(&ti)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetTxContext(thash, ti)
	})
}

func Fuzz_Nosy_StateDB_SlotInAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SlotInAccessList(addr, slot)
	})
}

func Fuzz_Nosy_StateDB_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Snapshot()
	})
}

func Fuzz_Nosy_StateDB_StartPrefetcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartPrefetcher(namespace)
	})
}

func Fuzz_Nosy_StateDB_StopPrefetcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StopPrefetcher()
	})
}

func Fuzz_Nosy_StateDB_SubBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if s == nil || amount == nil {
			return
		}

		s.SubBalance(addr, amount)
	})
}

func Fuzz_Nosy_StateDB_SubRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SubRefund(gas)
	})
}

func Fuzz_Nosy_StateDB_TxIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.TxIndex()
	})
}

func Fuzz_Nosy_StateDB_clearJournalAndRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.clearJournalAndRefund()
	})
}

func Fuzz_Nosy_StateDB_convertAccountSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var set map[common.Address]*types.StateAccount
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.convertAccountSet(set)
	})
}

func Fuzz_Nosy_StateDB_createObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.createObject(addr)
	})
}

func Fuzz_Nosy_StateDB_deleteStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var obj *stateObject
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		if s == nil || obj == nil {
			return
		}

		s.deleteStateObject(obj)
	})
}

func Fuzz_Nosy_StateDB_deleteStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
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
		if s == nil {
			return
		}

		s.deleteStorage(addr, addrHash, root)
	})
}

func Fuzz_Nosy_StateDB_fastDeleteStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.fastDeleteStorage(addrHash, root)
	})
}

func Fuzz_Nosy_StateDB_getDeletedStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getDeletedStateObject(addr)
	})
}

func Fuzz_Nosy_StateDB_getStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getStateObject(addr)
	})
}

func Fuzz_Nosy_StateDB_handleDestruction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var nodes *trienode.MergedNodeSet
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		if s == nil || nodes == nil {
			return
		}

		s.handleDestruction(nodes)
	})
}

// skipping Fuzz_Nosy_StateDB_setError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StateDB_setStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var object *stateObject
		fill_err = tp.Fill(&object)
		if fill_err != nil {
			return
		}
		if s == nil || object == nil {
			return
		}

		s.setStateObject(object)
	})
}

func Fuzz_Nosy_StateDB_setTransientState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.setTransientState(addr, key, value)
	})
}

func Fuzz_Nosy_StateDB_slowDeleteStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
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
		if s == nil {
			return
		}

		s.slowDeleteStorage(addr, addrHash, root)
	})
}

func Fuzz_Nosy_StateDB_updateStateObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var obj *stateObject
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		if s == nil || obj == nil {
			return
		}

		s.updateStateObject(obj)
	})
}

func Fuzz_Nosy_accessList_AddAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.AddAddress(address)
	})
}

func Fuzz_Nosy_accessList_AddSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.AddSlot(address, slot)
	})
}

func Fuzz_Nosy_accessList_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.Contains(address, slot)
	})
}

func Fuzz_Nosy_accessList_ContainsAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.ContainsAddress(address)
	})
}

func Fuzz_Nosy_accessList_DeleteAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.DeleteAddress(address)
	})
}

func Fuzz_Nosy_accessList_DeleteSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.DeleteSlot(address, slot)
	})
}

func Fuzz_Nosy_cachingDB_ContractCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.ContractCode(address, codeHash)
	})
}

func Fuzz_Nosy_cachingDB_ContractCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
		fill_err = tp.Fill(&db)
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
		if db == nil {
			return
		}

		db.ContractCodeSize(addr, codeHash)
	})
}

func Fuzz_Nosy_cachingDB_ContractCodeWithPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var codeHash common.Hash
		fill_err = tp.Fill(&codeHash)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.ContractCodeWithPrefix(address, codeHash)
	})
}

// skipping Fuzz_Nosy_cachingDB_CopyTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

func Fuzz_Nosy_cachingDB_DiskDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.DiskDB()
	})
}

// skipping Fuzz_Nosy_cachingDB_OpenStorageTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

func Fuzz_Nosy_cachingDB_OpenTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
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

		db.OpenTrie(root)
	})
}

func Fuzz_Nosy_cachingDB_TrieDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *cachingDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.TrieDB()
	})
}

// skipping Fuzz_Nosy_journal_append__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.journalEntry

func Fuzz_Nosy_journal_dirty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		j := newJournal()
		j.dirty(addr)
	})
}

func Fuzz_Nosy_journal_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statedb *StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var snapshot int
		fill_err = tp.Fill(&snapshot)
		if fill_err != nil {
			return
		}
		if statedb == nil {
			return
		}

		j := newJournal()
		j.revert(statedb, snapshot)
	})
}

func Fuzz_Nosy_nodeIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 *StateDB
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		if s1 == nil {
			return
		}

		it := newNodeIterator(s1)
		it.Next()
	})
}

func Fuzz_Nosy_nodeIterator_retrieve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 *StateDB
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		if s1 == nil {
			return
		}

		it := newNodeIterator(s1)
		it.retrieve()
	})
}

func Fuzz_Nosy_nodeIterator_step__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 *StateDB
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		if s1 == nil {
			return
		}

		it := newNodeIterator(s1)
		it.step()
	})
}

func Fuzz_Nosy_stateObject_AddBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil || amount == nil {
			return
		}

		s := newObject(db, address, acct)
		s.AddBalance(amount)
	})
}

func Fuzz_Nosy_stateObject_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.Address()
	})
}

func Fuzz_Nosy_stateObject_Balance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.Balance()
	})
}

func Fuzz_Nosy_stateObject_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.Code()
	})
}

func Fuzz_Nosy_stateObject_CodeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.CodeHash()
	})
}

func Fuzz_Nosy_stateObject_CodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.CodeSize()
	})
}

func Fuzz_Nosy_stateObject_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.EncodeRLP(w)
	})
}

func Fuzz_Nosy_stateObject_GetCommittedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.GetCommittedState(key)
	})
}

func Fuzz_Nosy_stateObject_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.GetState(key)
	})
}

func Fuzz_Nosy_stateObject_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.Nonce()
	})
}

func Fuzz_Nosy_stateObject_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.Root()
	})
}

func Fuzz_Nosy_stateObject_SetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil || amount == nil {
			return
		}

		s := newObject(db, address, acct)
		s.SetBalance(amount)
	})
}

func Fuzz_Nosy_stateObject_SetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
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
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.SetCode(codeHash, code)
	})
}

func Fuzz_Nosy_stateObject_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.SetNonce(nonce)
	})
}

func Fuzz_Nosy_stateObject_SetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.SetState(key, value)
	})
}

func Fuzz_Nosy_stateObject_SubBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil || amount == nil {
			return
		}

		s := newObject(db, address, acct)
		s.SubBalance(amount)
	})
}

func Fuzz_Nosy_stateObject_commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.commit()
	})
}

func Fuzz_Nosy_stateObject_deepCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *StateDB
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var d4 *StateDB
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if d1 == nil || acct == nil || d4 == nil {
			return
		}

		s := newObject(d1, address, acct)
		s.deepCopy(d4)
	})
}

func Fuzz_Nosy_stateObject_empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.empty()
	})
}

func Fuzz_Nosy_stateObject_finalise__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var prefetch bool
		fill_err = tp.Fill(&prefetch)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.finalise(prefetch)
	})
}

func Fuzz_Nosy_stateObject_getTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.getTrie()
	})
}

func Fuzz_Nosy_stateObject_markSelfdestructed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.markSelfdestructed()
	})
}

func Fuzz_Nosy_stateObject_setBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil || amount == nil {
			return
		}

		s := newObject(db, address, acct)
		s.setBalance(amount)
	})
}

func Fuzz_Nosy_stateObject_setCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
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
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.setCode(codeHash, code)
	})
}

func Fuzz_Nosy_stateObject_setNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.setNonce(nonce)
	})
}

func Fuzz_Nosy_stateObject_setState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.setState(key, value)
	})
}

func Fuzz_Nosy_stateObject_touch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.touch()
	})
}

func Fuzz_Nosy_stateObject_updateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.updateRoot()
	})
}

func Fuzz_Nosy_stateObject_updateTrie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *StateDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var acct *types.StateAccount
		fill_err = tp.Fill(&acct)
		if fill_err != nil {
			return
		}
		if db == nil || acct == nil {
			return
		}

		s := newObject(db, address, acct)
		s.updateTrie()
	})
}

func Fuzz_Nosy_subfetcher_abort__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sf *subfetcher
		fill_err = tp.Fill(&sf)
		if fill_err != nil {
			return
		}
		if sf == nil {
			return
		}

		sf.abort()
	})
}

func Fuzz_Nosy_subfetcher_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sf *subfetcher
		fill_err = tp.Fill(&sf)
		if fill_err != nil {
			return
		}
		if sf == nil {
			return
		}

		sf.loop()
	})
}

func Fuzz_Nosy_subfetcher_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sf *subfetcher
		fill_err = tp.Fill(&sf)
		if fill_err != nil {
			return
		}
		if sf == nil {
			return
		}

		sf.peek()
	})
}

func Fuzz_Nosy_subfetcher_schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sf *subfetcher
		fill_err = tp.Fill(&sf)
		if fill_err != nil {
			return
		}
		var keys [][]byte
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		if sf == nil {
			return
		}

		sf.schedule(keys)
	})
}

func Fuzz_Nosy_triePrefetcher_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.close()
	})
}

func Fuzz_Nosy_triePrefetcher_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.copy()
	})
}

func Fuzz_Nosy_triePrefetcher_prefetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
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
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var keys [][]byte
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.prefetch(owner, root, addr, keys)
	})
}

func Fuzz_Nosy_triePrefetcher_trie__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
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
		if p == nil {
			return
		}

		p.trie(owner, root)
	})
}

func Fuzz_Nosy_triePrefetcher_trieID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
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
		if p == nil {
			return
		}

		p.trieID(owner, root)
	})
}

func Fuzz_Nosy_triePrefetcher_used__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *triePrefetcher
		fill_err = tp.Fill(&p)
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
		var used [][]byte
		fill_err = tp.Fill(&used)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.used(owner, root, used)
	})
}

func Fuzz_Nosy_Code_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Code
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.String()
	})
}

// skipping Fuzz_Nosy_Database_ContractCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_ContractCodeSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_CopyTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_DiskDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_OpenStorageTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_OpenTrie__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_Database_TrieDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Database

// skipping Fuzz_Nosy_DumpCollector_OnAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.DumpCollector

// skipping Fuzz_Nosy_DumpCollector_OnRoot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.DumpCollector

func Fuzz_Nosy_Storage_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Storage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Copy()
	})
}

func Fuzz_Nosy_Storage_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Storage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}

// skipping Fuzz_Nosy_Trie_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_DeleteAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_DeleteStorage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_GetAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_GetKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_GetStorage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_NodeIterator__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_Prove__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_UpdateAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_UpdateContractCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

// skipping Fuzz_Nosy_Trie_UpdateStorage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.Trie

func Fuzz_Nosy_accessListAddAccountChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch accessListAddAccountChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_accessListAddAccountChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch accessListAddAccountChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_accessListAddSlotChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch accessListAddSlotChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_accessListAddSlotChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch accessListAddSlotChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_addLogChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch addLogChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_addLogChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch addLogChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_addPreimageChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch addPreimageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_addPreimageChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch addPreimageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_balanceChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch balanceChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_balanceChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch balanceChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_codeChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch codeChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_codeChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch codeChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_createObjectChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch createObjectChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_createObjectChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch createObjectChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_iterativeDump_OnAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d iterativeDump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var addr *common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var account DumpAccount
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if addr == nil {
			return
		}

		d.OnAccount(addr, account)
	})
}

func Fuzz_Nosy_iterativeDump_OnRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d iterativeDump
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}

		d.OnRoot(root)
	})
}

// skipping Fuzz_Nosy_journalEntry_dirtied__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.journalEntry

// skipping Fuzz_Nosy_journalEntry_revert__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/state.journalEntry

func Fuzz_Nosy_nonceChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch nonceChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_nonceChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch nonceChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_refundChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch refundChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_refundChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch refundChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_resetObjectChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch resetObjectChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_resetObjectChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch resetObjectChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_selfDestructChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch selfDestructChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_selfDestructChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch selfDestructChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_storageChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch storageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_storageChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch storageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_touchChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch touchChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_touchChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch touchChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_transientStorage_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		t1 := newTransientStorage()
		t1.Get(addr, key)
	})
}

func Fuzz_Nosy_transientStorage_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value common.Hash
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		t1 := newTransientStorage()
		t1.Set(addr, key, value)
	})
}

func Fuzz_Nosy_transientStorageChange_dirtied__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch transientStorageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		ch.dirtied()
	})
}

func Fuzz_Nosy_transientStorageChange_revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch transientStorageChange
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}
		var s *StateDB
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		ch.revert(s)
	})
}

func Fuzz_Nosy_copy2DSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var set map[k]map[common.Hash][]byte
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}

		copy2DSet(set)
	})
}

func Fuzz_Nosy_copySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var set map[k][]byte
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}

		copySet(set)
	})
}
