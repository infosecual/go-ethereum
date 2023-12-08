package ethapi

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/p2p"
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

func Fuzz_Nosy_AddrLocker_LockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AddrLocker
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.LockAddr(address)
	})
}

func Fuzz_Nosy_AddrLocker_UnlockAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AddrLocker
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.UnlockAddr(address)
	})
}

func Fuzz_Nosy_AddrLocker_lock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AddrLocker
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.lock(address)
	})
}

func Fuzz_Nosy_BlockChainAPI_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.BlockNumber()
	})
}

func Fuzz_Nosy_BlockChainAPI_Call__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var blockNrOrHash *rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		var overrides *StateOverride
		fill_err = tp.Fill(&overrides)
		if fill_err != nil {
			return
		}
		var blockOverrides *BlockOverrides
		fill_err = tp.Fill(&blockOverrides)
		if fill_err != nil {
			return
		}
		if s == nil || blockNrOrHash == nil || overrides == nil || blockOverrides == nil {
			return
		}

		s.Call(ctx, args, blockNrOrHash, overrides, blockOverrides)
	})
}

func Fuzz_Nosy_BlockChainAPI_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *BlockChainAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.ChainId()
	})
}

func Fuzz_Nosy_BlockChainAPI_CreateAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var blockNrOrHash *rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil || blockNrOrHash == nil {
			return
		}

		s.CreateAccessList(ctx, args, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var blockNrOrHash *rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		var overrides *StateOverride
		fill_err = tp.Fill(&overrides)
		if fill_err != nil {
			return
		}
		if s == nil || blockNrOrHash == nil || overrides == nil {
			return
		}

		s.EstimateGas(ctx, args, blockNrOrHash, overrides)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBalance(ctx, address, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		var fullTx bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlockByHash(ctx, hash, fullTx)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		var fullTx bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlockByNumber(ctx, number, fullTx)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetBlockReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlockReceipts(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetCode(ctx, address, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetHeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetHeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var storageKeys []string
		fill_err = tp.Fill(&storageKeys)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetProof(ctx, address, storageKeys, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var hexKey string
		fill_err = tp.Fill(&hexKey)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetStorageAt(ctx, address, hexKey, blockNrOrHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetUncleByBlockHashAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetUncleByBlockHashAndIndex(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetUncleByBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetUncleByBlockNumberAndIndex(ctx, blockNr, index)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetUncleCountByBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetUncleCountByBlockHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_BlockChainAPI_GetUncleCountByBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetUncleCountByBlockNumber(ctx, blockNr)
	})
}

func Fuzz_Nosy_BlockChainAPI_rpcMarshalBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var b *types.Block
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var inclTx bool
		fill_err = tp.Fill(&inclTx)
		if fill_err != nil {
			return
		}
		var fullTx bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		if s == nil || b == nil {
			return
		}

		s.rpcMarshalBlock(ctx, b, inclTx, fullTx)
	})
}

func Fuzz_Nosy_BlockChainAPI_rpcMarshalHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *BlockChainAPI
		fill_err = tp.Fill(&s)
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
		if s == nil || header == nil {
			return
		}

		s.rpcMarshalHeader(ctx, header)
	})
}

func Fuzz_Nosy_BlockOverrides_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var diff *BlockOverrides
		fill_err = tp.Fill(&diff)
		if fill_err != nil {
			return
		}
		var blockCtx *vm.BlockContext
		fill_err = tp.Fill(&blockCtx)
		if fill_err != nil {
			return
		}
		if diff == nil || blockCtx == nil {
			return
		}

		diff.Apply(blockCtx)
	})
}

func Fuzz_Nosy_ChainContext_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var context *ChainContext
		fill_err = tp.Fill(&context)
		if fill_err != nil {
			return
		}
		if context == nil {
			return
		}

		context.Engine()
	})
}

func Fuzz_Nosy_ChainContext_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var context *ChainContext
		fill_err = tp.Fill(&context)
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
		if context == nil {
			return
		}

		context.GetHeader(hash, number)
	})
}

func Fuzz_Nosy_DebugAPI_ChaindbCompact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.ChaindbCompact()
	})
}

func Fuzz_Nosy_DebugAPI_ChaindbProperty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.ChaindbProperty(property)
	})
}

func Fuzz_Nosy_DebugAPI_DbAncient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var kind string
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.DbAncient(kind, number)
	})
}

func Fuzz_Nosy_DebugAPI_DbAncients__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.DbAncients()
	})
}

func Fuzz_Nosy_DebugAPI_DbGet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.DbGet(key)
	})
}

func Fuzz_Nosy_DebugAPI_GetRawBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
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
		if api == nil {
			return
		}

		api.GetRawBlock(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_DebugAPI_GetRawHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
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
		if api == nil {
			return
		}

		api.GetRawHeader(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_DebugAPI_GetRawReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
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
		if api == nil {
			return
		}

		api.GetRawReceipts(ctx, blockNrOrHash)
	})
}

func Fuzz_Nosy_DebugAPI_GetRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DebugAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetRawTransaction(ctx, hash)
	})
}

