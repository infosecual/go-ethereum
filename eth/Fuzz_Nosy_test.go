package eth

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/protocols/eth"
	"github.com/ethereum/go-ethereum/eth/protocols/snap"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"
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

func Fuzz_Nosy_AdminAPI_ExportChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var first *uint64
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		var last *uint64
		fill_err = tp.Fill(&last)
		if fill_err != nil {
			return
		}
		if e1 == nil || first == nil || last == nil {
			return
		}

		api := NewAdminAPI(e1)
		api.ExportChain(file, first, last)
	})
}

func Fuzz_Nosy_AdminAPI_ImportChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewAdminAPI(e1)
		api.ImportChain(file)
	})
}

func Fuzz_Nosy_DebugAPI_AccountRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		var start hexutil.Bytes
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var maxResults int
		fill_err = tp.Fill(&maxResults)
		if fill_err != nil {
			return
		}
		var nocode bool
		fill_err = tp.Fill(&nocode)
		if fill_err != nil {
			return
		}
		var nostorage bool
		fill_err = tp.Fill(&nostorage)
		if fill_err != nil {
			return
		}
		var incompletes bool
		fill_err = tp.Fill(&incompletes)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.AccountRange(blockNrOrHash, start, maxResults, nocode, nostorage, incompletes)
	})
}

func Fuzz_Nosy_DebugAPI_DumpBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.DumpBlock(blockNr)
	})
}

func Fuzz_Nosy_DebugAPI_GetAccessibleState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var from rpc.BlockNumber
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to rpc.BlockNumber
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.GetAccessibleState(from, to)
	})
}

func Fuzz_Nosy_DebugAPI_GetBadBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.GetBadBlocks(ctx)
	})
}

func Fuzz_Nosy_DebugAPI_GetModifiedAccountsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var startHash common.Hash
		fill_err = tp.Fill(&startHash)
		if fill_err != nil {
			return
		}
		var endHash *common.Hash
		fill_err = tp.Fill(&endHash)
		if fill_err != nil {
			return
		}
		if e1 == nil || endHash == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.GetModifiedAccountsByHash(startHash, endHash)
	})
}

func Fuzz_Nosy_DebugAPI_GetModifiedAccountsByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var startNum uint64
		fill_err = tp.Fill(&startNum)
		if fill_err != nil {
			return
		}
		var endNum *uint64
		fill_err = tp.Fill(&endNum)
		if fill_err != nil {
			return
		}
		if e1 == nil || endNum == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.GetModifiedAccountsByNumber(startNum, endNum)
	})
}

func Fuzz_Nosy_DebugAPI_GetTrieFlushInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.GetTrieFlushInterval()
	})
}

func Fuzz_Nosy_DebugAPI_Preimage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
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
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.Preimage(ctx, hash)
	})
}

func Fuzz_Nosy_DebugAPI_SetTrieFlushInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var interval string
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.SetTrieFlushInterval(interval)
	})
}

func Fuzz_Nosy_DebugAPI_StorageRangeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		var txIndex int
		fill_err = tp.Fill(&txIndex)
		if fill_err != nil {
			return
		}
		var contractAddress common.Address
		fill_err = tp.Fill(&contractAddress)
		if fill_err != nil {
			return
		}
		var keyStart hexutil.Bytes
		fill_err = tp.Fill(&keyStart)
		if fill_err != nil {
			return
		}
		var maxResult int
		fill_err = tp.Fill(&maxResult)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.StorageRangeAt(ctx, blockNrOrHash, txIndex, contractAddress, keyStart, maxResult)
	})
}

func Fuzz_Nosy_DebugAPI_getModifiedAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *Ethereum
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var startBlock *types.Block
		fill_err = tp.Fill(&startBlock)
		if fill_err != nil {
			return
		}
		var endBlock *types.Block
		fill_err = tp.Fill(&endBlock)
		if fill_err != nil {
			return
		}
		if e1 == nil || startBlock == nil || endBlock == nil {
			return
		}

		api := NewDebugAPI(e1)
		api.getModifiedAccounts(startBlock, endBlock)
	})
}

