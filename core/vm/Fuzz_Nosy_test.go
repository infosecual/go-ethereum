package vm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
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

func Fuzz_Nosy_Contract_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Address()
	})
}

func Fuzz_Nosy_Contract_AsDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.AsDelegate()
	})
}

func Fuzz_Nosy_Contract_Caller__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Caller()
	})
}

func Fuzz_Nosy_Contract_GetOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetOp(n)
	})
}

func Fuzz_Nosy_Contract_SetCallCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr *common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var code []byte
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if c == nil || addr == nil {
			return
		}

		c.SetCallCode(addr, hash, code)
	})
}

func Fuzz_Nosy_Contract_SetCodeOptionalHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var addr *common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var codeAndHash *codeAndHash
		fill_err = tp.Fill(&codeAndHash)
		if fill_err != nil {
			return
		}
		if c == nil || addr == nil || codeAndHash == nil {
			return
		}

		c.SetCodeOptionalHash(addr, codeAndHash)
	})
}

func Fuzz_Nosy_Contract_UseGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UseGas(gas)
	})
}

func Fuzz_Nosy_Contract_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Value()
	})
}

func Fuzz_Nosy_Contract_isCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var udest uint64
		fill_err = tp.Fill(&udest)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.isCode(udest)
	})
}

func Fuzz_Nosy_Contract_validJumpdest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Contract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var dest *uint256.Int
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		if c == nil || dest == nil {
			return
		}

		c.validJumpdest(dest)
	})
}

// skipping Fuzz_Nosy_EVM_Call__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

// skipping Fuzz_Nosy_EVM_CallCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

func Fuzz_Nosy_EVM_Cancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		if evm == nil {
			return
		}

		evm.Cancel()
	})
}

func Fuzz_Nosy_EVM_Cancelled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		if evm == nil {
			return
		}

		evm.Cancelled()
	})
}

func Fuzz_Nosy_EVM_ChainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		if evm == nil {
			return
		}

		evm.ChainConfig()
	})
}

// skipping Fuzz_Nosy_EVM_Create__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

// skipping Fuzz_Nosy_EVM_Create2__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

// skipping Fuzz_Nosy_EVM_DelegateCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

func Fuzz_Nosy_EVM_Interpreter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		if evm == nil {
			return
		}

		evm.Interpreter()
	})
}

// skipping Fuzz_Nosy_EVM_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_EVM_StaticCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

// skipping Fuzz_Nosy_EVM_create__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

func Fuzz_Nosy_EVM_precompile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if evm == nil {
			return
		}

		evm.precompile(addr)
	})
}

func Fuzz_Nosy_EVMInterpreter_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var readOnly bool
		fill_err = tp.Fill(&readOnly)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil {
			return
		}

		in := NewEVMInterpreter(evm)
		in.Run(contract, input, readOnly)
	})
}

func Fuzz_Nosy_ErrInvalidOpCode_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ErrInvalidOpCode
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_ErrStackOverflow_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ErrStackOverflow
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_ErrStackUnderflow_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *ErrStackUnderflow
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_Memory_Copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst uint64, src uint64, len uint64) {
		m := NewMemory()
		m.Copy(dst, src, len)
	})
}

func Fuzz_Nosy_Memory_GetCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, offset int64, size int64) {
		m := NewMemory()
		m.GetCopy(offset, size)
	})
}

func Fuzz_Nosy_Memory_GetPtr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, offset int64, size int64) {
		m := NewMemory()
		m.GetPtr(offset, size)
	})
}

func Fuzz_Nosy_Memory_Resize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64) {
		m := NewMemory()
		m.Resize(size)
	})
}

func Fuzz_Nosy_Memory_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, offset uint64, size uint64, value []byte) {
		m := NewMemory()
		m.Set(offset, size, value)
	})
}

func Fuzz_Nosy_Memory_Set32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var val *uint256.Int
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if val == nil {
			return
		}

		m := NewMemory()
		m.Set32(offset, val)
	})
}

func Fuzz_Nosy_Stack_Back__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		st := newstack()
		st.Back(n)
	})
}

