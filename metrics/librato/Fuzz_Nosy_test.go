package librato

import (
	"testing"
	"time"

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

func Fuzz_Nosy_LibratoClient_PostMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *LibratoClient
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var batch Batch
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PostMetrics(batch)
	})
}

// skipping Fuzz_Nosy_Reporter_BuildRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_Reporter_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rep *Reporter
		fill_err = tp.Fill(&rep)
		if fill_err != nil {
			return
		}
		if rep == nil {
			return
		}

		rep.Run()
	})
}

// skipping Fuzz_Nosy_Librato__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_sumSquares__(f *testing.F) {
	f.Fuzz(func(t *testing.T, icount int64, mean float64, stDev float64) {
		sumSquares(icount, mean, stDev)
	})
}

// skipping Fuzz_Nosy_sumSquaresTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.TimerSnapshot

func Fuzz_Nosy_translateTimerAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		translateTimerAttributes(d)
	})
}