func Fuzz_Nosy_EthAPIBackend_AccountManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.AccountManager()
	})
}

func Fuzz_Nosy_EthAPIBackend_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_EthAPIBackend_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthAPIBackend_BlockByNumberOrHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BlockByNumberOrHash(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_EthAPIBackend_BloomStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BloomStatus()
	})
}

func Fuzz_Nosy_EthAPIBackend_ChainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ChainConfig()
	})
}

func Fuzz_Nosy_EthAPIBackend_ChainDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ChainDb()
	})
}

func Fuzz_Nosy_EthAPIBackend_CurrentBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.CurrentBlock()
	})
}

func Fuzz_Nosy_EthAPIBackend_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.CurrentHeader()
	})
}

func Fuzz_Nosy_EthAPIBackend_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Engine()
	})
}

func Fuzz_Nosy_EthAPIBackend_EventMux__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.EventMux()
	})
}

func Fuzz_Nosy_EthAPIBackend_ExtRPCEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.ExtRPCEnabled()
	})
}

func Fuzz_Nosy_EthAPIBackend_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockCount uint64
		fill_err = tp.Fill(&blockCount)
		if fill_err != nil {
			return
		}
		var lastBlock rpc.BlockNumber
		fill_err = tp.Fill(&lastBlock)
		if fill_err != nil {
			return
		}
		var rewardPercentiles []float64
		fill_err = tp.Fill(&rewardPercentiles)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		var number rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBody(ctx, hash, number)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetEVM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *core.Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var state *state.StateDB
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var vmConfig *vm.Config
		fill_err = tp.Fill(&vmConfig)
		if fill_err != nil {
			return
		}
		var blockCtx *vm.BlockContext
		fill_err = tp.Fill(&blockCtx)
		if fill_err != nil {
			return
		}
		if b == nil || msg == nil || state == nil || header == nil || vmConfig == nil || blockCtx == nil {
			return
		}

		b.GetEVM(ctx, msg, state, header, vmConfig, blockCtx)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetLogs(ctx, hash, number)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetPoolNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetPoolNonce(ctx, addr)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetPoolTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetPoolTransaction(hash)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetPoolTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetPoolTransactions()
	})
}

func Fuzz_Nosy_EthAPIBackend_GetReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.GetReceipts(ctx, hash)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetTd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.GetTd(ctx, hash)
	})
}

func Fuzz_Nosy_EthAPIBackend_GetTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetTransaction(ctx, txHash)
	})
}

func Fuzz_Nosy_EthAPIBackend_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_EthAPIBackend_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthAPIBackend_HeaderByNumberOrHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.HeaderByNumberOrHash(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_EthAPIBackend_Miner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Miner()
	})
}

func Fuzz_Nosy_EthAPIBackend_PendingBlockAndReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.PendingBlockAndReceipts()
	})
}

func Fuzz_Nosy_EthAPIBackend_RPCEVMTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RPCEVMTimeout()
	})
}

func Fuzz_Nosy_EthAPIBackend_RPCGasCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RPCGasCap()
	})
}

func Fuzz_Nosy_EthAPIBackend_RPCTxFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RPCTxFeeCap()
	})
}

func Fuzz_Nosy_EthAPIBackend_SendTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signedTx *types.Transaction
		fill_err = tp.Fill(&signedTx)
		if fill_err != nil {
			return
		}
		if b == nil || signedTx == nil {
			return
		}

		b.SendTx(ctx, signedTx)
	})
}

func Fuzz_Nosy_EthAPIBackend_ServiceFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var session *bloombits.MatcherSession
		fill_err = tp.Fill(&session)
		if fill_err != nil {
			return
		}
		if b == nil || session == nil {
			return
		}

		b.ServiceFilter(ctx, session)
	})
}

func Fuzz_Nosy_EthAPIBackend_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetHead(number)
	})
}