func Fuzz_Nosy_Stack_dup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		st := newstack()
		st.dup(n)
	})
}

func Fuzz_Nosy_Stack_push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *uint256.Int
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		st := newstack()
		st.push(d)
	})
}

func Fuzz_Nosy_Stack_swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		st := newstack()
		st.swap(n)
	})
}

func Fuzz_Nosy_bigModExp_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bigModExp
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bigModExp_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bigModExp
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bitvec_codeSegment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var code bitvec
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var bits bitvec
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var pos uint64
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		b1 := codeBitmapInternal(code, bits)
		b1.codeSegment(pos)
	})
}

func Fuzz_Nosy_blake2F_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *blake2F
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_blake2F_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *blake2F
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G1Add_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1Add
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G1Add_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1Add
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G1Mul_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1Mul
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G1Mul_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1Mul
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G1MultiExp_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1MultiExp
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G1MultiExp_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G1MultiExp
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G2Add_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2Add
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G2Add_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2Add
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G2Mul_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2Mul
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G2Mul_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2Mul
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381G2MultiExp_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2MultiExp
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381G2MultiExp_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381G2MultiExp
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381MapG1_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381MapG1
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381MapG1_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381MapG1
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381MapG2_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381MapG2
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381MapG2_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381MapG2
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bls12381Pairing_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381Pairing
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bls12381Pairing_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bls12381Pairing
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256AddByzantium_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256AddByzantium
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256AddByzantium_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256AddByzantium
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256AddIstanbul_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256AddIstanbul
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256AddIstanbul_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256AddIstanbul
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256PairingByzantium_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256PairingByzantium
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256PairingByzantium_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256PairingByzantium
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256PairingIstanbul_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256PairingIstanbul
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256PairingIstanbul_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256PairingIstanbul
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256ScalarMulByzantium_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256ScalarMulByzantium
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256ScalarMulByzantium_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256ScalarMulByzantium
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

		c.Run(input)
	})
}

func Fuzz_Nosy_bn256ScalarMulIstanbul_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256ScalarMulIstanbul
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_bn256ScalarMulIstanbul_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *bn256ScalarMulIstanbul
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

		c.Run(input)
	})
}

func Fuzz_Nosy_codeAndHash_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *codeAndHash
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Hash()
	})
}

func Fuzz_Nosy_dataCopy_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *dataCopy
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_dataCopy_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *dataCopy
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var in []byte
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Run(in)
	})
}

func Fuzz_Nosy_ecrecover_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ecrecover
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_ecrecover_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ecrecover
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

		c.Run(input)
	})
}

func Fuzz_Nosy_kzgPointEvaluation_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *kzgPointEvaluation
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.RequiredGas(input)
	})
}

func Fuzz_Nosy_kzgPointEvaluation_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *kzgPointEvaluation
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Run(input)
	})
}

func Fuzz_Nosy_operation_HasCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op *operation
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if op == nil {
			return
		}

		op.HasCost()
	})
}

func Fuzz_Nosy_operation_Stack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op *operation
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if op == nil {
			return
		}

		op.Stack()
	})
}

func Fuzz_Nosy_ripemd160hash_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ripemd160hash
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_ripemd160hash_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ripemd160hash
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

		c.Run(input)
	})
}

func Fuzz_Nosy_sha256hash_RequiredGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *sha256hash
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

		c.RequiredGas(input)
	})
}

func Fuzz_Nosy_sha256hash_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *sha256hash
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

		c.Run(input)
	})
}

func Fuzz_Nosy_AccountRef_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ar AccountRef
		fill_err = tp.Fill(&ar)
		if fill_err != nil {
			return
		}

		ar.Address()
	})
}

// skipping Fuzz_Nosy_CallContext_Call__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.CallContext

// skipping Fuzz_Nosy_CallContext_CallCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.CallContext

// skipping Fuzz_Nosy_CallContext_Create__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.CallContext

// skipping Fuzz_Nosy_CallContext_DelegateCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.CallContext

