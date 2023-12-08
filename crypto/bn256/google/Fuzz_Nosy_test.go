package bn256

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

func Fuzz_Nosy_G1_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *G1
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *G1
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil || b == nil {
			return
		}

		e.Add(a, b)
	})
}

func Fuzz_Nosy_G1_CurvePoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.CurvePoints()
	})
}

func Fuzz_Nosy_G1_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Marshal()
	})
}

func Fuzz_Nosy_G1_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *G1
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.Neg(a)
	})
}

func Fuzz_Nosy_G1_ScalarBaseMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if e == nil || k == nil {
			return
		}

		e.ScalarBaseMult(k)
	})
}

func Fuzz_Nosy_G1_ScalarMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *G1
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil || k == nil {
			return
		}

		e.ScalarMult(a, k)
	})
}

func Fuzz_Nosy_G1_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_G1_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G1
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var m []byte
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Unmarshal(m)
	})
}

func Fuzz_Nosy_G2_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *G2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *G2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil || b == nil {
			return
		}

		e.Add(a, b)
	})
}

func Fuzz_Nosy_G2_CurvePoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.CurvePoints()
	})
}

func Fuzz_Nosy_G2_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *G2
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Marshal()
	})
}

func Fuzz_Nosy_G2_ScalarBaseMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if e == nil || k == nil {
			return
		}

		e.ScalarBaseMult(k)
	})
}

func Fuzz_Nosy_G2_ScalarMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *G2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil || k == nil {
			return
		}

		e.ScalarMult(a, k)
	})
}

func Fuzz_Nosy_G2_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_G2_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *G2
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var m []byte
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Unmarshal(m)
	})
}

func Fuzz_Nosy_GT_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		var a *GT
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *GT
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil || a == nil || b == nil {
			return
		}

		e := Pair(g1, g2)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_GT_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil {
			return
		}

		n := Pair(g1, g2)
		n.Marshal()
	})
}

func Fuzz_Nosy_GT_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		var a *GT
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil || a == nil {
			return
		}

		e := Pair(g1, g2)
		e.Neg(a)
	})
}

func Fuzz_Nosy_GT_ScalarMult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		var a *GT
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil || a == nil || k == nil {
			return
		}

		e := Pair(g1, g2)
		e.ScalarMult(a, k)
	})
}

func Fuzz_Nosy_GT_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil {
			return
		}

		g := Pair(g1, g2)
		g.String()
	})
}

func Fuzz_Nosy_GT_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g1 *G1
		fill_err = tp.Fill(&g1)
		if fill_err != nil {
			return
		}
		var g2 *G2
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		var m []byte
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if g1 == nil || g2 == nil {
			return
		}

		e := Pair(g1, g2)
		e.Unmarshal(m)
	})
}

func Fuzz_Nosy_bnPool_Count__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Count()
	})
}

func Fuzz_Nosy_bnPool_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		pool.Get()
	})
}

func Fuzz_Nosy_bnPool_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var bn *big.Int
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		if pool == nil || bn == nil {
			return
		}

		pool.Put(bn)
	})
}

func Fuzz_Nosy_curvePoint_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		c := newCurvePoint(p1)
		c.Add(a, b, p4)
	})
}

func Fuzz_Nosy_curvePoint_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		c := newCurvePoint(p1)
		c.Double(a, p3)
	})
}

func Fuzz_Nosy_curvePoint_IsInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newCurvePoint(pool)
		c.IsInfinity()
	})
}

func Fuzz_Nosy_curvePoint_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newCurvePoint(pool)
		c.IsOnCurve()
	})
}

func Fuzz_Nosy_curvePoint_MakeAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		c := newCurvePoint(p1)
		c.MakeAffine(p2)
	})
}

func Fuzz_Nosy_curvePoint_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var scalar *big.Int
		fill_err = tp.Fill(&scalar)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || scalar == nil || p4 == nil {
			return
		}

		c := newCurvePoint(p1)
		c.Mul(a, scalar, p4)
	})
}

func Fuzz_Nosy_curvePoint_Negative__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		c := newCurvePoint(pool)
		c.Negative(a)
	})
}

func Fuzz_Nosy_curvePoint_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		c := newCurvePoint(p1)
		c.Put(p2)
	})
}

func Fuzz_Nosy_curvePoint_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		c := newCurvePoint(pool)
		c.Set(a)
	})
}

func Fuzz_Nosy_curvePoint_SetInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newCurvePoint(pool)
		c.SetInfinity()
	})
}

func Fuzz_Nosy_curvePoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newCurvePoint(pool)
		c.String()
	})
}

