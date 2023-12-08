package bls12381

import (
	"io"
	"math/big"
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

func Fuzz_Nosy_Engine_AddPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *PointG1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *PointG2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil {
			return
		}

		e := NewPairingEngine()
		e.AddPair(g1, g2)
	})
}

func Fuzz_Nosy_Engine_AddPairInv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *PointG1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *PointG2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil {
			return
		}

		e := NewPairingEngine()
		e.AddPairInv(g1, g2)
	})
}

func Fuzz_Nosy_Engine_additionStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coeff *[3]fe2
		fill_err = tp.Fill(&coeff)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var q *PointG2
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if coeff == nil || r == nil || q == nil {
			return
		}

		e := NewPairingEngine()
		e.additionStep(coeff, r, q)
	})
}

func Fuzz_Nosy_Engine_affine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p pair
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		e := NewPairingEngine()
		e.affine(p)
	})
}

func Fuzz_Nosy_Engine_doublingStep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coeff *[3]fe2
		fill_err = tp.Fill(&coeff)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if coeff == nil || r == nil {
			return
		}

		e := NewPairingEngine()
		e.doublingStep(coeff, r)
	})
}

func Fuzz_Nosy_Engine_exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := NewPairingEngine()
		e.exp(c, a)
	})
}

func Fuzz_Nosy_Engine_finalExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fe12
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := NewPairingEngine()
		e.finalExp(f1)
	})
}

func Fuzz_Nosy_Engine_isZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p pair
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		e := NewPairingEngine()
		e.isZero(p)
	})
}

func Fuzz_Nosy_Engine_millerLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fe12
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := NewPairingEngine()
		e.millerLoop(f1)
	})
}

func Fuzz_Nosy_Engine_preCompute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ellCoeffs *[68][3]fe2
		fill_err = tp.Fill(&ellCoeffs)
		if fill_err != nil {
			return
		}
		var twistPoint *PointG2
		fill_err = tp.Fill(&twistPoint)
		if fill_err != nil {
			return
		}
		if ellCoeffs == nil || twistPoint == nil {
			return
		}

		e := NewPairingEngine()
		e.preCompute(ellCoeffs, twistPoint)
	})
}

func Fuzz_Nosy_G1_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG1
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p1 *PointG1
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *PointG1
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if r == nil || p1 == nil || p2 == nil {
			return
		}

		g := NewG1()
		g.Add(r, p1, p2)
	})
}

func Fuzz_Nosy_G1_Affine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.Affine(p)
	})
}

func Fuzz_Nosy_G1_ClearCofactor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.ClearCofactor(p)
	})
}

func Fuzz_Nosy_G1_DecodePoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG1()
		g.DecodePoint(in)
	})
}

func Fuzz_Nosy_G1_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG1
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if r == nil || p == nil {
			return
		}

		g := NewG1()
		g.Double(r, p)
	})
}

func Fuzz_Nosy_G1_EncodePoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.EncodePoint(p)
	})
}

func Fuzz_Nosy_G1_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *PointG1
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *PointG1
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		g := NewG1()
		g.Equal(p1, p2)
	})
}

func Fuzz_Nosy_G1_FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG1()
		g.FromBytes(in)
	})
}

func Fuzz_Nosy_G1_InCorrectSubgroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.InCorrectSubgroup(p)
	})
}

func Fuzz_Nosy_G1_IsAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.IsAffine(p)
	})
}

func Fuzz_Nosy_G1_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.IsOnCurve(p)
	})
}

func Fuzz_Nosy_G1_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.IsZero(p)
	})
}

func Fuzz_Nosy_G1_MapToCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG1()
		g.MapToCurve(in)
	})
}

func Fuzz_Nosy_G1_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PointG1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var e *big.Int
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if c == nil || p == nil || e == nil {
			return
		}

		g := NewG1()
		g.MulScalar(c, p, e)
	})
}

func Fuzz_Nosy_G1_MultiExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG1
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var points []*PointG1
		fill_err = tp.Fill(&points)
		if fill_err != nil {
			return
		}
		var powers []*big.Int
		fill_err = tp.Fill(&powers)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		g := NewG1()
		g.MultiExp(r, points, powers)
	})
}

func Fuzz_Nosy_G1_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG1
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if r == nil || p == nil {
			return
		}

		g := NewG1()
		g.Neg(r, p)
	})
}

func Fuzz_Nosy_G1_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PointG1
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *PointG1
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *PointG1
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		g := NewG1()
		g.Sub(c, a, b)
	})
}

