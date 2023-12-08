package blobpool

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_BlobPool_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Content__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_ContentFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Filter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if p == nil || tx == nil {
			return
		}

		p.Filter(tx)
	})
}

func Fuzz_Nosy_BlobPool_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

// skipping Fuzz_Nosy_BlobPool_Init__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool.AddressReserver

func Fuzz_Nosy_BlobPool_Locals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Header
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if p == nil || oldHead == nil || newHead == nil {
			return
		}

		p.Reset(oldHead, newHead)
	})
}

func Fuzz_Nosy_BlobPool_SetGasTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

func Fuzz_Nosy_BlobPool_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
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

// skipping Fuzz_Nosy_BlobPool_SubscribeTransactions__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.NewTxsEvent

func Fuzz_Nosy_BlobPool_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if p == nil || tx == nil {
			return
		}

		p.add(tx)
	})
}

func Fuzz_Nosy_BlobPool_drop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.drop()
	})
}

func Fuzz_Nosy_BlobPool_offload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var inclusions map[common.Hash]uint64
		fill_err = tp.Fill(&inclusions)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.offload(addr, nonce, id, inclusions)
	})
}

func Fuzz_Nosy_BlobPool_parseTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var size uint32
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		var blob []byte
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.parseTransaction(id, size, blob)
	})
}

func Fuzz_Nosy_BlobPool_recheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var inclusions map[common.Hash]uint64
		fill_err = tp.Fill(&inclusions)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.recheck(addr, inclusions)
	})
}

func Fuzz_Nosy_BlobPool_reinject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var txhash common.Hash
		fill_err = tp.Fill(&txhash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.reinject(addr, txhash)
	})
}

func Fuzz_Nosy_BlobPool_reorg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Header
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if p == nil || oldHead == nil || newHead == nil {
			return
		}

		p.reorg(oldHead, newHead)
	})
}

func Fuzz_Nosy_BlobPool_updateLimboMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.updateLimboMetrics()
	})
}

func Fuzz_Nosy_BlobPool_updateStorageMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.updateStorageMetrics()
	})
}

func Fuzz_Nosy_BlobPool_validateTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlobPool
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if p == nil || tx == nil {
			return
		}

		p.validateTx(tx)
	})
}

func Fuzz_Nosy_Config_sanitize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		config.sanitize()
	})
}

func Fuzz_Nosy_evictHeap_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var basefee *uint256.Int
		fill_err = tp.Fill(&basefee)
		if fill_err != nil {
			return
		}
		var blobfee *uint256.Int
		fill_err = tp.Fill(&blobfee)
		if fill_err != nil {
			return
		}
		var index *map[common.Address][]*blobTxMeta
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if basefee == nil || blobfee == nil || index == nil {
			return
		}

		h := newPriceHeap(basefee, blobfee, index)
		h.Len()
	})
}

func Fuzz_Nosy_evictHeap_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var basefee *uint256.Int
		fill_err = tp.Fill(&basefee)
		if fill_err != nil {
			return
		}
		var blobfee *uint256.Int
		fill_err = tp.Fill(&blobfee)
		if fill_err != nil {
			return
		}
		var index *map[common.Address][]*blobTxMeta
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if basefee == nil || blobfee == nil || index == nil {
			return
		}

		h := newPriceHeap(basefee, blobfee, index)
		h.Less(i, j)
	})
}

func Fuzz_Nosy_evictHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var basefee *uint256.Int
		fill_err = tp.Fill(&basefee)
		if fill_err != nil {
			return
		}
		var blobfee *uint256.Int
		fill_err = tp.Fill(&blobfee)
		if fill_err != nil {
			return
		}
		var index *map[common.Address][]*blobTxMeta
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if basefee == nil || blobfee == nil || index == nil {
			return
		}

		h := newPriceHeap(basefee, blobfee, index)
		h.Pop()
	})
}

// skipping Fuzz_Nosy_evictHeap_Push__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_evictHeap_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var basefee *uint256.Int
		fill_err = tp.Fill(&basefee)
		if fill_err != nil {
			return
		}
		var blobfee *uint256.Int
		fill_err = tp.Fill(&blobfee)
		if fill_err != nil {
			return
		}
		var index *map[common.Address][]*blobTxMeta
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}
		if basefee == nil || blobfee == nil || index == nil {
			return
		}

		h := newPriceHeap(basefee, blobfee, index)
		h.Swap(i, j)
	})
}

func Fuzz_Nosy_evictHeap_reinit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b1 *uint256.Int
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		var b2 *uint256.Int
		fill_err = tp.Fill(&b2)
		if fill_err != nil {
			return
		}
		var index *map[common.Address][]*blobTxMeta
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var b4 *uint256.Int
		fill_err = tp.Fill(&b4)
		if fill_err != nil {
			return
		}
		var b5 *uint256.Int
		fill_err = tp.Fill(&b5)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if b1 == nil || b2 == nil || index == nil || b4 == nil || b5 == nil {
			return
		}

		h := newPriceHeap(b1, b2, index)
		h.reinit(b4, b5, force)
	})
}

func Fuzz_Nosy_limbo_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string) {
		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.Close()
	})
}

func Fuzz_Nosy_limbo_finalize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var final *types.Header
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		if final == nil {
			return
		}

		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.finalize(final)
	})
}

func Fuzz_Nosy_limbo_getAndDrop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string, id uint64) {
		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.getAndDrop(id)
	})
}

func Fuzz_Nosy_limbo_parseBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, datadir string, id uint64, d3 []byte) {
		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.parseBlob(id, d3)
	})
}

func Fuzz_Nosy_limbo_pull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var tx common.Hash
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}

		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.pull(tx)
	})
}

func Fuzz_Nosy_limbo_push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.push(tx, block)
	})
}

func Fuzz_Nosy_limbo_setAndIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.setAndIndex(tx, block)
	})
}

func Fuzz_Nosy_limbo_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var txhash common.Hash
		fill_err = tp.Fill(&txhash)
		if fill_err != nil {
			return
		}
		var block uint64
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}

		l, err := newLimbo(datadir)
		if err != nil {
			return
		}
		l.update(txhash, block)
	})
}

// skipping Fuzz_Nosy_BlockChain_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/blobpool.BlockChain

// skipping Fuzz_Nosy_BlockChain_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/blobpool.BlockChain

// skipping Fuzz_Nosy_BlockChain_CurrentFinalBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/blobpool.BlockChain

// skipping Fuzz_Nosy_BlockChain_GetBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/blobpool.BlockChain

// skipping Fuzz_Nosy_BlockChain_StateAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/txpool/blobpool.BlockChain

func Fuzz_Nosy_dynamicFeeJumps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fee *uint256.Int
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
		if fee == nil {
			return
		}

		dynamicFeeJumps(fee)
	})
}

func Fuzz_Nosy_evictionPriority__(f *testing.F) {
	f.Fuzz(func(t *testing.T, basefeeJumps float64, txBasefeeJumps float64, blobfeeJumps float64, txBlobfeeJumps float64) {
		evictionPriority(basefeeJumps, txBasefeeJumps, blobfeeJumps, txBlobfeeJumps)
	})
}

func Fuzz_Nosy_evictionPriority1D__(f *testing.F) {
	f.Fuzz(func(t *testing.T, basefeeJumps float64, txfeeJumps float64) {
		evictionPriority1D(basefeeJumps, txfeeJumps)
	})
}

func Fuzz_Nosy_intLog2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n uint) {
		intLog2(n)
	})
}