func Fuzz_Nosy_EthAPIBackend_StartMining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.StartMining()
	})
}

func Fuzz_Nosy_EthAPIBackend_StateAndHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		if b == nil {
			return
		}

		b.StateAndHeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_EthAPIBackend_StateAndHeaderByNumberOrHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.StateAndHeaderByNumberOrHash(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_EthAPIBackend_StateAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		var reexec uint64
		fill_err = tp.Fill(&reexec)
		if fill_err != nil {
			return
		}
		var base *state.StateDB
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var readOnly bool
		fill_err = tp.Fill(&readOnly)
		if fill_err != nil {
			return
		}
		var preferDisk bool
		fill_err = tp.Fill(&preferDisk)
		if fill_err != nil {
			return
		}
		if b == nil || block == nil || base == nil {
			return
		}

		b.StateAtBlock(ctx, block, reexec, base, readOnly, preferDisk)
	})
}

func Fuzz_Nosy_EthAPIBackend_StateAtTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
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
		var txIndex int
		fill_err = tp.Fill(&txIndex)
		if fill_err != nil {
			return
		}
		var reexec uint64
		fill_err = tp.Fill(&reexec)
		if fill_err != nil {
			return
		}
		if b == nil || block == nil {
			return
		}

		b.StateAtTransaction(ctx, block, txIndex, reexec)
	})
}

func Fuzz_Nosy_EthAPIBackend_Stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Stats()
	})
}

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeChainEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainEvent

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainHeadEvent

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeChainSideEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainSideEvent

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeNewTxsEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.NewTxsEvent

// skipping Fuzz_Nosy_EthAPIBackend_SubscribePendingLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_EthAPIBackend_SubscribeRemovedLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.RemovedLogsEvent

func Fuzz_Nosy_EthAPIBackend_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_EthAPIBackend_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SyncProgress()
	})
}

func Fuzz_Nosy_EthAPIBackend_TxPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TxPool()
	})
}

func Fuzz_Nosy_EthAPIBackend_TxPoolContent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TxPoolContent()
	})
}

func Fuzz_Nosy_EthAPIBackend_TxPoolContentFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TxPoolContentFrom(addr)
	})
}

func Fuzz_Nosy_EthAPIBackend_UnprotectedAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *EthAPIBackend
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnprotectedAllowed()
	})
}

func Fuzz_Nosy_Ethereum_APIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.APIs()
	})
}

func Fuzz_Nosy_Ethereum_AccountManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.AccountManager()
	})
}

func Fuzz_Nosy_Ethereum_ArchiveMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.ArchiveMode()
	})
}

func Fuzz_Nosy_Ethereum_BlockChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.BlockChain()
	})
}

func Fuzz_Nosy_Ethereum_BloomIndexer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.BloomIndexer()
	})
}

func Fuzz_Nosy_Ethereum_ChainDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.ChainDb()
	})
}

func Fuzz_Nosy_Ethereum_Downloader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Downloader()
	})
}

func Fuzz_Nosy_Ethereum_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Engine()
	})
}

func Fuzz_Nosy_Ethereum_Etherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Etherbase()
	})
}

func Fuzz_Nosy_Ethereum_EventMux__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.EventMux()
	})
}

func Fuzz_Nosy_Ethereum_IsListening__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.IsListening()
	})
}

func Fuzz_Nosy_Ethereum_IsMining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.IsMining()
	})
}

func Fuzz_Nosy_Ethereum_Merger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Merger()
	})
}

func Fuzz_Nosy_Ethereum_Miner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Miner()
	})
}

func Fuzz_Nosy_Ethereum_Protocols__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Protocols()
	})
}

func Fuzz_Nosy_Ethereum_ResetWithGenesisBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var gb *types.Block
		fill_err = tp.Fill(&gb)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || gb == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.ResetWithGenesisBlock(gb)
	})
}

func Fuzz_Nosy_Ethereum_SetEtherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var etherbase common.Address
		fill_err = tp.Fill(&etherbase)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.SetEtherbase(etherbase)
	})
}

