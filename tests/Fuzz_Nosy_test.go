package tests

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
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

// skipping Fuzz_Nosy_BlockTest_Run__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.EVMLogger

func Fuzz_Nosy_BlockTest_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *BlockTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var in []byte
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UnmarshalJSON(in)
	})
}

func Fuzz_Nosy_BlockTest_genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *BlockTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if t1 == nil || config == nil {
			return
		}

		t1.genesis(config)
	})
}

func Fuzz_Nosy_BlockTest_insertBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *BlockTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var blockchain *core.BlockChain
		fill_err = tp.Fill(&blockchain)
		if fill_err != nil {
			return
		}
		if t1 == nil || blockchain == nil {
			return
		}

		t1.insertBlocks(blockchain)
	})
}

func Fuzz_Nosy_BlockTest_validateImportedHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *BlockTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var cm *core.BlockChain
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var validBlocks []btBlock
		fill_err = tp.Fill(&validBlocks)
		if fill_err != nil {
			return
		}
		if t1 == nil || cm == nil {
			return
		}

		t1.validateImportedHeaders(cm, validBlocks)
	})
}

func Fuzz_Nosy_BlockTest_validatePostState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *BlockTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		if t1 == nil || statedb == nil {
			return
		}

		t1.validatePostState(statedb)
	})
}

func Fuzz_Nosy_DifficultyTest_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var test *DifficultyTest
		fill_err = tp.Fill(&test)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if test == nil || config == nil {
			return
		}

		test.Run(config)
	})
}

func Fuzz_Nosy_DifficultyTest_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *DifficultyTest
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_RLPTest_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *RLPTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Run()
	})
}

// skipping Fuzz_Nosy_StateTest_Run__ because parameters include func, chan, or unsupported interface: func(err error, snaps *github.com/ethereum/go-ethereum/core/state/snapshot.Tree, state *github.com/ethereum/go-ethereum/core/state.StateDB)

func Fuzz_Nosy_StateTest_RunNoVerify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StateTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var subtest StateSubtest
		fill_err = tp.Fill(&subtest)
		if fill_err != nil {
			return
		}
		var vmconfig vm.Config
		fill_err = tp.Fill(&vmconfig)
		if fill_err != nil {
			return
		}
		var snapshotter bool
		fill_err = tp.Fill(&snapshotter)
		if fill_err != nil {
			return
		}
		var scheme string
		fill_err = tp.Fill(&scheme)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RunNoVerify(subtest, vmconfig, snapshotter, scheme)
	})
}

func Fuzz_Nosy_StateTest_Subtests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StateTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Subtests()
	})
}

func Fuzz_Nosy_StateTest_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StateTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var in []byte
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UnmarshalJSON(in)
	})
}

// skipping Fuzz_Nosy_StateTest_checkError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_StateTest_gasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StateTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var subtest StateSubtest
		fill_err = tp.Fill(&subtest)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.gasLimit(subtest)
	})
}

func Fuzz_Nosy_StateTest_genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *StateTest
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if t1 == nil || config == nil {
			return
		}

		t1.genesis(config)
	})
}

func Fuzz_Nosy_TransactionTest_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TransactionTest
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if tt == nil || config == nil {
			return
		}

		tt.Run(config)
	})
}

func Fuzz_Nosy_btBlock_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bb *btBlock
		fill_err = tp.Fill(&bb)
		if fill_err != nil {
			return
		}
		if bb == nil {
			return
		}

		bb.decode()
	})
}

func Fuzz_Nosy_btHeader_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *btHeader
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

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_stEnv_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stEnv
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

func Fuzz_Nosy_stTransaction_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *stTransaction
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

func Fuzz_Nosy_stTransaction_toMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *stTransaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var ps stPostState
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if tx == nil || baseFee == nil {
			return
		}

		tx.toMessage(ps, baseFee)
	})
}

func Fuzz_Nosy_DifficultyTest_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DifficultyTest
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.MarshalJSON()
	})
}

func Fuzz_Nosy_UnsupportedForkError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e UnsupportedForkError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_btHeader_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b btHeader
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalJSON()
	})
}

func Fuzz_Nosy_stEnv_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s stEnv
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_stTransaction_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s stTransaction
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_FromHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		FromHex(s)
	})
}

func Fuzz_Nosy_GetChainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, forkString string) {
		GetChainConfig(forkString)
	})
}

// skipping Fuzz_Nosy_MakePreState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_translateJSON__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_u64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, val uint64) {
		u64(val)
	})
}
