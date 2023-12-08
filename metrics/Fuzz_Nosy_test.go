package metrics

import (
	"io"
	"math/rand"
	"runtime/metrics"
	"sync/atomic"
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

func Fuzz_Nosy_ExpDecaySample_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExpDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Clear()
	})
}

func Fuzz_Nosy_ExpDecaySample_SetRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExpDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var prng *rand.Rand
		fill_err = tp.Fill(&prng)
		if fill_err != nil {
			return
		}
		if s == nil || prng == nil {
			return
		}

		s.SetRand(prng)
	})
}

func Fuzz_Nosy_ExpDecaySample_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExpDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Snapshot()
	})
}

func Fuzz_Nosy_ExpDecaySample_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExpDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Update(v)
	})
}

func Fuzz_Nosy_ExpDecaySample_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *ExpDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.update(t2, v)
	})
}

func Fuzz_Nosy_OpenTSDBConfig_writeRegistry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *OpenTSDBConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var now int64
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
		var shortHostname string
		fill_err = tp.Fill(&shortHostname)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.writeRegistry(w, now, shortHostname)
	})
}

// skipping Fuzz_Nosy_PrefixedRegistry_Each__ because parameters include func, chan, or unsupported interface: func(string, interface{})

func Fuzz_Nosy_PrefixedRegistry_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PrefixedRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Get(name)
	})
}

func Fuzz_Nosy_PrefixedRegistry_GetAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PrefixedRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.GetAll()
	})
}

// skipping Fuzz_Nosy_PrefixedRegistry_GetOrRegister__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_PrefixedRegistry_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PrefixedRegistry
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarshalJSON()
	})
}

// skipping Fuzz_Nosy_PrefixedRegistry_Register__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_PrefixedRegistry_RunHealthchecks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PrefixedRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.RunHealthchecks()
	})
}

func Fuzz_Nosy_PrefixedRegistry_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PrefixedRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Unregister(name)
	})
}

func Fuzz_Nosy_StandardCounter_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Clear()
	})
}

func Fuzz_Nosy_StandardCounter_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Dec(i)
	})
}

func Fuzz_Nosy_StandardCounter_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Inc(i)
	})
}

func Fuzz_Nosy_StandardCounter_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Snapshot()
	})
}

func Fuzz_Nosy_StandardCounterFloat64_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounterFloat64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Clear()
	})
}

func Fuzz_Nosy_StandardCounterFloat64_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounterFloat64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v float64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Dec(v)
	})
}

func Fuzz_Nosy_StandardCounterFloat64_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounterFloat64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var v float64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Inc(v)
	})
}

func Fuzz_Nosy_StandardCounterFloat64_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *StandardCounterFloat64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Snapshot()
	})
}

func Fuzz_Nosy_StandardEWMA_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *StandardEWMA
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Snapshot()
	})
}

func Fuzz_Nosy_StandardEWMA_Tick__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *StandardEWMA
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Tick()
	})
}

func Fuzz_Nosy_StandardEWMA_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *StandardEWMA
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var n int64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.Update(n)
	})
}

func Fuzz_Nosy_StandardEWMA_fetchInstantRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *StandardEWMA
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.fetchInstantRate()
	})
}

func Fuzz_Nosy_StandardEWMA_updateRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *StandardEWMA
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var instantRate float64
		fill_err = tp.Fill(&instantRate)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.updateRate(instantRate)
	})
}

func Fuzz_Nosy_StandardGauge_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGauge
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Dec(i)
	})
}

func Fuzz_Nosy_StandardGauge_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGauge
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Inc(i)
	})
}

func Fuzz_Nosy_StandardGauge_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGauge
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Snapshot()
	})
}

func Fuzz_Nosy_StandardGauge_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGauge
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Update(v)
	})
}

func Fuzz_Nosy_StandardGauge_UpdateIfGt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGauge
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UpdateIfGt(v)
	})
}

func Fuzz_Nosy_StandardGaugeFloat64_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGaugeFloat64
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Snapshot()
	})
}

func Fuzz_Nosy_StandardGaugeFloat64_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGaugeFloat64
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var v float64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Update(v)
	})
}

func Fuzz_Nosy_StandardGaugeInfo_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGaugeInfo
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Snapshot()
	})
}

func Fuzz_Nosy_StandardGaugeInfo_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *StandardGaugeInfo
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var v GaugeInfoValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.Update(v)
	})
}

