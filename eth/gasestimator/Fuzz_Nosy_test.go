package gasestimator

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core"
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

func Fuzz_Nosy_Estimate__(f *testing.F) {
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
		var call *core.Message
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var opts *Options
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gasCap uint64
		fill_err = tp.Fill(&gasCap)
		if fill_err != nil {
			return
		}
		if call == nil || opts == nil {
			return
		}

		Estimate(ctx, call, opts, gasCap)
	})
}

func Fuzz_Nosy_execute__(f *testing.F) {
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
		var call *core.Message
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		var opts *Options
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if call == nil || opts == nil {
			return
		}

		execute(ctx, call, opts, gasLimit)
	})
}
