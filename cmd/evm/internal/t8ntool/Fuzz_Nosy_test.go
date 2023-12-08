package t8ntool

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_NumberedError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NumberedError
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Error()
	})
}

func Fuzz_Nosy_NumberedError_ExitCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *NumberedError
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ExitCode()
	})
}

// skipping Fuzz_Nosy_Prestate_Apply__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/evm/internal/t8ntool.txIterator

func Fuzz_Nosy_bbInput_SealBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if ctx == nil || block == nil {
			return
		}

		i, err := readInput(ctx)
		if err != nil {
			return
		}
		i.SealBlock(block)
	})
}

func Fuzz_Nosy_bbInput_ToBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		i, err := readInput(ctx)
		if err != nil {
			return
		}
		i.ToBlock()
	})
}

func Fuzz_Nosy_bbInput_sealClique__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if ctx == nil || block == nil {
			return
		}

		i, err := readInput(ctx)
		if err != nil {
			return
		}
		i.sealClique(block)
	})
}

func Fuzz_Nosy_cliqueInput_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cliqueInput
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

func Fuzz_Nosy_header_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_result_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *result
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.MarshalJSON()
	})
}

func Fuzz_Nosy_rlpTxIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *rlpTxIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_rlpTxIterator_Tx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *rlpTxIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Tx()
	})
}

func Fuzz_Nosy_sliceTxIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ait *sliceTxIterator
		fill_err = tp.Fill(&ait)
		if fill_err != nil {
			return
		}
		if ait == nil {
			return
		}

		ait.Next()
	})
}

func Fuzz_Nosy_sliceTxIterator_Tx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ait *sliceTxIterator
		fill_err = tp.Fill(&ait)
		if fill_err != nil {
			return
		}
		if ait == nil {
			return
		}

		ait.Tx()
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

func Fuzz_Nosy_txWithKey_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *txWithKey
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Alloc_OnAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g Alloc
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var addr *common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var dumpAccount state.DumpAccount
		fill_err = tp.Fill(&dumpAccount)
		if fill_err != nil {
			return
		}
		if addr == nil {
			return
		}

		g.OnAccount(addr, dumpAccount)
	})
}

func Fuzz_Nosy_Alloc_OnRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g Alloc
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var _x2 common.Hash
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		g.OnRoot(_x2)
	})
}

func Fuzz_Nosy_header_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.MarshalJSON()
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

func Fuzz_Nosy_txIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var transactions types.Transactions
		fill_err = tp.Fill(&transactions)
		if fill_err != nil {
			return
		}

		_x1 := newSliceTxIterator(transactions)
		_x1.Next()
	})
}

func Fuzz_Nosy_txIterator_Tx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var transactions types.Transactions
		fill_err = tp.Fill(&transactions)
		if fill_err != nil {
			return
		}

		_x1 := newSliceTxIterator(transactions)
		_x1.Tx()
	})
}

func Fuzz_Nosy_createBasedir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		createBasedir(ctx)
	})
}
