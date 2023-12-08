package core

import (
	"context"
	"io"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
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

func Fuzz_Nosy_BlockChain_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Config()
	})
}

func Fuzz_Nosy_BlockChain_ContractCodeWithPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.ContractCodeWithPrefix(hash)
	})
}

func Fuzz_Nosy_BlockChain_CurrentBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.CurrentBlock()
	})
}

func Fuzz_Nosy_BlockChain_CurrentFinalBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.CurrentFinalBlock()
	})
}

func Fuzz_Nosy_BlockChain_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.CurrentHeader()
	})
}

func Fuzz_Nosy_BlockChain_CurrentSafeBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.CurrentSafeBlock()
	})
}

func Fuzz_Nosy_BlockChain_CurrentSnapBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.CurrentSnapBlock()
	})
}

func Fuzz_Nosy_BlockChain_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Engine()
	})
}

func Fuzz_Nosy_BlockChain_Export__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Export(w)
	})
}

func Fuzz_Nosy_BlockChain_ExportN__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var first uint64
		fill_err = tp.Fill(&first)
		if fill_err != nil {
			return
		}
		var last uint64
		fill_err = tp.Fill(&last)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.ExportN(w, first, last)
	})
}

func Fuzz_Nosy_BlockChain_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GasLimit()
	})
}

func Fuzz_Nosy_BlockChain_Genesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Genesis()
	})
}

func Fuzz_Nosy_BlockChain_GetAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var ancestor uint64
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		var maxNonCanonical *uint64
		fill_err = tp.Fill(&maxNonCanonical)
		if fill_err != nil {
			return
		}
		if bc == nil || maxNonCanonical == nil {
			return
		}

		bc.GetAncestor(hash, number, ancestor, maxNonCanonical)
	})
}

func Fuzz_Nosy_BlockChain_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBlock(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_GetBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBlockByHash(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetBlockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBlockByNumber(number)
	})
}

func Fuzz_Nosy_BlockChain_GetBlocksFromHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBlocksFromHash(hash, n)
	})
}

func Fuzz_Nosy_BlockChain_GetBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBody(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetBodyRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetBodyRLP(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetCanonicalHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetCanonicalHash(number)
	})
}

func Fuzz_Nosy_BlockChain_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetHeader(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetHeaderByHash(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetHeaderByNumber(number)
	})
}

func Fuzz_Nosy_BlockChain_GetHeadersFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetHeadersFrom(number, count)
	})
}

func Fuzz_Nosy_BlockChain_GetReceiptsByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetReceiptsByHash(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetTd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetTd(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_GetTransactionLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetTransactionLookup(hash)
	})
}

func Fuzz_Nosy_BlockChain_GetTrieFlushInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetTrieFlushInterval()
	})
}

func Fuzz_Nosy_BlockChain_GetUnclesInChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var length int
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.GetUnclesInChain(block, length)
	})
}

func Fuzz_Nosy_BlockChain_GetVMConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetVMConfig()
	})
}

func Fuzz_Nosy_BlockChain_HasBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.HasBlock(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_HasBlockAndState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.HasBlockAndState(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_HasFastBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.HasFastBlock(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_HasHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.HasHeader(hash, number)
	})
}

func Fuzz_Nosy_BlockChain_HasState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.HasState(hash)
	})
}

func Fuzz_Nosy_BlockChain_InsertBlockWithoutSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.InsertBlockWithoutSetHead(block)
	})
}

func Fuzz_Nosy_BlockChain_InsertChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chain types.Blocks
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.InsertChain(chain)
	})
}

func Fuzz_Nosy_BlockChain_InsertHeaderChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chain []*types.Header
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.InsertHeaderChain(chain)
	})
}

func Fuzz_Nosy_BlockChain_InsertReceiptChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var blockChain types.Blocks
		fill_err = tp.Fill(&blockChain)
		if fill_err != nil {
			return
		}
		var receiptChain []types.Receipts
		fill_err = tp.Fill(&receiptChain)
		if fill_err != nil {
			return
		}
		var ancientLimit uint64
		fill_err = tp.Fill(&ancientLimit)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.InsertReceiptChain(blockChain, receiptChain, ancientLimit)
	})
}

