package consensus

import (
	"testing"

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

func Fuzz_Nosy_Merger_FinalizePoS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Merger
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.FinalizePoS()
	})
}

func Fuzz_Nosy_Merger_PoSFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Merger
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.PoSFinalized()
	})
}

func Fuzz_Nosy_Merger_ReachTTD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Merger
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ReachTTD()
	})
}

func Fuzz_Nosy_Merger_TDDReached__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Merger
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.TDDReached()
	})
}

// skipping Fuzz_Nosy_ChainHeaderReader_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainHeaderReader_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainHeaderReader_GetHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainHeaderReader_GetHeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainHeaderReader_GetHeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainHeaderReader_GetTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_ChainReader_GetBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainReader

// skipping Fuzz_Nosy_Engine_APIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_Author__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_CalcDifficulty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_FinalizeAndAssemble__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_Prepare__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_Seal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_SealHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_VerifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_VerifyHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_Engine_VerifyUncles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_PoW_Hashrate__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.PoW
