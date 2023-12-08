package ethclient

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_Client_BalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BalanceAt(c3, account, blockNumber)
	})
}

func Fuzz_Nosy_Client_BalanceAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BalanceAtHash(c3, account, blockHash)
	})
}

func Fuzz_Nosy_Client_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BlockByHash(c3, hash)
	})
}

func Fuzz_Nosy_Client_BlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BlockByNumber(c3, number)
	})
}

func Fuzz_Nosy_Client_BlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BlockNumber(c3)
	})
}

func Fuzz_Nosy_Client_BlockReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var blockNrOrHash rpc.BlockNumberOrHash
		fill_err = tp.Fill(&blockNrOrHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.BlockReceipts(c3, blockNrOrHash)
	})
}

func Fuzz_Nosy_Client_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.CallContract(c3, msg, blockNumber)
	})
}

func Fuzz_Nosy_Client_CallContractAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.CallContractAtHash(c3, msg, blockHash)
	})
}

func Fuzz_Nosy_Client_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.ChainID(c3)
	})
}

func Fuzz_Nosy_Client_Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(ctx, rawurl)
		if err != nil {
			return
		}
		ec.Client()
	})
}

func Fuzz_Nosy_Client_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(ctx, rawurl)
		if err != nil {
			return
		}
		ec.Close()
	})
}

func Fuzz_Nosy_Client_CodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.CodeAt(c3, account, blockNumber)
	})
}

func Fuzz_Nosy_Client_CodeAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.CodeAtHash(c3, account, blockHash)
	})
}

func Fuzz_Nosy_Client_EstimateGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.EstimateGas(c3, msg)
	})
}

func Fuzz_Nosy_Client_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var blockCount uint64
		fill_err = tp.Fill(&blockCount)
		if fill_err != nil {
			return
		}
		var lastBlock *big.Int
		fill_err = tp.Fill(&lastBlock)
		if fill_err != nil {
			return
		}
		var rewardPercentiles []float64
		fill_err = tp.Fill(&rewardPercentiles)
		if fill_err != nil {
			return
		}
		if lastBlock == nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.FeeHistory(c3, blockCount, lastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_Client_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.FilterLogs(c3, q)
	})
}

func Fuzz_Nosy_Client_HeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.HeaderByHash(c3, hash)
	})
}

func Fuzz_Nosy_Client_HeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.HeaderByNumber(c3, number)
	})
}

func Fuzz_Nosy_Client_NetworkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.NetworkID(c3)
	})
}

func Fuzz_Nosy_Client_NonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.NonceAt(c3, account, blockNumber)
	})
}

func Fuzz_Nosy_Client_NonceAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.NonceAtHash(c3, account, blockHash)
	})
}

func Fuzz_Nosy_Client_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PeerCount(c3)
	})
}

func Fuzz_Nosy_Client_PendingBalanceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingBalanceAt(c3, account)
	})
}

func Fuzz_Nosy_Client_PendingCallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingCallContract(c3, msg)
	})
}

func Fuzz_Nosy_Client_PendingCodeAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingCodeAt(c3, account)
	})
}

func Fuzz_Nosy_Client_PendingNonceAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingNonceAt(c3, account)
	})
}

func Fuzz_Nosy_Client_PendingStorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingStorageAt(c3, account, key)
	})
}

func Fuzz_Nosy_Client_PendingTransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.PendingTransactionCount(c3)
	})
}

func Fuzz_Nosy_Client_SendTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.SendTransaction(c3, tx)
	})
}

func Fuzz_Nosy_Client_StorageAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.StorageAt(c3, account, key, blockNumber)
	})
}

func Fuzz_Nosy_Client_StorageAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var account common.Address
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.StorageAtHash(c3, account, key, blockHash)
	})
}

// skipping Fuzz_Nosy_Client_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_Client_SubscribeNewHead__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Header

func Fuzz_Nosy_Client_SuggestGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.SuggestGasPrice(c3)
	})
}

func Fuzz_Nosy_Client_SuggestGasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.SuggestGasTipCap(c3)
	})
}

func Fuzz_Nosy_Client_SyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.SyncProgress(c3)
	})
}

func Fuzz_Nosy_Client_TransactionByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.TransactionByHash(c3, hash)
	})
}

func Fuzz_Nosy_Client_TransactionCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var blockHash common.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.TransactionCount(c3, blockHash)
	})
}

func Fuzz_Nosy_Client_TransactionInBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.TransactionInBlock(c3, blockHash, index)
	})
}

func Fuzz_Nosy_Client_TransactionReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var txHash common.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.TransactionReceipt(c3, txHash)
	})
}

func Fuzz_Nosy_Client_TransactionSender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var rawurl string
		fill_err = tp.Fill(&rawurl)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var block common.Hash
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		ec, err := DialContext(c1, rawurl)
		if err != nil {
			return
		}
		ec.TransactionSender(c3, tx, block, index)
	})
}

// skipping Fuzz_Nosy_Client_getBlock__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_rpcProgress_toSyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *rpcProgress
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.toSyncProgress()
	})
}

func Fuzz_Nosy_rpcTransaction_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *rpcTransaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.UnmarshalJSON(msg)
	})
}

func Fuzz_Nosy_senderFromServer_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *senderFromServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ChainID()
	})
}

// skipping Fuzz_Nosy_senderFromServer_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_senderFromServer_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *senderFromServer
		fill_err = tp.Fill(&s)
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

		s.Hash(tx)
	})
}

func Fuzz_Nosy_senderFromServer_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *senderFromServer
		fill_err = tp.Fill(&s)
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

		s.Sender(tx)
	})
}

func Fuzz_Nosy_senderFromServer_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *senderFromServer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil {
			return
		}

		s.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_setSenderFromServer__(f *testing.F) {
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
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var block common.Hash
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		setSenderFromServer(tx, addr, block)
	})
}

func Fuzz_Nosy_toBlockNumArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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

		toBlockNumArg(number)
	})
}

func Fuzz_Nosy_toCallArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		toCallArg(msg)
	})
}

func Fuzz_Nosy_toFilterArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}

		toFilterArg(q)
	})
}
