package forkid

import (
	"testing"

	"github.com/ethereum/go-ethereum/params"
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

// skipping Fuzz_Nosy_Blockchain_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/forkid.Blockchain

// skipping Fuzz_Nosy_Blockchain_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/forkid.Blockchain

// skipping Fuzz_Nosy_Blockchain_Genesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/forkid.Blockchain

func Fuzz_Nosy_checksumToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash uint32) {
		checksumToBytes(hash)
	})
}

func Fuzz_Nosy_checksumUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash uint32, fork uint64) {
		checksumUpdate(hash, fork)
	})
}

func Fuzz_Nosy_gatherForks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var genesis uint64
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		gatherForks(config, genesis)
	})
}
