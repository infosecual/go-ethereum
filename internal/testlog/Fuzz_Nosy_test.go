package testlog

import (
	"context"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"golang.org/x/exp/slog"
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

func Fuzz_Nosy_bufHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bufHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var lvl slog.Level
		fill_err = tp.Fill(&lvl)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Enabled(_x2, lvl)
	})
}

func Fuzz_Nosy_bufHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bufHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Handle(_x2, r)
	})
}

func Fuzz_Nosy_bufHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bufHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var attrs []slog.Attr
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.WithAttrs(attrs)
	})
}

func Fuzz_Nosy_bufHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bufHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.WithGroup(_x2)
	})
}

func Fuzz_Nosy_bufHandler_terminalFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bufHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.terminalFormat(r)
	})
}

// skipping Fuzz_Nosy_logger_Crit__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Debug__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_logger_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var level slog.Level
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Enabled(ctx, level)
	})
}

// skipping Fuzz_Nosy_logger_Error__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Info__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Log__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_New__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Trace__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Warn__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_With__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Write__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_logger_flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *logger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.flush()
	})
}