func Fuzz_Nosy_gfP12_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp12(pool)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_gfP12_Conjugate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp12(pool)
		e.Conjugate(a)
	})
}

func Fuzz_Nosy_gfP12_Exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power *big.Int
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || power == nil || p4 == nil {
			return
		}

		c := newGFp12(p1)
		c.Exp(a, power, p4)
	})
}

func Fuzz_Nosy_gfP12_Frobenius__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp12(p1)
		e.Frobenius(a, p3)
	})
}

func Fuzz_Nosy_gfP12_FrobeniusP2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp12(p1)
		e.FrobeniusP2(a, p3)
	})
}

func Fuzz_Nosy_gfP12_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp12(p1)
		e.Invert(a, p3)
	})
}

func Fuzz_Nosy_gfP12_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.IsOne()
	})
}

func Fuzz_Nosy_gfP12_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.IsZero()
	})
}

func Fuzz_Nosy_gfP12_Minimal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.Minimal()
	})
}

func Fuzz_Nosy_gfP12_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		e := newGFp12(p1)
		e.Mul(a, b, p4)
	})
}

func Fuzz_Nosy_gfP12_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		e := newGFp12(p1)
		e.MulScalar(a, b, p4)
	})
}

func Fuzz_Nosy_gfP12_Negative__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp12(pool)
		e.Negative(a)
	})
}

func Fuzz_Nosy_gfP12_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		e := newGFp12(p1)
		e.Put(p2)
	})
}

func Fuzz_Nosy_gfP12_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp12(pool)
		e.Set(a)
	})
}

func Fuzz_Nosy_gfP12_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.SetOne()
	})
}

func Fuzz_Nosy_gfP12_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.SetZero()
	})
}

func Fuzz_Nosy_gfP12_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp12(p1)
		e.Square(a, p3)
	})
}

func Fuzz_Nosy_gfP12_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp12(pool)
		e.String()
	})
}

func Fuzz_Nosy_gfP12_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP12
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP12
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp12(pool)
		e.Sub(a, b)
	})
}

func Fuzz_Nosy_gfP2_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp2(pool)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_gfP2_Conjugate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp2(pool)
		e.Conjugate(a)
	})
}

func Fuzz_Nosy_gfP2_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp2(pool)
		e.Double(a)
	})
}

func Fuzz_Nosy_gfP2_Exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var power *big.Int
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || power == nil || p4 == nil {
			return
		}

		c := newGFp2(p1)
		c.Exp(a, power, p4)
	})
}

func Fuzz_Nosy_gfP2_Imag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.Imag()
	})
}

func Fuzz_Nosy_gfP2_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp2(p1)
		e.Invert(a, p3)
	})
}

func Fuzz_Nosy_gfP2_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.IsOne()
	})
}

func Fuzz_Nosy_gfP2_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.IsZero()
	})
}

func Fuzz_Nosy_gfP2_Minimal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.Minimal()
	})
}

func Fuzz_Nosy_gfP2_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		e := newGFp2(p1)
		e.Mul(a, b, p4)
	})
}

func Fuzz_Nosy_gfP2_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp2(pool)
		e.MulScalar(a, b)
	})
}

func Fuzz_Nosy_gfP2_MulXi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp2(p1)
		e.MulXi(a, p3)
	})
}

func Fuzz_Nosy_gfP2_Negative__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp2(pool)
		e.Negative(a)
	})
}

func Fuzz_Nosy_gfP2_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		e := newGFp2(p1)
		e.Put(p2)
	})
}

func Fuzz_Nosy_gfP2_Real__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.Real()
	})
}

func Fuzz_Nosy_gfP2_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp2(pool)
		e.Set(a)
	})
}

func Fuzz_Nosy_gfP2_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.SetOne()
	})
}

func Fuzz_Nosy_gfP2_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.SetZero()
	})
}

func Fuzz_Nosy_gfP2_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp2(p1)
		e.Square(a, p3)
	})
}

func Fuzz_Nosy_gfP2_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp2(pool)
		e.String()
	})
}

func Fuzz_Nosy_gfP2_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp2(pool)
		e.Sub(a, b)
	})
}

func Fuzz_Nosy_gfP6_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp6(pool)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_gfP6_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp6(pool)
		e.Double(a)
	})
}

func Fuzz_Nosy_gfP6_Frobenius__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp6(p1)
		e.Frobenius(a, p3)
	})
}

func Fuzz_Nosy_gfP6_FrobeniusP2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp6(pool)
		e.FrobeniusP2(a)
	})
}

func Fuzz_Nosy_gfP6_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp6(p1)
		e.Invert(a, p3)
	})
}