func Fuzz_Nosy_G1_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG1()
		g.ToBytes(p)
	})
}

func Fuzz_Nosy_G1_fromBytesUnchecked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG1()
		g.fromBytesUnchecked(in)
	})
}

func Fuzz_Nosy_G2_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p1 *PointG2
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *PointG2
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if r == nil || p1 == nil || p2 == nil {
			return
		}

		g := NewG2()
		g.Add(r, p1, p2)
	})
}

func Fuzz_Nosy_G2_Affine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.Affine(p)
	})
}

func Fuzz_Nosy_G2_ClearCofactor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.ClearCofactor(p)
	})
}

func Fuzz_Nosy_G2_DecodePoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG2()
		g.DecodePoint(in)
	})
}

func Fuzz_Nosy_G2_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if r == nil || p == nil {
			return
		}

		g := NewG2()
		g.Double(r, p)
	})
}

func Fuzz_Nosy_G2_EncodePoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.EncodePoint(p)
	})
}

func Fuzz_Nosy_G2_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *PointG2
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *PointG2
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		g := NewG2()
		g.Equal(p1, p2)
	})
}

func Fuzz_Nosy_G2_FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG2()
		g.FromBytes(in)
	})
}

func Fuzz_Nosy_G2_InCorrectSubgroup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.InCorrectSubgroup(p)
	})
}

func Fuzz_Nosy_G2_IsAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.IsAffine(p)
	})
}

func Fuzz_Nosy_G2_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.IsOnCurve(p)
	})
}

func Fuzz_Nosy_G2_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.IsZero(p)
	})
}

func Fuzz_Nosy_G2_MapToCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG2()
		g.MapToCurve(in)
	})
}

func Fuzz_Nosy_G2_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PointG2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var e *big.Int
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if c == nil || p == nil || e == nil {
			return
		}

		g := NewG2()
		g.MulScalar(c, p, e)
	})
}

func Fuzz_Nosy_G2_MultiExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var points []*PointG2
		fill_err = tp.Fill(&points)
		if fill_err != nil {
			return
		}
		var powers []*big.Int
		fill_err = tp.Fill(&powers)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		g := NewG2()
		g.MultiExp(r, points, powers)
	})
}

func Fuzz_Nosy_G2_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *PointG2
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if r == nil || p == nil {
			return
		}

		g := NewG2()
		g.Neg(r, p)
	})
}

func Fuzz_Nosy_G2_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *PointG2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *PointG2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *PointG2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		g := NewG2()
		g.Sub(c, a, b)
	})
}

func Fuzz_Nosy_G2_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		g := NewG2()
		g.ToBytes(p)
	})
}

func Fuzz_Nosy_G2_fromBytesUnchecked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewG2()
		g.fromBytesUnchecked(in)
	})
}

func Fuzz_Nosy_GT_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		g := NewGT()
		g.Add(c, a, b)
	})
}

func Fuzz_Nosy_GT_Exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || s == nil {
			return
		}

		g := NewGT()
		g.Exp(c, a, s)
	})
}

func Fuzz_Nosy_GT_FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		g := NewGT()
		g.FromBytes(in)
	})
}

func Fuzz_Nosy_GT_Inverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		g := NewGT()
		g.Inverse(c, a)
	})
}

func Fuzz_Nosy_GT_IsValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		g := NewGT()
		g.IsValid(e)
	})
}

func Fuzz_Nosy_GT_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		g := NewGT()
		g.Mul(c, a, b)
	})
}

func Fuzz_Nosy_GT_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		g := NewGT()
		g.Square(c, a)
	})
}

func Fuzz_Nosy_GT_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		g := NewGT()
		g.Sub(c, a, b)
	})
}

func Fuzz_Nosy_GT_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		g := NewGT()
		g.ToBytes(e)
	})
}

func Fuzz_Nosy_PointG1_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var p2 *PointG1
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p == nil || p2 == nil {
			return
		}

		p.Set(p2)
	})
}

func Fuzz_Nosy_PointG1_Zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG1
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Zero()
	})
}

func Fuzz_Nosy_PointG2_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var p2 *PointG2
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p == nil || p2 == nil {
			return
		}

		p.Set(p2)
	})
}

func Fuzz_Nosy_PointG2_Zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *PointG2
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Zero()
	})
}

func Fuzz_Nosy_fe_big__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.big()
	})
}

func Fuzz_Nosy_fe_bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.bytes()
	})
}

