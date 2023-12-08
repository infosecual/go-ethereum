package lru

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
	
	// skipping Fuzz_Nosy_BasicLRU[K comparable, V any]_Add__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_BasicLRU[K comparable, V any]_Contains__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_BasicLRU[K comparable, V any]_Get__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_BasicLRU[K comparable, V any]_GetOldest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewBasicLRU(capacity)
	c.GetOldest()
	})
}

func Fuzz_Nosy_BasicLRU[K comparable, V any]_Keys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewBasicLRU(capacity)
	c.Keys()
	})
}

func Fuzz_Nosy_BasicLRU[K comparable, V any]_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewBasicLRU(capacity)
	c.Len()
	})
}

// skipping Fuzz_Nosy_BasicLRU[K comparable, V any]_Peek__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_BasicLRU[K comparable, V any]_Purge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewBasicLRU(capacity)
	c.Purge()
	})
}

// skipping Fuzz_Nosy_BasicLRU[K comparable, V any]_Remove__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_BasicLRU[K comparable, V any]_RemoveOldest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewBasicLRU(capacity)
	c.RemoveOldest()
	})
}

// skipping Fuzz_Nosy_Cache[K comparable, V any]_Add__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_Cache[K comparable, V any]_Contains__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_Cache[K comparable, V any]_Get__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_Cache[K comparable, V any]_Keys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewCache(capacity)
	c.Keys()
	})
}

func Fuzz_Nosy_Cache[K comparable, V any]_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewCache(capacity)
	c.Len()
	})
}

// skipping Fuzz_Nosy_Cache[K comparable, V any]_Peek__ because parameters include func, chan, or unsupported interface: K

func Fuzz_Nosy_Cache[K comparable, V any]_Purge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, capacity int) {
	c := NewCache(capacity)
	c.Purge()
	})
}

// skipping Fuzz_Nosy_Cache[K comparable, V any]_Remove__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_SizeConstrainedCache[K comparable, V blobType]_Add__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_SizeConstrainedCache[K comparable, V blobType]_Get__ because parameters include func, chan, or unsupported interface: K

// skipping Fuzz_Nosy_list[T any]_appendTo__ because parameters include func, chan, or unsupported interface: []T

func Fuzz_Nosy_list[T any]_moveToFront__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *listElem[T]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	l := newList()
	l.moveToFront(e)
	})
}

func Fuzz_Nosy_list[T any]_pushElem__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *listElem[T]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	l := newList()
	l.pushElem(e)
	})
}

func Fuzz_Nosy_list[T any]_remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
				var e *listElem[T]
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
	if e == nil {
		return
	}

	l := newList()
	l.remove(e)
	})
}