func Fuzz_Nosy_BlockChain_Processor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Processor()
	})
}

func Fuzz_Nosy_BlockChain_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Reset()
	})
}

func Fuzz_Nosy_BlockChain_ResetWithGenesisBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var genesis *types.Block
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		if bc == nil || genesis == nil {
			return
		}

		bc.ResetWithGenesisBlock(genesis)
	})
}

// skipping Fuzz_Nosy_BlockChain_SetBlockValidatorAndProcessorForTesting__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.Validator

func Fuzz_Nosy_BlockChain_SetCanonical__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var head *types.Block
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if bc == nil || head == nil {
			return
		}

		bc.SetCanonical(head)
	})
}

func Fuzz_Nosy_BlockChain_SetFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if bc == nil || header == nil {
			return
		}

		bc.SetFinalized(header)
	})
}

func Fuzz_Nosy_BlockChain_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.SetHead(head)
	})
}

func Fuzz_Nosy_BlockChain_SetHeadWithTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.SetHeadWithTimestamp(timestamp)
	})
}

func Fuzz_Nosy_BlockChain_SetSafe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if bc == nil || header == nil {
			return
		}

		bc.SetSafe(header)
	})
}

func Fuzz_Nosy_BlockChain_SetTrieFlushInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.SetTrieFlushInterval(interval)
	})
}

func Fuzz_Nosy_BlockChain_SetTxLookupLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.SetTxLookupLimit(limit)
	})
}

func Fuzz_Nosy_BlockChain_SnapSyncCommitHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.SnapSyncCommitHead(hash)
	})
}

func Fuzz_Nosy_BlockChain_Snapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Snapshots()
	})
}

func Fuzz_Nosy_BlockChain_State__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.State()
	})
}

func Fuzz_Nosy_BlockChain_StateAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.StateAt(root)
	})
}

func Fuzz_Nosy_BlockChain_StateCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.StateCache()
	})
}

func Fuzz_Nosy_BlockChain_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Stop()
	})
}

func Fuzz_Nosy_BlockChain_StopInsert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.StopInsert()
	})
}

// skipping Fuzz_Nosy_BlockChain_SubscribeBlockProcessingEvent__ because parameters include func, chan, or unsupported interface: chan<- bool

// skipping Fuzz_Nosy_BlockChain_SubscribeChainEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainEvent

// skipping Fuzz_Nosy_BlockChain_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainHeadEvent

// skipping Fuzz_Nosy_BlockChain_SubscribeChainSideEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.ChainSideEvent

// skipping Fuzz_Nosy_BlockChain_SubscribeLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_BlockChain_SubscribeRemovedLogsEvent__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core.RemovedLogsEvent

func Fuzz_Nosy_BlockChain_TrieDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.TrieDB()
	})
}

func Fuzz_Nosy_BlockChain_TxLookupLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.TxLookupLimit()
	})
}

func Fuzz_Nosy_BlockChain_Validator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.Validator()
	})
}

func Fuzz_Nosy_BlockChain_WriteBlockAndSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var logs []*types.Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}
		var state *state.StateDB
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var emitHeadEvent bool
		fill_err = tp.Fill(&emitHeadEvent)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil || state == nil {
			return
		}

		bc.WriteBlockAndSetHead(block, receipts, logs, state, emitHeadEvent)
	})
}

func Fuzz_Nosy_BlockChain_addFutureBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.addFutureBlock(block)
	})
}

func Fuzz_Nosy_BlockChain_collectLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var b *types.Block
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var removed bool
		fill_err = tp.Fill(&removed)
		if fill_err != nil {
			return
		}
		if bc == nil || b == nil {
			return
		}

		bc.collectLogs(b, removed)
	})
}

func Fuzz_Nosy_BlockChain_empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.empty()
	})
}

