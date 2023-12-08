package types

import (
	"bytes"
	"io"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
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

func Fuzz_Nosy_AccessListTx_accessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.accessList()
	})
}

func Fuzz_Nosy_AccessListTx_chainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.chainID()
	})
}

func Fuzz_Nosy_AccessListTx_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.copy()
	})
}

func Fuzz_Nosy_AccessListTx_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.data()
	})
}

func Fuzz_Nosy_AccessListTx_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.decode(input)
	})
}

func Fuzz_Nosy_AccessListTx_effectiveGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var dst *big.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if tx == nil || dst == nil || baseFee == nil {
			return
		}

		tx.effectiveGasPrice(dst, baseFee)
	})
}

func Fuzz_Nosy_AccessListTx_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tx == nil || b == nil {
			return
		}

		tx.encode(b)
	})
}

func Fuzz_Nosy_AccessListTx_gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gas()
	})
}

func Fuzz_Nosy_AccessListTx_gasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasFeeCap()
	})
}

func Fuzz_Nosy_AccessListTx_gasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasPrice()
	})
}

func Fuzz_Nosy_AccessListTx_gasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasTipCap()
	})
}

func Fuzz_Nosy_AccessListTx_nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.nonce()
	})
}

func Fuzz_Nosy_AccessListTx_rawSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.rawSignatureValues()
	})
}

func Fuzz_Nosy_AccessListTx_setSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r *big.Int
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil || v == nil || r == nil || s == nil {
			return
		}

		tx.setSignatureValues(chainID, v, r, s)
	})
}

func Fuzz_Nosy_AccessListTx_to__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.to()
	})
}

func Fuzz_Nosy_AccessListTx_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.txType()
	})
}

func Fuzz_Nosy_AccessListTx_value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *AccessListTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.value()
	})
}

func Fuzz_Nosy_AccessTuple_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AccessTuple
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

func Fuzz_Nosy_BlobTx_accessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.accessList()
	})
}

func Fuzz_Nosy_BlobTx_blobGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.blobGas()
	})
}

func Fuzz_Nosy_BlobTx_chainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.chainID()
	})
}

func Fuzz_Nosy_BlobTx_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.copy()
	})
}

func Fuzz_Nosy_BlobTx_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.data()
	})
}

func Fuzz_Nosy_BlobTx_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.decode(input)
	})
}

func Fuzz_Nosy_BlobTx_effectiveGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var dst *big.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if tx == nil || dst == nil || baseFee == nil {
			return
		}

		tx.effectiveGasPrice(dst, baseFee)
	})
}

func Fuzz_Nosy_BlobTx_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tx == nil || b == nil {
			return
		}

		tx.encode(b)
	})
}

func Fuzz_Nosy_BlobTx_gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gas()
	})
}

func Fuzz_Nosy_BlobTx_gasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasFeeCap()
	})
}

func Fuzz_Nosy_BlobTx_gasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasPrice()
	})
}

func Fuzz_Nosy_BlobTx_gasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasTipCap()
	})
}

func Fuzz_Nosy_BlobTx_nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.nonce()
	})
}

func Fuzz_Nosy_BlobTx_rawSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.rawSignatureValues()
	})
}

func Fuzz_Nosy_BlobTx_setSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r *big.Int
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil || v == nil || r == nil || s == nil {
			return
		}

		tx.setSignatureValues(chainID, v, r, s)
	})
}

func Fuzz_Nosy_BlobTx_to__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.to()
	})
}

func Fuzz_Nosy_BlobTx_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.txType()
	})
}

func Fuzz_Nosy_BlobTx_value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.value()
	})
}

func Fuzz_Nosy_BlobTx_withoutSidecar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *BlobTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.withoutSidecar()
	})
}

func Fuzz_Nosy_BlobTxSidecar_BlobHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *BlobTxSidecar
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.BlobHashes()
	})
}

func Fuzz_Nosy_BlobTxSidecar_encodedSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *BlobTxSidecar
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.encodedSize()
	})
}

func Fuzz_Nosy_Block_BaseFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.BaseFee()
	})
}

