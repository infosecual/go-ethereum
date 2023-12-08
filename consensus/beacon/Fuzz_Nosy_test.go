package beacon

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
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

// skipping Fuzz_Nosy_Beacon_APIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Beacon_Author__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if b1 == nil || header == nil {
			return
		}

		b1.Author(header)
	})
}

// skipping Fuzz_Nosy_Beacon_CalcDifficulty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Beacon_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		if b1 == nil {
			return
		}

		b1.Close()
	})
}

// skipping Fuzz_Nosy_Beacon_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_FinalizeAndAssemble__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Beacon_InnerEngine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		if b1 == nil {
			return
		}

		b1.InnerEngine()
	})
}

func Fuzz_Nosy_Beacon_IsPoSHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if b1 == nil || header == nil {
			return
		}

		b1.IsPoSHeader(header)
	})
}

// skipping Fuzz_Nosy_Beacon_Prepare__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_Seal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Beacon_SealHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if b1 == nil || header == nil {
			return
		}

		b1.SealHash(header)
	})
}

func Fuzz_Nosy_Beacon_SetThreads__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *Beacon
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		var threads int
		fill_err = tp.Fill(&threads)
		if fill_err != nil {
			return
		}
		if b1 == nil {
			return
		}

		b1.SetThreads(threads)
	})
}

// skipping Fuzz_Nosy_Beacon_VerifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_VerifyHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_VerifyUncles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainReader

// skipping Fuzz_Nosy_Beacon_splitHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_verifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Beacon_verifyHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_faker_CalcDifficulty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_threaded_SetThreads__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus/beacon.threaded

// skipping Fuzz_Nosy_IsTTDReached__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_errOut__ because parameters include func, chan, or unsupported interface: error
