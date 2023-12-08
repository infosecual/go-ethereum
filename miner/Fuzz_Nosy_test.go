package miner

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_BuildPayloadArgs_Id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *BuildPayloadArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.Id()
	})
}

func Fuzz_Nosy_Miner_BuildPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var args *BuildPayloadArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if m1 == nil || args == nil {
			return
		}

		m1.BuildPayload(args)
	})
}

func Fuzz_Nosy_Miner_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Close()
	})
}

func Fuzz_Nosy_Miner_Hashrate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Hashrate()
	})
}

func Fuzz_Nosy_Miner_Mining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Mining()
	})
}

func Fuzz_Nosy_Miner_Pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Pending()
	})
}

func Fuzz_Nosy_Miner_PendingBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingBlock()
	})
}

func Fuzz_Nosy_Miner_PendingBlockAndReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.PendingBlockAndReceipts()
	})
}

func Fuzz_Nosy_Miner_SetEtherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SetEtherbase(addr)
	})
}

func Fuzz_Nosy_Miner_SetExtra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var extra []byte
		fill_err = tp.Fill(&extra)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SetExtra(extra)
	})
}

func Fuzz_Nosy_Miner_SetGasCeil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var ceil uint64
		fill_err = tp.Fill(&ceil)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SetGasCeil(ceil)
	})
}

func Fuzz_Nosy_Miner_SetRecommitInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.SetRecommitInterval(interval)
	})
}

func Fuzz_Nosy_Miner_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Start()
	})
}

func Fuzz_Nosy_Miner_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.Stop()
	})
}

// skipping Fuzz_Nosy_Miner_SubscribePendingLogs__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

func Fuzz_Nosy_Miner_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m1 *Miner
		fill_err = tp.Fill(&m1)
		if fill_err != nil {
			return
		}
		if m1 == nil {
			return
		}

		m1.update()
	})
}

func Fuzz_Nosy_Payload_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var empty *types.Block
		fill_err = tp.Fill(&empty)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if empty == nil {
			return
		}

		payload := newPayload(empty, id)
		payload.Resolve()
	})
}

func Fuzz_Nosy_Payload_ResolveEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var empty *types.Block
		fill_err = tp.Fill(&empty)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if empty == nil {
			return
		}

		payload := newPayload(empty, id)
		payload.ResolveEmpty()
	})
}

func Fuzz_Nosy_Payload_ResolveFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var empty *types.Block
		fill_err = tp.Fill(&empty)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if empty == nil {
			return
		}

		payload := newPayload(empty, id)
		payload.ResolveFull()
	})
}

func Fuzz_Nosy_Payload_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var empty *types.Block
		fill_err = tp.Fill(&empty)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var r *newPayloadResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		if empty == nil || r == nil {
			return
		}

		payload := newPayload(empty, id)
		payload.update(r, elapsed)
	})
}

func Fuzz_Nosy_environment_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		if env == nil {
			return
		}

		env.copy()
	})
}

func Fuzz_Nosy_environment_discard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		if env == nil {
			return
		}

		env.discard()
	})
}

func Fuzz_Nosy_transactionsByPriceAndNonce_Peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *transactionsByPriceAndNonce
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Peek()
	})
}

func Fuzz_Nosy_transactionsByPriceAndNonce_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *transactionsByPriceAndNonce
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Pop()
	})
}

func Fuzz_Nosy_transactionsByPriceAndNonce_Shift__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *transactionsByPriceAndNonce
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Shift()
	})
}

func Fuzz_Nosy_txByPriceAndTime_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *txByPriceAndTime
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Pop()
	})
}

// skipping Fuzz_Nosy_txByPriceAndTime_Push__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_worker_applyTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if w == nil || env == nil || tx == nil {
			return
		}

		w.applyTransaction(env, tx)
	})
}

func Fuzz_Nosy_worker_buildPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var args *BuildPayloadArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if w == nil || args == nil {
			return
		}

		w.buildPayload(args)
	})
}

func Fuzz_Nosy_worker_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.close()
	})
}

// skipping Fuzz_Nosy_worker_commit__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_worker_commitBlobTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if w == nil || env == nil || tx == nil {
			return
		}

		w.commitBlobTransaction(env, tx)
	})
}

func Fuzz_Nosy_worker_commitTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if w == nil || env == nil || tx == nil {
			return
		}

		w.commitTransaction(env, tx)
	})
}

