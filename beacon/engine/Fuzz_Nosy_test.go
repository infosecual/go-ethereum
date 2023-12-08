package engine

import (
	"math/big"
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

func Fuzz_Nosy_EngineAPIError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineAPIError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_EngineAPIError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineAPIError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_EngineAPIError_ErrorData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *EngineAPIError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorData()
	})
}

// skipping Fuzz_Nosy_EngineAPIError_With__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ExecutableData_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ExecutableData
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_ExecutionPayloadEnvelope_UnmarshalJSON__(f *testing.F) {
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
		var fees *big.Int
		fill_err = tp.Fill(&fees)
		if fill_err != nil {
			return
		}
		var sidecars []*types.BlobTxSidecar
		fill_err = tp.Fill(&sidecars)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if block == nil || fees == nil {
			return
		}

		e := BlockToExecutableData(block, fees, sidecars)
		e.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_PayloadAttributes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PayloadAttributes
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_PayloadID_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *PayloadID
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

func Fuzz_Nosy_ExecutableData_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e ExecutableData
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.MarshalJSON()
	})
}

func Fuzz_Nosy_ExecutionPayloadEnvelope_MarshalJSON__(f *testing.F) {
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
		var fees *big.Int
		fill_err = tp.Fill(&fees)
		if fill_err != nil {
			return
		}
		var sidecars []*types.BlobTxSidecar
		fill_err = tp.Fill(&sidecars)
		if fill_err != nil {
			return
		}
		if block == nil || fees == nil {
			return
		}

		e := BlockToExecutableData(block, fees, sidecars)
		e.MarshalJSON()
	})
}

func Fuzz_Nosy_PayloadAttributes_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p PayloadAttributes
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.MarshalJSON()
	})
}

func Fuzz_Nosy_PayloadID_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b PayloadID
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_PayloadID_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b PayloadID
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_decodeTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var enc [][]byte
		fill_err = tp.Fill(&enc)
		if fill_err != nil {
			return
		}

		decodeTransactions(enc)
	})
}

func Fuzz_Nosy_encodeTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		encodeTransactions(txs)
	})
}