func Fuzz_Nosy_DebugAPI_PrintBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.PrintBlock(ctx, number)
	})
}

func Fuzz_Nosy_DebugAPI_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *DebugAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var number hexutil.Uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SetHead(number)
	})
}

func Fuzz_Nosy_EthereumAPI_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EthereumAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockCount math.HexOrDecimal64
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
		if s == nil {
			return
		}

		s.FeeHistory(ctx, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_EthereumAPI_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EthereumAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GasPrice(ctx)
	})
}

func Fuzz_Nosy_EthereumAPI_MaxPriorityFeePerGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EthereumAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MaxPriorityFeePerGas(ctx)
	})
}

func Fuzz_Nosy_EthereumAPI_Syncing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *EthereumAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Syncing()
	})
}

func Fuzz_Nosy_EthereumAccountAPI_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *accounts.Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		s := NewEthereumAccountAPI(am)
		s.Accounts()
	})
}

func Fuzz_Nosy_NetAPI_Listening__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *p2p.Server
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var networkVersion uint64
		fill_err = tp.Fill(&networkVersion)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		s := NewNetAPI(net, networkVersion)
		s.Listening()
	})
}

func Fuzz_Nosy_NetAPI_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *p2p.Server
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var networkVersion uint64
		fill_err = tp.Fill(&networkVersion)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		s := NewNetAPI(net, networkVersion)
		s.PeerCount()
	})
}

func Fuzz_Nosy_NetAPI_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *p2p.Server
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var networkVersion uint64
		fill_err = tp.Fill(&networkVersion)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		s := NewNetAPI(net, networkVersion)
		s.Version()
	})
}

func Fuzz_Nosy_PersonalAccountAPI_DeriveAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var pin *bool
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}
		if s == nil || pin == nil {
			return
		}

		s.DeriveAccount(url, path, pin)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_EcRecover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 hexutil.Bytes
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var sig hexutil.Bytes
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.EcRecover(ctx, d3, sig)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_ImportRawKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var privkey string
		fill_err = tp.Fill(&privkey)
		if fill_err != nil {
			return
		}
		var password string
		fill_err = tp.Fill(&password)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ImportRawKey(privkey, password)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_InitializeWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.InitializeWallet(ctx, url)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_ListAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListAccounts()
	})
}

func Fuzz_Nosy_PersonalAccountAPI_ListWallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ListWallets()
	})
}

func Fuzz_Nosy_PersonalAccountAPI_LockAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
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

		s.LockAccount(addr)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_NewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var password string
		fill_err = tp.Fill(&password)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NewAccount(password)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_OpenWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var passphrase *string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if s == nil || passphrase == nil {
			return
		}

		s.OpenWallet(url, passphrase)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var passwd string
		fill_err = tp.Fill(&passwd)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SendTransaction(ctx, args, passwd)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 hexutil.Bytes
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var passwd string
		fill_err = tp.Fill(&passwd)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Sign(ctx, d3, addr, passwd)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_SignTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var passwd string
		fill_err = tp.Fill(&passwd)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SignTransaction(ctx, args, passwd)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_UnlockAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
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
		var password string
		fill_err = tp.Fill(&password)
		if fill_err != nil {
			return
		}
		var duration *uint64
		fill_err = tp.Fill(&duration)
		if fill_err != nil {
			return
		}
		if s == nil || duration == nil {
			return
		}

		s.UnlockAccount(ctx, addr, password, duration)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_Unpair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var pin string
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Unpair(ctx, url, pin)
	})
}

func Fuzz_Nosy_PersonalAccountAPI_signTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *PersonalAccountAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var passwd string
		fill_err = tp.Fill(&passwd)
		if fill_err != nil {
			return
		}
		if s == nil || args == nil {
			return
		}

		s.signTransaction(ctx, args, passwd)
	})
}

func Fuzz_Nosy_StateOverride_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var diff *StateOverride
		fill_err = tp.Fill(&diff)
		if fill_err != nil {
			return
		}
		var state *state.StateDB
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if diff == nil || state == nil {
			return
		}

		diff.Apply(state)
	})
}

func Fuzz_Nosy_TransactionAPI_FillTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.FillTransaction(ctx, args)
	})
}

func Fuzz_Nosy_TransactionAPI_GetBlockTransactionCountByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetBlockTransactionCountByHash(ctx, blockHash)
	})
}

func Fuzz_Nosy_TransactionAPI_GetBlockTransactionCountByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetBlockTransactionCountByNumber(ctx, blockNr)
	})
}

func Fuzz_Nosy_TransactionAPI_GetRawTransactionByBlockHashAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetRawTransactionByBlockHashAndIndex(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_TransactionAPI_GetRawTransactionByBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetRawTransactionByBlockNumberAndIndex(ctx, blockNr, index)
	})
}

func Fuzz_Nosy_TransactionAPI_GetRawTransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetRawTransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_TransactionAPI_GetTransactionByBlockHashAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetTransactionByBlockHashAndIndex(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_TransactionAPI_GetTransactionByBlockNumberAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNr rpc.BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		var index hexutil.Uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetTransactionByBlockNumberAndIndex(ctx, blockNr, index)
	})
}