func Fuzz_Nosy_StandardHealthcheck_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHealthcheck
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Check()
	})
}

func Fuzz_Nosy_StandardHealthcheck_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHealthcheck
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Error()
	})
}

func Fuzz_Nosy_StandardHealthcheck_Healthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHealthcheck
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Healthy()
	})
}

// skipping Fuzz_Nosy_StandardHealthcheck_Unhealthy__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StandardHistogram_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHistogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Clear()
	})
}

func Fuzz_Nosy_StandardHistogram_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHistogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Snapshot()
	})
}

func Fuzz_Nosy_StandardHistogram_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StandardHistogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Update(v)
	})
}

func Fuzz_Nosy_StandardMeter_Mark__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int64) {
		m := newStandardMeter()
		m.Mark(n)
	})
}

// skipping Fuzz_Nosy_StandardRegistry_Each__ because parameters include func, chan, or unsupported interface: func(string, interface{})

func Fuzz_Nosy_StandardRegistry_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Get(name)
	})
}

func Fuzz_Nosy_StandardRegistry_GetAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.GetAll()
	})
}

// skipping Fuzz_Nosy_StandardRegistry_GetOrRegister__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_StandardRegistry_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.MarshalJSON()
	})
}

// skipping Fuzz_Nosy_StandardRegistry_Register__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_StandardRegistry_RunHealthchecks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.RunHealthchecks()
	})
}

func Fuzz_Nosy_StandardRegistry_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Unregister(name)
	})
}

// skipping Fuzz_Nosy_StandardRegistry_loadOrRegister__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_StandardRegistry_registered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.registered()
	})
}

func Fuzz_Nosy_StandardRegistry_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *StandardRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.stop(name)
	})
}

func Fuzz_Nosy_StandardResettingTimer_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardResettingTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Snapshot()
	})
}

// skipping Fuzz_Nosy_StandardResettingTimer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_StandardResettingTimer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardResettingTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Update(d)
	})
}

func Fuzz_Nosy_StandardResettingTimer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardResettingTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ts time.Time
		fill_err = tp.Fill(&ts)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateSince(ts)
	})
}

func Fuzz_Nosy_StandardTimer_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Snapshot()
	})
}

func Fuzz_Nosy_StandardTimer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Stop()
	})
}

// skipping Fuzz_Nosy_StandardTimer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_StandardTimer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var d time.Duration
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Update(d)
	})
}

func Fuzz_Nosy_StandardTimer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StandardTimer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ts time.Time
		fill_err = tp.Fill(&ts)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UpdateSince(ts)
	})
}

func Fuzz_Nosy_UniformSample_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *UniformSample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Clear()
	})
}

func Fuzz_Nosy_UniformSample_SetRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *UniformSample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var prng *rand.Rand
		fill_err = tp.Fill(&prng)
		if fill_err != nil {
			return
		}
		if s == nil || prng == nil {
			return
		}

		s.SetRand(prng)
	})
}

func Fuzz_Nosy_UniformSample_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *UniformSample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Snapshot()
	})
}

func Fuzz_Nosy_UniformSample_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *UniformSample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Update(v)
	})
}

func Fuzz_Nosy_emptySnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Count()
	})
}

func Fuzz_Nosy_emptySnapshot_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Max()
	})
}

func Fuzz_Nosy_emptySnapshot_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Mean()
	})
}

func Fuzz_Nosy_emptySnapshot_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Min()
	})
}

func Fuzz_Nosy_emptySnapshot_Percentile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var p float64
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Percentile(p)
	})
}

func Fuzz_Nosy_emptySnapshot_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ps []float64
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Percentiles(ps)
	})
}

func Fuzz_Nosy_emptySnapshot_Rate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Rate()
	})
}

func Fuzz_Nosy_emptySnapshot_Rate1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Rate1()
	})
}

func Fuzz_Nosy_emptySnapshot_Rate15__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Rate15()
	})
}

func Fuzz_Nosy_emptySnapshot_Rate5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Rate5()
	})
}

func Fuzz_Nosy_emptySnapshot_RateMean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RateMean()
	})
}

func Fuzz_Nosy_emptySnapshot_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Size()
	})
}

func Fuzz_Nosy_emptySnapshot_StdDev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.StdDev()
	})
}