func Fuzz_Nosy_worker_commitTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var txs *transactionsByPriceAndNonce
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var interrupt *atomic.Int32
		fill_err = tp.Fill(&interrupt)
		if fill_err != nil {
			return
		}
		if w == nil || env == nil || txs == nil || interrupt == nil {
			return
		}

		w.commitTransactions(env, txs, interrupt)
	})
}

func Fuzz_Nosy_worker_commitWork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var interrupt *atomic.Int32
		fill_err = tp.Fill(&interrupt)
		if fill_err != nil {
			return
		}
		var timestamp int64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if w == nil || interrupt == nil {
			return
		}

		w.commitWork(interrupt, timestamp)
	})
}

func Fuzz_Nosy_worker_etherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.etherbase()
	})
}

func Fuzz_Nosy_worker_fillTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var interrupt *atomic.Int32
		fill_err = tp.Fill(&interrupt)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		if w == nil || interrupt == nil || env == nil {
			return
		}

		w.fillTransactions(interrupt, env)
	})
}

func Fuzz_Nosy_worker_generateWork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var params *generateParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if w == nil || params == nil {
			return
		}

		w.generateWork(params)
	})
}

func Fuzz_Nosy_worker_getSealingBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var params *generateParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if w == nil || params == nil {
			return
		}

		w.getSealingBlock(params)
	})
}

func Fuzz_Nosy_worker_isRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.isRunning()
	})
}

func Fuzz_Nosy_worker_isTTDReached__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if w == nil || header == nil {
			return
		}

		w.isTTDReached(header)
	})
}

func Fuzz_Nosy_worker_mainLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.mainLoop()
	})
}

func Fuzz_Nosy_worker_makeEnv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var parent *types.Header
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var coinbase common.Address
		fill_err = tp.Fill(&coinbase)
		if fill_err != nil {
			return
		}
		if w == nil || parent == nil || header == nil {
			return
		}

		w.makeEnv(parent, header, coinbase)
	})
}

func Fuzz_Nosy_worker_newWorkLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var recommit time.Duration
		fill_err = tp.Fill(&recommit)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.newWorkLoop(recommit)
	})
}

func Fuzz_Nosy_worker_pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.pending()
	})
}

func Fuzz_Nosy_worker_pendingBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.pendingBlock()
	})
}

func Fuzz_Nosy_worker_pendingBlockAndReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.pendingBlockAndReceipts()
	})
}

func Fuzz_Nosy_worker_prepareWork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var genParams *generateParams
		fill_err = tp.Fill(&genParams)
		if fill_err != nil {
			return
		}
		if w == nil || genParams == nil {
			return
		}

		w.prepareWork(genParams)
	})
}

func Fuzz_Nosy_worker_resultLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.resultLoop()
	})
}

func Fuzz_Nosy_worker_setEtherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.setEtherbase(addr)
	})
}

func Fuzz_Nosy_worker_setExtra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var extra []byte
		fill_err = tp.Fill(&extra)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.setExtra(extra)
	})
}

func Fuzz_Nosy_worker_setGasCeil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var ceil uint64
		fill_err = tp.Fill(&ceil)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.setGasCeil(ceil)
	})
}

func Fuzz_Nosy_worker_setRecommitInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.setRecommitInterval(interval)
	})
}

func Fuzz_Nosy_worker_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.start()
	})
}

func Fuzz_Nosy_worker_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.stop()
	})
}

func Fuzz_Nosy_worker_taskLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.taskLoop()
	})
}

func Fuzz_Nosy_worker_updateSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *worker
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var env *environment
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		if w == nil || env == nil {
			return
		}

		w.updateSnapshot(env)
	})
}

// skipping Fuzz_Nosy_Backend_BlockChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/miner.Backend

// skipping Fuzz_Nosy_Backend_TxPool__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/miner.Backend

func Fuzz_Nosy_txByPriceAndTime_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s txByPriceAndTime
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_txByPriceAndTime_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s txByPriceAndTime
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

		s.Less(i, j)
	})
}

func Fuzz_Nosy_txByPriceAndTime_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s txByPriceAndTime
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

		s.Swap(i, j)
	})
}

func Fuzz_Nosy_copyReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}

		copyReceipts(receipts)
	})
}
