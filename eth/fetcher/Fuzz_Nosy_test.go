package fetcher

import (
	"testing"
	"time"

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

func Fuzz_Nosy_BlockFetcher_Enqueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil || block == nil {
			return
		}

		f1.Enqueue(peer, block)
	})
}

func Fuzz_Nosy_BlockFetcher_FilterBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var transactions [][]*types.Transaction
		fill_err = tp.Fill(&transactions)
		if fill_err != nil {
			return
		}
		var uncles [][]*types.Header
		fill_err = tp.Fill(&uncles)
		if fill_err != nil {
			return
		}
		var time time.Time
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterBodies(peer, transactions, uncles, time)
	})
}

func Fuzz_Nosy_BlockFetcher_FilterHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var time time.Time
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.FilterHeaders(peer, headers, time)
	})
}

// skipping Fuzz_Nosy_BlockFetcher_Notify__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/fetcher.headerRequesterFn

func Fuzz_Nosy_BlockFetcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Start()
	})
}

func Fuzz_Nosy_BlockFetcher_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Stop()
	})
}

func Fuzz_Nosy_BlockFetcher_enqueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil || header == nil || block == nil {
			return
		}

		f1.enqueue(peer, header, block)
	})
}

func Fuzz_Nosy_BlockFetcher_forgetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.forgetBlock(hash)
	})
}

func Fuzz_Nosy_BlockFetcher_forgetHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.forgetHash(hash)
	})
}

func Fuzz_Nosy_BlockFetcher_importBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if f1 == nil || block == nil {
			return
		}

		f1.importBlocks(peer, block)
	})
}

func Fuzz_Nosy_BlockFetcher_importHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if f1 == nil || header == nil {
			return
		}

		f1.importHeaders(peer, header)
	})
}

func Fuzz_Nosy_BlockFetcher_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.loop()
	})
}

func Fuzz_Nosy_BlockFetcher_rescheduleComplete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var complete *time.Timer
		fill_err = tp.Fill(&complete)
		if fill_err != nil {
			return
		}
		if f1 == nil || complete == nil {
			return
		}

		f1.rescheduleComplete(complete)
	})
}

func Fuzz_Nosy_BlockFetcher_rescheduleFetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BlockFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var fetch *time.Timer
		fill_err = tp.Fill(&fetch)
		if fill_err != nil {
			return
		}
		if f1 == nil || fetch == nil {
			return
		}

		f1.rescheduleFetch(fetch)
	})
}

func Fuzz_Nosy_TxFetcher_Drop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Drop(peer)
	})
}

func Fuzz_Nosy_TxFetcher_Enqueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var direct bool
		fill_err = tp.Fill(&direct)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Enqueue(peer, txs, direct)
	})
}

func Fuzz_Nosy_TxFetcher_Notify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var types []byte
		fill_err = tp.Fill(&types)
		if fill_err != nil {
			return
		}
		var sizes []uint32
		fill_err = tp.Fill(&sizes)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Notify(peer, types, sizes, hashes)
	})
}

func Fuzz_Nosy_TxFetcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Start()
	})
}

func Fuzz_Nosy_TxFetcher_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Stop()
	})
}

// skipping Fuzz_Nosy_TxFetcher_forEachAnnounce__ because parameters include func, chan, or unsupported interface: func(hash github.com/ethereum/go-ethereum/common.Hash, meta *github.com/ethereum/go-ethereum/eth/fetcher.txMetadata) bool

// skipping Fuzz_Nosy_TxFetcher_forEachPeer__ because parameters include func, chan, or unsupported interface: func(peer string)

func Fuzz_Nosy_TxFetcher_isKnownUnderpriced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.isKnownUnderpriced(hash)
	})
}

func Fuzz_Nosy_TxFetcher_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TxFetcher
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.loop()
	})
}

// skipping Fuzz_Nosy_TxFetcher_rescheduleTimeout__ because parameters include func, chan, or unsupported interface: *github.com/ethereum/go-ethereum/common/mclock.Timer

// skipping Fuzz_Nosy_TxFetcher_rescheduleWait__ because parameters include func, chan, or unsupported interface: *github.com/ethereum/go-ethereum/common/mclock.Timer

// skipping Fuzz_Nosy_TxFetcher_scheduleFetches__ because parameters include func, chan, or unsupported interface: *github.com/ethereum/go-ethereum/common/mclock.Timer

func Fuzz_Nosy_blockOrHeaderInject_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inject *blockOrHeaderInject
		fill_err = tp.Fill(&inject)
		if fill_err != nil {
			return
		}
		if inject == nil {
			return
		}

		inject.hash()
	})
}

func Fuzz_Nosy_blockOrHeaderInject_number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inject *blockOrHeaderInject
		fill_err = tp.Fill(&inject)
		if fill_err != nil {
			return
		}
		if inject == nil {
			return
		}

		inject.number()
	})
}

func Fuzz_Nosy_rotateHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slice []common.Hash
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		rotateHashes(slice, n)
	})
}

func Fuzz_Nosy_rotateStrings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slice []string
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		rotateStrings(slice, n)
	})
}

func Fuzz_Nosy_sortHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slice []common.Hash
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}

		sortHashes(slice)
	})
}
