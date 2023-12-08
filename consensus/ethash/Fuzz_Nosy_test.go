package ethash

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
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

// skipping Fuzz_Nosy_Ethash_APIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Ethash_Author__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		e1 := NewFullFaker()
		e1.Author(header)
	})
}

// skipping Fuzz_Nosy_Ethash_CalcDifficulty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_FinalizeAndAssemble__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_Prepare__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_Seal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Ethash_SealHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		e1 := NewFullFaker()
		e1.SealHash(header)
	})
}

// skipping Fuzz_Nosy_Ethash_VerifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_VerifyHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Ethash_VerifyUncles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainReader

// skipping Fuzz_Nosy_Ethash_verifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_MakeDifficultyCalculatorU256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bombDelay *big.Int
		fill_err = tp.Fill(&bombDelay)
		if fill_err != nil {
			return
		}
		if bombDelay == nil {
			return
		}

		MakeDifficultyCalculatorU256(bombDelay)
	})
}

func Fuzz_Nosy_accumulateRewards__(f *testing.F) {
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
		var uncles []*types.Header
		fill_err = tp.Fill(&uncles)
		if fill_err != nil {
			return
		}
		if config == nil || state == nil || header == nil {
			return
		}

		accumulateRewards(config, state, header, uncles)
	})
}

func Fuzz_Nosy_makeDifficultyCalculator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bombDelay *big.Int
		fill_err = tp.Fill(&bombDelay)
		if fill_err != nil {
			return
		}
		if bombDelay == nil {
			return
		}

		makeDifficultyCalculator(bombDelay)
	})
}
