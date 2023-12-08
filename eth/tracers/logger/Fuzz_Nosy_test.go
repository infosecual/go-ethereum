package logger

import (
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_AccessListTracer_AccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
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
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}

		a := NewAccessListTracer(acl, from, to, precompiles)
		a.AccessList()
	})
}

// skipping Fuzz_Nosy_AccessListTracer_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_AccessListTracer_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
		if fill_err != nil {
			return
		}
		var f2 common.Address
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var typ vm.OpCode
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var f6 common.Address
		fill_err = tp.Fill(&f6)
		if fill_err != nil {
			return
		}
		var t7 common.Address
		fill_err = tp.Fill(&t7)
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
		if value == nil {
			return
		}

		_x1 := NewAccessListTracer(acl, f2, t3, precompiles)
		_x1.CaptureEnter(typ, f6, t7, input, gas, value)
	})
}

// skipping Fuzz_Nosy_AccessListTracer_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_AccessListTracer_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_AccessListTracer_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
		if fill_err != nil {
			return
		}
		var f2 common.Address
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		var t3 common.Address
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var env *vm.EVM
		fill_err = tp.Fill(&env)
		if fill_err != nil {
			return
		}
		var f6 common.Address
		fill_err = tp.Fill(&f6)
		if fill_err != nil {
			return
		}
		var t7 common.Address
		fill_err = tp.Fill(&t7)
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
		if env == nil || value == nil {
			return
		}

		a := NewAccessListTracer(acl, f2, t3, precompiles)
		a.CaptureStart(env, f6, t7, create, input, gas, value)
	})
}

// skipping Fuzz_Nosy_AccessListTracer_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_AccessListTracer_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
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
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}

		_x1 := NewAccessListTracer(acl, from, to, precompiles)
		_x1.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_AccessListTracer_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
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
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}

		_x1 := NewAccessListTracer(acl, from, to, precompiles)
		_x1.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_AccessListTracer_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var acl types.AccessList
		fill_err = tp.Fill(&acl)
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
		var precompiles []common.Address
		fill_err = tp.Fill(&precompiles)
		if fill_err != nil {
			return
		}
		var other *AccessListTracer
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if other == nil {
			return
		}

		a := NewAccessListTracer(acl, from, to, precompiles)
		a.Equal(other)
	})
}

// skipping Fuzz_Nosy_JSONLogger_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_JSONLogger_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
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
		if cfg == nil || value == nil {
			return
		}

		l := NewJSONLogger(cfg, writer)
		l.CaptureEnter(typ, from, to, input, gas, value)
	})
}

// skipping Fuzz_Nosy_JSONLogger_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_JSONLogger_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_JSONLogger_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
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
		if cfg == nil || env == nil || value == nil {
			return
		}

		l := NewJSONLogger(cfg, writer)
		l.CaptureStart(env, from, to, create, input, gas, value)
	})
}

// skipping Fuzz_Nosy_JSONLogger_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_JSONLogger_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewJSONLogger(cfg, writer)
		l.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_JSONLogger_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewJSONLogger(cfg, writer)
		l.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_StructLog_ErrorString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StructLog
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ErrorString()
	})
}

func Fuzz_Nosy_StructLog_OpName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StructLog
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OpName()
	})
}

func Fuzz_Nosy_StructLog_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *StructLog
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnmarshalJSON(input)
	})
}

// skipping Fuzz_Nosy_StructLogger_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StructLogger_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
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
		if cfg == nil || value == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.CaptureEnter(typ, from, to, input, gas, value)
	})
}

// skipping Fuzz_Nosy_StructLogger_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_StructLogger_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StructLogger_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
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
		if cfg == nil || env == nil || value == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.CaptureStart(env, from, to, create, input, gas, value)
	})
}

// skipping Fuzz_Nosy_StructLogger_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StructLogger_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_StructLogger_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_StructLogger_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.Error()
	})
}

func Fuzz_Nosy_StructLogger_GetResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.GetResult()
	})
}

func Fuzz_Nosy_StructLogger_Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.Output()
	})
}

func Fuzz_Nosy_StructLogger_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.Reset()
	})
}

// skipping Fuzz_Nosy_StructLogger_Stop__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StructLogger_StructLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		l := NewStructLogger(cfg)
		l.StructLogs()
	})
}

// skipping Fuzz_Nosy_mdLogger_CaptureEnd__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_mdLogger_CaptureEnter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
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
		if cfg == nil || value == nil {
			return
		}

		t1 := NewMarkdownLogger(cfg, writer)
		t1.CaptureEnter(typ, from, to, input, gas, value)
	})
}

// skipping Fuzz_Nosy_mdLogger_CaptureExit__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_mdLogger_CaptureFault__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_mdLogger_CaptureStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
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
		if cfg == nil || env == nil || value == nil {
			return
		}

		t1 := NewMarkdownLogger(cfg, writer)
		t1.CaptureStart(env, from, to, create, input, gas, value)
	})
}

// skipping Fuzz_Nosy_mdLogger_CaptureState__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_mdLogger_CaptureTxEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var restGas uint64
		fill_err = tp.Fill(&restGas)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewMarkdownLogger(cfg, writer)
		_x1.CaptureTxEnd(restGas)
	})
}

func Fuzz_Nosy_mdLogger_CaptureTxStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		_x1 := NewMarkdownLogger(cfg, writer)
		_x1.CaptureTxStart(gasLimit)
	})
}

func Fuzz_Nosy_Storage_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Storage
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Copy()
	})
}

func Fuzz_Nosy_StructLog_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s StructLog
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_accessList_addAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.addAddress(address)
	})
}

func Fuzz_Nosy_accessList_addSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var slot common.Hash
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.addSlot(address, slot)
	})
}

func Fuzz_Nosy_accessList_equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var other accessList
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		al := newAccessList()
		al.equal(other)
	})
}

func Fuzz_Nosy_WriteLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var logs []*types.Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		WriteLogs(writer, logs)
	})
}

func Fuzz_Nosy_WriteTrace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var writer io.Writer
		fill_err = tp.Fill(&writer)
		if fill_err != nil {
			return
		}
		var logs []StructLog
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		WriteTrace(writer, logs)
	})
}

func Fuzz_Nosy_formatLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var logs []StructLog
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		formatLogs(logs)
	})
}
