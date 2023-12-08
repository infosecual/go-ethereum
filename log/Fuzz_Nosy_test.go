package log

import (
	"bytes"
	"context"
	"io"
	"math/big"
	"testing"
	"time"

	"github.com/holiman/uint256"
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

func Fuzz_Nosy_GlogHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		h.Enabled(ctx, lvl)
	})
}

func Fuzz_Nosy_GlogHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
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

func Fuzz_Nosy_GlogHandler_Verbosity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var level slog.Level
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Verbosity(level)
	})
}

func Fuzz_Nosy_GlogHandler_Vmodule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ruleset string
		fill_err = tp.Fill(&ruleset)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Vmodule(ruleset)
	})
}

func Fuzz_Nosy_GlogHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
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

func Fuzz_Nosy_GlogHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *GlogHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.WithGroup(name)
	})
}

func Fuzz_Nosy_TerminalHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var _x3 context.Context
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var level slog.Level
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.Enabled(_x3, level)
	})
}

func Fuzz_Nosy_TerminalHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var _x3 context.Context
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.Handle(_x3, r)
	})
}

func Fuzz_Nosy_TerminalHandler_ResetFieldPadding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}

		t1 := NewTerminalHandler(wr, useColor)
		t1.ResetFieldPadding()
	})
}

func Fuzz_Nosy_TerminalHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var attrs []slog.Attr
		fill_err = tp.Fill(&attrs)
		if fill_err != nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.WithAttrs(attrs)
	})
}

func Fuzz_Nosy_TerminalHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.WithGroup(name)
	})
}

func Fuzz_Nosy_TerminalHandler_format__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var usecolor bool
		fill_err = tp.Fill(&usecolor)
		if fill_err != nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.format(buf, r, usecolor)
	})
}

func Fuzz_Nosy_TerminalHandler_formatAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wr io.Writer
		fill_err = tp.Fill(&wr)
		if fill_err != nil {
			return
		}
		var useColor bool
		fill_err = tp.Fill(&useColor)
		if fill_err != nil {
			return
		}
		var buf *bytes.Buffer
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		var r slog.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var color string
		fill_err = tp.Fill(&color)
		if fill_err != nil {
			return
		}
		if buf == nil {
			return
		}

		h := NewTerminalHandler(wr, useColor)
		h.formatAttributes(buf, r, color)
	})
}

func Fuzz_Nosy_discardHandler_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *discardHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var level slog.Level
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Enabled(_x2, level)
	})
}

func Fuzz_Nosy_discardHandler_Handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *discardHandler
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

func Fuzz_Nosy_discardHandler_WithAttrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *discardHandler
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

func Fuzz_Nosy_discardHandler_WithGroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *discardHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.WithGroup(name)
	})
}

func Fuzz_Nosy_leveler_Level__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *leveler
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Level()
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

// skipping Fuzz_Nosy_logger_Log__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_New__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Trace__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Warn__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_logger_With__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_logger_Write__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_Logger_Crit__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Debug__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Logger_Enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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

		_x1 := Root()
		_x1.Enabled(ctx, level)
	})
}

// skipping Fuzz_Nosy_Logger_Error__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Info__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Log__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_New__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Trace__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Warn__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_With__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Logger_Write__ because parameters include func, chan, or unsupported interface: []any

// skipping Fuzz_Nosy_TerminalStringer_TerminalString__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.TerminalStringer

// skipping Fuzz_Nosy_Crit__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Debug__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Error__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_FormatLogfmtUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n uint64) {
		FormatLogfmtUint64(n)
	})
}

func Fuzz_Nosy_FormatSlogValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v slog.Value
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var tmp []byte
		fill_err = tp.Fill(&tmp)
		if fill_err != nil {
			return
		}

		FormatSlogValue(v, tmp)
	})
}

// skipping Fuzz_Nosy_Info__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_LevelAlignedString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l slog.Level
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		LevelAlignedString(l)
	})
}

func Fuzz_Nosy_LevelString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l slog.Level
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		LevelString(l)
	})
}

// skipping Fuzz_Nosy_SetDefault__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/log.Logger

// skipping Fuzz_Nosy_Trace__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Warn__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_appendBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var n *big.Int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		appendBigInt(dst, n)
	})
}

func Fuzz_Nosy_appendEscapeString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, s string) {
		appendEscapeString(dst, s)
	})
}

func Fuzz_Nosy_appendInt64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, n int64) {
		appendInt64(dst, n)
	})
}

func Fuzz_Nosy_appendU256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var n *uint256.Int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		appendU256(dst, n)
	})
}

func Fuzz_Nosy_appendUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, n uint64, neg bool) {
		appendUint64(dst, n, neg)
	})
}

func Fuzz_Nosy_escapeMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		escapeMessage(s)
	})
}

func Fuzz_Nosy_writePosIntWidth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var width int
		fill_err = tp.Fill(&width)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		writePosIntWidth(b, i, width)
	})
}

func Fuzz_Nosy_writeTimeTermFormat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var buf *bytes.Buffer
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if buf == nil {
			return
		}

		writeTimeTermFormat(buf, t2)
	})
}
