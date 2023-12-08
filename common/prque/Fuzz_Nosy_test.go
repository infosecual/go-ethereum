package prque

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
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
	
	func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Empty()
	})
}

// skipping Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_MultiPop__ because parameters include func, chan, or unsupported interface: func(data V, priority P) bool

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Pop()
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_PopItem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.PopItem()
	})
}

// skipping Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Push__ because parameters include func, chan, or unsupported interface: V

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Refresh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Refresh()
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Remove(index)
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Reset()
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Size()
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.Update(index)
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_peekIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.peekIndex()
	})
}

func Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_refresh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var q *LazyQueue[P, V]
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var now mclock.AbsTime
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}
	if q == nil {
		return
	}

	q.refresh(now)
	})
}

// skipping Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_setIndex0__ because parameters include func, chan, or unsupported interface: V

// skipping Fuzz_Nosy_LazyQueue[P constraints.Ordered, V any]_setIndex1__ because parameters include func, chan, or unsupported interface: V

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Empty()
	})
}

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Peek()
	})
}

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Pop()
	})
}

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_PopItem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.PopItem()
	})
}

// skipping Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Push__ because parameters include func, chan, or unsupported interface: V

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Remove(i)
	})
}

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Reset()
	})
}

func Fuzz_Nosy_Prque[P constraints.Ordered, V any]_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var p *Prque[P, V]
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
	if p == nil {
		return
	}

	p.Size()
	})
}

func Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *sstack[P, V]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Len()
	})
}

func Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *sstack[P, V]
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
	if s == nil {
		return
	}

	s.Less(i, j)
	})
}

func Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *sstack[P, V]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Pop()
	})
}

// skipping Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Push__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *sstack[P, V]
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Reset()
	})
}

func Fuzz_Nosy_sstack[P constraints.Ordered, V any]_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *sstack[P, V]
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
	if s == nil {
		return
	}

	s.Swap(i, j)
	})
}