func Fuzz_Nosy_TransactionAPI_GetTransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetTransactionByHash(ctx, hash)
	})
}

func Fuzz_Nosy_TransactionAPI_GetTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetTransactionCount(ctx, address, blockNrOrHash)
	})
}

func Fuzz_Nosy_TransactionAPI_GetTransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
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
		if s == nil {
			return
		}

		s.GetTransactionReceipt(ctx, hash)
	})
}

func Fuzz_Nosy_TransactionAPI_PendingTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PendingTransactions()
	})
}

func Fuzz_Nosy_TransactionAPI_Resend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sendArgs TransactionArgs
		fill_err = tp.Fill(&sendArgs)
		if fill_err != nil {
			return
		}
		var gasPrice *hexutil.Big
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var gasLimit *hexutil.Uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if s == nil || gasPrice == nil || gasLimit == nil {
			return
		}

		s.Resend(ctx, sendArgs, gasPrice, gasLimit)
	})
}

func Fuzz_Nosy_TransactionAPI_SendRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var input hexutil.Bytes
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SendRawTransaction(ctx, input)
	})
}

func Fuzz_Nosy_TransactionAPI_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SendTransaction(ctx, args)
	})
}

func Fuzz_Nosy_TransactionAPI_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var d3 hexutil.Bytes
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Sign(addr, d3)
	})
}

func Fuzz_Nosy_TransactionAPI_SignTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SignTransaction(ctx, args)
	})
}

func Fuzz_Nosy_TransactionAPI_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TransactionAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil {
			return
		}

		s.sign(addr, tx)
	})
}

func Fuzz_Nosy_TransactionArgs_ToMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var globalGasCap uint64
		fill_err = tp.Fill(&globalGasCap)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if args == nil || baseFee == nil {
			return
		}

		args.ToMessage(globalGasCap, baseFee)
	})
}

func Fuzz_Nosy_TransactionArgs_ToTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.ToTransaction()
	})
}

func Fuzz_Nosy_TransactionArgs_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.data()
	})
}

func Fuzz_Nosy_TransactionArgs_from__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.from()
	})
}

// skipping Fuzz_Nosy_TransactionArgs_setDefaults__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_TransactionArgs_setFeeDefaults__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_TransactionArgs_setLondonFeeDefaults__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

func Fuzz_Nosy_TransactionArgs_toTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *TransactionArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.toTransaction()
	})
}

func Fuzz_Nosy_TxPoolAPI_Content__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxPoolAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Content()
	})
}

func Fuzz_Nosy_TxPoolAPI_ContentFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxPoolAPI
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

		s.ContentFrom(addr)
	})
}

func Fuzz_Nosy_TxPoolAPI_Inspect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxPoolAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Inspect()
	})
}

func Fuzz_Nosy_TxPoolAPI_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *TxPoolAPI
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Status()
	})
}

func Fuzz_Nosy_proofList_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *proofList
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

func Fuzz_Nosy_proofList_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *proofList
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

func Fuzz_Nosy_revertError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, revert []byte) {
		e := newRevertError(revert)
		e.ErrorCode()
	})
}

func Fuzz_Nosy_revertError_ErrorData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, revert []byte) {
		e := newRevertError(revert)
		e.ErrorData()
	})
}

// skipping Fuzz_Nosy_Backend_AccountManager__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_BlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_BlockByNumberOrHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_BloomStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_ChainConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_ChainDb__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_ExtRPCEnabled__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_FeeHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetEVM__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetPoolNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetPoolTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetPoolTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_GetTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_HeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_HeaderByNumberOrHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_PendingBlockAndReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_RPCEVMTimeout__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_RPCGasCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_RPCTxFeeCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SendTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_ServiceFilter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SetHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_StateAndHeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_StateAndHeaderByNumberOrHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_Stats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeChainEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeChainSideEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeNewTxsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribePendingLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SubscribeRemovedLogsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_SyncProgress__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_TxPoolContent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_TxPoolContentFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_Backend_UnprotectedAllowed__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_ChainContextBackend_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.ChainContextBackend

// skipping Fuzz_Nosy_ChainContextBackend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.ChainContextBackend

// skipping Fuzz_Nosy_AccessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

// skipping Fuzz_Nosy_GetAPIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

func Fuzz_Nosy_RPCMarshalBlock__(f *testing.F) {
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
		var inclTx bool
		fill_err = tp.Fill(&inclTx)
		if fill_err != nil {
			return
		}
		var fullTx bool
		fill_err = tp.Fill(&fullTx)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if block == nil || config == nil {
			return
		}

		RPCMarshalBlock(block, inclTx, fullTx, config)
	})
}

func Fuzz_Nosy_RPCMarshalHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if head == nil {
			return
		}

		RPCMarshalHeader(head)
	})
}

func Fuzz_Nosy_decodeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		decodeHash(s)
	})
}

// skipping Fuzz_Nosy_marshalReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer
