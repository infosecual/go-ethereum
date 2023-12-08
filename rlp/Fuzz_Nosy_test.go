package rlp

import (
	"io"
	"math/big"
	"reflect"
	"testing"

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

func Fuzz_Nosy_EncoderBuffer_AppendToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 io.Writer
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(d1)
		w.AppendToBytes(d2)
	})
}

func Fuzz_Nosy_EncoderBuffer_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.Flush()
	})
}

func Fuzz_Nosy_EncoderBuffer_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 io.Writer
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var d2 io.Writer
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(d1)
		w.Reset(d2)
	})
}

func Fuzz_Nosy_EncoderBuffer_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.ToBytes()
	})
}

func Fuzz_Nosy_Stream_BigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.BigInt()
	})
}

func Fuzz_Nosy_Stream_Bool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Bool()
	})
}

func Fuzz_Nosy_Stream_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Bytes()
	})
}

// skipping Fuzz_Nosy_Stream_Decode__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Stream_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Kind()
	})
}

func Fuzz_Nosy_Stream_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.List()
	})
}

func Fuzz_Nosy_Stream_ListEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.ListEnd()
	})
}

func Fuzz_Nosy_Stream_MoreDataInList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.MoreDataInList()
	})
}

func Fuzz_Nosy_Stream_Raw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Raw()
	})
}

func Fuzz_Nosy_Stream_ReadBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.ReadBytes(b)
	})
}

func Fuzz_Nosy_Stream_ReadUint256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var dst *uint256.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if dst == nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.ReadUint256(dst)
	})
}

func Fuzz_Nosy_Stream_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 io.Reader
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		var i2 uint64
		fill_err = tp.Fill(&i2)
		if fill_err != nil {
			return
		}
		var r3 io.Reader
		fill_err = tp.Fill(&r3)
		if fill_err != nil {
			return
		}
		var i4 uint64
		fill_err = tp.Fill(&i4)
		if fill_err != nil {
			return
		}

		s := NewStream(r1, i2)
		s.Reset(r3, i4)
	})
}

func Fuzz_Nosy_Stream_Uint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Uint()
	})
}

func Fuzz_Nosy_Stream_Uint16__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Uint16()
	})
}

func Fuzz_Nosy_Stream_Uint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Uint32()
	})
}

func Fuzz_Nosy_Stream_Uint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Uint64()
	})
}

func Fuzz_Nosy_Stream_Uint8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.Uint8()
	})
}

func Fuzz_Nosy_Stream_decodeBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var dst *big.Int
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if dst == nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.decodeBigInt(dst)
	})
}

func Fuzz_Nosy_Stream_listLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.listLimit()
	})
}

func Fuzz_Nosy_Stream_readByte__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.readByte()
	})
}

func Fuzz_Nosy_Stream_readFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.readFull(buf)
	})
}

func Fuzz_Nosy_Stream_readKind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.readKind()
	})
}

func Fuzz_Nosy_Stream_readUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var size byte
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.readUint(size)
	})
}

func Fuzz_Nosy_Stream_uint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var maxbits int
		fill_err = tp.Fill(&maxbits)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.uint(maxbits)
	})
}

func Fuzz_Nosy_Stream_willRead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var inputLimit uint64
		fill_err = tp.Fill(&inputLimit)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		s := NewStream(r, inputLimit)
		s.willRead(n)
	})
}

func Fuzz_Nosy_decodeError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *decodeError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_encBuffer_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		buf := getEncBuffer()
		buf.Write(b)
	})
}

func Fuzz_Nosy_encBuffer_copyTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte) {
		buf := getEncBuffer()
		buf.copyTo(dst)
	})
}

// skipping Fuzz_Nosy_encBuffer_encode__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encBuffer_encodeStringHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
		buf := getEncBuffer()
		buf.encodeStringHeader(size)
	})
}

func Fuzz_Nosy_encBuffer_listEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, index int) {
		buf := getEncBuffer()
		buf.listEnd(index)
	})
}

func Fuzz_Nosy_encBuffer_writeBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *big.Int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		buf := getEncBuffer()
		buf.writeBigInt(i)
	})
}

func Fuzz_Nosy_encBuffer_writeBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b bool) {
		buf := getEncBuffer()
		buf.writeBool(b)
	})
}

func Fuzz_Nosy_encBuffer_writeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		buf := getEncBuffer()
		buf.writeBytes(b)
	})
}

func Fuzz_Nosy_encBuffer_writeString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		buf := getEncBuffer()
		buf.writeString(s)
	})
}

func Fuzz_Nosy_encBuffer_writeTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		buf := getEncBuffer()
		buf.writeTo(w)
	})
}

func Fuzz_Nosy_encBuffer_writeUint256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *uint256.Int
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if z == nil {
			return
		}

		buf := getEncBuffer()
		buf.writeUint256(z)
	})
}

