package msgrate

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

func Fuzz_Nosy_Tracker_Capacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var caps map[uint64]float64
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var rtt time.Duration
		fill_err = tp.Fill(&rtt)
		if fill_err != nil {
			return
		}
		var kind uint64
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var targetRTT time.Duration
		fill_err = tp.Fill(&targetRTT)
		if fill_err != nil {
			return
		}

		t1 := NewTracker(caps, rtt)
		t1.Capacity(kind, targetRTT)
	})
}

func Fuzz_Nosy_Tracker_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var caps map[uint64]float64
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var rtt time.Duration
		fill_err = tp.Fill(&rtt)
		if fill_err != nil {
			return
		}
		var kind uint64
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}

		t1 := NewTracker(caps, rtt)
		t1.Update(kind, elapsed, items)
	})
}

func Fuzz_Nosy_Trackers_Capacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var kind uint64
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var targetRTT time.Duration
		fill_err = tp.Fill(&targetRTT)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Capacity(id, kind, targetRTT)
	})
}

func Fuzz_Nosy_Trackers_MeanCapacities__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.MeanCapacities()
	})
}

func Fuzz_Nosy_Trackers_MedianRoundTrip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.MedianRoundTrip()
	})
}

func Fuzz_Nosy_Trackers_TargetRoundTrip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.TargetRoundTrip()
	})
}

func Fuzz_Nosy_Trackers_TargetTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.TargetTimeout()
	})
}

func Fuzz_Nosy_Trackers_Track__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var tracker *Tracker
		fill_err = tp.Fill(&tracker)
		if fill_err != nil {
			return
		}
		if t1 == nil || tracker == nil {
			return
		}

		t1.Track(id, tracker)
	})
}

func Fuzz_Nosy_Trackers_Untrack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Untrack(id)
	})
}

func Fuzz_Nosy_Trackers_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var kind uint64
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Update(id, kind, elapsed, items)
	})
}

func Fuzz_Nosy_Trackers_detune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.detune()
	})
}

func Fuzz_Nosy_Trackers_meanCapacities__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.meanCapacities()
	})
}

func Fuzz_Nosy_Trackers_medianRoundTrip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.medianRoundTrip()
	})
}

func Fuzz_Nosy_Trackers_targetTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.targetTimeout()
	})
}

func Fuzz_Nosy_Trackers_tune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Trackers
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.tune()
	})
}

func Fuzz_Nosy_roundCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, cap float64) {
		roundCapacity(cap)
	})
}
