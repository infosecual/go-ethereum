package influxdb

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

func Fuzz_Nosy_reporter_makeClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.makeClient()
	})
}

func Fuzz_Nosy_reporter_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.run()
	})
}

func Fuzz_Nosy_reporter_send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var tstamp int64
		fill_err = tp.Fill(&tstamp)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.send(tstamp)
	})
}

func Fuzz_Nosy_v2Reporter_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *v2Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.run()
	})
}

func Fuzz_Nosy_v2Reporter_send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *v2Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var tstamp int64
		fill_err = tp.Fill(&tstamp)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.send(tstamp)
	})
}

// skipping Fuzz_Nosy_InfluxDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_InfluxDBV2WithTags__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_InfluxDBWithTags__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_readMeter__ because parameters include func, chan, or unsupported interface: interface{}
