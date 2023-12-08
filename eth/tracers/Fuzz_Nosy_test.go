package tracers

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/params"
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

func Fuzz_Nosy_API_IntermediateRoots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.IntermediateRoots(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_StandardTraceBadBlockToFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *StdTraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.StandardTraceBadBlockToFile(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_StandardTraceBlockToFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *StdTraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.StandardTraceBlockToFile(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_TraceBadBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceBadBlock(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_TraceBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blob hexutil.Bytes
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceBlock(ctx, blob, config)
	})
}

func Fuzz_Nosy_API_TraceBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceBlockByHash(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_TraceBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceBlockByNumber(ctx, number, config)
	})
}

func Fuzz_Nosy_API_TraceBlockFromFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceBlockFromFile(ctx, file, config)
	})
}

func Fuzz_Nosy_API_TraceCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args ethapi.TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		var config *TraceCallConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceCall(ctx, args, blockNrOrHash, config)
	})
}

func Fuzz_Nosy_API_TraceChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var start rpc.BlockNumber
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end rpc.BlockNumber
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceChain(ctx, start, end, config)
	})
}

func Fuzz_Nosy_API_TraceTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || config == nil {
			return
		}

		api.TraceTransaction(ctx, hash, config)
	})
}

func Fuzz_Nosy_API_blockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.blockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_API_blockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.blockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_API_blockByNumberAndHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.blockByNumberAndHash(ctx, number, hash)
	})
}

func Fuzz_Nosy_API_chainContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.chainContext(ctx)
	})
}

func Fuzz_Nosy_API_standardTraceBlockToFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var config *StdTraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || block == nil || config == nil {
			return
		}

		api.standardTraceBlockToFile(ctx, block, config)
	})
}

func Fuzz_Nosy_API_traceBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || block == nil || config == nil {
			return
		}

		api.traceBlock(ctx, block, config)
	})
}

func Fuzz_Nosy_API_traceBlockParallel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || block == nil || statedb == nil || config == nil {
			return
		}

		api.traceBlockParallel(ctx, block, statedb, config)
	})
}

// skipping Fuzz_Nosy_API_traceChain__ because parameters include func, chan, or unsupported interface: <-chan interface{}

func Fuzz_Nosy_API_traceTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var message *core.Message
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		var txctx *Context
		fill_err = tp.Fill(&txctx)
		if fill_err != nil {
			return
		}
		var vmctx vm.BlockContext
		fill_err = tp.Fill(&vmctx)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var config *TraceConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if api == nil || message == nil || txctx == nil || statedb == nil || config == nil {
			return
		}

		api.traceTx(ctx, message, txctx, vmctx, statedb, config)
	})
}

func Fuzz_Nosy_directory_IsJS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *directory
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.IsJS(name)
	})
}

func Fuzz_Nosy_directory_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *directory
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ctx *Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg json.RawMessage
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if d == nil || ctx == nil {
			return
		}

		d.New(name, ctx, cfg)
	})
}

// skipping Fuzz_Nosy_directory_Register__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.ctorFn

// skipping Fuzz_Nosy_directory_RegisterJSEval__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.jsCtorFn

func Fuzz_Nosy_stateTracker_callReleases__(f *testing.F) {
	f.Fuzz(func(t *testing.T, limit int, oldest uint64) {
		t1 := newStateTracker(limit, oldest)
		t1.callReleases()
	})
}

// skipping Fuzz_Nosy_stateTracker_releaseState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.StateReleaseFunc

func Fuzz_Nosy_stateTracker_wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, limit int, oldest uint64, number uint64) {
		t1 := newStateTracker(limit, oldest)
		t1.wait(number)
	})
}

// skipping Fuzz_Nosy_Backend_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_ChainConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_ChainDb__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_GetTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_RPCGasCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_StateAtBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Backend_StateAtTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

// skipping Fuzz_Nosy_Tracer_GetResult__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Tracer

// skipping Fuzz_Nosy_Tracer_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Tracer

// skipping Fuzz_Nosy_APIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/tracers.Backend

func Fuzz_Nosy_GetMemoryCopyPadded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *vm.Memory
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var offset int64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var size int64
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		GetMemoryCopyPadded(m, offset, size)
	})
}

func Fuzz_Nosy_containsTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		containsTx(block, hash)
	})
}

func Fuzz_Nosy_overrideConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var original *params.ChainConfig
		fill_err = tp.Fill(&original)
		if fill_err != nil {
			return
		}
		var override *params.ChainConfig
		fill_err = tp.Fill(&override)
		if fill_err != nil {
			return
		}
		if original == nil || override == nil {
			return
		}

		overrideConfig(original, override)
	})
}