func Fuzz_Nosy_fe_cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in string
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var fe2 *fe
		fill_err = tp.Fill(&fe2)
		if fill_err != nil {
			return
		}
		if fe2 == nil {
			return
		}

		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.cmp(fe2)
	})
}

func Fuzz_Nosy_fe_div2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string, e uint64) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.div2(e)
	})
}

func Fuzz_Nosy_fe_equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in string
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var fe2 *fe
		fill_err = tp.Fill(&fe2)
		if fill_err != nil {
			return
		}
		if fe2 == nil {
			return
		}

		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.equal(fe2)
	})
}

func Fuzz_Nosy_fe_isEven__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.isEven()
	})
}

func Fuzz_Nosy_fe_isOdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.isOdd()
	})
}

func Fuzz_Nosy_fe_isOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.isOne()
	})
}

func Fuzz_Nosy_fe_isValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.isValid()
	})
}

func Fuzz_Nosy_fe_isZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.isZero()
	})
}

func Fuzz_Nosy_fe_mul2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.mul2()
	})
}

func Fuzz_Nosy_fe_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.one()
	})
}

func Fuzz_Nosy_fe_rand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in string
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.rand(r)
	})
}

func Fuzz_Nosy_fe_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in string
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var fe2 *fe
		fill_err = tp.Fill(&fe2)
		if fill_err != nil {
			return
		}
		if fe2 == nil {
			return
		}

		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.set(fe2)
	})
}

func Fuzz_Nosy_fe_setBig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in string
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *big.Int
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.setBig(a)
	})
}

func Fuzz_Nosy_fe_setBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		fe, err := fromString(i1)
		if err != nil {
			return
		}
		fe.setBytes(i2)
	})
}

func Fuzz_Nosy_fe_setString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string, s string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.setString(s)
	})
}

func Fuzz_Nosy_fe_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		e, err := fromString(in)
		if err != nil {
			return
		}
		e.sign()
	})
}

func Fuzz_Nosy_fe_string__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.string()
	})
}

func Fuzz_Nosy_fe_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in string) {
		fe, err := fromString(in)
		if err != nil {
			return
		}
		fe.zero()
	})
}

func Fuzz_Nosy_fe12_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *fe12
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		var g2 *fe12
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g == nil || g2 == nil {
			return
		}

		g.Equal(g2)
	})
}

func Fuzz_Nosy_fe12_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.IsOne()
	})
}

func Fuzz_Nosy_fe12_One__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.One()
	})
}

func Fuzz_Nosy_fe12_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe12
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.Set(e2)
	})
}

func Fuzz_Nosy_fe12_equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe12
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.equal(e2)
	})
}

func Fuzz_Nosy_fe12_isOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isOne()
	})
}

func Fuzz_Nosy_fe12_isZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isZero()
	})
}

func Fuzz_Nosy_fe12_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.one()
	})
}

func Fuzz_Nosy_fe12_rand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.rand(r)
	})
}

func Fuzz_Nosy_fe12_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe12
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.set(e2)
	})
}

func Fuzz_Nosy_fe12_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe12
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.zero()
	})
}

func Fuzz_Nosy_fe2_equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe2
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.equal(e2)
	})
}

func Fuzz_Nosy_fe2_isOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isOne()
	})
}

func Fuzz_Nosy_fe2_isZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isZero()
	})
}

func Fuzz_Nosy_fe2_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.one()
	})
}

func Fuzz_Nosy_fe2_rand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.rand(r)
	})
}

func Fuzz_Nosy_fe2_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe2
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.set(e2)
	})
}

func Fuzz_Nosy_fe2_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.sign()
	})
}

func Fuzz_Nosy_fe2_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.zero()
	})
}

func Fuzz_Nosy_fe6_equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe6
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.equal(e2)
	})
}

func Fuzz_Nosy_fe6_isOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isOne()
	})
}

func Fuzz_Nosy_fe6_isZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.isZero()
	})
}

func Fuzz_Nosy_fe6_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.one()
	})
}

func Fuzz_Nosy_fe6_rand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.rand(r)
	})
}

func Fuzz_Nosy_fe6_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var e2 *fe6
		fill_err = tp.Fill(&e2)
		if fill_err != nil {
			return
		}
		if e == nil || e2 == nil {
			return
		}

		e.set(e2)
	})
}

func Fuzz_Nosy_fe6_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.zero()
	})
}

func Fuzz_Nosy_fp12_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp12(fp6)
		e.add(c, a, b)
	})
}

func Fuzz_Nosy_fp12_conjugate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.conjugate(c, a)
	})
}