// skipping Fuzz_Nosy_BlockChain_indexBlocks__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_BlockChain_insertChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chain types.Blocks
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var setHead bool
		fill_err = tp.Fill(&setHead)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.insertChain(chain, setHead)
	})
}

func Fuzz_Nosy_BlockChain_insertSideChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil || it == nil {
			return
		}

		bc.insertSideChain(block, it)
	})
}

func Fuzz_Nosy_BlockChain_insertStopped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.insertStopped()
	})
}

func Fuzz_Nosy_BlockChain_loadLastState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.loadLastState()
	})
}

func Fuzz_Nosy_BlockChain_maintainTxIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.maintainTxIndex()
	})
}

func Fuzz_Nosy_BlockChain_procFutureBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.procFutureBlocks()
	})
}

func Fuzz_Nosy_BlockChain_recoverAncestors__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.recoverAncestors(block)
	})
}

func Fuzz_Nosy_BlockChain_reorg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var oldHead *types.Header
		fill_err = tp.Fill(&oldHead)
		if fill_err != nil {
			return
		}
		var newHead *types.Block
		fill_err = tp.Fill(&newHead)
		if fill_err != nil {
			return
		}
		if bc == nil || oldHead == nil || newHead == nil {
			return
		}

		bc.reorg(oldHead, newHead)
	})
}

// skipping Fuzz_Nosy_BlockChain_reportBlock__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_BlockChain_setHeadBeyondRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var repair bool
		fill_err = tp.Fill(&repair)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.setHeadBeyondRoot(head, time, root, repair)
	})
}

// skipping Fuzz_Nosy_BlockChain_skipBlock__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_BlockChain_stateRecoverable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.stateRecoverable(root)
	})
}

func Fuzz_Nosy_BlockChain_stopWithoutSaving__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.stopWithoutSaving()
	})
}

func Fuzz_Nosy_BlockChain_updateFutureBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.updateFutureBlocks()
	})
}

func Fuzz_Nosy_BlockChain_writeBlockAndSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var logs []*types.Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}
		var state *state.StateDB
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var emitHeadEvent bool
		fill_err = tp.Fill(&emitHeadEvent)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil || state == nil {
			return
		}

		bc.writeBlockAndSetHead(block, receipts, logs, state, emitHeadEvent)
	})
}

func Fuzz_Nosy_BlockChain_writeBlockWithState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var receipts []*types.Receipt
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var state *state.StateDB
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil || state == nil {
			return
		}

		bc.writeBlockWithState(block, receipts, state)
	})
}

func Fuzz_Nosy_BlockChain_writeBlockWithoutState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil || td == nil {
			return
		}

		bc.writeBlockWithoutState(block, td)
	})
}

func Fuzz_Nosy_BlockChain_writeHeadBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.writeHeadBlock(block)
	})
}

func Fuzz_Nosy_BlockChain_writeKnownBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bc == nil || block == nil {
			return
		}

		bc.writeKnownBlock(block)
	})
}

func Fuzz_Nosy_BlockGen_AddTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.AddTx(tx)
	})
}

func Fuzz_Nosy_BlockGen_AddTxWithChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || bc == nil || tx == nil {
			return
		}

		b.AddTxWithChain(bc, tx)
	})
}

func Fuzz_Nosy_BlockGen_AddTxWithVMConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var config vm.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.AddTxWithVMConfig(tx, config)
	})
}

func Fuzz_Nosy_BlockGen_AddUncheckedReceipt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var receipt *types.Receipt
		fill_err = tp.Fill(&receipt)
		if fill_err != nil {
			return
		}
		if b == nil || receipt == nil {
			return
		}

		b.AddUncheckedReceipt(receipt)
	})
}

func Fuzz_Nosy_BlockGen_AddUncheckedTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || tx == nil {
			return
		}

		b.AddUncheckedTx(tx)
	})
}

func Fuzz_Nosy_BlockGen_AddUncle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var h *types.Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if b == nil || h == nil {
			return
		}

		b.AddUncle(h)
	})
}