func Fuzz_Nosy_Block_BeaconRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.BeaconRoot()
	})
}

func Fuzz_Nosy_Block_BlobGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.BlobGasUsed()
	})
}

func Fuzz_Nosy_Block_Bloom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Bloom()
	})
}

func Fuzz_Nosy_Block_Body__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Body()
	})
}

func Fuzz_Nosy_Block_Coinbase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Coinbase()
	})
}

func Fuzz_Nosy_Block_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if header == nil || s == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.DecodeRLP(s)
	})
}

func Fuzz_Nosy_Block_Difficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Difficulty()
	})
}

func Fuzz_Nosy_Block_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.EncodeRLP(w)
	})
}

func Fuzz_Nosy_Block_ExcessBlobGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.ExcessBlobGas()
	})
}

func Fuzz_Nosy_Block_Extra__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Extra()
	})
}

func Fuzz_Nosy_Block_GasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.GasLimit()
	})
}

func Fuzz_Nosy_Block_GasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.GasUsed()
	})
}

func Fuzz_Nosy_Block_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Hash()
	})
}

func Fuzz_Nosy_Block_Header__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Header()
	})
}

func Fuzz_Nosy_Block_MixDigest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.MixDigest()
	})
}

func Fuzz_Nosy_Block_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Nonce()
	})
}

func Fuzz_Nosy_Block_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Number()
	})
}

func Fuzz_Nosy_Block_NumberU64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.NumberU64()
	})
}

func Fuzz_Nosy_Block_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.ParentHash()
	})
}

func Fuzz_Nosy_Block_ReceiptHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.ReceiptHash()
	})
}

func Fuzz_Nosy_Block_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Root()
	})
}

func Fuzz_Nosy_Block_SanityCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.SanityCheck()
	})
}

func Fuzz_Nosy_Block_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Size()
	})
}

func Fuzz_Nosy_Block_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Time()
	})
}

func Fuzz_Nosy_Block_Transaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Transaction(hash)
	})
}

func Fuzz_Nosy_Block_Transactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Transactions()
	})
}

func Fuzz_Nosy_Block_TxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.TxHash()
	})
}

func Fuzz_Nosy_Block_UncleHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.UncleHash()
	})
}

func Fuzz_Nosy_Block_Uncles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Uncles()
	})
}

func Fuzz_Nosy_Block_WithBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var transactions []*Transaction
		fill_err = tp.Fill(&transactions)
		if fill_err != nil {
			return
		}
		var uncles []*Header
		fill_err = tp.Fill(&uncles)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.WithBody(transactions, uncles)
	})
}

func Fuzz_Nosy_Block_WithSeal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h1 *Header
		fill_err = tp.Fill(&h1)
		if fill_err != nil {
			return
		}
		var h2 *Header
		fill_err = tp.Fill(&h2)
		if fill_err != nil {
			return
		}
		if h1 == nil || h2 == nil {
			return
		}

		b := NewBlockWithHeader(h1)
		b.WithSeal(h2)
	})
}

func Fuzz_Nosy_Block_WithWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var withdrawals []*Withdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.WithWithdrawals(withdrawals)
	})
}

func Fuzz_Nosy_Block_Withdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		b := NewBlockWithHeader(header)
		b.Withdrawals()
	})
}

func Fuzz_Nosy_BlockNonce_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64, input []byte) {
		n := EncodeNonce(i)
		n.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Bloom_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, d []byte) {
		b1 := BytesToBloom(b)
		b1.Add(d)
	})
}

func Fuzz_Nosy_Bloom_SetBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, d []byte) {
		b1 := BytesToBloom(b)
		b1.SetBytes(d)
	})
}

func Fuzz_Nosy_Bloom_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, input []byte) {
		b1 := BytesToBloom(b)
		b1.UnmarshalText(input)
	})
}

func Fuzz_Nosy_Bloom_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, d []byte, buf []byte) {
		b1 := BytesToBloom(b)
		b1.add(d, buf)
	})
}

func Fuzz_Nosy_DynamicFeeTx_accessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.accessList()
	})
}

