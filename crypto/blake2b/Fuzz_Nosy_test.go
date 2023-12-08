package blake2b

import (
	"testing"

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

func Fuzz_Nosy_digest_BlockSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.BlockSize()
	})
}

func Fuzz_Nosy_digest_MarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.MarshalBinary()
	})
}

func Fuzz_Nosy_digest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.Reset()
	})
}

func Fuzz_Nosy_digest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.Size()
	})
}

func Fuzz_Nosy_digest_Sum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte, sum []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.Sum(sum)
	})
}

func Fuzz_Nosy_digest_UnmarshalBinary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte, b []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.UnmarshalBinary(b)
	})
}

func Fuzz_Nosy_digest_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hashSize int, key []byte, p []byte) {
		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.Write(p)
	})
}

func Fuzz_Nosy_digest_finalize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hashSize int
		fill_err = tp.Fill(&hashSize)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var hash *[64]byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hash == nil {
			return
		}

		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.finalize(hash)
	})
}

func Fuzz_Nosy_digest_initConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hashSize int
		fill_err = tp.Fill(&hashSize)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var cfg *[64]byte
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d, err := newDigest(hashSize, key)
		if err != nil {
			return
		}
		d.initConfig(cfg)
	})
}

func Fuzz_Nosy_xof_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *xof
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Clone()
	})
}

func Fuzz_Nosy_xof_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *xof
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Read(p)
	})
}

func Fuzz_Nosy_xof_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *xof
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Reset()
	})
}

func Fuzz_Nosy_xof_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *xof
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var p []byte
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if x == nil {
			return
		}

		x.Write(p)
	})
}

func Fuzz_Nosy_XOF_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint32, key []byte) {
		_x1, err := NewXOF(size, key)
		if err != nil {
			return
		}
		_x1.Clone()
	})
}

func Fuzz_Nosy_XOF_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size uint32, key []byte) {
		_x1, err := NewXOF(size, key)
		if err != nil {
			return
		}
		_x1.Reset()
	})
}

func Fuzz_Nosy_F__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m [16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c [2]uint64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var final bool
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		var rounds uint32
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		F(h, m, c, final, rounds)
	})
}

func Fuzz_Nosy_Sum256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		Sum256(d1)
	})
}

func Fuzz_Nosy_Sum384__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		Sum384(d1)
	})
}

func Fuzz_Nosy_Sum512__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		Sum512(d1)
	})
}

func Fuzz_Nosy_appendUint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, x uint32) {
		appendUint32(b, x)
	})
}

func Fuzz_Nosy_appendUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, x uint64) {
		appendUint64(b, x)
	})
}

func Fuzz_Nosy_checkSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sum *[64]byte
		fill_err = tp.Fill(&sum)
		if fill_err != nil {
			return
		}
		var hashSize int
		fill_err = tp.Fill(&hashSize)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if sum == nil {
			return
		}

		checkSum(sum, hashSize, d3)
	})
}

func Fuzz_Nosy_consumeUint32__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		consumeUint32(b)
	})
}

func Fuzz_Nosy_consumeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		consumeUint64(b)
	})
}

func Fuzz_Nosy_f__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m *[16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c0 uint64
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 uint64
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var rounds uint64
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil || m == nil {
			return
		}

		f(h, m, c0, c1, flag, rounds)
	})
}

func Fuzz_Nosy_fAVX__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m *[16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c0 uint64
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 uint64
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var rounds uint64
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil || m == nil {
			return
		}

		fAVX(h, m, c0, c1, flag, rounds)
	})
}

func Fuzz_Nosy_fAVX2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m *[16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c0 uint64
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 uint64
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var rounds uint64
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil || m == nil {
			return
		}

		fAVX2(h, m, c0, c1, flag, rounds)
	})
}

func Fuzz_Nosy_fGeneric__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m *[16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c0 uint64
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 uint64
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var rounds uint64
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil || m == nil {
			return
		}

		fGeneric(h, m, c0, c1, flag, rounds)
	})
}

func Fuzz_Nosy_fSSE4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var m *[16]uint64
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var c0 uint64
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 uint64
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var rounds uint64
		fill_err = tp.Fill(&rounds)
		if fill_err != nil {
			return
		}
		if h == nil || m == nil {
			return
		}

		fSSE4(h, m, c0, c1, flag, rounds)
	})
}

func Fuzz_Nosy_hashBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var c *[2]uint64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var blocks []byte
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if h == nil || c == nil {
			return
		}

		hashBlocks(h, c, flag, blocks)
	})
}

func Fuzz_Nosy_hashBlocksGeneric__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *[8]uint64
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var c *[2]uint64
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var flag uint64
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var blocks []byte
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if h == nil || c == nil {
			return
		}

		hashBlocksGeneric(h, c, flag, blocks)
	})
}