func Fuzz_Nosy_BlockGen_AddWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var w *types.Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if b == nil || w == nil {
			return
		}

		b.AddWithdrawal(w)
	})
}

func Fuzz_Nosy_BlockGen_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.BaseFee()
	})
}

func Fuzz_Nosy_BlockGen_Difficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Difficulty()
	})
}

func Fuzz_Nosy_BlockGen_Gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Gas()
	})
}

func Fuzz_Nosy_BlockGen_GetBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBalance(addr)
	})
}

func Fuzz_Nosy_BlockGen_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Number()
	})
}

func Fuzz_Nosy_BlockGen_OffsetTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var seconds int64
		fill_err = tp.Fill(&seconds)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.OffsetTime(seconds)
	})
}

func Fuzz_Nosy_BlockGen_PrevBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.PrevBlock(index)
	})
}

func Fuzz_Nosy_BlockGen_SetCoinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetCoinbase(addr)
	})
}

func Fuzz_Nosy_BlockGen_SetDifficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var diff *big.Int
		fill_err = tp.Fill(&diff)
		if fill_err != nil {
			return
		}
		if b == nil || diff == nil {
			return
		}

		b.SetDifficulty(diff)
	})
}

func Fuzz_Nosy_BlockGen_SetExtra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetExtra(d2)
	})
}

func Fuzz_Nosy_BlockGen_SetNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var nonce types.BlockNonce
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetNonce(nonce)
	})
}

func Fuzz_Nosy_BlockGen_SetParentBeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetParentBeaconRoot(root)
	})
}

func Fuzz_Nosy_BlockGen_SetPoS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetPoS()
	})
}

func Fuzz_Nosy_BlockGen_Signer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Signer()
	})
}

func Fuzz_Nosy_BlockGen_Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Timestamp()
	})
}

func Fuzz_Nosy_BlockGen_TxNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.TxNonce(addr)
	})
}

func Fuzz_Nosy_BlockGen_addTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var bc *BlockChain
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var vmConfig vm.Config
		fill_err = tp.Fill(&vmConfig)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if b == nil || bc == nil || tx == nil {
			return
		}

		b.addTx(bc, vmConfig, tx)
	})
}

func Fuzz_Nosy_BlockGen_nextWithdrawalIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlockGen
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.nextWithdrawalIndex()
	})
}

func Fuzz_Nosy_BlockValidator_ValidateBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *BlockValidator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if v == nil || block == nil {
			return
		}

		v.ValidateBody(block)
	})
}

func Fuzz_Nosy_BlockValidator_ValidateState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *BlockValidator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var receipts types.Receipts
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		var usedGas uint64
		fill_err = tp.Fill(&usedGas)
		if fill_err != nil {
			return
		}
		if v == nil || block == nil || statedb == nil {
			return
		}

		v.ValidateState(block, statedb, receipts, usedGas)
	})
}

func Fuzz_Nosy_BloomIndexer_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BloomIndexer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Commit()
	})
}

func Fuzz_Nosy_BloomIndexer_Process__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BloomIndexer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if b == nil || header == nil {
			return
		}

		b.Process(ctx, header)
	})
}

func Fuzz_Nosy_BloomIndexer_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BloomIndexer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var threshold uint64
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Prune(threshold)
	})
}

func Fuzz_Nosy_BloomIndexer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BloomIndexer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		var lastSectionHead common.Hash
		fill_err = tp.Fill(&lastSectionHead)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Reset(ctx, section, lastSectionHead)
	})
}

func Fuzz_Nosy_CacheConfig_triedbConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, scheme string) {
		c := DefaultCacheConfigWithScheme(scheme)
		c.triedbConfig()
	})
}

func Fuzz_Nosy_ChainIndexer_AddCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		var shead common.Hash
		fill_err = tp.Fill(&shead)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.AddCheckpoint(section, shead)
	})
}

func Fuzz_Nosy_ChainIndexer_AddChildIndexer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var indexer *ChainIndexer
		fill_err = tp.Fill(&indexer)
		if fill_err != nil {
			return
		}
		if c == nil || indexer == nil {
			return
		}

		c.AddChildIndexer(indexer)
	})
}