// skipping Fuzz_Nosy_ContractRef_Address__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.ContractRef

// skipping Fuzz_Nosy_EVMLogger_CaptureEnd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureEnter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureExit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureFault__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureStart__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureTxEnd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

// skipping Fuzz_Nosy_EVMLogger_CaptureTxStart__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

func Fuzz_Nosy_OpCode_IsPush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		op := StringToOp(str)
		op.IsPush()
	})
}

func Fuzz_Nosy_OpCode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string) {
		op := StringToOp(str)
		op.String()
	})
}

// skipping Fuzz_Nosy_PrecompiledContract_RequiredGas__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.PrecompiledContract

// skipping Fuzz_Nosy_PrecompiledContract_Run__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.PrecompiledContract

// skipping Fuzz_Nosy_StateDB_AddAddressToAccessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddBalance__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddLog__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddPreimage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddRefund__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddSlotToAccessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_AddressInAccessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_CreateAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_Empty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_Exist__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetBalance__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetCodeHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetCodeSize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetCommittedState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetRefund__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_GetTransientState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_HasSelfDestructed__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_Prepare__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_RevertToSnapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SelfDestruct__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_Selfdestruct6780__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SetCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SetNonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SetState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SetTransientState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SlotInAccessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_Snapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SubBalance__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_StateDB_SubRefund__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

func Fuzz_Nosy_bitvec_set1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var code bitvec
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var bits bitvec
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var pos uint64
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		b1 := codeBitmapInternal(code, bits)
		b1.set1(pos)
	})
}

func Fuzz_Nosy_bitvec_set16__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var code bitvec
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var bits bitvec
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var pos uint64
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		b1 := codeBitmapInternal(code, bits)
		b1.set16(pos)
	})
}

func Fuzz_Nosy_bitvec_set8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var code bitvec
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var bits bitvec
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var pos uint64
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		b1 := codeBitmapInternal(code, bits)
		b1.set8(pos)
	})
}

func Fuzz_Nosy_bitvec_setN__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var code bitvec
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var bits bitvec
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var flag uint16
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var pos uint64
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		b1 := codeBitmapInternal(code, bits)
		b1.setN(flag, pos)
	})
}

func Fuzz_Nosy_ActivePrecompiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rules params.Rules
		fill_err = tp.Fill(&rules)
		if fill_err != nil {
			return
		}

		ActivePrecompiles(rules)
	})
}

// skipping Fuzz_Nosy_RunPrecompiledContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.PrecompiledContract

func Fuzz_Nosy_ValidEip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, eipNum int) {
		ValidEip(eipNum)
	})
}

func Fuzz_Nosy_allZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		allZero(b)
	})
}

func Fuzz_Nosy_calcMemSize64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var off *uint256.Int
		fill_err = tp.Fill(&off)
		if fill_err != nil {
			return
		}
		var l *uint256.Int
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if off == nil || l == nil {
			return
		}

		calcMemSize64(off, l)
	})
}

func Fuzz_Nosy_calcMemSize64WithUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var off *uint256.Int
		fill_err = tp.Fill(&off)
		if fill_err != nil {
			return
		}
		var length64 uint64
		fill_err = tp.Fill(&length64)
		if fill_err != nil {
			return
		}
		if off == nil {
			return
		}

		calcMemSize64WithUint(off, length64)
	})
}

func Fuzz_Nosy_callGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var isEip150 bool
		fill_err = tp.Fill(&isEip150)
		if fill_err != nil {
			return
		}
		var availableGas uint64
		fill_err = tp.Fill(&availableGas)
		if fill_err != nil {
			return
		}
		var base uint64
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}
		var callCost *uint256.Int
		fill_err = tp.Fill(&callCost)
		if fill_err != nil {
			return
		}
		if callCost == nil {
			return
		}

		callGas(isEip150, availableGas, base, callCost)
	})
}

func Fuzz_Nosy_decodeBLS12381FieldElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		decodeBLS12381FieldElement(in)
	})
}

func Fuzz_Nosy_enable1153__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable1153(jt)
	})
}

func Fuzz_Nosy_enable1344__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable1344(jt)
	})
}

func Fuzz_Nosy_enable1884__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable1884(jt)
	})
}

func Fuzz_Nosy_enable2200__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable2200(jt)
	})
}

func Fuzz_Nosy_enable2929__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable2929(jt)
	})
}

func Fuzz_Nosy_enable3198__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable3198(jt)
	})
}

func Fuzz_Nosy_enable3529__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable3529(jt)
	})
}

func Fuzz_Nosy_enable3855__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable3855(jt)
	})
}

func Fuzz_Nosy_enable3860__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable3860(jt)
	})
}

func Fuzz_Nosy_enable4844__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable4844(jt)
	})
}

func Fuzz_Nosy_enable5656__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable5656(jt)
	})
}

func Fuzz_Nosy_enable6780__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable6780(jt)
	})
}

func Fuzz_Nosy_enable7516__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var jt *JumpTable
		fill_err = tp.Fill(&jt)
		if fill_err != nil {
			return
		}
		if jt == nil {
			return
		}

		enable7516(jt)
	})
}

func Fuzz_Nosy_gasCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasCall(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasCallCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasCallCode(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasCreate2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasCreate2(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasCreate2Eip3860__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasCreate2Eip3860(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasCreateEip3860__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasCreateEip3860(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasDelegateCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasDelegateCall(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasEip2929AccountCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasEip2929AccountCheck(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasExpEIP158__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasExpEIP158(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasExpFrontier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasExpFrontier(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasExtCodeCopyEIP2929__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasExtCodeCopyEIP2929(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasKeccak256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasKeccak256(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasSLoadEIP2929__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasSLoadEIP2929(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasSStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasSStore(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasSStoreEIP2200__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasSStoreEIP2200(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasSelfdestruct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasSelfdestruct(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_gasStaticCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		gasStaticCall(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_getData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, start uint64, size uint64) {
		getData(d1, start, size)
	})
}

func Fuzz_Nosy_maxDupStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		maxDupStack(n)
	})
}

func Fuzz_Nosy_maxStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pop int, push int) {
		maxStack(pop, push)
	})
}

func Fuzz_Nosy_maxSwapStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		maxSwapStack(n)
	})
}

func Fuzz_Nosy_memoryCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryCall(stack)
	})
}

func Fuzz_Nosy_memoryCallDataCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryCallDataCopy(stack)
	})
}

func Fuzz_Nosy_memoryCodeCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryCodeCopy(stack)
	})
}

func Fuzz_Nosy_memoryCreate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryCreate(stack)
	})
}

func Fuzz_Nosy_memoryCreate2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryCreate2(stack)
	})
}

func Fuzz_Nosy_memoryDelegateCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryDelegateCall(stack)
	})
}

func Fuzz_Nosy_memoryExtCodeCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryExtCodeCopy(stack)
	})
}

func Fuzz_Nosy_memoryGasCost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var newMemSize uint64
		fill_err = tp.Fill(&newMemSize)
		if fill_err != nil {
			return
		}
		if mem == nil {
			return
		}

		memoryGasCost(mem, newMemSize)
	})
}

func Fuzz_Nosy_memoryKeccak256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryKeccak256(stack)
	})
}

func Fuzz_Nosy_memoryLog__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryLog(stack)
	})
}

func Fuzz_Nosy_memoryMLoad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryMLoad(stack)
	})
}

func Fuzz_Nosy_memoryMStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryMStore(stack)
	})
}

func Fuzz_Nosy_memoryMStore8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryMStore8(stack)
	})
}

func Fuzz_Nosy_memoryMcopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryMcopy(stack)
	})
}

func Fuzz_Nosy_memoryReturn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryReturn(stack)
	})
}

func Fuzz_Nosy_memoryReturnDataCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryReturnDataCopy(stack)
	})
}

func Fuzz_Nosy_memoryRevert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryRevert(stack)
	})
}