func Fuzz_Nosy_fp12_cyclotomicExp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil || s == nil {
			return
		}

		e := newFp12(fp6)
		e.cyclotomicExp(c, a, s)
	})
}

func Fuzz_Nosy_fp12_cyclotomicSquare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.cyclotomicSquare(c, a)
	})
}

func Fuzz_Nosy_fp12_double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.double(c, a)
	})
}

func Fuzz_Nosy_fp12_exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil || s == nil {
			return
		}

		e := newFp12(fp6)
		e.exp(c, a, s)
	})
}

func Fuzz_Nosy_fp12_fp2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		if fp6 == nil {
			return
		}

		e := newFp12(fp6)
		e.fp2()
	})
}

func Fuzz_Nosy_fp12_fp4Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c0 *fe2
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 *fe2
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var a0 *fe2
		fill_err = tp.Fill(&a0)
		if fill_err != nil {
			return
		}
		var a1 *fe2
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c0 == nil || c1 == nil || a0 == nil || a1 == nil {
			return
		}

		e := newFp12(fp6)
		e.fp4Square(c0, c1, a0, a1)
	})
}

func Fuzz_Nosy_fp12_frobeniusMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.frobeniusMap(c, a, power)
	})
}

func Fuzz_Nosy_fp12_frobeniusMapAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if fp6 == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.frobeniusMapAssign(a, power)
	})
}

func Fuzz_Nosy_fp12_fromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var in []byte
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if fp6 == nil {
			return
		}

		e := newFp12(fp6)
		e.fromBytes(in)
	})
}

func Fuzz_Nosy_fp12_inverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.inverse(c, a)
	})
}

func Fuzz_Nosy_fp12_mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp12(fp6)
		e.mul(c, a, b)
	})
}

func Fuzz_Nosy_fp12_mulAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp6 == nil || a == nil || b == nil {
			return
		}

		e := newFp12(fp6)
		e.mulAssign(a, b)
	})
}

func Fuzz_Nosy_fp12_mulBy014Assign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var c0 *fe2
		fill_err = tp.Fill(&c0)
		if fill_err != nil {
			return
		}
		var c1 *fe2
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var c4 *fe2
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		if fp6 == nil || a == nil || c0 == nil || c1 == nil || c4 == nil {
			return
		}

		e := newFp12(fp6)
		e.mulBy014Assign(a, c0, c1, c4)
	})
}

func Fuzz_Nosy_fp12_neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.neg(c, a)
	})
}

func Fuzz_Nosy_fp12_new__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		if fp6 == nil {
			return
		}

		e := newFp12(fp6)
		e.new()
	})
}

func Fuzz_Nosy_fp12_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		if fp6 == nil {
			return
		}

		e := newFp12(fp6)
		e.one()
	})
}

func Fuzz_Nosy_fp12_square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.square(c, a)
	})
}

func Fuzz_Nosy_fp12_sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var c *fe12
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp6 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp12(fp6)
		e.sub(c, a, b)
	})
}

func Fuzz_Nosy_fp12_toBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		var a *fe12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if fp6 == nil || a == nil {
			return
		}

		e := newFp12(fp6)
		e.toBytes(a)
	})
}

func Fuzz_Nosy_fp12_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp6 *fp6
		fill_err = tp.Fill(&fp6)
		if fill_err != nil {
			return
		}
		if fp6 == nil {
			return
		}

		e := newFp12(fp6)
		e.zero()
	})
}

func Fuzz_Nosy_fp2_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		e := newFp2()
		e.add(c, a, b)
	})
}

func Fuzz_Nosy_fp2_addAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := newFp2()
		e.addAssign(a, b)
	})
}

func Fuzz_Nosy_fp2_double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.double(c, a)
	})
}

func Fuzz_Nosy_fp2_doubleAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		e := newFp2()
		e.doubleAssign(a)
	})
}

func Fuzz_Nosy_fp2_exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || s == nil {
			return
		}

		e := newFp2()
		e.exp(c, a, s)
	})
}

func Fuzz_Nosy_fp2_frobeniusMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.frobeniusMap(c, a, power)
	})
}

func Fuzz_Nosy_fp2_frobeniusMapAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		e := newFp2()
		e.frobeniusMapAssign(a, power)
	})
}

func Fuzz_Nosy_fp2_fromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		e := newFp2()
		e.fromBytes(in)
	})
}

func Fuzz_Nosy_fp2_inverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.inverse(c, a)
	})
}

