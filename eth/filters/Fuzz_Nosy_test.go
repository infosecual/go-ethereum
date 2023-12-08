package filters

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
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

// skipping Fuzz_Nosy_EventSystem_SubscribeLogs__ because parameters include func, chan, or unsupported interface: chan []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_EventSystem_SubscribeNewHeads__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/types.Header

// skipping Fuzz_Nosy_EventSystem_SubscribePendingTxs__ because parameters include func, chan, or unsupported interface: chan []*github.com/ethereum/go-ethereum/core/types.Transaction

func Fuzz_Nosy_EventSystem_eventLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.eventLoop()
	})
}

func Fuzz_Nosy_EventSystem_handleChainEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var f3 filterIndex
		fill_err = tp.Fill(&f3)
		if fill_err != nil {
			return
		}
		var ev core.ChainEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.handleChainEvent(f3, ev)
	})
}

func Fuzz_Nosy_EventSystem_handleLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var f3 filterIndex
		fill_err = tp.Fill(&f3)
		if fill_err != nil {
			return
		}
		var ev []*types.Log
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.handleLogs(f3, ev)
	})
}

func Fuzz_Nosy_EventSystem_handlePendingLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var f3 filterIndex
		fill_err = tp.Fill(&f3)
		if fill_err != nil {
			return
		}
		var ev []*types.Log
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.handlePendingLogs(f3, ev)
	})
}

func Fuzz_Nosy_EventSystem_handleTxsEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var f3 filterIndex
		fill_err = tp.Fill(&f3)
		if fill_err != nil {
			return
		}
		var ev core.NewTxsEvent
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.handleTxsEvent(f3, ev)
	})
}

func Fuzz_Nosy_EventSystem_lightFilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		var remove bool
		fill_err = tp.Fill(&remove)
		if fill_err != nil {
			return
		}
		if sys == nil || header == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.lightFilterLogs(header, addresses, topics, remove)
	})
}

// skipping Fuzz_Nosy_EventSystem_lightFilterNewHead__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum/go-ethereum/core/types.Header, bool)

func Fuzz_Nosy_EventSystem_subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var sub *subscription
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
		if sys == nil || sub == nil {
			return
		}

		es := NewEventSystem(sys, lightMode)
		es.subscribe(sub)
	})
}

// skipping Fuzz_Nosy_EventSystem_subscribeLogs__ because parameters include func, chan, or unsupported interface: chan []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_EventSystem_subscribeMinedPendingLogs__ because parameters include func, chan, or unsupported interface: chan []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_EventSystem_subscribePendingLogs__ because parameters include func, chan, or unsupported interface: chan []*github.com/ethereum/go-ethereum/core/types.Log

func Fuzz_Nosy_Filter_Logs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		f1 := newFilter(sys, addresses, topics)
		f1.Logs(ctx)
	})
}

func Fuzz_Nosy_Filter_blockLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if sys == nil || header == nil {
			return
		}

		f1 := newFilter(sys, addresses, topics)
		f1.blockLogs(ctx, header)
	})
}

func Fuzz_Nosy_Filter_checkMatches__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if sys == nil || header == nil {
			return
		}

		f1 := newFilter(sys, addresses, topics)
		f1.checkMatches(ctx, header)
	})
}

// skipping Fuzz_Nosy_Filter_indexedLogs__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/types.Log

func Fuzz_Nosy_Filter_pendingLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		f1 := newFilter(sys, addresses, topics)
		f1.pendingLogs()
	})
}

func Fuzz_Nosy_Filter_rangeLogsAsync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		f1 := newFilter(sys, addresses, topics)
		f1.rangeLogsAsync(ctx)
	})
}

// skipping Fuzz_Nosy_Filter_unindexedLogs__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/types.Log

func Fuzz_Nosy_FilterAPI_GetFilterChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var id rpc.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.GetFilterChanges(id)
	})
}

func Fuzz_Nosy_FilterAPI_GetFilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id rpc.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.GetFilterLogs(ctx, id)
	})
}

func Fuzz_Nosy_FilterAPI_GetLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var crit FilterCriteria
		fill_err = tp.Fill(&crit)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.GetLogs(ctx, crit)
	})
}

