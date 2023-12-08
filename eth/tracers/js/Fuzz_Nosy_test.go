package js

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
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

func Fuzz_Nosy_callframe_GetFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetFrom()
	})
}

func Fuzz_Nosy_callframe_GetGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetGas()
	})
}

func Fuzz_Nosy_callframe_GetInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetInput()
	})
}

func Fuzz_Nosy_callframe_GetTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetTo()
	})
}

func Fuzz_Nosy_callframe_GetType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetType()
	})
}

func Fuzz_Nosy_callframe_GetValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetValue()
	})
}

func Fuzz_Nosy_callframe_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *callframe
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.setupObject()
	})
}

func Fuzz_Nosy_callframeResult_GetError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *callframeResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.GetError()
	})
}

func Fuzz_Nosy_callframeResult_GetGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *callframeResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.GetGasUsed()
	})
}

func Fuzz_Nosy_callframeResult_GetOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *callframeResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.GetOutput()
	})
}

func Fuzz_Nosy_callframeResult_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *callframeResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.setupObject()
	})
}

func Fuzz_Nosy_contractObj_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *contractObj
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.GetAddress()
	})
}

func Fuzz_Nosy_contractObj_GetCaller__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *contractObj
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.GetCaller()
	})
}

func Fuzz_Nosy_contractObj_GetInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *contractObj
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.GetInput()
	})
}

func Fuzz_Nosy_contractObj_GetValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var co *contractObj
		fill_err = tp.Fill(&co)
		if fill_err != nil {
			return
		}
		if co == nil {
			return
		}

		co.GetValue()
	})
}

func Fuzz_Nosy_contractObj_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *contractObj
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.setupObject()
	})
}

// skipping Fuzz_Nosy_dbObj_Exists__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

// skipping Fuzz_Nosy_dbObj_GetBalance__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

// skipping Fuzz_Nosy_dbObj_GetCode__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

// skipping Fuzz_Nosy_dbObj_GetNonce__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

// skipping Fuzz_Nosy_dbObj_GetState__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

func Fuzz_Nosy_dbObj_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var do *dbObj
		fill_err = tp.Fill(&do)
		if fill_err != nil {
			return
		}
		if do == nil {
			return
		}

		do.setupObject()
	})
}

// skipping Fuzz_Nosy_jsTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_jsTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var typ vm.OpCode
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if t1 == nil || value == nil {
			return
		}

		t1.CaptureEnter(typ, from, to, input, gas, value)
	})
}

// skipping Fuzz_Nosy_jsTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_jsTracer_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_jsTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var env *vm.EVM
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var from common.Address
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to common.Address
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var create bool
		fill_err = tp.Fill(&create)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if t1 == nil || env == nil || value == nil {
			return
		}

		t1.CaptureStart(env, from, to, create, input, gas, value)
	})
}

// skipping Fuzz_Nosy_jsTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_jsTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_jsTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_jsTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.GetResult()
	})
}

// skipping Fuzz_Nosy_jsTracer_Stop__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_jsTracer_onError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_jsTracer_setBuiltinFunctions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.setBuiltinFunctions()
	})
}

func Fuzz_Nosy_jsTracer_setTypeConverters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *jsTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.setTypeConverters()
	})
}

func Fuzz_Nosy_memoryObj_GetUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mo *memoryObj
		fill_err = tp.Fill(&mo)
		if fill_err != nil {
			return
		}
		var addr int64
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if mo == nil {
			return
		}

		mo.GetUint(addr)
	})
}

func Fuzz_Nosy_memoryObj_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mo *memoryObj
		fill_err = tp.Fill(&mo)
		if fill_err != nil {
			return
		}
		if mo == nil {
			return
		}

		mo.Length()
	})
}

func Fuzz_Nosy_memoryObj_Slice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mo *memoryObj
		fill_err = tp.Fill(&mo)
		if fill_err != nil {
			return
		}
		var begin int64
		fill_err = tp.Fill(&begin)
		if fill_err != nil {
			return
		}
		var end int64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if mo == nil {
			return
		}

		mo.Slice(begin, end)
	})
}

func Fuzz_Nosy_memoryObj_getUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mo *memoryObj
		fill_err = tp.Fill(&mo)
		if fill_err != nil {
			return
		}
		var addr int64
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if mo == nil {
			return
		}

		mo.getUint(addr)
	})
}

func Fuzz_Nosy_memoryObj_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *memoryObj
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.setupObject()
	})
}

func Fuzz_Nosy_memoryObj_slice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mo *memoryObj
		fill_err = tp.Fill(&mo)
		if fill_err != nil {
			return
		}
		var begin int64
		fill_err = tp.Fill(&begin)
		if fill_err != nil {
			return
		}
		var end int64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if mo == nil {
			return
		}

		mo.slice(begin, end)
	})
}

func Fuzz_Nosy_opObj_IsPush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *opObj
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.IsPush()
	})
}

func Fuzz_Nosy_opObj_ToNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *opObj
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.ToNumber()
	})
}

func Fuzz_Nosy_opObj_ToString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *opObj
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.ToString()
	})
}

func Fuzz_Nosy_opObj_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var o *opObj
		fill_err = tp.Fill(&o)
		if fill_err != nil {
			return
		}
		if o == nil {
			return
		}

		o.setupObject()
	})
}

func Fuzz_Nosy_stackObj_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stackObj
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Length()
	})
}

func Fuzz_Nosy_stackObj_Peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stackObj
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Peek(idx)
	})
}

func Fuzz_Nosy_stackObj_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stackObj
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.peek(idx)
	})
}

func Fuzz_Nosy_stackObj_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stackObj
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.setupObject()
	})
}

func Fuzz_Nosy_steplog_GetCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetCost()
	})
}

func Fuzz_Nosy_steplog_GetDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetDepth()
	})
}

func Fuzz_Nosy_steplog_GetError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetError()
	})
}

func Fuzz_Nosy_steplog_GetGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetGas()
	})
}

func Fuzz_Nosy_steplog_GetPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetPC()
	})
}

func Fuzz_Nosy_steplog_GetRefund__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.GetRefund()
	})
}

func Fuzz_Nosy_steplog_setupObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *steplog
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.setupObject()
	})
}

// skipping Fuzz_Nosy_fromBuf__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value
