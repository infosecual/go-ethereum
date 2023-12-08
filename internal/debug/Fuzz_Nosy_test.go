package debug

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

func Fuzz_Nosy_HandlerT_BlockProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nsec uint
		fill_err = tp.Fill(&nsec)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.BlockProfile(file, nsec)
	})
}

func Fuzz_Nosy_HandlerT_CpuProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nsec uint
		fill_err = tp.Fill(&nsec)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.CpuProfile(file, nsec)
	})
}

func Fuzz_Nosy_HandlerT_FreeOSMemory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.FreeOSMemory()
	})
}

func Fuzz_Nosy_HandlerT_GcStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.GcStats()
	})
}

func Fuzz_Nosy_HandlerT_GoTrace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nsec uint
		fill_err = tp.Fill(&nsec)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.GoTrace(file, nsec)
	})
}

func Fuzz_Nosy_HandlerT_MemStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.MemStats()
	})
}

func Fuzz_Nosy_HandlerT_MutexProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nsec uint
		fill_err = tp.Fill(&nsec)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.MutexProfile(file, nsec)
	})
}

func Fuzz_Nosy_HandlerT_SetBlockProfileRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var rate int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SetBlockProfileRate(rate)
	})
}

func Fuzz_Nosy_HandlerT_SetGCPercent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var v int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SetGCPercent(v)
	})
}

func Fuzz_Nosy_HandlerT_SetMutexProfileFraction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var rate int
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SetMutexProfileFraction(rate)
	})
}

func Fuzz_Nosy_HandlerT_Stacks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var filter *string
		fill_err = tp.Fill(&filter)
		if fill_err != nil {
			return
		}
		if _x1 == nil || filter == nil {
			return
		}

		_x1.Stacks(filter)
	})
}

func Fuzz_Nosy_HandlerT_StartCPUProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.StartCPUProfile(file)
	})
}

func Fuzz_Nosy_HandlerT_StartGoTrace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.StartGoTrace(file)
	})
}

func Fuzz_Nosy_HandlerT_StopCPUProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.StopCPUProfile()
	})
}

func Fuzz_Nosy_HandlerT_StopGoTrace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *HandlerT
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.StopGoTrace()
	})
}

func Fuzz_Nosy_HandlerT_Verbosity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var level int
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Verbosity(level)
	})
}

func Fuzz_Nosy_HandlerT_Vmodule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var pattern string
		fill_err = tp.Fill(&pattern)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Vmodule(pattern)
	})
}

func Fuzz_Nosy_HandlerT_WriteBlockProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.WriteBlockProfile(file)
	})
}

func Fuzz_Nosy_HandlerT_WriteMemProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.WriteMemProfile(file)
	})
}

func Fuzz_Nosy_HandlerT_WriteMutexProfile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *HandlerT
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.WriteMutexProfile(file)
	})
}

// skipping Fuzz_Nosy_LoudPanic__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_StartPProf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, address string, withMetrics bool) {
		StartPProf(address, withMetrics)
	})
}

func Fuzz_Nosy_expandHome__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		expandHome(p)
	})
}