func Fuzz_Nosy_memoryStaticCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if stack == nil {
			return
		}

		memoryStaticCall(stack)
	})
}

func Fuzz_Nosy_minDupStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		minDupStack(n)
	})
}

func Fuzz_Nosy_minStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pops int, push int) {
		minStack(pops, push)
	})
}

func Fuzz_Nosy_minSwapStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		minSwapStack(n)
	})
}

func Fuzz_Nosy_opAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opAdd(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opAddmod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opAddmod(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opAddress(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opAnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opAnd(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opBalance(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opBaseFee(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opBlobBaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opBlobBaseFee(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opBlobHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opBlobHash(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opBlockhash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opBlockhash(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opByte__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opByte(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCall(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCallCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCallCode(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCallDataCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCallDataCopy(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCallDataLoad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCallDataLoad(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCallDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCallDataSize(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCallValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCallValue(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCaller__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCaller(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opChainID(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCodeCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCodeCopy(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCodeSize(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCoinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCoinbase(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCreate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCreate(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opCreate2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opCreate2(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opDelegateCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opDelegateCall(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opDifficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opDifficulty(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opDiv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opDiv(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opEq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opEq(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opExp(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opExtCodeCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opExtCodeCopy(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opExtCodeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opExtCodeHash(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opExtCodeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opExtCodeSize(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opGas(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opGasLimit(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opGasprice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opGasprice(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opGt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opGt(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opIszero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opIszero(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opJump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opJump(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opJumpdest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opJumpdest(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opJumpi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opJumpi(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opKeccak256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opKeccak256(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opLt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opLt(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMcopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMcopy(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMload(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMod(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMsize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMsize(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMstore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMstore(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMstore8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMstore8(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMul(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opMulmod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opMulmod(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opNot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opNot(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opNumber(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opOr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opOr(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opOrigin(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opPc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opPc(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opPop(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opPush0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opPush0(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opPush1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opPush1(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opRandom(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opReturn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opReturn(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opReturnDataCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opReturnDataCopy(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opReturnDataSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opReturnDataSize(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opRevert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opRevert(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSAR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSAR(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSHL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSHL(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSHR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSHR(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSdiv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSdiv(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSelfBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSelfBalance(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSelfdestruct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSelfdestruct(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSelfdestruct6780__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSelfdestruct6780(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSgt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSgt(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSignExtend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSignExtend(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSload(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSlt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSlt(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSmod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSmod(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSstore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSstore(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opStaticCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opStaticCall(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opStop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opStop(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opSub(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opTimestamp(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opTload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opTload(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opTstore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opTstore(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opUndefined__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opUndefined(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_opXor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc *uint64
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var interpreter *EVMInterpreter
		fill_err = tp.Fill(&interpreter)
		if fill_err != nil {
			return
		}
		var scope *ScopeContext
		fill_err = tp.Fill(&scope)
		if fill_err != nil {
			return
		}
		if pc == nil || interpreter == nil || scope == nil {
			return
		}

		opXor(pc, interpreter, scope)
	})
}

func Fuzz_Nosy_pureMemoryGascost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var contract *Contract
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var stack *Stack
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var mem *Memory
		fill_err = tp.Fill(&mem)
		if fill_err != nil {
			return
		}
		var memorySize uint64
		fill_err = tp.Fill(&memorySize)
		if fill_err != nil {
			return
		}
		if evm == nil || contract == nil || stack == nil || mem == nil {
			return
		}

		pureMemoryGascost(evm, contract, stack, mem, memorySize)
	})
}

func Fuzz_Nosy_returnStack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Stack
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		returnStack(s)
	})
}

func Fuzz_Nosy_runBn256Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		runBn256Add(input)
	})
}

func Fuzz_Nosy_runBn256Pairing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		runBn256Pairing(input)
	})
}

func Fuzz_Nosy_runBn256ScalarMul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		runBn256ScalarMul(input)
	})
}

func Fuzz_Nosy_toWordSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64) {
		toWordSize(size)
	})
}