func Fuzz_Nosy_DynamicFeeTx_chainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.chainID()
	})
}

func Fuzz_Nosy_DynamicFeeTx_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.copy()
	})
}

func Fuzz_Nosy_DynamicFeeTx_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.data()
	})
}

func Fuzz_Nosy_DynamicFeeTx_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.decode(input)
	})
}

func Fuzz_Nosy_DynamicFeeTx_effectiveGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var dst *big.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if tx == nil || dst == nil || baseFee == nil {
			return
		}

		tx.effectiveGasPrice(dst, baseFee)
	})
}

func Fuzz_Nosy_DynamicFeeTx_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tx == nil || b == nil {
			return
		}

		tx.encode(b)
	})
}

func Fuzz_Nosy_DynamicFeeTx_gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gas()
	})
}

func Fuzz_Nosy_DynamicFeeTx_gasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasFeeCap()
	})
}

func Fuzz_Nosy_DynamicFeeTx_gasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasPrice()
	})
}

func Fuzz_Nosy_DynamicFeeTx_gasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasTipCap()
	})
}

func Fuzz_Nosy_DynamicFeeTx_nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.nonce()
	})
}

func Fuzz_Nosy_DynamicFeeTx_rawSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.rawSignatureValues()
	})
}

func Fuzz_Nosy_DynamicFeeTx_setSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r *big.Int
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil || v == nil || r == nil || s == nil {
			return
		}

		tx.setSignatureValues(chainID, v, r, s)
	})
}

func Fuzz_Nosy_DynamicFeeTx_to__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.to()
	})
}

func Fuzz_Nosy_DynamicFeeTx_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.txType()
	})
}

func Fuzz_Nosy_DynamicFeeTx_value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *DynamicFeeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.value()
	})
}

func Fuzz_Nosy_Header_EmptyBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.EmptyBody()
	})
}

func Fuzz_Nosy_Header_EmptyReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.EmptyReceipts()
	})
}

func Fuzz_Nosy_Header_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		obj := CopyHeader(h)
		obj.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_Header_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.Hash()
	})
}

func Fuzz_Nosy_Header_SanityCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.SanityCheck()
	})
}

func Fuzz_Nosy_Header_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.Size()
	})
}

func Fuzz_Nosy_Header_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
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

		h1 := CopyHeader(h)
		h1.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_LegacyTx_accessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.accessList()
	})
}

func Fuzz_Nosy_LegacyTx_chainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.chainID()
	})
}

func Fuzz_Nosy_LegacyTx_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.copy()
	})
}

func Fuzz_Nosy_LegacyTx_data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.data()
	})
}

func Fuzz_Nosy_LegacyTx_decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.decode(_x2)
	})
}

func Fuzz_Nosy_LegacyTx_effectiveGasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var dst *big.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if tx == nil || dst == nil || baseFee == nil {
			return
		}

		tx.effectiveGasPrice(dst, baseFee)
	})
}

func Fuzz_Nosy_LegacyTx_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var _x2 *bytes.Buffer
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if tx == nil || _x2 == nil {
			return
		}

		tx.encode(_x2)
	})
}

func Fuzz_Nosy_LegacyTx_gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gas()
	})
}

func Fuzz_Nosy_LegacyTx_gasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasFeeCap()
	})
}

func Fuzz_Nosy_LegacyTx_gasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasPrice()
	})
}

func Fuzz_Nosy_LegacyTx_gasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.gasTipCap()
	})
}

func Fuzz_Nosy_LegacyTx_nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.nonce()
	})
}

func Fuzz_Nosy_LegacyTx_rawSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.rawSignatureValues()
	})
}

func Fuzz_Nosy_LegacyTx_setSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var v *big.Int
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var r *big.Int
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil || v == nil || r == nil || s == nil {
			return
		}

		tx.setSignatureValues(chainID, v, r, s)
	})
}

func Fuzz_Nosy_LegacyTx_to__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.to()
	})
}

func Fuzz_Nosy_LegacyTx_txType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.txType()
	})
}

func Fuzz_Nosy_LegacyTx_value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *LegacyTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.value()
	})
}