func Fuzz_Nosy_emptySnapshot_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Sum()
	})
}

func Fuzz_Nosy_emptySnapshot_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Value()
	})
}

func Fuzz_Nosy_emptySnapshot_Values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Values()
	})
}

func Fuzz_Nosy_emptySnapshot_Variance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *emptySnapshot
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Variance()
	})
}

func Fuzz_Nosy_expDecaySampleHeap_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.Clear()
	})
}

func Fuzz_Nosy_expDecaySampleHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.Pop()
	})
}

func Fuzz_Nosy_expDecaySampleHeap_Push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reservoirSize int
		fill_err = tp.Fill(&reservoirSize)
		if fill_err != nil {
			return
		}
		var s expDecaySample
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		h := newExpDecaySampleHeap(reservoirSize)
		h.Push(s)
	})
}

func Fuzz_Nosy_expDecaySampleHeap_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.Size()
	})
}

func Fuzz_Nosy_expDecaySampleHeap_Values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.Values()
	})
}

func Fuzz_Nosy_expDecaySampleHeap_down__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int, i int, n int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.down(i, n)
	})
}

func Fuzz_Nosy_expDecaySampleHeap_up__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int, j int) {
		h := newExpDecaySampleHeap(reservoirSize)
		h.up(j)
	})
}

func Fuzz_Nosy_meterArbiter_tick__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ma *meterArbiter
		fill_err = tp.Fill(&ma)
		if fill_err != nil {
			return
		}
		if ma == nil {
			return
		}

		ma.tick()
	})
}

func Fuzz_Nosy_meterArbiter_tickMeters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ma *meterArbiter
		fill_err = tp.Fill(&ma)
		if fill_err != nil {
			return
		}
		if ma == nil {
			return
		}

		ma.tickMeters()
	})
}

func Fuzz_Nosy_meterSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meterSnapshot
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Count()
	})
}

func Fuzz_Nosy_meterSnapshot_Rate1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meterSnapshot
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Rate1()
	})
}

func Fuzz_Nosy_meterSnapshot_Rate15__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meterSnapshot
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Rate15()
	})
}

func Fuzz_Nosy_meterSnapshot_Rate5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meterSnapshot
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Rate5()
	})
}

func Fuzz_Nosy_meterSnapshot_RateMean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *meterSnapshot
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.RateMean()
	})
}

// skipping Fuzz_Nosy_orderedRegistry_Each__ because parameters include func, chan, or unsupported interface: func(string, interface{})

func Fuzz_Nosy_resettingSample_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs *resettingSample
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		if rs == nil {
			return
		}

		rs.Snapshot()
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Count()
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Max()
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Mean()
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Min()
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var percentiles []float64
		fill_err = tp.Fill(&percentiles)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Percentiles(percentiles)
	})
}

func Fuzz_Nosy_resettingTimerSnapshot_calc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *resettingTimerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var percentiles []float64
		fill_err = tp.Fill(&percentiles)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.calc(percentiles)
	})
}

func Fuzz_Nosy_runtimeHistogram_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, scale float64) {
		h := newRuntimeHistogram(scale)
		h.Clear()
	})
}

func Fuzz_Nosy_runtimeHistogram_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, scale float64) {
		h := newRuntimeHistogram(scale)
		h.Snapshot()
	})
}

func Fuzz_Nosy_runtimeHistogram_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, scale float64, _x2 int64) {
		h := newRuntimeHistogram(scale)
		h.Update(_x2)
	})
}

func Fuzz_Nosy_runtimeHistogram_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var scale float64
		fill_err = tp.Fill(&scale)
		if fill_err != nil {
			return
		}
		var mh *metrics.Float64Histogram
		fill_err = tp.Fill(&mh)
		if fill_err != nil {
			return
		}
		if mh == nil {
			return
		}

		h := newRuntimeHistogram(scale)
		h.update(mh)
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Count()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Max()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Mean()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Min()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Percentile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var p float64
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Percentile(p)
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ps []float64
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Percentiles(ps)
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Size()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_StdDev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.StdDev()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Sum()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_Variance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.Variance()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_calc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.calc()
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_computePercentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var thresh []float64
		fill_err = tp.Fill(&thresh)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.computePercentiles(thresh)
	})
}

func Fuzz_Nosy_runtimeHistogramSnapshot_midpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *metrics.Float64Histogram
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var bucket int
		fill_err = tp.Fill(&bucket)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := newRuntimeHistogramSnapshot(h)
		h1.midpoint(bucket)
	})
}