func Fuzz_Nosy_fp2_isQuadraticNonResidue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		e := newFp2()
		e.isQuadraticNonResidue(a)
	})
}

func Fuzz_Nosy_fp2_ladd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		e := newFp2()
		e.ladd(c, a, b)
	})
}

func Fuzz_Nosy_fp2_ldouble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.ldouble(c, a)
	})
}

func Fuzz_Nosy_fp2_mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		e := newFp2()
		e.mul(c, a, b)
	})
}

func Fuzz_Nosy_fp2_mulAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := newFp2()
		e.mulAssign(a, b)
	})
}

func Fuzz_Nosy_fp2_mulByB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.mulByB(c, a)
	})
}

func Fuzz_Nosy_fp2_mulByFq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		e := newFp2()
		e.mulByFq(c, a, b)
	})
}

func Fuzz_Nosy_fp2_mulByNonResidue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.mulByNonResidue(c, a)
	})
}

func Fuzz_Nosy_fp2_neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.neg(c, a)
	})
}

func Fuzz_Nosy_fp2_sqrt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.sqrt(c, a)
	})
}

func Fuzz_Nosy_fp2_square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.square(c, a)
	})
}

func Fuzz_Nosy_fp2_squareAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		e := newFp2()
		e.squareAssign(a)
	})
}

func Fuzz_Nosy_fp2_sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		e := newFp2()
		e.sub(c, a, b)
	})
}

func Fuzz_Nosy_fp2_subAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		e := newFp2()
		e.subAssign(c, a)
	})
}

func Fuzz_Nosy_fp2_toBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *fe2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		e := newFp2()
		e.toBytes(a)
	})
}

func Fuzz_Nosy_fp6_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.add(c, a, b)
	})
}

func Fuzz_Nosy_fp6_addAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.addAssign(a, b)
	})
}

func Fuzz_Nosy_fp6_double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.double(c, a)
	})
}

func Fuzz_Nosy_fp6_doubleAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.doubleAssign(a)
	})
}

func Fuzz_Nosy_fp6_exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || s == nil {
			return
		}

		e := newFp6(f1)
		e.exp(c, a, s)
	})
}

func Fuzz_Nosy_fp6_frobeniusMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.frobeniusMap(c, a, power)
	})
}

func Fuzz_Nosy_fp6_frobeniusMapAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power uint
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.frobeniusMapAssign(a, power)
	})
}

func Fuzz_Nosy_fp6_fromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := newFp6(f1)
		e.fromBytes(b)
	})
}

func Fuzz_Nosy_fp6_inverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.inverse(c, a)
	})
}

func Fuzz_Nosy_fp6_mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.mul(c, a, b)
	})
}

func Fuzz_Nosy_fp6_mulAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.mulAssign(a, b)
	})
}

func Fuzz_Nosy_fp6_mulBy01__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b0 *fe2
		fill_err = tp.Fill(&b0)
		if fill_err != nil {
			return
		}
		var b1 *fe2
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b0 == nil || b1 == nil {
			return
		}

		e := newFp6(f1)
		e.mulBy01(c, a, b0, b1)
	})
}

func Fuzz_Nosy_fp6_mulBy01Assign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b0 *fe2
		fill_err = tp.Fill(&b0)
		if fill_err != nil {
			return
		}
		var b1 *fe2
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil || b0 == nil || b1 == nil {
			return
		}

		e := newFp6(f1)
		e.mulBy01Assign(a, b0, b1)
	})
}

func Fuzz_Nosy_fp6_mulBy1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b1 *fe2
		fill_err = tp.Fill(&b1)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b1 == nil {
			return
		}

		e := newFp6(f1)
		e.mulBy1(c, a, b1)
	})
}

func Fuzz_Nosy_fp6_mulByBaseField__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.mulByBaseField(c, a, b)
	})
}

func Fuzz_Nosy_fp6_mulByNonResidue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.mulByNonResidue(c, a)
	})
}

func Fuzz_Nosy_fp6_neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.neg(c, a)
	})
}

func Fuzz_Nosy_fp6_new__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := newFp6(f1)
		e.new()
	})
}

func Fuzz_Nosy_fp6_one__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := newFp6(f1)
		e.one()
	})
}

func Fuzz_Nosy_fp6_square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.square(c, a)
	})
}

func Fuzz_Nosy_fp6_sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var c *fe6
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || c == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.sub(c, a, b)
	})
}

func Fuzz_Nosy_fp6_subAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *fe6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil || b == nil {
			return
		}

		e := newFp6(f1)
		e.subAssign(a, b)
	})
}