func Fuzz_Nosy_ChainIndexer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_ChainIndexer_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var threshold uint64
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Prune(threshold)
	})
}

func Fuzz_Nosy_ChainIndexer_SectionHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SectionHead(section)
	})
}

func Fuzz_Nosy_ChainIndexer_Sections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Sections()
	})
}

// skipping Fuzz_Nosy_ChainIndexer_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerChain

// skipping Fuzz_Nosy_ChainIndexer_eventLoop__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum/go-ethereum/core.ChainHeadEvent

func Fuzz_Nosy_ChainIndexer_loadValidSections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.loadValidSections()
	})
}

func Fuzz_Nosy_ChainIndexer_newHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var reorg bool
		fill_err = tp.Fill(&reorg)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.newHead(head, reorg)
	})
}

func Fuzz_Nosy_ChainIndexer_processSection__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		var lastHead common.Hash
		fill_err = tp.Fill(&lastHead)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.processSection(section, lastHead)
	})
}

func Fuzz_Nosy_ChainIndexer_removeSectionHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.removeSectionHead(section)
	})
}

func Fuzz_Nosy_ChainIndexer_setSectionHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var section uint64
		fill_err = tp.Fill(&section)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.setSectionHead(section, hash)
	})
}

func Fuzz_Nosy_ChainIndexer_setValidSections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sections uint64
		fill_err = tp.Fill(&sections)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.setValidSections(sections)
	})
}

func Fuzz_Nosy_ChainIndexer_updateLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.updateLoop()
	})
}

func Fuzz_Nosy_ChainIndexer_verifyLastHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainIndexer
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.verifyLastHead()
	})
}

func Fuzz_Nosy_ExecutionResult_Failed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		result, err := ApplyMessage(evm, msg, gp)
		if err != nil {
			return
		}
		result.Failed()
	})
}

func Fuzz_Nosy_ExecutionResult_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		result, err := ApplyMessage(evm, msg, gp)
		if err != nil {
			return
		}
		result.Return()
	})
}

func Fuzz_Nosy_ExecutionResult_Revert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		result, err := ApplyMessage(evm, msg, gp)
		if err != nil {
			return
		}
		result.Revert()
	})
}

func Fuzz_Nosy_ExecutionResult_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		result, err := ApplyMessage(evm, msg, gp)
		if err != nil {
			return
		}
		result.Unwrap()
	})
}

func Fuzz_Nosy_ForkChoice_ReorgNeeded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *ForkChoice
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var current *types.Header
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}
		var extern *types.Header
		fill_err = tp.Fill(&extern)
		if fill_err != nil {
			return
		}
		if f1 == nil || current == nil || extern == nil {
			return
		}

		f1.ReorgNeeded(current, extern)
	})
}

func Fuzz_Nosy_GasPool_AddGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if gp == nil {
			return
		}

		gp.AddGas(amount)
	})
}

func Fuzz_Nosy_GasPool_Gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if gp == nil {
			return
		}

		gp.Gas()
	})
}

func Fuzz_Nosy_GasPool_SetGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		var gas uint64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}
		if gp == nil {
			return
		}

		gp.SetGas(gas)
	})
}

func Fuzz_Nosy_GasPool_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if gp == nil {
			return
		}

		gp.String()
	})
}

func Fuzz_Nosy_GasPool_SubGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		var amount uint64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if gp == nil {
			return
		}

		gp.SubGas(amount)
	})
}

// skipping Fuzz_Nosy_Genesis_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_Genesis_MustCommit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_Genesis_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		g := DefaultGenesisBlock()
		g.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Genesis_configOrDefault__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ghash common.Hash
		fill_err = tp.Fill(&ghash)
		if fill_err != nil {
			return
		}

		g := DefaultGenesisBlock()
		g.configOrDefault(ghash)
	})
}

func Fuzz_Nosy_GenesisAccount_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *GenesisAccount
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_GenesisAlloc_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 string, d2 []byte) {
		ga := decodePrealloc(d1)
		ga.UnmarshalJSON(d2)
	})
}

