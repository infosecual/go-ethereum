package legacypool

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
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
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		config.sanitize()
	})
}

func Fuzz_Nosy_LegacyPool_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var sync bool
		fill_err = tp.Fill(&sync)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Add(txs, local, sync)
	})
}

func Fuzz_Nosy_LegacyPool_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Close()
	})
}

func Fuzz_Nosy_LegacyPool_Content__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Content()
	})
}

func Fuzz_Nosy_LegacyPool_ContentFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.ContentFrom(addr)
	})
}

func Fuzz_Nosy_LegacyPool_Filter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.Filter(tx)
	})
}

func Fuzz_Nosy_LegacyPool_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Get(hash)
	})
}

func Fuzz_Nosy_LegacyPool_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Has(hash)
	})
}

// skipping Fuzz_Nosy_LegacyPool_Init__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.AddressReserver

func Fuzz_Nosy_LegacyPool_Locals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Locals()
	})
}

func Fuzz_Nosy_LegacyPool_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Nonce(addr)
	})
}

func Fuzz_Nosy_LegacyPool_Pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var enforceTips bool
		fill_err = tp.Fill(&enforceTips)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Pending(enforceTips)
	})
}

func Fuzz_Nosy_LegacyPool_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Header
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if pool == nil || oldHead == nil || newHead == nil {
			return
		}

		pool.Reset(oldHead, newHead)
	})
}

func Fuzz_Nosy_LegacyPool_SetGasTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tip *big.Int
		fill_err = tp.Fill(&tip)
		if fill_err != nil {
			return
		}
		if pool == nil || tip == nil {
			return
		}

		pool.SetGasTip(tip)
	})
}

func Fuzz_Nosy_LegacyPool_Stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Stats()
	})
}

func Fuzz_Nosy_LegacyPool_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Status(hash)
	})
}

// skipping Fuzz_Nosy_LegacyPool_SubscribeTransactions__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.NewTxsEvent

func Fuzz_Nosy_LegacyPool_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.add(tx, local)
	})
}

func Fuzz_Nosy_LegacyPool_addLocal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.addLocal(tx)
	})
}

func Fuzz_Nosy_LegacyPool_addLocals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.addLocals(txs)
	})
}

func Fuzz_Nosy_LegacyPool_addRemote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.addRemote(tx)
	})
}

func Fuzz_Nosy_LegacyPool_addRemoteSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.addRemoteSync(tx)
	})
}

func Fuzz_Nosy_LegacyPool_addRemotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.addRemotes(txs)
	})
}

func Fuzz_Nosy_LegacyPool_addRemotesSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.addRemotesSync(txs)
	})
}

func Fuzz_Nosy_LegacyPool_addTxsLocked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.addTxsLocked(txs, local)
	})
}

func Fuzz_Nosy_LegacyPool_demoteUnexecutables__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.demoteUnexecutables()
	})
}

func Fuzz_Nosy_LegacyPool_enqueueTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var addAll bool
		fill_err = tp.Fill(&addAll)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.enqueueTx(hash, tx, local, addAll)
	})
}

func Fuzz_Nosy_LegacyPool_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.get(hash)
	})
}

func Fuzz_Nosy_LegacyPool_isGapped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.isGapped(from, tx)
	})
}

func Fuzz_Nosy_LegacyPool_journalTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.journalTx(from, tx)
	})
}

func Fuzz_Nosy_LegacyPool_local__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.local()
	})
}

func Fuzz_Nosy_LegacyPool_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.loop()
	})
}

func Fuzz_Nosy_LegacyPool_promoteExecutables__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var accounts []common.Address
		fill_err = tp.Fill(&accounts)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.promoteExecutables(accounts)
	})
}

func Fuzz_Nosy_LegacyPool_promoteTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
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
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.promoteTx(addr, hash, tx)
	})
}

func Fuzz_Nosy_LegacyPool_queueTxEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.queueTxEvent(tx)
	})
}