func Fuzz_Nosy_sampleSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Count()
	})
}

func Fuzz_Nosy_sampleSnapshot_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Max()
	})
}

func Fuzz_Nosy_sampleSnapshot_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Mean()
	})
}

func Fuzz_Nosy_sampleSnapshot_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Min()
	})
}

func Fuzz_Nosy_sampleSnapshot_Percentile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}
		var p float64
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Percentile(p)
	})
}

func Fuzz_Nosy_sampleSnapshot_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}
		var ps []float64
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Percentiles(ps)
	})
}

func Fuzz_Nosy_sampleSnapshot_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Size()
	})
}

func Fuzz_Nosy_sampleSnapshot_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Snapshot()
	})
}

func Fuzz_Nosy_sampleSnapshot_StdDev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.StdDev()
	})
}

func Fuzz_Nosy_sampleSnapshot_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Sum()
	})
}

func Fuzz_Nosy_sampleSnapshot_Values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Values()
	})
}

func Fuzz_Nosy_sampleSnapshot_Variance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var count int64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var min int64
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int64
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var sum int64
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}

		s := newSampleSnapshotPrecalculated(count, values, min, max, sum)
		s.Variance()
	})
}

func Fuzz_Nosy_timerSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Count()
	})
}

func Fuzz_Nosy_timerSnapshot_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Max()
	})
}

func Fuzz_Nosy_timerSnapshot_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Mean()
	})
}

func Fuzz_Nosy_timerSnapshot_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Min()
	})
}

func Fuzz_Nosy_timerSnapshot_Percentile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p float64
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Percentile(p)
	})
}

func Fuzz_Nosy_timerSnapshot_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ps []float64
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Percentiles(ps)
	})
}

func Fuzz_Nosy_timerSnapshot_Rate1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Rate1()
	})
}

func Fuzz_Nosy_timerSnapshot_Rate15__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Rate15()
	})
}

func Fuzz_Nosy_timerSnapshot_Rate5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Rate5()
	})
}

func Fuzz_Nosy_timerSnapshot_RateMean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RateMean()
	})
}

func Fuzz_Nosy_timerSnapshot_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Size()
	})
}

func Fuzz_Nosy_timerSnapshot_StdDev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.StdDev()
	})
}

func Fuzz_Nosy_timerSnapshot_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Sum()
	})
}

func Fuzz_Nosy_timerSnapshot_Variance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *timerSnapshot
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Variance()
	})
}

func Fuzz_Nosy_Counter_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewCounterForced()
		_x1.Dec(_x1)
	})
}

func Fuzz_Nosy_Counter_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewCounterForced()
		_x1.Inc(_x1)
	})
}

func Fuzz_Nosy_CounterFloat64_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 float64) {
		_x1 := NewCounterFloat64()
		_x1.Dec(_x1)
	})
}

func Fuzz_Nosy_CounterFloat64_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 float64) {
		_x1 := NewCounterFloat64()
		_x1.Inc(_x1)
	})
}

// skipping Fuzz_Nosy_CounterFloat64Snapshot_Count__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterFloat64Snapshot

// skipping Fuzz_Nosy_CounterSnapshot_Count__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.CounterSnapshot

func Fuzz_Nosy_DuplicateMetric_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err DuplicateMetric
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_EWMA_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, alpha float64) {
		_x1 := NewEWMA(alpha)
		_x1.Snapshot()
	})
}

func Fuzz_Nosy_EWMA_Tick__(f *testing.F) {
	f.Fuzz(func(t *testing.T, alpha float64) {
		_x1 := NewEWMA(alpha)
		_x1.Tick()
	})
}

func Fuzz_Nosy_EWMA_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, alpha float64, _x2 int64) {
		_x1 := NewEWMA(alpha)
		_x1.Update(_x2)
	})
}

// skipping Fuzz_Nosy_EWMASnapshot_Rate__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.EWMASnapshot

func Fuzz_Nosy_Gauge_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewGauge()
		_x1.Dec(_x1)
	})
}

func Fuzz_Nosy_Gauge_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewGauge()
		_x1.Inc(_x1)
	})
}

func Fuzz_Nosy_Gauge_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewGauge()
		_x1.Update(_x1)
	})
}