func Fuzz_Nosy_FilterAPI_Logs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var crit FilterCriteria
		fill_err = tp.Fill(&crit)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.Logs(ctx, crit)
	})
}

func Fuzz_Nosy_FilterAPI_NewBlockFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.NewBlockFilter()
	})
}

func Fuzz_Nosy_FilterAPI_NewFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var crit FilterCriteria
		fill_err = tp.Fill(&crit)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.NewFilter(crit)
	})
}

func Fuzz_Nosy_FilterAPI_NewHeads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.NewHeads(ctx)
	})
}

func Fuzz_Nosy_FilterAPI_NewPendingTransactionFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var fullTx *bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if system == nil || fullTx == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.NewPendingTransactionFilter(fullTx)
	})
}

func Fuzz_Nosy_FilterAPI_NewPendingTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fullTx *bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if system == nil || fullTx == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.NewPendingTransactions(ctx, fullTx)
	})
}

func Fuzz_Nosy_FilterAPI_UninstallFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var id rpc.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.UninstallFilter(id)
	})
}

func Fuzz_Nosy_FilterAPI_timeoutLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var system *FilterSystem
		fill_err = tp.Fill(&system)
		if fill_err != nil {
			return
		}
		var lightMode bool
		fill_err = tp.Fill(&lightMode)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if system == nil {
			return
		}

		api := NewFilterAPI(system, lightMode)
		api.timeoutLoop(timeout)
	})
}

func Fuzz_Nosy_FilterCriteria_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *FilterCriteria
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_FilterSystem_NewBlockFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var block common.Hash
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.NewBlockFilter(block, addresses, topics)
	})
}

func Fuzz_Nosy_FilterSystem_NewRangeFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var begin int64
		fill_err = tp.Fill(&begin)
		if fill_err != nil {
			return
		}
		var end int64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.NewRangeFilter(begin, end, addresses, topics)
	})
}

func Fuzz_Nosy_FilterSystem_cachedGetBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var elem *logCacheElem
		fill_err = tp.Fill(&elem)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if sys == nil || elem == nil {
			return
		}

		sys.cachedGetBody(ctx, elem, hash, number)
	})
}

func Fuzz_Nosy_FilterSystem_cachedLogElem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sys *FilterSystem
		fill_err = tp.Fill(&sys)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if sys == nil {
			return
		}

		sys.cachedLogElem(ctx, blockHash, number)
	})
}

func Fuzz_Nosy_Subscription_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sub *Subscription
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
		if sub == nil {
			return
		}

		sub.Err()
	})
}

func Fuzz_Nosy_Subscription_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sub *Subscription
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
		if sub == nil {
			return
		}

		sub.Unsubscribe()
	})
}

// skipping Fuzz_Nosy_Backend_BloomStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_ChainConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_ChainDb__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_GetBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_GetLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_GetReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_PendingBlockAndReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_ServiceFilter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_SubscribeChainEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_SubscribeLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_SubscribeNewTxsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_SubscribePendingLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

// skipping Fuzz_Nosy_Backend_SubscribeRemovedLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/filters.Backend

func Fuzz_Nosy_Config_withDefaults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.withDefaults()
	})
}

func Fuzz_Nosy_bloomFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bloom types.Bloom
		fill_err = tp.Fill(&bloom)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}

		bloomFilter(bloom, addresses, topics)
	})
}

func Fuzz_Nosy_filterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var logs []*types.Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}
		var fromBlock *big.Int
		fill_err = tp.Fill(&fromBlock)
		if fill_err != nil {
			return
		}
		var toBlock *big.Int
		fill_err = tp.Fill(&toBlock)
		if fill_err != nil {
			return
		}
		var addresses []common.Address
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var topics [][]common.Hash
		fill_err = tp.Fill(&topics)
		if fill_err != nil {
			return
		}
		if fromBlock == nil || toBlock == nil {
			return
		}

		filterLogs(logs, fromBlock, toBlock, addresses, topics)
	})
}

// skipping Fuzz_Nosy_includes__ because parameters include func, chan, or unsupported interface: []T

func Fuzz_Nosy_returnHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

		returnHashes(hashes)
	})
}

func Fuzz_Nosy_returnLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var logs []*types.Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		returnLogs(logs)
	})
}