func Fuzz_Nosy_LegacyPool_removeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var outofbound bool
		fill_err = tp.Fill(&outofbound)
		if fill_err != nil {
			return
		}
		var unreserve bool
		fill_err = tp.Fill(&unreserve)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.removeTx(hash, outofbound, unreserve)
	})
}

func Fuzz_Nosy_LegacyPool_requestPromoteExecutables__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var set *accountSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if pool == nil || set == nil {
			return
		}

		pool.requestPromoteExecutables(set)
	})
}

func Fuzz_Nosy_LegacyPool_requestReset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Header
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if pool == nil || oldHead == nil || newHead == nil {
			return
		}

		pool.requestReset(oldHead, newHead)
	})
}

func Fuzz_Nosy_LegacyPool_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Header
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if pool == nil || oldHead == nil || newHead == nil {
			return
		}

		pool.reset(oldHead, newHead)
	})
}

// skipping Fuzz_Nosy_LegacyPool_runReorg__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_LegacyPool_scheduleReorgLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.scheduleReorgLoop()
	})
}

func Fuzz_Nosy_LegacyPool_stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.stats()
	})
}

func Fuzz_Nosy_LegacyPool_truncatePending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.truncatePending()
	})
}

func Fuzz_Nosy_LegacyPool_truncateQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.truncateQueue()
	})
}

func Fuzz_Nosy_LegacyPool_validateTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.validateTx(tx, local)
	})
}

func Fuzz_Nosy_LegacyPool_validateTxBasics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *LegacyPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if pool == nil || tx == nil {
			return
		}

		pool.validateTxBasics(tx, local)
	})
}

func Fuzz_Nosy_accountSet_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.add(addr)
	})
}

func Fuzz_Nosy_accountSet_addTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if as == nil || tx == nil {
			return
		}

		as.addTx(tx)
	})
}

func Fuzz_Nosy_accountSet_contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.contains(addr)
	})
}

func Fuzz_Nosy_accountSet_containsTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if as == nil || tx == nil {
			return
		}

		as.containsTx(tx)
	})
}

func Fuzz_Nosy_accountSet_flatten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.flatten()
	})
}

func Fuzz_Nosy_accountSet_merge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *accountSet
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		var other *accountSet
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if as == nil || other == nil {
			return
		}

		as.merge(other)
	})
}

func Fuzz_Nosy_devNull_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *devNull
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Close()
	})
}

func Fuzz_Nosy_devNull_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *devNull
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Write(p)
	})
}

func Fuzz_Nosy_journal_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		journal := newTxJournal(path)
		journal.close()
	})
}

func Fuzz_Nosy_journal_insert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		journal := newTxJournal(path)
		journal.insert(tx)
	})
}

// skipping Fuzz_Nosy_journal_load__ because parameters include func, chan, or unsupported interface: func([]*github.com/ethereum/go-ethereum/core/types.Transaction) []error

func Fuzz_Nosy_journal_rotate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var all map[common.Address]types.Transactions
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}

		journal := newTxJournal(path)
		journal.rotate(all)
	})
}

func Fuzz_Nosy_list_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var strict bool
		fill_err = tp.Fill(&strict)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var priceBump uint64
		fill_err = tp.Fill(&priceBump)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		l := newList(strict)
		l.Add(tx, priceBump)
	})
}

func Fuzz_Nosy_list_Cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool, threshold int) {
		l := newList(strict)
		l.Cap(threshold)
	})
}

func Fuzz_Nosy_list_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool, nonce uint64) {
		l := newList(strict)
		l.Contains(nonce)
	})
}

func Fuzz_Nosy_list_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool) {
		l := newList(strict)
		l.Empty()
	})
}

func Fuzz_Nosy_list_Filter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var strict bool
		fill_err = tp.Fill(&strict)
		if fill_err != nil {
			return
		}
		var costLimit *big.Int
		fill_err = tp.Fill(&costLimit)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if costLimit == nil {
			return
		}

		l := newList(strict)
		l.Filter(costLimit, gasLimit)
	})
}