// skipping Fuzz_Nosy_GenesisAlloc_flush__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_GenesisAlloc_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 string, isVerkle bool) {
		ga := decodePrealloc(d1)
		ga.hash(isVerkle)
	})
}

func Fuzz_Nosy_GenesisMismatchError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *GenesisMismatchError
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

func Fuzz_Nosy_HeaderChain_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.Config()
	})
}

func Fuzz_Nosy_HeaderChain_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.CurrentHeader()
	})
}

func Fuzz_Nosy_HeaderChain_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.Engine()
	})
}

func Fuzz_Nosy_HeaderChain_GetAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var ancestor uint64
		fill_err = tp.Fill(&ancestor)
		if fill_err != nil {
			return
		}
		var maxNonCanonical *uint64
		fill_err = tp.Fill(&maxNonCanonical)
		if fill_err != nil {
			return
		}
		if hc == nil || maxNonCanonical == nil {
			return
		}

		hc.GetAncestor(hash, number, ancestor, maxNonCanonical)
	})
}

func Fuzz_Nosy_HeaderChain_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetBlock(hash, number)
	})
}

func Fuzz_Nosy_HeaderChain_GetBlockNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetBlockNumber(hash)
	})
}

func Fuzz_Nosy_HeaderChain_GetCanonicalHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetCanonicalHash(number)
	})
}

func Fuzz_Nosy_HeaderChain_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetHeader(hash, number)
	})
}

func Fuzz_Nosy_HeaderChain_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetHeaderByHash(hash)
	})
}

func Fuzz_Nosy_HeaderChain_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetHeaderByNumber(number)
	})
}

func Fuzz_Nosy_HeaderChain_GetHeadersFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var count uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetHeadersFrom(number, count)
	})
}

func Fuzz_Nosy_HeaderChain_GetTd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.GetTd(hash, number)
	})
}

func Fuzz_Nosy_HeaderChain_HasHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.HasHeader(hash, number)
	})
}

func Fuzz_Nosy_HeaderChain_InsertHeaderChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var chain []*types.Header
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var forker *ForkChoice
		fill_err = tp.Fill(&forker)
		if fill_err != nil {
			return
		}
		if hc == nil || forker == nil {
			return
		}

		hc.InsertHeaderChain(chain, start, forker)
	})
}

func Fuzz_Nosy_HeaderChain_Reorg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.Reorg(headers)
	})
}

func Fuzz_Nosy_HeaderChain_SetCurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if hc == nil || head == nil {
			return
		}

		hc.SetCurrentHeader(head)
	})
}

func Fuzz_Nosy_HeaderChain_SetGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if hc == nil || head == nil {
			return
		}

		hc.SetGenesis(head)
	})
}

// skipping Fuzz_Nosy_HeaderChain_SetHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.UpdateHeadBlocksCallback

// skipping Fuzz_Nosy_HeaderChain_SetHeadWithTimestamp__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.UpdateHeadBlocksCallback

func Fuzz_Nosy_HeaderChain_ValidateHeaderChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var chain []*types.Header
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.ValidateHeaderChain(chain)
	})
}

func Fuzz_Nosy_HeaderChain_WriteHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.WriteHeaders(headers)
	})
}

// skipping Fuzz_Nosy_HeaderChain_setHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.UpdateHeadBlocksCallback

func Fuzz_Nosy_HeaderChain_writeHeadersAndSetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *HeaderChain
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var forker *ForkChoice
		fill_err = tp.Fill(&forker)
		if fill_err != nil {
			return
		}
		if hc == nil || forker == nil {
			return
		}

		hc.writeHeadersAndSetHead(headers, forker)
	})
}

func Fuzz_Nosy_StateProcessor_Process__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *StateProcessor
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var cfg vm.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if p == nil || block == nil || statedb == nil {
			return
		}

		p.Process(block, statedb, cfg)
	})
}

func Fuzz_Nosy_StateTransition_TransitionDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.TransitionDb()
	})
}

