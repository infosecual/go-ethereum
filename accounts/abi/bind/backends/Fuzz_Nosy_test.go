package backends

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/bloombits"
	"github.com/ethereum/go-ethereum/core/state"
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

func Fuzz_Nosy_SimulatedBackend_AdjustTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var adjustment time.Duration
		fill_err = tp.Fill(&adjustment)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.AdjustTime(adjustment)
	})
}

func Fuzz_Nosy_SimulatedBackend_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.BalanceAt(ctx, contract, blockNumber)
	})
}

func Fuzz_Nosy_SimulatedBackend_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_SimulatedBackend_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.BlockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_SimulatedBackend_Blockchain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.Blockchain()
	})
}

func Fuzz_Nosy_SimulatedBackend_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.CallContract(ctx, call, blockNumber)
	})
}

func Fuzz_Nosy_SimulatedBackend_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.CallContractAtHash(ctx, call, blockHash)
	})
}

func Fuzz_Nosy_SimulatedBackend_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.Close()
	})
}

func Fuzz_Nosy_SimulatedBackend_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.CodeAt(ctx, contract, blockNumber)
	})
}

func Fuzz_Nosy_SimulatedBackend_CodeAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.CodeAtHash(ctx, contract, blockHash)
	})
}

func Fuzz_Nosy_SimulatedBackend_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.Commit()
	})
}

func Fuzz_Nosy_SimulatedBackend_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.EstimateGas(ctx, call)
	})
}

func Fuzz_Nosy_SimulatedBackend_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query ethereum.FilterQuery
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.FilterLogs(ctx, query)
	})
}

func Fuzz_Nosy_SimulatedBackend_Fork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var parent common.Hash
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.Fork(ctx, parent)
	})
}

func Fuzz_Nosy_SimulatedBackend_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_SimulatedBackend_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *big.Int
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.HeaderByNumber(ctx, block)
	})
}

func Fuzz_Nosy_SimulatedBackend_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.NonceAt(ctx, contract, blockNumber)
	})
}

func Fuzz_Nosy_SimulatedBackend_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.PendingCallContract(ctx, call)
	})
}

func Fuzz_Nosy_SimulatedBackend_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.PendingCodeAt(ctx, contract)
	})
}

func Fuzz_Nosy_SimulatedBackend_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.PendingNonceAt(ctx, account)
	})
}

func Fuzz_Nosy_SimulatedBackend_Rollback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.Rollback()
	})
}

func Fuzz_Nosy_SimulatedBackend_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.SendTransaction(ctx, tx)
	})
}

func Fuzz_Nosy_SimulatedBackend_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var contract common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.StorageAt(ctx, contract, key, blockNumber)
	})
}

// skipping Fuzz_Nosy_SimulatedBackend_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_SimulatedBackend_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_SimulatedBackend_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.SuggestGasPrice(ctx)
	})
}

func Fuzz_Nosy_SimulatedBackend_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.SuggestGasTipCap(ctx)
	})
}

func Fuzz_Nosy_SimulatedBackend_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.TransactionByHash(ctx, txHash)
	})
}

func Fuzz_Nosy_SimulatedBackend_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.TransactionCount(ctx, blockHash)
	})
}

func Fuzz_Nosy_SimulatedBackend_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.TransactionInBlock(ctx, blockHash, index)
	})
}

func Fuzz_Nosy_SimulatedBackend_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.TransactionReceipt(ctx, txHash)
	})
}

func Fuzz_Nosy_SimulatedBackend_blockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
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

		b := NewSimulatedBackend(alloc, gasLimit)
		b.blockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_SimulatedBackend_blockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.blockByNumber(ctx, number)
	})
}

func Fuzz_Nosy_SimulatedBackend_callContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var stateDB *state.StateDB
		fill_err = tp.Fill(&stateDB)
		if fill_err != nil {
			return
		}
		if header == nil || stateDB == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.callContract(ctx, call, header, stateDB)
	})
}

func Fuzz_Nosy_SimulatedBackend_callContractAtHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var call ethereum.CallMsg
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.callContractAtHead(ctx, call)
	})
}

func Fuzz_Nosy_SimulatedBackend_headerByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.headerByHash(hash)
	})
}

func Fuzz_Nosy_SimulatedBackend_rollback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var parent *types.Block
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if parent == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.rollback(parent)
	})
}

func Fuzz_Nosy_SimulatedBackend_stateByBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var alloc core.GenesisAlloc
		fill_err = tp.Fill(&alloc)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if blockNumber == nil {
			return
		}

		b := NewSimulatedBackend(alloc, gasLimit)
		b.stateByBlockNumber(ctx, blockNumber)
	})
}

func Fuzz_Nosy_filterBackend_BloomStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.BloomStatus()
	})
}

func Fuzz_Nosy_filterBackend_ChainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.ChainConfig()
	})
}

func Fuzz_Nosy_filterBackend_ChainDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.ChainDb()
	})
}

func Fuzz_Nosy_filterBackend_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.CurrentHeader()
	})
}

func Fuzz_Nosy_filterBackend_EventMux__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.EventMux()
	})
}

func Fuzz_Nosy_filterBackend_GetBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
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
		if fb == nil {
			return
		}

		fb.GetBody(ctx, hash, number)
	})
}

func Fuzz_Nosy_filterBackend_GetLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
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
		if fb == nil {
			return
		}

		fb.GetLogs(ctx, hash, number)
	})
}

func Fuzz_Nosy_filterBackend_GetReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
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
		if fb == nil {
			return
		}

		fb.GetReceipts(ctx, hash)
	})
}

func Fuzz_Nosy_filterBackend_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
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
		if fb == nil {
			return
		}

		fb.HeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_filterBackend_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
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
		if fb == nil {
			return
		}

		fb.HeaderByNumber(ctx, number)
	})
}

func Fuzz_Nosy_filterBackend_PendingBlockAndReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		if fb == nil {
			return
		}

		fb.PendingBlockAndReceipts()
	})
}

func Fuzz_Nosy_filterBackend_ServiceFilter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fb *filterBackend
		fill_err = tp.Fill(&fb)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ms *bloombits.MatcherSession
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		if fb == nil || ms == nil {
			return
		}

		fb.ServiceFilter(ctx, ms)
	})
}

// skipping Fuzz_Nosy_filterBackend_SubscribeChainEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainEvent

// skipping Fuzz_Nosy_filterBackend_SubscribeLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_filterBackend_SubscribeNewTxsEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.NewTxsEvent

// skipping Fuzz_Nosy_filterBackend_SubscribePendingLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_filterBackend_SubscribeRemovedLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.RemovedLogsEvent

func Fuzz_Nosy_revertError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var result *core.ExecutionResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if result == nil {
			return
		}

		e := newRevertError(result)
		e.ErrorCode()
	})
}

func Fuzz_Nosy_revertError_ErrorData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var result *core.ExecutionResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if result == nil {
			return
		}

		e := newRevertError(result)
		e.ErrorData()
	})
}
