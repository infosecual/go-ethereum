package gasprice

import (
	"context"
	"testing"

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

func Fuzz_Nosy_Oracle_FeeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oracle *Oracle
		fill_err = tp.Fill(&oracle)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var blocks uint64
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		var unresolvedLastBlock rpc.BlockNumber
		fill_err = tp.Fill(&unresolvedLastBlock)
		if fill_err != nil {
			return
		}
		var rewardPercentiles []float64
		fill_err = tp.Fill(&rewardPercentiles)
		if fill_err != nil {
			return
		}
		if oracle == nil {
			return
		}

		oracle.FeeHistory(ctx, blocks, unresolvedLastBlock, rewardPercentiles)
	})
}

func Fuzz_Nosy_Oracle_SuggestTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oracle *Oracle
		fill_err = tp.Fill(&oracle)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if oracle == nil {
			return
		}

		oracle.SuggestTipCap(ctx)
	})
}

// skipping Fuzz_Nosy_Oracle_getBlockValues__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum/go-ethereum/eth/gasprice.results

func Fuzz_Nosy_Oracle_processBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oracle *Oracle
		fill_err = tp.Fill(&oracle)
		if fill_err != nil {
			return
		}
		var bf *blockFees
		fill_err = tp.Fill(&bf)
		if fill_err != nil {
			return
		}
		var percentiles []float64
		fill_err = tp.Fill(&percentiles)
		if fill_err != nil {
			return
		}
		if oracle == nil || bf == nil {
			return
		}

		oracle.processBlock(bf, percentiles)
	})
}

func Fuzz_Nosy_Oracle_resolveBlockRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var oracle *Oracle
		fill_err = tp.Fill(&oracle)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var reqEnd rpc.BlockNumber
		fill_err = tp.Fill(&reqEnd)
		if fill_err != nil {
			return
		}
		var blocks uint64
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if oracle == nil {
			return
		}

		oracle.resolveBlockRange(ctx, reqEnd, blocks)
	})
}

// skipping Fuzz_Nosy_OracleBackend_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend

// skipping Fuzz_Nosy_OracleBackend_ChainConfig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend

// skipping Fuzz_Nosy_OracleBackend_GetReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend

// skipping Fuzz_Nosy_OracleBackend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend

// skipping Fuzz_Nosy_OracleBackend_PendingBlockAndReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend

// skipping Fuzz_Nosy_OracleBackend_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/gasprice.OracleBackend