func Fuzz_Nosy_StateTransition_blobGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.blobGasUsed()
	})
}

func Fuzz_Nosy_StateTransition_buyGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.buyGas()
	})
}

func Fuzz_Nosy_StateTransition_gasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.gasUsed()
	})
}

func Fuzz_Nosy_StateTransition_preCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.preCheck()
	})
}

func Fuzz_Nosy_StateTransition_refundGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		var refundQuotient uint64
		fill_err = tp.Fill(&refundQuotient)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.refundGas(refundQuotient)
	})
}

func Fuzz_Nosy_StateTransition_to__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evm *vm.EVM
		fill_err = tp.Fill(&evm)
		if fill_err != nil {
			return
		}
		var msg *Message
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var gp *GasPool
		fill_err = tp.Fill(&gp)
		if fill_err != nil {
			return
		}
		if evm == nil || msg == nil || gp == nil {
			return
		}

		st := NewStateTransition(evm, msg, gp)
		st.to()
	})
}

func Fuzz_Nosy_chainMaker_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.Config()
	})
}

func Fuzz_Nosy_chainMaker_CurrentHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.CurrentHeader()
	})
}

func Fuzz_Nosy_chainMaker_Engine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.Engine()
	})
}

func Fuzz_Nosy_chainMaker_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.GetBlock(hash, number)
	})
}

func Fuzz_Nosy_chainMaker_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.GetHeader(hash, number)
	})
}

func Fuzz_Nosy_chainMaker_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.GetHeaderByHash(hash)
	})
}

func Fuzz_Nosy_chainMaker_GetHeaderByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.GetHeaderByNumber(number)
	})
}

func Fuzz_Nosy_chainMaker_GetTd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.GetTd(hash, number)
	})
}

func Fuzz_Nosy_chainMaker_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var b *types.Block
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var r []*types.Receipt
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if cm == nil || b == nil {
			return
		}

		cm.add(b, r)
	})
}

func Fuzz_Nosy_chainMaker_blockByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMaker
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.blockByNumber(number)
	})
}

// skipping Fuzz_Nosy_chainMaker_makeHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

func Fuzz_Nosy_insertIterator_current__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.current()
	})
}

func Fuzz_Nosy_insertIterator_first__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.first()
	})
}

func Fuzz_Nosy_insertIterator_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.next()
	})
}

func Fuzz_Nosy_insertIterator_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.peek()
	})
}

func Fuzz_Nosy_insertIterator_previous__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.previous()
	})
}

func Fuzz_Nosy_insertIterator_processed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.processed()
	})
}

func Fuzz_Nosy_insertIterator_remaining__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *insertIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.remaining()
	})
}

func Fuzz_Nosy_insertStats_report__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var st *insertStats
		fill_err = tp.Fill(&st)
		if fill_err != nil {
			return
		}
		var chain []*types.Block
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var snapDiffItems common.StorageSize
		fill_err = tp.Fill(&snapDiffItems)
		if fill_err != nil {
			return
		}
		var snapBufItems common.StorageSize
		fill_err = tp.Fill(&snapBufItems)
		if fill_err != nil {
			return
		}
		var trieDiffNodes common.StorageSize
		fill_err = tp.Fill(&trieDiffNodes)
		if fill_err != nil {
			return
		}
		var triebufNodes common.StorageSize
		fill_err = tp.Fill(&triebufNodes)
		if fill_err != nil {
			return
		}
		var setHead bool
		fill_err = tp.Fill(&setHead)
		if fill_err != nil {
			return
		}
		if st == nil {
			return
		}

		st.report(chain, index, snapDiffItems, snapBufItems, trieDiffNodes, triebufNodes, setHead)
	})
}

func Fuzz_Nosy_statePrefetcher_Prefetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *statePrefetcher
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		var cfg vm.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var interrupt *atomic.Bool
		fill_err = tp.Fill(&interrupt)
		if fill_err != nil {
			return
		}
		if p == nil || block == nil || statedb == nil || interrupt == nil {
			return
		}

		p.Prefetch(block, statedb, cfg, interrupt)
	})
}