func Fuzz_Nosy_list_Flatten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool) {
		l := newList(strict)
		l.Flatten()
	})
}

func Fuzz_Nosy_list_Forward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool, threshold uint64) {
		l := newList(strict)
		l.Forward(threshold)
	})
}

func Fuzz_Nosy_list_LastElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool) {
		l := newList(strict)
		l.LastElement()
	})
}

func Fuzz_Nosy_list_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool) {
		l := newList(strict)
		l.Len()
	})
}

func Fuzz_Nosy_list_Ready__(f *testing.F) {
	f.Fuzz(func(t *testing.T, strict bool, start uint64) {
		l := newList(strict)
		l.Ready(start)
	})
}

func Fuzz_Nosy_list_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var strict bool
		fill_err = tp.Fill(&strict)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		l := newList(strict)
		l.Remove(tx)
	})
}

func Fuzz_Nosy_list_subTotalCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var strict bool
		fill_err = tp.Fill(&strict)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		l := newList(strict)
		l.subTotalCost(txs)
	})
}

func Fuzz_Nosy_lookup_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		t1 := newLookup()
		t1.Add(tx, local)
	})
}

func Fuzz_Nosy_lookup_Get__(f *testing.F) {
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

		t1 := newLookup()
		t1.Get(hash)
	})
}

func Fuzz_Nosy_lookup_GetLocal__(f *testing.F) {
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

		t1 := newLookup()
		t1.GetLocal(hash)
	})
}

func Fuzz_Nosy_lookup_GetRemote__(f *testing.F) {
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

		t1 := newLookup()
		t1.GetRemote(hash)
	})
}

// skipping Fuzz_Nosy_lookup_Range__ because parameters include func, chan, or unsupported interface: func(hash github.com/ethereum/go-ethereum/common.Hash, tx *github.com/ethereum/go-ethereum/core/types.Transaction, local bool) bool

func Fuzz_Nosy_lookup_RemoteToLocals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var locals *accountSet
		fill_err = tp.Fill(&locals)
		if fill_err != nil {
			return
		}
		if locals == nil {
			return
		}

		t1 := newLookup()
		t1.RemoteToLocals(locals)
	})
}

func Fuzz_Nosy_lookup_RemotesBelowTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var threshold *big.Int
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		if threshold == nil {
			return
		}

		t1 := newLookup()
		t1.RemotesBelowTip(threshold)
	})
}

func Fuzz_Nosy_lookup_Remove__(f *testing.F) {
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

		t1 := newLookup()
		t1.Remove(hash)
	})
}

func Fuzz_Nosy_nonceHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *nonceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Pop()
	})
}

// skipping Fuzz_Nosy_nonceHeap_Push__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_noncer_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if statedb == nil {
			return
		}

		txn := newNoncer(statedb)
		txn.get(addr)
	})
}

func Fuzz_Nosy_noncer_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
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
		if statedb == nil {
			return
		}

		txn := newNoncer(statedb)
		txn.set(addr, nonce)
	})
}

func Fuzz_Nosy_noncer_setAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var all map[common.Address]uint64
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		if statedb == nil {
			return
		}

		txn := newNoncer(statedb)
		txn.setAll(all)
	})
}

func Fuzz_Nosy_noncer_setIfLower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
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
		if statedb == nil {
			return
		}

		txn := newNoncer(statedb)
		txn.setIfLower(addr, nonce)
	})
}

func Fuzz_Nosy_priceHeap_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Len()
	})
}

func Fuzz_Nosy_priceHeap_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
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
		if h == nil {
			return
		}

		h.Less(i, j)
	})
}

func Fuzz_Nosy_priceHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Pop()
	})
}

// skipping Fuzz_Nosy_priceHeap_Push__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_priceHeap_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
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
		if h == nil {
			return
		}

		h.Swap(i, j)
	})
}

