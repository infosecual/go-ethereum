package enr

import (
	"io"
	"testing"

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

func Fuzz_Nosy_IP_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *IP
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if v == nil || s == nil {
			return
		}

		v.DecodeRLP(s)
	})
}

func Fuzz_Nosy_IPv4_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *IPv4
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if v == nil || s == nil {
			return
		}

		v.DecodeRLP(s)
	})
}

func Fuzz_Nosy_IPv6_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *IPv6
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if v == nil || s == nil {
			return
		}

		v.DecodeRLP(s)
	})
}

func Fuzz_Nosy_KeyError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *KeyError
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

func Fuzz_Nosy_KeyError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *KeyError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Unwrap()
	})
}

// skipping Fuzz_Nosy_Record_AppendElements__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Record_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
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

func Fuzz_Nosy_Record_IdentityScheme__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.IdentityScheme()
	})
}

// skipping Fuzz_Nosy_Record_Load__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_Record_Seq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Seq()
	})
}

// skipping Fuzz_Nosy_Record_Set__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_Record_SetSeq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var s uint64
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.SetSeq(s)
	})
}

// skipping Fuzz_Nosy_Record_SetSig__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.IdentityScheme

func Fuzz_Nosy_Record_Signature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Signature()
	})
}

func Fuzz_Nosy_Record_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Size()
	})
}

// skipping Fuzz_Nosy_Record_VerifySignature__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.IdentityScheme

func Fuzz_Nosy_Record_encode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.encode(sig)
	})
}

func Fuzz_Nosy_Record_invalidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.invalidate()
	})
}

func Fuzz_Nosy_generic_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *generic
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if g == nil || s == nil {
			return
		}

		g.DecodeRLP(s)
	})
}

// skipping Fuzz_Nosy_Entry_ENRKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_ID_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v ID
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_IP_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IP
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_IP_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IP
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		v.EncodeRLP(w)
	})
}

func Fuzz_Nosy_IPv4_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IPv4
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_IPv4_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IPv4
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		v.EncodeRLP(w)
	})
}

func Fuzz_Nosy_IPv6_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IPv6
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_IPv6_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v IPv6
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		v.EncodeRLP(w)
	})
}

// skipping Fuzz_Nosy_IdentityScheme_NodeAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.IdentityScheme

// skipping Fuzz_Nosy_IdentityScheme_Verify__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.IdentityScheme

func Fuzz_Nosy_Record_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		r.EncodeRLP(w)
	})
}

// skipping Fuzz_Nosy_SchemeMap_NodeAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.SchemeMap

// skipping Fuzz_Nosy_SchemeMap_Verify__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.SchemeMap

func Fuzz_Nosy_TCP_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v TCP
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_TCP6_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v TCP6
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_UDP_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v UDP
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_UDP6_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v UDP6
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_generic_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g generic
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}

		g.ENRKey()
	})
}

func Fuzz_Nosy_generic_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g generic
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		g.EncodeRLP(w)
	})
}

// skipping Fuzz_Nosy_IsNotFound__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_computeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		computeSize(r)
	})
}

func Fuzz_Nosy_decodeRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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

		decodeRecord(s)
	})
}