func Fuzz_Nosy_gfP6_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.IsOne()
	})
}

func Fuzz_Nosy_gfP6_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.IsZero()
	})
}

func Fuzz_Nosy_gfP6_Minimal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.Minimal()
	})
}

func Fuzz_Nosy_gfP6_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		e := newGFp6(p1)
		e.Mul(a, b, p4)
	})
}

func Fuzz_Nosy_gfP6_MulGFP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp6(pool)
		e.MulGFP(a, b)
	})
}

func Fuzz_Nosy_gfP6_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		e := newGFp6(p1)
		e.MulScalar(a, b, p4)
	})
}

func Fuzz_Nosy_gfP6_MulTau__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp6(p1)
		e.MulTau(a, p3)
	})
}

func Fuzz_Nosy_gfP6_Negative__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp6(pool)
		e.Negative(a)
	})
}

func Fuzz_Nosy_gfP6_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		e := newGFp6(p1)
		e.Put(p2)
	})
}

func Fuzz_Nosy_gfP6_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		e := newGFp6(pool)
		e.Set(a)
	})
}

func Fuzz_Nosy_gfP6_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.SetOne()
	})
}

func Fuzz_Nosy_gfP6_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.SetZero()
	})
}

func Fuzz_Nosy_gfP6_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		e := newGFp6(p1)
		e.Square(a, p3)
	})
}

func Fuzz_Nosy_gfP6_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		e := newGFp6(pool)
		e.String()
	})
}

func Fuzz_Nosy_gfP6_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP6
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil || b == nil {
			return
		}

		e := newGFp6(pool)
		e.Sub(a, b)
	})
}

func Fuzz_Nosy_twistPoint_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *twistPoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || b == nil || p4 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.Add(a, b, p4)
	})
}

func Fuzz_Nosy_twistPoint_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.Double(a, p3)
	})
}

func Fuzz_Nosy_twistPoint_IsInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newTwistPoint(pool)
		c.IsInfinity()
	})
}

func Fuzz_Nosy_twistPoint_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newTwistPoint(pool)
		c.IsOnCurve()
	})
}

func Fuzz_Nosy_twistPoint_MakeAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.MakeAffine(p2)
	})
}

func Fuzz_Nosy_twistPoint_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var scalar *big.Int
		fill_err = tp.Fill(&scalar)
		if fill_err != nil {
			return
		}
		var p4 *bnPool
		fill_err = tp.Fill(&p4)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || scalar == nil || p4 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.Mul(a, scalar, p4)
	})
}

func Fuzz_Nosy_twistPoint_Negative__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var p3 *bnPool
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		if p1 == nil || a == nil || p3 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.Negative(a, p3)
	})
}

func Fuzz_Nosy_twistPoint_Put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 *bnPool
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *bnPool
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if p1 == nil || p2 == nil {
			return
		}

		c := newTwistPoint(p1)
		c.Put(p2)
	})
}

func Fuzz_Nosy_twistPoint_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if pool == nil || a == nil {
			return
		}

		c := newTwistPoint(pool)
		c.Set(a)
	})
}

func Fuzz_Nosy_twistPoint_SetInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newTwistPoint(pool)
		c.SetInfinity()
	})
}

func Fuzz_Nosy_twistPoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if pool == nil {
			return
		}

		c := newTwistPoint(pool)
		c.String()
	})
}

func Fuzz_Nosy_PairingCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a []*G1
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b []*G2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		PairingCheck(a, b)
	})
}

func Fuzz_Nosy_RandomG1__(f *testing.F) {
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

		RandomG1(r)
	})
}

func Fuzz_Nosy_RandomG2__(f *testing.F) {
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

		RandomG2(r)
	})
}

func Fuzz_Nosy_lineFunctionAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *twistPoint
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var p *twistPoint
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var q *curvePoint
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var r2 *gfP2
		fill_err = tp.Fill(&r2)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if r == nil || p == nil || q == nil || r2 == nil || pool == nil {
			return
		}

		lineFunctionAdd(r, p, q, r2, pool)
	})
}

func Fuzz_Nosy_lineFunctionDouble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *twistPoint
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var q *curvePoint
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if r == nil || q == nil || pool == nil {
			return
		}

		lineFunctionDouble(r, q, pool)
	})
}

func Fuzz_Nosy_mulLine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ret *gfP12
		fill_err = tp.Fill(&ret)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP2
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var c *gfP2
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var pool *bnPool
		fill_err = tp.Fill(&pool)
		if fill_err != nil {
			return
		}
		if ret == nil || a == nil || b == nil || c == nil || pool == nil {
			return
		}

		mulLine(ret, a, b, c, pool)
	})
}