func Fuzz_Nosy_Ethereum_SetSynced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.SetSynced()
	})
}

func Fuzz_Nosy_Ethereum_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Start()
	})
}

func Fuzz_Nosy_Ethereum_StartMining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.StartMining()
	})
}

func Fuzz_Nosy_Ethereum_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Stop()
	})
}

func Fuzz_Nosy_Ethereum_StopMining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.StopMining()
	})
}

func Fuzz_Nosy_Ethereum_SyncMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.SyncMode()
	})
}

func Fuzz_Nosy_Ethereum_Synced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.Synced()
	})
}

func Fuzz_Nosy_Ethereum_TxPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.TxPool()
	})
}

func Fuzz_Nosy_Ethereum_hashState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
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
		var reexec uint64
		fill_err = tp.Fill(&reexec)
		if fill_err != nil {
			return
		}
		var base *state.StateDB
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var readOnly bool
		fill_err = tp.Fill(&readOnly)
		if fill_err != nil {
			return
		}
		var preferDisk bool
		fill_err = tp.Fill(&preferDisk)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || block == nil || base == nil {
			return
		}

		e1, err := New(stack, config)
		if err != nil {
			return
		}
		e1.hashState(ctx, block, reexec, base, readOnly, preferDisk)
	})
}

func Fuzz_Nosy_Ethereum_isLocalBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || header == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.isLocalBlock(header)
	})
}

func Fuzz_Nosy_Ethereum_pathState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || block == nil {
			return
		}

		e1, err := New(stack, config)
		if err != nil {
			return
		}
		e1.pathState(block)
	})
}

func Fuzz_Nosy_Ethereum_shouldPreserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || header == nil {
			return
		}

		s, err := New(stack, config)
		if err != nil {
			return
		}
		s.shouldPreserve(header)
	})
}

func Fuzz_Nosy_Ethereum_startBloomHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sectionSize uint64
		fill_err = tp.Fill(&sectionSize)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil {
			return
		}

		e1, err := New(stack, config)
		if err != nil {
			return
		}
		e1.startBloomHandlers(sectionSize)
	})
}

func Fuzz_Nosy_Ethereum_stateAtBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
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
		var reexec uint64
		fill_err = tp.Fill(&reexec)
		if fill_err != nil {
			return
		}
		var base *state.StateDB
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var readOnly bool
		fill_err = tp.Fill(&readOnly)
		if fill_err != nil {
			return
		}
		var preferDisk bool
		fill_err = tp.Fill(&preferDisk)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || block == nil || base == nil {
			return
		}

		e1, err := New(stack, config)
		if err != nil {
			return
		}
		e1.stateAtBlock(ctx, block, reexec, base, readOnly, preferDisk)
	})
}

func Fuzz_Nosy_Ethereum_stateAtTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var config *ethconfig.Config
		fill_err = tp.Fill(&config)
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
		var txIndex int
		fill_err = tp.Fill(&txIndex)
		if fill_err != nil {
			return
		}
		var reexec uint64
		fill_err = tp.Fill(&reexec)
		if fill_err != nil {
			return
		}
		if stack == nil || config == nil || block == nil {
			return
		}

		e1, err := New(stack, config)
		if err != nil {
			return
		}
		e1.stateAtTransaction(ctx, block, txIndex, reexec)
	})
}

func Fuzz_Nosy_EthereumAPI_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewEthereumAPI(e)
		api.Coinbase()
	})
}

func Fuzz_Nosy_EthereumAPI_Etherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewEthereumAPI(e)
		api.Etherbase()
	})
}

func Fuzz_Nosy_EthereumAPI_Hashrate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewEthereumAPI(e)
		api.Hashrate()
	})
}

func Fuzz_Nosy_EthereumAPI_Mining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewEthereumAPI(e)
		api.Mining()
	})
}