func Fuzz_Nosy_Log_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var obj *Log
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}
		if obj == nil {
			return
		}

		obj.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_Log_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *Log
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Receipt_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root []byte
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var failed bool
		fill_err = tp.Fill(&failed)
		if fill_err != nil {
			return
		}
		var cumulativeGasUsed uint64
		fill_err = tp.Fill(&cumulativeGasUsed)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.DecodeRLP(s)
	})
}

func Fuzz_Nosy_Receipt_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root []byte
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var failed bool
		fill_err = tp.Fill(&failed)
		if fill_err != nil {
			return
		}
		var cumulativeGasUsed uint64
		fill_err = tp.Fill(&cumulativeGasUsed)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.EncodeRLP(w)
	})
}

func Fuzz_Nosy_Receipt_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.MarshalBinary()
	})
}

func Fuzz_Nosy_Receipt_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.Size()
	})
}

func Fuzz_Nosy_Receipt_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64, b []byte) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.UnmarshalBinary(b)
	})
}

func Fuzz_Nosy_Receipt_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64, input []byte) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Receipt_decodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64, b []byte) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.decodeTyped(b)
	})
}

func Fuzz_Nosy_Receipt_encodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root []byte
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var failed bool
		fill_err = tp.Fill(&failed)
		if fill_err != nil {
			return
		}
		var cumulativeGasUsed uint64
		fill_err = tp.Fill(&cumulativeGasUsed)
		if fill_err != nil {
			return
		}
		var d4 *receiptRLP
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if d4 == nil || w == nil {
			return
		}

		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.encodeTyped(d4, w)
	})
}

func Fuzz_Nosy_Receipt_setFromRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root []byte
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var failed bool
		fill_err = tp.Fill(&failed)
		if fill_err != nil {
			return
		}
		var cumulativeGasUsed uint64
		fill_err = tp.Fill(&cumulativeGasUsed)
		if fill_err != nil {
			return
		}
		var d4 receiptRLP
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}

		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.setFromRLP(d4)
	})
}

func Fuzz_Nosy_Receipt_setStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64, postStateOrStatus []byte) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.setStatus(postStateOrStatus)
	})
}

func Fuzz_Nosy_Receipt_statusEncoding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.statusEncoding()
	})
}

func Fuzz_Nosy_ReceiptForStorage_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *ReceiptForStorage
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if r == nil || s == nil {
			return
		}

		r.DecodeRLP(s)
	})
}

func Fuzz_Nosy_ReceiptForStorage_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *ReceiptForStorage
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_StateAccount_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}

		obj := NewEmptyStateAccount()
		obj.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_Transaction_AccessList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.AccessList()
	})
}

func Fuzz_Nosy_Transaction_BlobGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobGas()
	})
}

func Fuzz_Nosy_Transaction_BlobGasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobGasFeeCap()
	})
}

func Fuzz_Nosy_Transaction_BlobGasFeeCapCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *Transaction
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobGasFeeCapCmp(other)
	})
}

func Fuzz_Nosy_Transaction_BlobGasFeeCapIntCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *big.Int
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobGasFeeCapIntCmp(other)
	})
}

func Fuzz_Nosy_Transaction_BlobHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobHashes()
	})
}

func Fuzz_Nosy_Transaction_BlobTxSidecar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.BlobTxSidecar()
	})
}

func Fuzz_Nosy_Transaction_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.ChainId()
	})
}

func Fuzz_Nosy_Transaction_Cost__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Cost()
	})
}

func Fuzz_Nosy_Transaction_Data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Data()
	})
}

func Fuzz_Nosy_Transaction_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || s == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.DecodeRLP(s)
	})
}

func Fuzz_Nosy_Transaction_EffectiveGasTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || baseFee == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.EffectiveGasTip(baseFee)
	})
}

func Fuzz_Nosy_Transaction_EffectiveGasTipCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *Transaction
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil || baseFee == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.EffectiveGasTipCmp(other, baseFee)
	})
}

func Fuzz_Nosy_Transaction_EffectiveGasTipIntCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *big.Int
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil || baseFee == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.EffectiveGasTipIntCmp(other, baseFee)
	})
}

