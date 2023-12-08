package event

import (
	"testing"
	"github.com/ethereum/go-ethereum"
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
	
	// skipping Fuzz_Nosy_Feed_Send__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Feed_Subscribe__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Feed_init__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_Feed_remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *Feed
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var sub *feedSub
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if f1 == nil|| sub == nil {
		return
	}

	f1.remove(sub)
	})
}

// skipping Fuzz_Nosy_FeedOf[T any]_Send__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_FeedOf[T any]_Subscribe__ because parameters include func, chan, or unsupported interface: chan<- T

func Fuzz_Nosy_FeedOf[T any]_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *FeedOf[T]
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
	if f1 == nil {
		return
	}

	f1.init()
	})
}

func Fuzz_Nosy_FeedOf[T any]_remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var f1 *FeedOf[T]
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var sub *feedOfSub[T]
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if f1 == nil|| sub == nil {
		return
	}

	f1.remove(sub)
	})
}

func Fuzz_Nosy_SubscriptionScope_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sc *SubscriptionScope
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
	if sc == nil {
		return
	}

	sc.Close()
	})
}

func Fuzz_Nosy_SubscriptionScope_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sc *SubscriptionScope
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
	if sc == nil {
		return
	}

	sc.Count()
	})
}

// skipping Fuzz_Nosy_SubscriptionScope_Track__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/event.Subscription

// skipping Fuzz_Nosy_TypeMux_Post__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_TypeMux_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
	if mux == nil {
		return
	}

	mux.Stop()
	})
}

// skipping Fuzz_Nosy_TypeMux_Subscribe__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_TypeMux_del__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
		var s *TypeMuxSubscription
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if mux == nil|| s == nil {
		return
	}

	mux.del(s)
	})
}

func Fuzz_Nosy_TypeMuxSubscription_Chan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
	if mux == nil {
		return
	}

	s := newsub(mux)
	s.Chan()
	})
}

func Fuzz_Nosy_TypeMuxSubscription_Closed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
	if mux == nil {
		return
	}

	s := newsub(mux)
	s.Closed()
	})
}

func Fuzz_Nosy_TypeMuxSubscription_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
	if mux == nil {
		return
	}

	s := newsub(mux)
	s.Unsubscribe()
	})
}

func Fuzz_Nosy_TypeMuxSubscription_closewait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
	if mux == nil {
		return
	}

	s := newsub(mux)
	s.closewait()
	})
}

func Fuzz_Nosy_TypeMuxSubscription_deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var mux *TypeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
		var e2 *TypeMuxEvent
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
	if mux == nil|| e2 == nil {
		return
	}

	s := newsub(mux)
	s.deliver(e2)
	})
}

func Fuzz_Nosy_feedOfSub[T any]_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sub *feedOfSub[T]
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if sub == nil {
		return
	}

	sub.Err()
	})
}

func Fuzz_Nosy_feedOfSub[T any]_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sub *feedOfSub[T]
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if sub == nil {
		return
	}

	sub.Unsubscribe()
	})
}

func Fuzz_Nosy_feedSub_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sub *feedSub
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if sub == nil {
		return
	}

	sub.Err()
	})
}

func Fuzz_Nosy_feedSub_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var sub *feedSub
		fill_err = tp.Fill(&sub)
		if fill_err != nil {
			return
		}
	if sub == nil {
		return
	}

	sub.Unsubscribe()
	})
}

func Fuzz_Nosy_funcSub_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *funcSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Err()
	})
}

func Fuzz_Nosy_funcSub_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *funcSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Unsubscribe()
	})
}

func Fuzz_Nosy_resubscribeSub_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *resubscribeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Err()
	})
}

func Fuzz_Nosy_resubscribeSub_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *resubscribeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Unsubscribe()
	})
}

func Fuzz_Nosy_resubscribeSub_backoffWait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *resubscribeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.backoffWait()
	})
}

func Fuzz_Nosy_resubscribeSub_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *resubscribeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.loop()
	})
}

func Fuzz_Nosy_resubscribeSub_subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *resubscribeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.subscribe()
	})
}

// skipping Fuzz_Nosy_resubscribeSub_waitForError__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/event.Subscription

func Fuzz_Nosy_scopeSub_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *scopeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Err()
	})
}

func Fuzz_Nosy_scopeSub_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var s *scopeSub
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
	if s == nil {
		return
	}

	s.Unsubscribe()
	})
}

// skipping Fuzz_Nosy_Subscription_Err__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/event.Subscription

// skipping Fuzz_Nosy_Subscription_Unsubscribe__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/event.Subscription

func Fuzz_Nosy_caseList_deactivate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cs caseList
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

	cs.deactivate(index)
	})
}

func Fuzz_Nosy_caseList_delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var cs caseList
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

	cs.delete(index)
	})
}

// skipping Fuzz_Nosy_caseList_find__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_feedTypeError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e feedTypeError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

	e.Error()
	})
}

func Fuzz_Nosy_find__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var slice []*TypeMuxSubscription
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}
		var item *TypeMuxSubscription
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}
	if item == nil {
		return
	}

	find(slice, item)
	})
}

func Fuzz_Nosy_posdelete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var slice []*TypeMuxSubscription
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}
		var pos int
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

	posdelete(slice, pos)
	})
}