func Fuzz_Nosy_MinerAPI_SetEtherbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var etherbase common.Address
		fill_err = tp.Fill(&etherbase)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.SetEtherbase(etherbase)
	})
}

func Fuzz_Nosy_MinerAPI_SetExtra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var extra string
		fill_err = tp.Fill(&extra)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.SetExtra(extra)
	})
}

func Fuzz_Nosy_MinerAPI_SetGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var gasLimit hexutil.Uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.SetGasLimit(gasLimit)
	})
}

func Fuzz_Nosy_MinerAPI_SetGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var gasPrice hexutil.Big
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.SetGasPrice(gasPrice)
	})
}

func Fuzz_Nosy_MinerAPI_SetRecommitInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var interval int
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.SetRecommitInterval(interval)
	})
}

func Fuzz_Nosy_MinerAPI_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.Start()
	})
}

func Fuzz_Nosy_MinerAPI_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Ethereum
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		api := NewMinerAPI(e)
		api.Stop()
	})
}

func Fuzz_Nosy_chainSyncer_handlePeerEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var handler *handler
		fill_err = tp.Fill(&handler)
		if fill_err != nil {
			return
		}
		if handler == nil {
			return
		}

		cs := newChainSyncer(handler)
		cs.handlePeerEvent()
	})
}

func Fuzz_Nosy_chainSyncer_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var handler *handler
		fill_err = tp.Fill(&handler)
		if fill_err != nil {
			return
		}
		if handler == nil {
			return
		}

		cs := newChainSyncer(handler)
		cs.loop()
	})
}

func Fuzz_Nosy_chainSyncer_modeAndLocalHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var handler *handler
		fill_err = tp.Fill(&handler)
		if fill_err != nil {
			return
		}
		if handler == nil {
			return
		}

		cs := newChainSyncer(handler)
		cs.modeAndLocalHead()
	})
}

func Fuzz_Nosy_chainSyncer_nextSyncOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var handler *handler
		fill_err = tp.Fill(&handler)
		if fill_err != nil {
			return
		}
		if handler == nil {
			return
		}

		cs := newChainSyncer(handler)
		cs.nextSyncOp()
	})
}

func Fuzz_Nosy_chainSyncer_startSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var handler *handler
		fill_err = tp.Fill(&handler)
		if fill_err != nil {
			return
		}
		var op *chainSyncOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if handler == nil || op == nil {
			return
		}

		cs := newChainSyncer(handler)
		cs.startSync(op)
	})
}

func Fuzz_Nosy_ethHandler_AcceptTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.AcceptTxs()
	})
}

func Fuzz_Nosy_ethHandler_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Chain()
	})
}

// skipping Fuzz_Nosy_ethHandler_Handle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Packet

func Fuzz_Nosy_ethHandler_PeerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PeerInfo(id)
	})
}

// skipping Fuzz_Nosy_ethHandler_RunPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Handler

func Fuzz_Nosy_ethHandler_TxPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.TxPool()
	})
}

func Fuzz_Nosy_ethHandler_handleBlockAnnounces__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var peer *eth.Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var numbers []uint64
		fill_err = tp.Fill(&numbers)
		if fill_err != nil {
			return
		}
		if h == nil || peer == nil {
			return
		}

		h.handleBlockAnnounces(peer, hashes, numbers)
	})
}

func Fuzz_Nosy_ethHandler_handleBlockBroadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ethHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var peer *eth.Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if h == nil || peer == nil || block == nil || td == nil {
			return
		}

		h.handleBlockBroadcast(peer, block, td)
	})
}

func Fuzz_Nosy_ethPeer_info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ethPeer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.info()
	})
}

func Fuzz_Nosy_handler_BroadcastBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var propagate bool
		fill_err = tp.Fill(&propagate)
		if fill_err != nil {
			return
		}
		if config == nil || block == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.BroadcastBlock(block, propagate)
	})
}

func Fuzz_Nosy_handler_BroadcastTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var txs types.Transactions
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.BroadcastTransactions(txs)
	})
}