func Fuzz_Nosy_priceHeap_cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var a *types.Transaction
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *types.Transaction
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if h == nil || a == nil || b == nil {
			return
		}

		h.cmp(a, b)
	})
}

func Fuzz_Nosy_pricedList_Discard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var slots int
		fill_err = tp.Fill(&slots)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if all == nil {
			return
		}

		l := newPricedList(all)
		l.Discard(slots, force)
	})
}

func Fuzz_Nosy_pricedList_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		if all == nil || tx == nil {
			return
		}

		l := newPricedList(all)
		l.Put(tx, local)
	})
}

func Fuzz_Nosy_pricedList_Reheap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		if all == nil {
			return
		}

		l := newPricedList(all)
		l.Reheap()
	})
}

func Fuzz_Nosy_pricedList_Removed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if all == nil {
			return
		}

		l := newPricedList(all)
		l.Removed(count)
	})
}

func Fuzz_Nosy_pricedList_SetBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if all == nil || baseFee == nil {
			return
		}

		l := newPricedList(all)
		l.SetBaseFee(baseFee)
	})
}

func Fuzz_Nosy_pricedList_Underpriced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if all == nil || tx == nil {
			return
		}

		l := newPricedList(all)
		l.Underpriced(tx)
	})
}

func Fuzz_Nosy_pricedList_underpricedFor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var all *lookup
		fill_err = tp.Fill(&all)
		if fill_err != nil {
			return
		}
		var h *priceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if all == nil || h == nil || tx == nil {
			return
		}

		l := newPricedList(all)
		l.underpricedFor(h, tx)
	})
}

func Fuzz_Nosy_sortedMap_Cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, threshold int) {
		m := newSortedMap()
		m.Cap(threshold)
	})
}

// skipping Fuzz_Nosy_sortedMap_Filter__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum/go-ethereum/core/types.Transaction) bool

func Fuzz_Nosy_sortedMap_Forward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, threshold uint64) {
		m := newSortedMap()
		m.Forward(threshold)
	})
}

func Fuzz_Nosy_sortedMap_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nonce uint64) {
		m := newSortedMap()
		m.Get(nonce)
	})
}

func Fuzz_Nosy_sortedMap_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		m := newSortedMap()
		m.Put(tx)
	})
}

func Fuzz_Nosy_sortedMap_Ready__(f *testing.F) {
	f.Fuzz(func(t *testing.T, start uint64) {
		m := newSortedMap()
		m.Ready(start)
	})
}

func Fuzz_Nosy_sortedMap_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nonce uint64) {
		m := newSortedMap()
		m.Remove(nonce)
	})
}

// skipping Fuzz_Nosy_sortedMap_filter__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum/go-ethereum/core/types.Transaction) bool

// skipping Fuzz_Nosy_BlockChain_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/legacypool.BlockChain

// skipping Fuzz_Nosy_BlockChain_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/legacypool.BlockChain

// skipping Fuzz_Nosy_BlockChain_GetBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/legacypool.BlockChain

// skipping Fuzz_Nosy_BlockChain_StateAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/legacypool.BlockChain

func Fuzz_Nosy_addressesByHeartbeat_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a addressesByHeartbeat
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.Len()
	})
}

func Fuzz_Nosy_addressesByHeartbeat_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a addressesByHeartbeat
		fill_err = tp.Fill(&a)
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

		a.Less(i, j)
	})
}

func Fuzz_Nosy_addressesByHeartbeat_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a addressesByHeartbeat
		fill_err = tp.Fill(&a)
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

		a.Swap(i, j)
	})
}

func Fuzz_Nosy_nonceHeap_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h nonceHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Len()
	})
}

func Fuzz_Nosy_nonceHeap_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h nonceHeap
		fill_err = tp.Fill(&h)
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

		h.Less(i, j)
	})
}

func Fuzz_Nosy_nonceHeap_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h nonceHeap
		fill_err = tp.Fill(&h)
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

		h.Swap(i, j)
	})
}

func Fuzz_Nosy_numSlots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		numSlots(tx)
	})
}