func Fuzz_Nosy_Gauge_UpdateIfGt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewGauge()
		_x1.UpdateIfGt(_x1)
	})
}

func Fuzz_Nosy_GaugeFloat64_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 float64) {
		_x1 := NewGaugeFloat64()
		_x1.Update(_x1)
	})
}

// skipping Fuzz_Nosy_GaugeFloat64Snapshot_Value__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeFloat64Snapshot

func Fuzz_Nosy_GaugeInfo_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 GaugeInfoValue
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1 := NewGaugeInfo()
		_x1.Update(_x1)
	})
}

// skipping Fuzz_Nosy_GaugeInfoSnapshot_Value__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeInfoSnapshot

func Fuzz_Nosy_GaugeInfoValue_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var val GaugeInfoValue
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}

		val.String()
	})
}

// skipping Fuzz_Nosy_GaugeSnapshot_Value__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.GaugeSnapshot

// skipping Fuzz_Nosy_Healthcheck_Check__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Healthcheck

// skipping Fuzz_Nosy_Healthcheck_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Healthcheck

// skipping Fuzz_Nosy_Healthcheck_Healthy__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Healthcheck

// skipping Fuzz_Nosy_Healthcheck_Unhealthy__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Healthcheck

// skipping Fuzz_Nosy_Histogram_Clear__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Histogram

// skipping Fuzz_Nosy_Histogram_Snapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Histogram

// skipping Fuzz_Nosy_Histogram_Update__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Histogram

// skipping Fuzz_Nosy_Logger_Printf__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Logger

func Fuzz_Nosy_Meter_Mark__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 int64) {
		_x1 := NewMeter()
		_x1.Mark(_x1)
	})
}

// skipping Fuzz_Nosy_MeterSnapshot_Count__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

// skipping Fuzz_Nosy_MeterSnapshot_Rate1__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

// skipping Fuzz_Nosy_MeterSnapshot_Rate15__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

// skipping Fuzz_Nosy_MeterSnapshot_Rate5__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

// skipping Fuzz_Nosy_MeterSnapshot_RateMean__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.MeterSnapshot

func Fuzz_Nosy_NilCounter_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Clear()
	})
}

func Fuzz_Nosy_NilCounter_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Dec(i)
	})
}

func Fuzz_Nosy_NilCounter_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Inc(i)
	})
}

func Fuzz_Nosy_NilCounter_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilCounterFloat64_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounterFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Clear()
	})
}

func Fuzz_Nosy_NilCounterFloat64_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounterFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Count()
	})
}

func Fuzz_Nosy_NilCounterFloat64_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounterFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i float64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Dec(i)
	})
}

func Fuzz_Nosy_NilCounterFloat64_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounterFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i float64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Inc(i)
	})
}

func Fuzz_Nosy_NilCounterFloat64_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilCounterFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilEWMA_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilEWMA
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilEWMA_Tick__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilEWMA
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Tick()
	})
}

func Fuzz_Nosy_NilEWMA_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilEWMA
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var n int64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		_x1.Update(n)
	})
}

func Fuzz_Nosy_NilGauge_Dec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGauge
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Dec(i)
	})
}

func Fuzz_Nosy_NilGauge_Inc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGauge
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var i int64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		_x1.Inc(i)
	})
}

func Fuzz_Nosy_NilGauge_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGauge
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilGauge_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGauge
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.Update(v)
	})
}

func Fuzz_Nosy_NilGauge_UpdateIfGt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGauge
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.UpdateIfGt(v)
	})
}

func Fuzz_Nosy_NilGaugeFloat64_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilGaugeFloat64_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v float64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.Update(v)
	})
}

func Fuzz_Nosy_NilGaugeFloat64_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeFloat64
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Value()
	})
}

func Fuzz_Nosy_NilGaugeInfo_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilGaugeInfo_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v GaugeInfoValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.Update(v)
	})
}

func Fuzz_Nosy_NilGaugeInfo_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilGaugeInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Value()
	})
}

func Fuzz_Nosy_NilHealthcheck_Check__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHealthcheck
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Check()
	})
}

func Fuzz_Nosy_NilHealthcheck_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHealthcheck
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Error()
	})
}

func Fuzz_Nosy_NilHealthcheck_Healthy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHealthcheck
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Healthy()
	})
}

