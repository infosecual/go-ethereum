package exp

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

// skipping Fuzz_Nosy_exp_expHandler__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_exp_getFloat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *exp
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		e1.getFloat(name)
	})
}

func Fuzz_Nosy_exp_getInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *exp
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		e1.getInfo(name)
	})
}

func Fuzz_Nosy_exp_getInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *exp
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		e1.getInt(name)
	})
}

// skipping Fuzz_Nosy_exp_publishCounter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterSnapshot

// skipping Fuzz_Nosy_exp_publishCounterFloat64__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterFloat64Snapshot

// skipping Fuzz_Nosy_exp_publishGauge__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeSnapshot

// skipping Fuzz_Nosy_exp_publishGaugeFloat64__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeFloat64Snapshot

// skipping Fuzz_Nosy_exp_publishGaugeInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeInfoSnapshot

// skipping Fuzz_Nosy_exp_publishHistogram__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Histogram

// skipping Fuzz_Nosy_exp_publishMeter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Meter

// skipping Fuzz_Nosy_exp_publishResettingTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimer

// skipping Fuzz_Nosy_exp_publishTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Timer

func Fuzz_Nosy_exp_syncToExpvar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e1 *exp
		fill_err = tp.Fill(&e1)
		if fill_err != nil {
			return
		}
		if e1 == nil {
			return
		}

		e1.syncToExpvar()
	})
}

// skipping Fuzz_Nosy_Exp__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_Setup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address string) {
		Setup(address)
	})
}