func Fuzz_Nosy_encBuffer_writeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		buf := getEncBuffer()
		buf.writeUint64(i)
	})
}

func Fuzz_Nosy_encReader_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *encReader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Read(b)
	})
}

func Fuzz_Nosy_encReader_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *encReader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.next()
	})
}

func Fuzz_Nosy_listIterator_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 RawValue
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		it, err := NewListIterator(d1)
		if err != nil {
			return
		}
		it.Err()
	})
}

func Fuzz_Nosy_listIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 RawValue
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		it, err := NewListIterator(d1)
		if err != nil {
			return
		}
		it.Next()
	})
}

func Fuzz_Nosy_listIterator_Value__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 RawValue
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}

		it, err := NewListIterator(d1)
		if err != nil {
			return
		}
		it.Value()
	})
}

func Fuzz_Nosy_listhead_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var head *listhead
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if head == nil {
			return
		}

		head.encode(buf)
	})
}

func Fuzz_Nosy_sliceReader_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sr *sliceReader
		fill_err = tp.Fill(&sr)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if sr == nil {
			return
		}

		sr.Read(b)
	})
}

func Fuzz_Nosy_sliceReader_ReadByte__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sr *sliceReader
		fill_err = tp.Fill(&sr)
		if fill_err != nil {
			return
		}
		if sr == nil {
			return
		}

		sr.ReadByte()
	})
}

// skipping Fuzz_Nosy_typeCache_generate__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_typeCache_info__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_typeCache_infoWhileGenerating__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_typeinfo_generate__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_Decoder_DecodeRLP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp.Decoder

// skipping Fuzz_Nosy_Encoder_EncodeRLP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp.Encoder

func Fuzz_Nosy_EncoderBuffer_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.List()
	})
}

func Fuzz_Nosy_EncoderBuffer_ListEnd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.ListEnd(index)
	})
}

func Fuzz_Nosy_EncoderBuffer_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.Write(b)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var i *big.Int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteBigInt(i)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var b bool
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteBool(b)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteBytes(b)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteString(s)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteUint256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var i *uint256.Int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteUint256(i)
	})
}

func Fuzz_Nosy_EncoderBuffer_WriteUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dst io.Writer
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		var i uint64
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		w := NewEncoderBuffer(dst)
		w.WriteUint64(i)
	})
}

func Fuzz_Nosy_Kind_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Kind
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.String()
	})
}

func Fuzz_Nosy_structFieldError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e structFieldError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_AppendUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, i uint64) {
		AppendUint64(b, i)
	})
}

func Fuzz_Nosy_BytesSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		BytesSize(b)
	})
}

func Fuzz_Nosy_CountValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		CountValues(b)
	})
}

// skipping Fuzz_Nosy_EncodeToBytes__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_EncodeToReader__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_IntSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		IntSize(x)
	})
}

func Fuzz_Nosy_ListSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, contentSize uint64) {
		ListSize(contentSize)
	})
}

func Fuzz_Nosy_Split__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		Split(b)
	})
}

func Fuzz_Nosy_SplitList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		SplitList(b)
	})
}

func Fuzz_Nosy_SplitString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		SplitString(b)
	})
}

func Fuzz_Nosy_SplitUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		SplitUint64(b)
	})
}

func Fuzz_Nosy_StringSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		StringSize(s)
	})
}

func Fuzz_Nosy_byteArrayBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v reflect.Value
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var length int
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}

		byteArrayBytes(v, length)
	})
}

func Fuzz_Nosy_firstOptionalField__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fields []field
		fill_err = tp.Fill(&fields)
		if fill_err != nil {
			return
		}

		firstOptionalField(fields)
	})
}

func Fuzz_Nosy_headsize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint64) {
		headsize(size)
	})
}

func Fuzz_Nosy_intsize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i uint64) {
		intsize(i)
	})
}

// skipping Fuzz_Nosy_isByte__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_isUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k reflect.Kind
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		isUint(k)
	})
}

func Fuzz_Nosy_puthead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte, smalltag byte, largetag byte, size uint64) {
		puthead(buf, smalltag, largetag, size)
	})
}

func Fuzz_Nosy_putint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, i uint64) {
		putint(b, i)
	})
}

func Fuzz_Nosy_readKind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, buf []byte) {
		readKind(buf)
	})
}

func Fuzz_Nosy_readSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, slen byte) {
		readSize(b, slen)
	})
}

// skipping Fuzz_Nosy_structFields__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_zeroFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var structval reflect.Value
		fill_err = tp.Fill(&structval)
		if fill_err != nil {
			return
		}
		var fields []field
		fill_err = tp.Fill(&fields)
		if fill_err != nil {
			return
		}

		zeroFields(structval, fields)
	})
}
