package gethclient

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_Client_CallContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var overrides *map[common.Address]OverrideAccount
		fill_err = tp.Fill(&overrides)
		if fill_err != nil {
			return
		}
		if c == nil || blockNumber == nil || overrides == nil {
			return
		}

		ec := New(c)
		ec.CallContract(ctx, msg, blockNumber, overrides)
	})
}

func Fuzz_Nosy_Client_CallContractWithBlockOverrides__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var overrides *map[common.Address]OverrideAccount
		fill_err = tp.Fill(&overrides)
		if fill_err != nil {
			return
		}
		var blockOverrides BlockOverrides
		fill_err = tp.Fill(&blockOverrides)
		if fill_err != nil {
			return
		}
		if c == nil || blockNumber == nil || overrides == nil {
			return
		}

		ec := New(c)
		ec.CallContractWithBlockOverrides(ctx, msg, blockNumber, overrides, blockOverrides)
	})
}

func Fuzz_Nosy_Client_CreateAccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg ethereum.CallMsg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		ec := New(c)
		ec.CreateAccessList(ctx, msg)
	})
}

func Fuzz_Nosy_Client_GCStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		ec := New(c)
		ec.GCStats(ctx)
	})
}

func Fuzz_Nosy_Client_GetNodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		ec := New(c)
		ec.GetNodeInfo(ctx)
	})
}

func Fuzz_Nosy_Client_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
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
		var keys []string
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		var blockNumber *big.Int
		fill_err = tp.Fill(&blockNumber)
		if fill_err != nil {
			return
		}
		if c == nil || blockNumber == nil {
			return
		}

		ec := New(c)
		ec.GetProof(ctx, account, keys, blockNumber)
	})
}

func Fuzz_Nosy_Client_MemStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		ec := New(c)
		ec.MemStats(ctx)
	})
}

func Fuzz_Nosy_Client_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *rpc.Client
		fill_err = tp.Fill(&c)
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
		if c == nil || number == nil {
			return
		}

		ec := New(c)
		ec.SetHead(ctx, number)
	})
}

// skipping Fuzz_Nosy_Client_SubscribeFullPendingTransactions__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/core/types.Transaction

// skipping Fuzz_Nosy_Client_SubscribePendingTransactions__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/common.Hash

func Fuzz_Nosy_BlockOverrides_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o BlockOverrides
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}

		o.MarshalJSON()
	})
}

func Fuzz_Nosy_OverrideAccount_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a OverrideAccount
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.MarshalJSON()
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
