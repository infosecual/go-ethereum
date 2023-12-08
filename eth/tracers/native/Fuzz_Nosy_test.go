package native

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/tracers"
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

func Fuzz_Nosy_account_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_account_exists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.exists()
	})
}

func Fuzz_Nosy_callFrame_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *callFrame
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnmarshalJSON(input)
	})
}

// skipping Fuzz_Nosy_callFrame_processOutput__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_callTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_callTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *callTracer
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

// skipping Fuzz_Nosy_callTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_callTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *callTracer
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

// skipping Fuzz_Nosy_callTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_callTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *callTracer
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

func Fuzz_Nosy_callTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *callTracer
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

func Fuzz_Nosy_callTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *callTracer
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

// skipping Fuzz_Nosy_callTracer_Stop__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_flatCallAction_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *flatCallAction
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_flatCallResult_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *flatCallResult
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UnmarshalJSON(input)
	})
}

// skipping Fuzz_Nosy_flatCallTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_flatCallTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
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

// skipping Fuzz_Nosy_flatCallTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_flatCallTracer_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_flatCallTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
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

// skipping Fuzz_Nosy_flatCallTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_flatCallTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
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

func Fuzz_Nosy_flatCallTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
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

func Fuzz_Nosy_flatCallTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
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

// skipping Fuzz_Nosy_flatCallTracer_Stop__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_flatCallTracer_isPrecompiled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *flatCallTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.isPrecompiled(addr)
	})
}

func Fuzz_Nosy_fourByteTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *fourByteTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var op vm.OpCode
		fill_err = tp.Fill(&op)
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

		t1.CaptureEnter(op, from, to, input, gas, value)
	})
}

func Fuzz_Nosy_fourByteTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *fourByteTracer
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

func Fuzz_Nosy_fourByteTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *fourByteTracer
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

// skipping Fuzz_Nosy_fourByteTracer_Stop__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_fourByteTracer_isPrecompiled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *fourByteTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.isPrecompiled(addr)
	})
}

func Fuzz_Nosy_fourByteTracer_store__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *fourByteTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.store(id, size)
	})
}

// skipping Fuzz_Nosy_muxTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_muxTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *muxTracer
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

// skipping Fuzz_Nosy_muxTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_muxTracer_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_muxTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *muxTracer
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

// skipping Fuzz_Nosy_muxTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_muxTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *muxTracer
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

func Fuzz_Nosy_muxTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *muxTracer
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

func Fuzz_Nosy_muxTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *muxTracer
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

// skipping Fuzz_Nosy_muxTracer_Stop__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_noopTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_noopTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *noopTracer
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

// skipping Fuzz_Nosy_noopTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_noopTracer_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_noopTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *noopTracer
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

// skipping Fuzz_Nosy_noopTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_noopTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *noopTracer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_noopTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *noopTracer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_noopTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *noopTracer
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

// skipping Fuzz_Nosy_noopTracer_Stop__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_prestateTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_prestateTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
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

// skipping Fuzz_Nosy_prestateTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_prestateTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
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

func Fuzz_Nosy_prestateTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
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

func Fuzz_Nosy_prestateTracer_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
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

// skipping Fuzz_Nosy_prestateTracer_Stop__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_prestateTracer_lookupAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupAccount(addr)
	})
}

func Fuzz_Nosy_prestateTracer_lookupStorage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *prestateTracer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var key common.Hash
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupStorage(addr, key)
	})
}

func Fuzz_Nosy_account_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.MarshalJSON()
	})
}

func Fuzz_Nosy_callFrame_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c callFrame
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.MarshalJSON()
	})
}

func Fuzz_Nosy_callFrame_TypeString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 callFrame
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.TypeString()
	})
}

func Fuzz_Nosy_callFrame_failed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 callFrame
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.failed()
	})
}

func Fuzz_Nosy_flatCallAction_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 flatCallAction
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.MarshalJSON()
	})
}

func Fuzz_Nosy_flatCallResult_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 flatCallResult
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.MarshalJSON()
	})
}

func Fuzz_Nosy_bytesToHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		bytesToHex(s)
	})
}

func Fuzz_Nosy_childTraceAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a []int
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		childTraceAddress(a, i)
	})
}

func Fuzz_Nosy_clearFailedLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cf *callFrame
		fill_err = tp.Fill(&cf)
		if fill_err != nil {
			return
		}
		var parentFailed bool
		fill_err = tp.Fill(&parentFailed)
		if fill_err != nil {
			return
		}
		if cf == nil {
			return
		}

		clearFailedLogs(cf, parentFailed)
	})
}

func Fuzz_Nosy_convertErrorToParity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var call *flatCallFrame
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if call == nil {
			return
		}

		convertErrorToParity(call)
	})
}

func Fuzz_Nosy_fillCallFrameFromContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var callFrame *flatCallFrame
		fill_err = tp.Fill(&callFrame)
		if fill_err != nil {
			return
		}
		var ctx *tracers.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if callFrame == nil || ctx == nil {
			return
		}

		fillCallFrameFromContext(callFrame, ctx)
	})
}

func Fuzz_Nosy_flatFromNested__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var input *callFrame
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var traceAddress []int
		fill_err = tp.Fill(&traceAddress)
		if fill_err != nil {
			return
		}
		var convertErrs bool
		fill_err = tp.Fill(&convertErrs)
		if fill_err != nil {
			return
		}
		var ctx *tracers.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if input == nil || ctx == nil {
			return
		}

		flatFromNested(input, traceAddress, convertErrs, ctx)
	})
}
