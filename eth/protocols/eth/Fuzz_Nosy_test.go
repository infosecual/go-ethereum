package eth

import (
	"io"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/rlp"
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

func Fuzz_Nosy_BlockBodiesResponse_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockBodiesResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_BlockBodiesResponse_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockBodiesResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_BlockBodiesResponse_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *BlockBodiesResponse
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Unpack()
	})
}

func Fuzz_Nosy_BlockHeadersRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockHeadersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_BlockHeadersRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockHeadersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetBlockBodiesRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetBlockBodiesRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetBlockBodiesRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetBlockBodiesRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetBlockHeadersRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetBlockHeadersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetBlockHeadersRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetBlockHeadersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetPooledTransactionsRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetPooledTransactionsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetPooledTransactionsRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetPooledTransactionsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_GetReceiptsRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetReceiptsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_GetReceiptsRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GetReceiptsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_HashOrNumber_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hn *HashOrNumber
		fill_err = tp.Fill(&hn)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if hn == nil || s == nil {
			return
		}

		hn.DecodeRLP(s)
	})
}

func Fuzz_Nosy_HashOrNumber_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hn *HashOrNumber
		fill_err = tp.Fill(&hn)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if hn == nil {
			return
		}

		hn.EncodeRLP(w)
	})
}

func Fuzz_Nosy_NewBlockHashesPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewBlockHashesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_NewBlockHashesPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewBlockHashesPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_NewBlockHashesPacket_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *NewBlockHashesPacket
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Unpack()
	})
}

func Fuzz_Nosy_NewBlockPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewBlockPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_NewBlockPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewBlockPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_NewBlockPacket_sanityCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *NewBlockPacket
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		request.sanityCheck()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashesPacket67_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewPooledTransactionHashesPacket67
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashesPacket67_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewPooledTransactionHashesPacket67
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashesPacket68_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewPooledTransactionHashesPacket68
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashesPacket68_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *NewPooledTransactionHashesPacket68
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Peer_AsyncSendNewBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
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
		if p == nil || block == nil || td == nil {
			return
		}

		p.AsyncSendNewBlock(block, td)
	})
}

func Fuzz_Nosy_Peer_AsyncSendNewBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if p == nil || block == nil {
			return
		}

		p.AsyncSendNewBlockHash(block)
	})
}

func Fuzz_Nosy_Peer_AsyncSendPooledTransactionHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AsyncSendPooledTransactionHashes(hashes)
	})
}

func Fuzz_Nosy_Peer_AsyncSendTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.AsyncSendTransactions(hashes)
	})
}

func Fuzz_Nosy_Peer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
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

// skipping Fuzz_Nosy_Peer_Handshake__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/forkid.Filter

func Fuzz_Nosy_Peer_Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Head()
	})
}

func Fuzz_Nosy_Peer_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ID()
	})
}

func Fuzz_Nosy_Peer_KnownBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
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

		p.KnownBlock(hash)
	})
}

func Fuzz_Nosy_Peer_KnownTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
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

		p.KnownTransaction(hash)
	})
}

func Fuzz_Nosy_Peer_ReplyBlockBodiesRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var bodies []rlp.RawValue
		fill_err = tp.Fill(&bodies)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReplyBlockBodiesRLP(id, bodies)
	})
}

func Fuzz_Nosy_Peer_ReplyBlockHeadersRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var headers []rlp.RawValue
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReplyBlockHeadersRLP(id, headers)
	})
}

func Fuzz_Nosy_Peer_ReplyPooledTransactionsRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var txs []rlp.RawValue
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReplyPooledTransactionsRLP(id, hashes, txs)
	})
}

func Fuzz_Nosy_Peer_ReplyReceiptsRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var receipts []rlp.RawValue
		fill_err = tp.Fill(&receipts)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReplyReceiptsRLP(id, receipts)
	})
}

// skipping Fuzz_Nosy_Peer_RequestBodies__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

// skipping Fuzz_Nosy_Peer_RequestHeadersByHash__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

// skipping Fuzz_Nosy_Peer_RequestHeadersByNumber__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

// skipping Fuzz_Nosy_Peer_RequestOneHeader__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

// skipping Fuzz_Nosy_Peer_RequestReceipts__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

func Fuzz_Nosy_Peer_RequestTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestTxs(hashes)
	})
}

func Fuzz_Nosy_Peer_SendNewBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
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
		if p == nil || block == nil || td == nil {
			return
		}

		p.SendNewBlock(block, td)
	})
}

func Fuzz_Nosy_Peer_SendNewBlockHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var numbers []uint64
		fill_err = tp.Fill(&numbers)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SendNewBlockHashes(hashes, numbers)
	})
}

func Fuzz_Nosy_Peer_SendTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var txs types.Transactions
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SendTransactions(txs)
	})
}