func Fuzz_Nosy_storageJSON_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *storageJSON
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.UnmarshalText(text)
	})
}

// skipping Fuzz_Nosy_txSenderCacher_Recover__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

// skipping Fuzz_Nosy_txSenderCacher_RecoverFromBlocks__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_txSenderCacher_cache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, threads int) {
		cacher := newTxSenderCacher(threads)
		cacher.cache()
	})
}

// skipping Fuzz_Nosy_ChainContext_Engine__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainContext

// skipping Fuzz_Nosy_ChainContext_GetHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainContext

// skipping Fuzz_Nosy_ChainIndexerBackend_Commit__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerBackend

// skipping Fuzz_Nosy_ChainIndexerBackend_Process__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerBackend

// skipping Fuzz_Nosy_ChainIndexerBackend_Prune__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerBackend

// skipping Fuzz_Nosy_ChainIndexerBackend_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerBackend

// skipping Fuzz_Nosy_ChainIndexerChain_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerChain

// skipping Fuzz_Nosy_ChainIndexerChain_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainIndexerChain

// skipping Fuzz_Nosy_ChainReader_Config__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainReader

// skipping Fuzz_Nosy_ChainReader_GetTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainReader

func Fuzz_Nosy_GenesisAccount_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g GenesisAccount
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}

		g.MarshalJSON()
	})
}

// skipping Fuzz_Nosy_Prefetcher_Prefetch__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.Prefetcher

// skipping Fuzz_Nosy_Processor_Process__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.Processor

// skipping Fuzz_Nosy_Validator_ValidateBody__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.Validator

// skipping Fuzz_Nosy_Validator_ValidateState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.Validator

// skipping Fuzz_Nosy_codeReader_ContractCodeWithPrefix__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.codeReader

func Fuzz_Nosy_storageJSON_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h storageJSON
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.MarshalText()
	})
}

func Fuzz_Nosy_CalcGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, parentGasLimit uint64, desiredLimit uint64) {
		CalcGasLimit(parentGasLimit, desiredLimit)
	})
}

// skipping Fuzz_Nosy_CanTransfer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_GenerateChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_GenerateChainWithGenesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_GetHashFn__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core.ChainContext

func Fuzz_Nosy_IntrinsicGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 []byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var accessList types.AccessList
		fill_err = tp.Fill(&accessList)
		if fill_err != nil {
			return
		}
		var isContractCreation bool
		fill_err = tp.Fill(&isContractCreation)
		if fill_err != nil {
			return
		}
		var isHomestead bool
		fill_err = tp.Fill(&isHomestead)
		if fill_err != nil {
			return
		}
		var isEIP2028 bool
		fill_err = tp.Fill(&isEIP2028)
		if fill_err != nil {
			return
		}
		var isEIP3860 bool
		fill_err = tp.Fill(&isEIP3860)
		if fill_err != nil {
			return
		}

		IntrinsicGas(d1, accessList, isContractCreation, isHomestead, isEIP2028, isEIP3860)
	})
}

func Fuzz_Nosy_ProcessBeaconBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var beaconRoot common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		var vmenv *vm.EVM
		fill_err = tp.Fill(&vmenv)
		if fill_err != nil {
			return
		}
		var statedb *state.StateDB
		fill_err = tp.Fill(&statedb)
		if fill_err != nil {
			return
		}
		if vmenv == nil || statedb == nil {
			return
		}

		ProcessBeaconBlockRoot(beaconRoot, vmenv, statedb)
	})
}

// skipping Fuzz_Nosy_SetupGenesisBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_SetupGenesisBlockWithOverride__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

// skipping Fuzz_Nosy_Transfer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/vm.StateDB

// skipping Fuzz_Nosy_makeBlockChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_makeBlockChainWithGenesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_makeHeaderChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_makeHeaderChainWithGenesis__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.Engine

// skipping Fuzz_Nosy_summarizeBadBlock__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_toWordSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64) {
		toWordSize(size)
	})
}