func Fuzz_Nosy_Transaction_EffectiveGasTipValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || baseFee == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.EffectiveGasTipValue(baseFee)
	})
}

func Fuzz_Nosy_Transaction_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.EncodeRLP(w)
	})
}

func Fuzz_Nosy_Transaction_Gas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Gas()
	})
}

func Fuzz_Nosy_Transaction_GasFeeCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasFeeCap()
	})
}

func Fuzz_Nosy_Transaction_GasFeeCapCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *Transaction
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasFeeCapCmp(other)
	})
}

func Fuzz_Nosy_Transaction_GasFeeCapIntCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *big.Int
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasFeeCapIntCmp(other)
	})
}

func Fuzz_Nosy_Transaction_GasPrice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasPrice()
	})
}

func Fuzz_Nosy_Transaction_GasTipCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasTipCap()
	})
}

func Fuzz_Nosy_Transaction_GasTipCapCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *Transaction
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasTipCapCmp(other)
	})
}

func Fuzz_Nosy_Transaction_GasTipCapIntCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var other *big.Int
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || other == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.GasTipCapIntCmp(other)
	})
}

func Fuzz_Nosy_Transaction_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Hash()
	})
}

func Fuzz_Nosy_Transaction_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.MarshalBinary()
	})
}

func Fuzz_Nosy_Transaction_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.MarshalJSON()
	})
}

func Fuzz_Nosy_Transaction_Nonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Nonce()
	})
}

func Fuzz_Nosy_Transaction_Protected__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Protected()
	})
}

func Fuzz_Nosy_Transaction_RawSignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.RawSignatureValues()
	})
}

func Fuzz_Nosy_Transaction_SetTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var t6 time.Time
		fill_err = tp.Fill(&t6)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.SetTime(t6)
	})
}

func Fuzz_Nosy_Transaction_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Size()
	})
}

func Fuzz_Nosy_Transaction_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Time()
	})
}

func Fuzz_Nosy_Transaction_To__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.To()
	})
}

func Fuzz_Nosy_Transaction_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Type()
	})
}

func Fuzz_Nosy_Transaction_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.UnmarshalBinary(b)
	})
}

func Fuzz_Nosy_Transaction_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Transaction_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.Value()
	})
}

// skipping Fuzz_Nosy_Transaction_WithSignature__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_Transaction_WithoutBlobTxSidecar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.WithoutBlobTxSidecar()
	})
}

func Fuzz_Nosy_Transaction_decodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.decodeTyped(b)
	})
}

func Fuzz_Nosy_Transaction_encodeTyped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nonce uint64
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var amount *big.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var gasLimit uint64
		fill_err = tp.Fill(&gasLimit)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if amount == nil || gasPrice == nil || w == nil {
			return
		}

		tx := NewContractCreation(nonce, amount, gasLimit, gasPrice, d5)
		tx.encodeTyped(w)
	})
}

// skipping Fuzz_Nosy_Transaction_setDecoded__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

func Fuzz_Nosy_Withdrawal_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var obj *Withdrawal
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}
		if obj == nil {
			return
		}

		obj.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_Withdrawal_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_txJSON_yParityValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *txJSON
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.yParityValue()
	})
}

func Fuzz_Nosy_writeCounter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *writeCounter
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Write(b)
	})
}

func Fuzz_Nosy_AccessList_StorageKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var al AccessList
		fill_err = tp.Fill(&al)
		if fill_err != nil {
			return
		}

		al.StorageKeys()
	})
}

func Fuzz_Nosy_AccessTuple_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AccessTuple
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.MarshalJSON()
	})
}

func Fuzz_Nosy_BlockNonce_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		n := EncodeNonce(i)
		n.MarshalText()
	})
}

func Fuzz_Nosy_BlockNonce_Uint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		n := EncodeNonce(i)
		n.Uint64()
	})
}

func Fuzz_Nosy_Bloom_Big__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		b1 := BytesToBloom(b)
		b1.Big()
	})
}

