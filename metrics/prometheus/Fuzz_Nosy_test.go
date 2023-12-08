package prometheus

import (
	"testing"

	"github.com/ethereum/go-ethereum/metrics"
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

// skipping Fuzz_Nosy_collector_Add__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_collector_addCounter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterSnapshot

// skipping Fuzz_Nosy_collector_addCounterFloat64__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterFloat64Snapshot

// skipping Fuzz_Nosy_collector_addGauge__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeSnapshot

// skipping Fuzz_Nosy_collector_addGaugeFloat64__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeFloat64Snapshot

// skipping Fuzz_Nosy_collector_addGaugeInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeInfoSnapshot

// skipping Fuzz_Nosy_collector_addHistogram__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.HistogramSnapshot

// skipping Fuzz_Nosy_collector_addMeter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

// skipping Fuzz_Nosy_collector_addResettingTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

// skipping Fuzz_Nosy_collector_addTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.TimerSnapshot

// skipping Fuzz_Nosy_collector_writeGaugeCounter__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_collector_writeGaugeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var value metrics.GaugeInfoValue
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		c := newCollector()
		c.writeGaugeInfo(name, value)
	})
}

// skipping Fuzz_Nosy_collector_writeSummaryCounter__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_collector_writeSummaryPercentile__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_mutateKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key string) {
		mutateKey(key)
	})
}