func Fuzz_Nosy_handler_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var maxPeers int
		fill_err = tp.Fill(&maxPeers)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.Start(maxPeers)
	})
}

func Fuzz_Nosy_handler_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.Stop()
	})
}

func Fuzz_Nosy_handler_decHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.decHandlers()
	})
}

func Fuzz_Nosy_handler_doSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var op *chainSyncOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if config == nil || op == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.doSync(op)
	})
}

func Fuzz_Nosy_handler_enableSyncedFeatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.enableSyncedFeatures()
	})
}

func Fuzz_Nosy_handler_incHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.incHandlers()
	})
}

func Fuzz_Nosy_handler_minedBroadcastLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.minedBroadcastLoop()
	})
}

func Fuzz_Nosy_handler_protoTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.protoTracker()
	})
}

func Fuzz_Nosy_handler_removePeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.removePeer(id)
	})
}

// skipping Fuzz_Nosy_handler_runEthPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Handler

// skipping Fuzz_Nosy_handler_runSnapExtension__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Handler

func Fuzz_Nosy_handler_syncTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var p *eth.Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if config == nil || p == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.syncTransactions(p)
	})
}

func Fuzz_Nosy_handler_txBroadcastLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.txBroadcastLoop()
	})
}

func Fuzz_Nosy_handler_unregisterPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *handlerConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		h, err := newHandler(config)
		if err != nil {
			return
		}
		h.unregisterPeer(id)
	})
}

func Fuzz_Nosy_peerSet_peer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		ps := newPeerSet()
		ps.peer(id)
	})
}

func Fuzz_Nosy_peerSet_peersWithoutBlock__(f *testing.F) {
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

		ps := newPeerSet()
		ps.peersWithoutBlock(hash)
	})
}

func Fuzz_Nosy_peerSet_peersWithoutTransaction__(f *testing.F) {
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

		ps := newPeerSet()
		ps.peersWithoutTransaction(hash)
	})
}

func Fuzz_Nosy_peerSet_registerPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var peer *eth.Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var ext *snap.Peer
		fill_err = tp.Fill(&ext)
		if fill_err != nil {
			return
		}
		if peer == nil || ext == nil {
			return
		}

		ps := newPeerSet()
		ps.registerPeer(peer, ext)
	})
}

func Fuzz_Nosy_peerSet_registerSnapExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var peer *snap.Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if peer == nil {
			return
		}

		ps := newPeerSet()
		ps.registerSnapExtension(peer)
	})
}

func Fuzz_Nosy_peerSet_unregisterPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		ps := newPeerSet()
		ps.unregisterPeer(id)
	})
}

func Fuzz_Nosy_peerSet_waitSnapExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var peer *eth.Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if peer == nil {
			return
		}

		ps := newPeerSet()
		ps.waitSnapExtension(peer)
	})
}

func Fuzz_Nosy_snapHandler_Chain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *snapHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Chain()
	})
}

// skipping Fuzz_Nosy_snapHandler_Handle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Packet

func Fuzz_Nosy_snapHandler_PeerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *snapHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PeerInfo(id)
	})
}

// skipping Fuzz_Nosy_snapHandler_RunPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Handler

func Fuzz_Nosy_snapPeer_info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *snapPeer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.info()
	})
}

// skipping Fuzz_Nosy_threaded_SetThreads__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.threaded

// skipping Fuzz_Nosy_txPool_Add__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.txPool

// skipping Fuzz_Nosy_txPool_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.txPool

// skipping Fuzz_Nosy_txPool_Has__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.txPool

// skipping Fuzz_Nosy_txPool_Pending__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.txPool

// skipping Fuzz_Nosy_txPool_SubscribeTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth.txPool

func Fuzz_Nosy_hasAllBlocks__(f *testing.F) {
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
		var bs []*types.Block
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		hasAllBlocks(chain, bs)
	})
}

func Fuzz_Nosy_makeExtraData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, extra []byte) {
		makeExtraData(extra)
	})
}