func Fuzz_Nosy_Bloom_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		b1 := BytesToBloom(b)
		b1.Bytes()
	})
}

func Fuzz_Nosy_Bloom_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		b1 := BytesToBloom(b)
		b1.MarshalText()
	})
}

func Fuzz_Nosy_Bloom_Test__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, topic []byte) {
		b1 := BytesToBloom(b)
		b1.Test(topic)
	})
}

// skipping Fuzz_Nosy_DerivableList_EncodeIndex__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.DerivableList

// skipping Fuzz_Nosy_DerivableList_Len__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.DerivableList

func Fuzz_Nosy_EIP155Signer_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if chainId == nil {
			return
		}

		s := NewEIP155Signer(chainId)
		s.ChainID()
	})
}

// skipping Fuzz_Nosy_EIP155Signer_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_EIP155Signer_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		s := NewEIP155Signer(chainId)
		s.Hash(tx)
	})
}

func Fuzz_Nosy_EIP155Signer_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		s := NewEIP155Signer(chainId)
		s.Sender(tx)
	})
}

func Fuzz_Nosy_EIP155Signer_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		s := NewEIP155Signer(chainId)
		s.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_FrontierSigner_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s FrontierSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ChainID()
	})
}

// skipping Fuzz_Nosy_FrontierSigner_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_FrontierSigner_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fs FrontierSigner
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		fs.Hash(tx)
	})
}

func Fuzz_Nosy_FrontierSigner_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fs FrontierSigner
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		fs.Sender(tx)
	})
}

func Fuzz_Nosy_FrontierSigner_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fs FrontierSigner
		fill_err = tp.Fill(&fs)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		fs.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_Header_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h1 := CopyHeader(h)
		h1.MarshalJSON()
	})
}

func Fuzz_Nosy_HomesteadSigner_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s HomesteadSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ChainID()
	})
}

// skipping Fuzz_Nosy_HomesteadSigner_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_HomesteadSigner_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hs HomesteadSigner
		fill_err = tp.Fill(&hs)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		hs.Sender(tx)
	})
}

func Fuzz_Nosy_HomesteadSigner_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hs HomesteadSigner
		fill_err = tp.Fill(&hs)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		hs.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_Log_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l Log
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.MarshalJSON()
	})
}

func Fuzz_Nosy_Receipt_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, root []byte, failed bool, cumulativeGasUsed uint64) {
		r := NewReceipt(root, failed, cumulativeGasUsed)
		r.MarshalJSON()
	})
}

func Fuzz_Nosy_Receipts_DeriveFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs Receipts
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var config *params.ChainConfig
		fill_err = tp.Fill(&config)
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
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		var baseFee *big.Int
		fill_err = tp.Fill(&baseFee)
		if fill_err != nil {
			return
		}
		var blobGasPrice *big.Int
		fill_err = tp.Fill(&blobGasPrice)
		if fill_err != nil {
			return
		}
		var txs []*Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if config == nil || baseFee == nil || blobGasPrice == nil {
			return
		}

		rs.DeriveFields(config, hash, number, time, baseFee, blobGasPrice, txs)
	})
}

func Fuzz_Nosy_Receipts_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs Receipts
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		rs.EncodeIndex(i, w)
	})
}

func Fuzz_Nosy_Receipts_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs Receipts
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}

		rs.Len()
	})
}

func Fuzz_Nosy_Signer_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		if chainId == nil {
			return
		}

		_x1 := NewLondonSigner(chainId)
		_x1.ChainID()
	})
}

// skipping Fuzz_Nosy_Signer_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_Signer_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		_x1 := NewLondonSigner(chainId)
		_x1.Hash(tx)
	})
}

func Fuzz_Nosy_Signer_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		_x1 := NewLondonSigner(chainId)
		_x1.Sender(tx)
	})
}

func Fuzz_Nosy_Signer_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainId *big.Int
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if chainId == nil || tx == nil {
			return
		}

		_x1 := NewLondonSigner(chainId)
		_x1.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_Transactions_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a Transactions
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b Transactions
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		s := TxDifference(a, b)
		s.EncodeIndex(i, w)
	})
}

