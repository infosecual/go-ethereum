package txpool

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_LazyTransaction_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ltx *LazyTransaction
		fill_err = tp.Fill(&ltx)
		if fill_err != nil {
			return
		}
		if ltx == nil {
			return
		}

		ltx.Resolve()
	})
}

func Fuzz_Nosy_TxPool_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var local bool
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var sync bool
		fill_err = tp.Fill(&sync)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Add(txs, local, sync)
	})
}

func Fuzz_Nosy_TxPool_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_TxPool_Content__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Content()
	})
}

func Fuzz_Nosy_TxPool_ContentFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ContentFrom(addr)
	})
}

func Fuzz_Nosy_TxPool_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Get(hash)
	})
}

func Fuzz_Nosy_TxPool_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Has(hash)
	})
}

func Fuzz_Nosy_TxPool_Locals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Locals()
	})
}

func Fuzz_Nosy_TxPool_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Nonce(addr)
	})
}

func Fuzz_Nosy_TxPool_Pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var enforceTips bool
		fill_err = tp.Fill(&enforceTips)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Pending(enforceTips)
	})
}

func Fuzz_Nosy_TxPool_SetGasTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var tip *big.Int
		fill_err = tp.Fill(&tip)
		if fill_err != nil {
			return
		}
		if p == nil || tip == nil {
			return
		}

		p.SetGasTip(tip)
	})
}

func Fuzz_Nosy_TxPool_Stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Stats()
	})
}

func Fuzz_Nosy_TxPool_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TxPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Status(hash)
	})
}

// skipping Fuzz_Nosy_TxPool_SubscribeTransactions__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.NewTxsEvent

// skipping Fuzz_Nosy_TxPool_loop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.BlockChain

// skipping Fuzz_Nosy_TxPool_reserver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_BlockChain_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.BlockChain

// skipping Fuzz_Nosy_BlockChain_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.BlockChain

// skipping Fuzz_Nosy_LazyResolver_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.LazyResolver

// skipping Fuzz_Nosy_SubPool_Add__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Content__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_ContentFrom__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Filter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Has__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Init__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Locals__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Nonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Pending__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_SetGasTip__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Stats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_Status__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool

// skipping Fuzz_Nosy_SubPool_SubscribeTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.SubPool