func Fuzz_Nosy_fp6_toBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var a *fe6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if f1 == nil || a == nil {
			return
		}

		e := newFp6(f1)
		e.toBytes(a)
	})
}

func Fuzz_Nosy_fp6_zero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *fp2
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		e := newFp6(f1)
		e.zero()
	})
}

func Fuzz_Nosy_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil || y == nil {
			return
		}

		add(z, x, y)
	})
}

func Fuzz_Nosy_addAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		addAssign(x, y)
	})
}

func Fuzz_Nosy_decodeFieldElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		decodeFieldElement(in)
	})
}

func Fuzz_Nosy_double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		double(z, x)
	})
}

func Fuzz_Nosy_doubleAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		if z == nil {
			return
		}

		doubleAssign(z)
	})
}

func Fuzz_Nosy_exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var e *big.Int
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || e == nil {
			return
		}

		exp(c, a, e)
	})
}

func Fuzz_Nosy_fromMont__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		fromMont(c, a)
	})
}

func Fuzz_Nosy_inverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inv *fe
		fill_err = tp.Fill(&inv)
		if fill_err != nil {
			return
		}
		var e *fe
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if inv == nil || e == nil {
			return
		}

		inverse(inv, e)
	})
}

func Fuzz_Nosy_isQuadraticNonResidue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var elem *fe
		fill_err = tp.Fill(&elem)
		if fill_err != nil {
			return
		}
		if elem == nil {
			return
		}

		isQuadraticNonResidue(elem)
	})
}

func Fuzz_Nosy_isogenyMapG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		isogenyMapG1(x, y)
	})
}

func Fuzz_Nosy_isogenyMapG2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fp2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var x *fe2
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe2
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if e == nil || x == nil || y == nil {
			return
		}

		isogenyMapG2(e, x, y)
	})
}

func Fuzz_Nosy_ladd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil || y == nil {
			return
		}

		ladd(z, x, y)
	})
}

func Fuzz_Nosy_laddAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		laddAssign(x, y)
	})
}

func Fuzz_Nosy_ldouble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		ldouble(z, x)
	})
}

func Fuzz_Nosy_lsubAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		lsubAssign(z, x)
	})
}

func Fuzz_Nosy_madd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, t3 uint64, u uint64, v uint64) {
		madd(a, b, t3, u, v)
	})
}

func Fuzz_Nosy_madd0__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64) {
		madd0(a, b, c)
	})
}

func Fuzz_Nosy_madd1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64) {
		madd1(a, b, c)
	})
}

func Fuzz_Nosy_madd1s__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, d uint64, e uint64) {
		madd1s(a, b, d, e)
	})
}

func Fuzz_Nosy_madd1sb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, e uint64) {
		madd1sb(a, b, e)
	})
}

func Fuzz_Nosy_madd2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64, d uint64) {
		madd2(a, b, c, d)
	})
}

func Fuzz_Nosy_madd2s__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64, d uint64, e uint64) {
		madd2s(a, b, c, d, e)
	})
}

func Fuzz_Nosy_madd2sb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64, e uint64) {
		madd2sb(a, b, c, e)
	})
}

func Fuzz_Nosy_madd3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, a uint64, b uint64, c uint64, d uint64, e uint64) {
		madd3(a, b, c, d, e)
	})
}

func Fuzz_Nosy_mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil || y == nil {
			return
		}

		mul(z, x, y)
	})
}

func Fuzz_Nosy_neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		neg(z, x)
	})
}

func Fuzz_Nosy_sqrt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		sqrt(c, a)
	})
}

func Fuzz_Nosy_square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		square(z, x)
	})
}

func Fuzz_Nosy_sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *fe
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil || y == nil {
			return
		}

		sub(z, x, y)
	})
}

func Fuzz_Nosy_subAssign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var z *fe
		fill_err = tp.Fill(&z)
		if fill_err != nil {
			return
		}
		var x *fe
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if z == nil || x == nil {
			return
		}

		subAssign(z, x)
	})
}

func Fuzz_Nosy_swuMapG1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *fe
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		swuMapG1(u)
	})
}

func Fuzz_Nosy_swuMapG2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fp2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var u *fe2
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if e == nil || u == nil {
			return
		}

		swuMapG2(e, u)
	})
}

func Fuzz_Nosy_toBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		toBytes(e)
	})
}

func Fuzz_Nosy_toMont__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *fe
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *fe
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		toMont(c, a)
	})
}

func Fuzz_Nosy_toString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *fe
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		toString(e)
	})
}
