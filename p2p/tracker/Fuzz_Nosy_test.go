package tracker

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

func Fuzz_Nosy_Tracker_Fulfil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var version uint
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		t1 := New(protocol, timeout)
		t1.Fulfil(peer, version, code, id)
	})
}

func Fuzz_Nosy_Tracker_Track__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var version uint
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var reqCode uint64
		fill_err = tp.Fill(&reqCode)
		if fill_err != nil {
			return
		}
		var resCode uint64
		fill_err = tp.Fill(&resCode)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		t1 := New(protocol, timeout)
		t1.Track(peer, version, reqCode, resCode, id)
	})
}

func Fuzz_Nosy_Tracker_clean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		t1 := New(protocol, timeout)
		t1.clean()
	})
}

func Fuzz_Nosy_Tracker_schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		t1 := New(protocol, timeout)
		t1.schedule()
	})
}