// skipping Fuzz_Nosy_NilHealthcheck_Unhealthy__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_NilHistogram_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHistogram
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Clear()
	})
}

func Fuzz_Nosy_NilHistogram_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHistogram
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilHistogram_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilHistogram
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.Update(v)
	})
}

func Fuzz_Nosy_NilMeter_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilMeter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Count()
	})
}

func Fuzz_Nosy_NilMeter_Mark__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilMeter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var n int64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		_x1.Mark(n)
	})
}

func Fuzz_Nosy_NilMeter_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilMeter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilMeter_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilMeter
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Stop()
	})
}

func Fuzz_Nosy_NilResettingTimer_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Count()
	})
}

func Fuzz_Nosy_NilResettingTimer_Max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Max()
	})
}

func Fuzz_Nosy_NilResettingTimer_Mean__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Mean()
	})
}

func Fuzz_Nosy_NilResettingTimer_Min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Min()
	})
}

func Fuzz_Nosy_NilResettingTimer_Percentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 []float64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.Percentiles(_x2)
	})
}

func Fuzz_Nosy_NilResettingTimer_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n NilResettingTimer
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.Snapshot()
	})
}

// skipping Fuzz_Nosy_NilResettingTimer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_NilResettingTimer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Duration
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.Update(_x2)
	})
}

func Fuzz_Nosy_NilResettingTimer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Time
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.UpdateSince(_x2)
	})
}

func Fuzz_Nosy_NilResettingTimer_Values__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilResettingTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Values()
	})
}

func Fuzz_Nosy_NilSample_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilSample
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Clear()
	})
}

func Fuzz_Nosy_NilSample_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilSample
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilSample_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilSample
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v int64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		_x1.Update(v)
	})
}

func Fuzz_Nosy_NilTimer_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Snapshot()
	})
}

func Fuzz_Nosy_NilTimer_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Stop()
	})
}

// skipping Fuzz_Nosy_NilTimer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_NilTimer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Duration
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.Update(_x2)
	})
}

func Fuzz_Nosy_NilTimer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NilTimer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 time.Time
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.UpdateSince(_x2)
	})
}

// skipping Fuzz_Nosy_Registry_Each__ because parameters include func, chan, or unsupported interface: func(string, interface{})

func Fuzz_Nosy_Registry_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 string) {
		_x1 := NewRegistry()
		_x1.Get(_x1)
	})
}

// skipping Fuzz_Nosy_Registry_GetOrRegister__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Registry_Register__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Registry_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, _x1 string) {
		_x1 := NewRegistry()
		_x1.Unregister(_x1)
	})
}

// skipping Fuzz_Nosy_ResettingTimer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_ResettingTimer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 time.Duration
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1 := NewResettingTimer()
		_x1.Update(_x1)
	})
}

func Fuzz_Nosy_ResettingTimer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 time.Time
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1 := NewResettingTimer()
		_x1.UpdateSince(_x1)
	})
}

// skipping Fuzz_Nosy_ResettingTimerSnapshot_Count__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

// skipping Fuzz_Nosy_ResettingTimerSnapshot_Max__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

// skipping Fuzz_Nosy_ResettingTimerSnapshot_Mean__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

// skipping Fuzz_Nosy_ResettingTimerSnapshot_Min__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

// skipping Fuzz_Nosy_ResettingTimerSnapshot_Percentiles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.ResettingTimerSnapshot

func Fuzz_Nosy_Sample_Clear__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int, alpha float64) {
		_x1 := NewExpDecaySample(reservoirSize, alpha)
		_x1.Clear()
	})
}

func Fuzz_Nosy_Sample_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int, alpha float64) {
		_x1 := NewExpDecaySample(reservoirSize, alpha)
		_x1.Snapshot()
	})
}

func Fuzz_Nosy_Sample_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, reservoirSize int, alpha float64, _x3 int64) {
		_x1 := NewExpDecaySample(reservoirSize, alpha)
		_x1.Update(_x3)
	})
}

// skipping Fuzz_Nosy_SampleSnapshot_Count__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Max__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Mean__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Min__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Percentile__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Percentiles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Size__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_StdDev__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Sum__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_SampleSnapshot_Variance__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.SampleSnapshot

// skipping Fuzz_Nosy_Stoppable_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Stoppable