func Fuzz_Nosy_Peer_SetHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if p == nil || td == nil {
			return
		}

		p.SetHead(hash, td)
	})
}

func Fuzz_Nosy_Peer_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Version()
	})
}

func Fuzz_Nosy_Peer_announceTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.announceTransactions()
	})
}

func Fuzz_Nosy_Peer_broadcastBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.broadcastBlocks()
	})
}

func Fuzz_Nosy_Peer_broadcastTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.broadcastTransactions()
	})
}

func Fuzz_Nosy_Peer_dispatchRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var req *Request
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if p == nil || req == nil {
			return
		}

		p.dispatchRequest(req)
	})
}

// skipping Fuzz_Nosy_Peer_dispatchResponse__ because parameters include func, chan, or unsupported interface: func() interface{}

func Fuzz_Nosy_Peer_dispatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.dispatcher()
	})
}

func Fuzz_Nosy_Peer_markBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
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

		p.markBlock(hash)
	})
}

func Fuzz_Nosy_Peer_markTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
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

		p.markTransaction(hash)
	})
}

// skipping Fuzz_Nosy_Peer_readStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/forkid.Filter

func Fuzz_Nosy_Peer_sendPooledTransactionHashes66__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.sendPooledTransactionHashes66(hashes)
	})
}

func Fuzz_Nosy_Peer_sendPooledTransactionHashes68__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
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
		if p == nil {
			return
		}

		p.sendPooledTransactionHashes68(hashes, types, sizes)
	})
}

func Fuzz_Nosy_PooledTransactionsResponse_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *PooledTransactionsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_PooledTransactionsResponse_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *PooledTransactionsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_ReceiptsResponse_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ReceiptsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_ReceiptsResponse_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ReceiptsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Request_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Close()
	})
}

func Fuzz_Nosy_StatusPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StatusPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_StatusPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StatusPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_TransactionsPacket_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TransactionsPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_TransactionsPacket_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TransactionsPacket
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_bidirectionalMeters_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *bidirectionalMeters
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ingress bool
		fill_err = tp.Fill(&ingress)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.get(ingress)
	})
}

func Fuzz_Nosy_knownCache_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}

		k := newKnownCache(max)
		k.Add(hashes...)
	})
}

func Fuzz_Nosy_knownCache_Cardinality__(f *testing.F) {
	f.Fuzz(func(t *testing.T, max int) {
		k := newKnownCache(max)
		k.Cardinality()
	})
}

func Fuzz_Nosy_knownCache_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		k := newKnownCache(max)
		k.Contains(hash)
	})
}

// skipping Fuzz_Nosy_Backend_AcceptTxs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Backend_Chain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Backend_Handle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Backend_PeerInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Backend_RunPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Backend_TxPool__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_Decoder_Decode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Decoder

// skipping Fuzz_Nosy_Decoder_Time__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Decoder

// skipping Fuzz_Nosy_Packet_Kind__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Packet

// skipping Fuzz_Nosy_Packet_Name__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Packet

// skipping Fuzz_Nosy_TxPool_Get__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.TxPool

func Fuzz_Nosy_enrEntry_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		e := currentENREntry(chain)
		e.ENRKey()
	})
}

// skipping Fuzz_Nosy_MakeProtocols__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

func Fuzz_Nosy_ServiceGetBlockBodiesQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var query GetBlockBodiesRequest
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		ServiceGetBlockBodiesQuery(chain, query)
	})
}

func Fuzz_Nosy_ServiceGetBlockHeadersQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var query *GetBlockHeadersRequest
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var peer *Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if chain == nil || query == nil || peer == nil {
			return
		}

		ServiceGetBlockHeadersQuery(chain, query, peer)
	})
}

func Fuzz_Nosy_ServiceGetReceiptsQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var query GetReceiptsRequest
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		ServiceGetReceiptsQuery(chain, query)
	})
}

func Fuzz_Nosy_StartENRUpdater__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var ln *enode.LocalNode
		fill_err = tp.Fill(&ln)
		if fill_err != nil {
			return
		}
		if chain == nil || ln == nil {
			return
		}

		StartENRUpdater(chain, ln)
	})
}

// skipping Fuzz_Nosy_answerGetPooledTransactions__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/eth.Backend

// skipping Fuzz_Nosy_markError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_max__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		max(a, b)
	})
}

func Fuzz_Nosy_serviceContiguousBlockHeaderQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var query *GetBlockHeadersRequest
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if chain == nil || query == nil {
			return
		}

		serviceContiguousBlockHeaderQuery(chain, query)
	})
}

func Fuzz_Nosy_serviceNonContiguousBlockHeaderQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var query *GetBlockHeadersRequest
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var peer *Peer
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if chain == nil || query == nil || peer == nil {
			return
		}

		serviceNonContiguousBlockHeaderQuery(chain, query, peer)
	})
}