func Fuzz_Nosy_Transactions_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a Transactions
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b Transactions
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		s := TxDifference(a, b)
		s.Len()
	})
}

// skipping Fuzz_Nosy_TrieHasher_Hash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TrieHasher

// skipping Fuzz_Nosy_TrieHasher_Reset__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TrieHasher

// skipping Fuzz_Nosy_TrieHasher_Update__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TrieHasher

func Fuzz_Nosy_TxByNonce_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s TxByNonce
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_TxByNonce_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s TxByNonce
		fill_err = tp.Fill(&s)
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

		s.Less(i, j)
	})
}

func Fuzz_Nosy_TxByNonce_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s TxByNonce
		fill_err = tp.Fill(&s)
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

		s.Swap(i, j)
	})
}

// skipping Fuzz_Nosy_TxData_accessList__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_chainID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_copy__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_data__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_decode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_effectiveGasPrice__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_gas__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_gasFeeCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_gasPrice__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_gasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_nonce__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_rawSignatureValues__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_setSignatureValues__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_to__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_txType__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

// skipping Fuzz_Nosy_TxData_value__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.TxData

func Fuzz_Nosy_Withdrawal_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		w.MarshalJSON()
	})
}

func Fuzz_Nosy_Withdrawals_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Withdrawals
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var w *bytes.Buffer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		s.EncodeIndex(i, w)
	})
}

func Fuzz_Nosy_Withdrawals_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s Withdrawals
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

// skipping Fuzz_Nosy_bytesBacked_Bytes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.bytesBacked

// skipping Fuzz_Nosy_cancunSigner_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_cancunSigner_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s cancunSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Hash(tx)
	})
}

func Fuzz_Nosy_cancunSigner_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s cancunSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Sender(tx)
	})
}

func Fuzz_Nosy_cancunSigner_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s cancunSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_eip2930Signer_ChainID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s eip2930Signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.ChainID()
	})
}

// skipping Fuzz_Nosy_eip2930Signer_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_eip2930Signer_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s eip2930Signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Hash(tx)
	})
}

func Fuzz_Nosy_eip2930Signer_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s eip2930Signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Sender(tx)
	})
}

func Fuzz_Nosy_eip2930Signer_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s eip2930Signer
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.SignatureValues(tx, sig)
	})
}

// skipping Fuzz_Nosy_londonSigner_Equal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.Signer

func Fuzz_Nosy_londonSigner_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s londonSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Hash(tx)
	})
}

func Fuzz_Nosy_londonSigner_Sender__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s londonSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.Sender(tx)
	})
}

func Fuzz_Nosy_londonSigner_SignatureValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s londonSigner
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		s.SignatureValues(tx, sig)
	})
}

func Fuzz_Nosy_Bloom9__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		Bloom9(d1)
	})
}

// skipping Fuzz_Nosy_BloomLookup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.bytesBacked

func Fuzz_Nosy_FullAccountRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		FullAccountRLP(d1)
	})
}

func Fuzz_Nosy_HashDifference__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a []common.Hash
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b []common.Hash
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		HashDifference(a, b)
	})
}

func Fuzz_Nosy_LogsBloom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var logs []*Log
		fill_err = tp.Fill(&logs)
		if fill_err != nil {
			return
		}

		LogsBloom(logs)
	})
}

func Fuzz_Nosy_SlimAccountRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var account StateAccount
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		SlimAccountRLP(account)
	})
}

func Fuzz_Nosy_bloomValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, hashbuf []byte) {
		bloomValues(d1, hashbuf)
	})
}

func Fuzz_Nosy_decodeSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sig []byte) {
		decodeSignature(sig)
	})
}

// skipping Fuzz_Nosy_encodeForDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/core/types.DerivableList

func Fuzz_Nosy_getPooledBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64) {
		getPooledBuffer(size)
	})
}

func Fuzz_Nosy_isProtectedV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var V *big.Int
		fill_err = tp.Fill(&V)
		if fill_err != nil {
			return
		}
		if V == nil {
			return
		}

		isProtectedV(V)
	})
}