// skipping Fuzz_Nosy_Timer_Time__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_Timer_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 time.Duration
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1 := NewTimer()
		_x1.Update(_x1)
	})
}

func Fuzz_Nosy_Timer_UpdateSince__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 time.Time
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1 := NewTimer()
		_x1.UpdateSince(_x1)
	})
}

func Fuzz_Nosy_counterFloat64Snapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c counterFloat64Snapshot
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Count()
	})
}

func Fuzz_Nosy_counterSnapshot_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c counterSnapshot
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Count()
	})
}

func Fuzz_Nosy_ewmaSnapshot_Rate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ewmaSnapshot
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.Rate()
	})
}

func Fuzz_Nosy_floatsAscendingKeepingIndex_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsAscendingKeepingIndex
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_floatsAscendingKeepingIndex_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsAscendingKeepingIndex
		fill_err = tp.Fill(&s)
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

		s.Less(i, j)
	})
}

func Fuzz_Nosy_floatsAscendingKeepingIndex_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsAscendingKeepingIndex
		fill_err = tp.Fill(&s)
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

		s.Swap(i, j)
	})
}

func Fuzz_Nosy_floatsByIndex_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsByIndex
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_floatsByIndex_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsByIndex
		fill_err = tp.Fill(&s)
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

		s.Less(i, j)
	})
}

func Fuzz_Nosy_floatsByIndex_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s floatsByIndex
		fill_err = tp.Fill(&s)
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

		s.Swap(i, j)
	})
}

func Fuzz_Nosy_gaugeFloat64Snapshot_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g gaugeFloat64Snapshot
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}

		g.Value()
	})
}

func Fuzz_Nosy_gaugeInfoSnapshot_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g gaugeInfoSnapshot
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}

		g.Value()
	})
}

func Fuzz_Nosy_gaugeSnapshot_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g gaugeSnapshot
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}

		g.Value()
	})
}

func Fuzz_Nosy_namedMetric_cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m namedMetric
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var other namedMetric
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		m.cmp(other)
	})
}

func Fuzz_Nosy_CalculatePercentiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var ps []float64
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}

		CalculatePercentiles(values, ps)
	})
}

// skipping Fuzz_Nosy_CaptureDebugGCStats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_CaptureDebugGCStatsOnce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_CollectProcessMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var refresh time.Duration
		fill_err = tp.Fill(&refresh)
		if fill_err != nil {
			return
		}

		CollectProcessMetrics(refresh)
	})
}

// skipping Fuzz_Nosy_Each__ because parameters include func, chan, or unsupported interface: func(string, interface{})

func Fuzz_Nosy_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		Get(name)
	})
}

// skipping Fuzz_Nosy_GetOrRegister__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Graphite__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_GraphiteWithConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c GraphiteConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		GraphiteWithConfig(c)
	})
}

// skipping Fuzz_Nosy_Log__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_LogScaled__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_MustRegister__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_OpenTSDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_OpenTSDBWithConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c OpenTSDBConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		OpenTSDBWithConfig(c)
	})
}

func Fuzz_Nosy_ReadCPUStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stats *CPUStats
		fill_err = tp.Fill(&stats)
		if fill_err != nil {
			return
		}
		if stats == nil {
			return
		}

		ReadCPUStats(stats)
	})
}

// skipping Fuzz_Nosy_RegisterDebugGCStats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_SamplePercentile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}
		var p float64
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		SamplePercentile(values, p)
	})
}

func Fuzz_Nosy_SampleVariance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mean float64
		fill_err = tp.Fill(&mean)
		if fill_err != nil {
			return
		}
		var values []int64
		fill_err = tp.Fill(&values)
		if fill_err != nil {
			return
		}

		SampleVariance(mean, values)
	})
}

// skipping Fuzz_Nosy_Syslog__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		Unregister(name)
	})
}

// skipping Fuzz_Nosy_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_WriteJSON__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_WriteJSONOnce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

// skipping Fuzz_Nosy_WriteOnce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_atomicAddFloat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fbits *atomic.Uint64
		fill_err = tp.Fill(&fbits)
		if fill_err != nil {
			return
		}
		var v float64
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if fbits == nil {
			return
		}

		atomicAddFloat(fbits, v)
	})
}

// skipping Fuzz_Nosy_findPrefix__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Registry

func Fuzz_Nosy_readRuntimeStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *runtimeStats
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		readRuntimeStats(v)
	})
}
