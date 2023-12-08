package mclock

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

func Fuzz_Nosy_Alarm_C__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Alarm
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.C()
	})
}

func Fuzz_Nosy_Alarm_Schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Alarm
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var time AbsTime
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Schedule(time)
	})
}

func Fuzz_Nosy_Alarm_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Alarm
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Stop()
	})
}

func Fuzz_Nosy_Alarm_schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Alarm
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var now AbsTime
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var newDeadline AbsTime
		fill_err = tp.Fill(&newDeadline)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.schedule(now, newDeadline)
	})
}

func Fuzz_Nosy_Alarm_send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Alarm
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.send()
	})
}

func Fuzz_Nosy_Simulated_ActiveTimers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ActiveTimers()
	})
}

func Fuzz_Nosy_Simulated_After__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.After(d)
	})
}

// skipping Fuzz_Nosy_Simulated_AfterFunc__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_Simulated_NewTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NewTimer(d)
	})
}

func Fuzz_Nosy_Simulated_Now__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Now()
	})
}

func Fuzz_Nosy_Simulated_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Run(d)
	})
}

func Fuzz_Nosy_Simulated_Sleep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Sleep(d)
	})
}

func Fuzz_Nosy_Simulated_WaitForTimers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.WaitForTimers(n)
	})
}

func Fuzz_Nosy_Simulated_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Simulated
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.init()
	})
}

// skipping Fuzz_Nosy_Simulated_schedule__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_simTimer_C__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *simTimer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.C()
	})
}

func Fuzz_Nosy_simTimer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *simTimer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.Reset(d)
	})
}

func Fuzz_Nosy_simTimer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *simTimer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.Stop()
	})
}

func Fuzz_Nosy_simTimerHeap_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *simTimerHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Len()
	})
}

func Fuzz_Nosy_simTimerHeap_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *simTimerHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Less(i, j)
	})
}

func Fuzz_Nosy_simTimerHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *simTimerHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Pop()
	})
}

// skipping Fuzz_Nosy_simTimerHeap_Push__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_simTimerHeap_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *simTimerHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Swap(i, j)
	})
}

func Fuzz_Nosy_systemTimer_C__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *systemTimer
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.C()
	})
}

func Fuzz_Nosy_systemTimer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *systemTimer
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.Reset(d)
	})
}

func Fuzz_Nosy_AbsTime_Add__(f *testing.F) {
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

		t1 := Now()
		t1.Add(d)
	})
}

func Fuzz_Nosy_AbsTime_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t2 AbsTime
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		t1 := Now()
		t1.Sub(t2)
	})
}

// skipping Fuzz_Nosy_ChanTimer_C__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.ChanTimer

// skipping Fuzz_Nosy_ChanTimer_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.ChanTimer

// skipping Fuzz_Nosy_Clock_After__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Clock

// skipping Fuzz_Nosy_Clock_AfterFunc__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Clock

// skipping Fuzz_Nosy_Clock_NewTimer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Clock

// skipping Fuzz_Nosy_Clock_Now__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Clock

// skipping Fuzz_Nosy_Clock_Sleep__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Clock

func Fuzz_Nosy_System_After__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c System
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		c.After(d)
	})
}

// skipping Fuzz_Nosy_System_AfterFunc__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_System_NewTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c System
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		c.NewTimer(d)
	})
}

func Fuzz_Nosy_System_Now__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c System
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Now()
	})
}

func Fuzz_Nosy_System_Sleep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c System
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		c.Sleep(d)
	})
}

// skipping Fuzz_Nosy_Timer_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/common/mclock.Timer
