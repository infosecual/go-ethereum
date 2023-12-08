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

func Fuzz_Nosy_G1_Set__(f *testing.F) {
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

		e.Set(a)
	})
}

func Fuzz_Nosy_G1_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var g *G1
		fill_err = tp.Fill(&g)
		if fill_err != nil {
			return
		}
		if g == nil {
			return
		}

		g.String()
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

func Fuzz_Nosy_G2_Marshal__(f *testing.F) {
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

		e.Marshal()
	})
}

func Fuzz_Nosy_G2_Neg__(f *testing.F) {
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
		if e == nil || a == nil {
			return
		}

		e.Neg(a)
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

func Fuzz_Nosy_G2_Set__(f *testing.F) {
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
		if e == nil || a == nil {
			return
		}

		e.Set(a)
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

		e := Miller(g1, g2)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_GT_Finalize__(f *testing.F) {
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

		e := Miller(g1, g2)
		e.Finalize()
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

		e := Miller(g1, g2)
		e.Marshal()
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

		e := Miller(g1, g2)
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

		e := Miller(g1, g2)
		e.ScalarMult(a, k)
	})
}

func Fuzz_Nosy_GT_Set__(f *testing.F) {
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

		e := Miller(g1, g2)
		e.Set(a)
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

		g := Miller(g1, g2)
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

		e := Miller(g1, g2)
		e.Unmarshal(m)
	})
}

func Fuzz_Nosy_curvePoint_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
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
		if c == nil || a == nil || b == nil {
			return
		}

		c.Add(a, b)
	})
}

func Fuzz_Nosy_curvePoint_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Double(a)
	})
}

func Fuzz_Nosy_curvePoint_IsInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsInfinity()
	})
}

func Fuzz_Nosy_curvePoint_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsOnCurve()
	})
}

func Fuzz_Nosy_curvePoint_MakeAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.MakeAffine()
	})
}

func Fuzz_Nosy_curvePoint_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
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
		if c == nil || a == nil || scalar == nil {
			return
		}

		c.Mul(a, scalar)
	})
}

func Fuzz_Nosy_curvePoint_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Neg(a)
	})
}

func Fuzz_Nosy_curvePoint_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *curvePoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Set(a)
	})
}

func Fuzz_Nosy_curvePoint_SetInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetInfinity()
	})
}

func Fuzz_Nosy_curvePoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *curvePoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_gfP_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x int64
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var f2 *gfP
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if f2 == nil {
			return
		}

		e := newGFp(x)
		e.Invert(f2)
	})
}

func Fuzz_Nosy_gfP_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64, out []byte) {
		e := newGFp(x)
		e.Marshal(out)
	})
}

func Fuzz_Nosy_gfP_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x int64
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var f2 *gfP
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if f2 == nil {
			return
		}

		e := newGFp(x)
		e.Set(f2)
	})
}

func Fuzz_Nosy_gfP_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64) {
		e := newGFp(x)
		e.String()
	})
}

func Fuzz_Nosy_gfP_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int64, in []byte) {
		e := newGFp(x)
		e.Unmarshal(in)
	})
}

func Fuzz_Nosy_gfP12_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b2 *curvePoint
		fill_err = tp.Fill(&b2)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var b4 *gfP12
		fill_err = tp.Fill(&b4)
		if fill_err != nil {
			return
		}
		if a1 == nil || b2 == nil || a3 == nil || b4 == nil {
			return
		}

		e := optimalAte(a1, b2)
		e.Add(a3, b4)
	})
}

func Fuzz_Nosy_gfP12_Conjugate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Conjugate(a3)
	})
}

func Fuzz_Nosy_gfP12_Exp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var power *big.Int
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil || power == nil {
			return
		}

		c := optimalAte(a1, b)
		c.Exp(a3, power)
	})
}

func Fuzz_Nosy_gfP12_Frobenius__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Frobenius(a3)
	})
}

func Fuzz_Nosy_gfP12_FrobeniusP2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.FrobeniusP2(a3)
	})
}

func Fuzz_Nosy_gfP12_FrobeniusP4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.FrobeniusP4(a3)
	})
}

func Fuzz_Nosy_gfP12_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Invert(a3)
	})
}

func Fuzz_Nosy_gfP12_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := optimalAte(a, b)
		e.IsOne()
	})
}

func Fuzz_Nosy_gfP12_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := optimalAte(a, b)
		e.IsZero()
	})
}

func Fuzz_Nosy_gfP12_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b2 *curvePoint
		fill_err = tp.Fill(&b2)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var b4 *gfP12
		fill_err = tp.Fill(&b4)
		if fill_err != nil {
			return
		}
		if a1 == nil || b2 == nil || a3 == nil || b4 == nil {
			return
		}

		e := optimalAte(a1, b2)
		e.Mul(a3, b4)
	})
}

func Fuzz_Nosy_gfP12_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b2 *curvePoint
		fill_err = tp.Fill(&b2)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var b4 *gfP6
		fill_err = tp.Fill(&b4)
		if fill_err != nil {
			return
		}
		if a1 == nil || b2 == nil || a3 == nil || b4 == nil {
			return
		}

		e := optimalAte(a1, b2)
		e.MulScalar(a3, b4)
	})
}

func Fuzz_Nosy_gfP12_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Neg(a3)
	})
}

func Fuzz_Nosy_gfP12_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Set(a3)
	})
}

func Fuzz_Nosy_gfP12_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := optimalAte(a, b)
		e.SetOne()
	})
}

func Fuzz_Nosy_gfP12_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := optimalAte(a, b)
		e.SetZero()
	})
}

func Fuzz_Nosy_gfP12_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		if a1 == nil || b == nil || a3 == nil {
			return
		}

		e := optimalAte(a1, b)
		e.Square(a3)
	})
}

func Fuzz_Nosy_gfP12_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *curvePoint
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if a == nil || b == nil {
			return
		}

		e := optimalAte(a, b)
		e.String()
	})
}

func Fuzz_Nosy_gfP12_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *twistPoint
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var b2 *curvePoint
		fill_err = tp.Fill(&b2)
		if fill_err != nil {
			return
		}
		var a3 *gfP12
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var b4 *gfP12
		fill_err = tp.Fill(&b4)
		if fill_err != nil {
			return
		}
		if a1 == nil || b2 == nil || a3 == nil || b4 == nil {
			return
		}

		e := optimalAte(a1, b2)
		e.Sub(a3, b4)
	})
}

func Fuzz_Nosy_gfP2_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
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
		if in == nil || a == nil || b == nil {
			return
		}

		e := gfP2Decode(in)
		e.Add(a, b)
	})
}

func Fuzz_Nosy_gfP2_Conjugate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.Conjugate(a)
	})
}

func Fuzz_Nosy_gfP2_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.Invert(a)
	})
}

func Fuzz_Nosy_gfP2_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		e := gfP2Decode(in)
		e.IsOne()
	})
}

func Fuzz_Nosy_gfP2_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		e := gfP2Decode(in)
		e.IsZero()
	})
}

func Fuzz_Nosy_gfP2_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
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
		if in == nil || a == nil || b == nil {
			return
		}

		e := gfP2Decode(in)
		e.Mul(a, b)
	})
}

func Fuzz_Nosy_gfP2_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil || b == nil {
			return
		}

		e := gfP2Decode(in)
		e.MulScalar(a, b)
	})
}

func Fuzz_Nosy_gfP2_MulXi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.MulXi(a)
	})
}

func Fuzz_Nosy_gfP2_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.Neg(a)
	})
}

func Fuzz_Nosy_gfP2_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.Set(a)
	})
}

func Fuzz_Nosy_gfP2_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		e := gfP2Decode(in)
		e.SetOne()
	})
}

func Fuzz_Nosy_gfP2_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		e := gfP2Decode(in)
		e.SetZero()
	})
}

func Fuzz_Nosy_gfP2_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var a *gfP2
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if in == nil || a == nil {
			return
		}

		e := gfP2Decode(in)
		e.Square(a)
	})
}

func Fuzz_Nosy_gfP2_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		if in == nil {
			return
		}

		e := gfP2Decode(in)
		e.String()
	})
}

func Fuzz_Nosy_gfP2_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in *gfP2
		fill_err = tp.Fill(&in)
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
		if in == nil || a == nil || b == nil {
			return
		}

		e := gfP2Decode(in)
		e.Sub(a, b)
	})
}

func Fuzz_Nosy_gfP6_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
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
		if e == nil || a == nil || b == nil {
			return
		}

		e.Add(a, b)
	})
}

func Fuzz_Nosy_gfP6_Frobenius__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.Frobenius(a)
	})
}

func Fuzz_Nosy_gfP6_FrobeniusP2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.FrobeniusP2(a)
	})
}

func Fuzz_Nosy_gfP6_FrobeniusP4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.FrobeniusP4(a)
	})
}

func Fuzz_Nosy_gfP6_Invert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.Invert(a)
	})
}

func Fuzz_Nosy_gfP6_IsOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
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

func Fuzz_Nosy_gfP6_IsZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.IsZero()
	})
}

func Fuzz_Nosy_gfP6_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
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
		if e == nil || a == nil || b == nil {
			return
		}

		e.Mul(a, b)
	})
}

func Fuzz_Nosy_gfP6_MulGFP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil || b == nil {
			return
		}

		e.MulGFP(a, b)
	})
}

func Fuzz_Nosy_gfP6_MulScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
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
		if e == nil || a == nil || b == nil {
			return
		}

		e.MulScalar(a, b)
	})
}

func Fuzz_Nosy_gfP6_MulTau__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.MulTau(a)
	})
}

func Fuzz_Nosy_gfP6_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
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

func Fuzz_Nosy_gfP6_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.Set(a)
	})
}

func Fuzz_Nosy_gfP6_SetOne__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetOne()
	})
}

func Fuzz_Nosy_gfP6_SetZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.SetZero()
	})
}

func Fuzz_Nosy_gfP6_Square__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var a *gfP6
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if e == nil || a == nil {
			return
		}

		e.Square(a)
	})
}

func Fuzz_Nosy_gfP6_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
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

func Fuzz_Nosy_gfP6_Sub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *gfP6
		fill_err = tp.Fill(&e)
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
		if e == nil || a == nil || b == nil {
			return
		}

		e.Sub(a, b)
	})
}

func Fuzz_Nosy_lattice_Multi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lattice
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var scalar *big.Int
		fill_err = tp.Fill(&scalar)
		if fill_err != nil {
			return
		}
		if l == nil || scalar == nil {
			return
		}

		l.Multi(scalar)
	})
}

// skipping Fuzz_Nosy_lattice_Precompute__ because parameters include func, chan, or unsupported interface: func(i uint, j uint)

func Fuzz_Nosy_lattice_decompose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lattice
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var k *big.Int
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if l == nil || k == nil {
			return
		}

		l.decompose(k)
	})
}

func Fuzz_Nosy_twistPoint_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
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
		if c == nil || a == nil || b == nil {
			return
		}

		c.Add(a, b)
	})
}

func Fuzz_Nosy_twistPoint_Double__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Double(a)
	})
}

func Fuzz_Nosy_twistPoint_IsInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsInfinity()
	})
}

func Fuzz_Nosy_twistPoint_IsOnCurve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IsOnCurve()
	})
}

func Fuzz_Nosy_twistPoint_MakeAffine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.MakeAffine()
	})
}

func Fuzz_Nosy_twistPoint_Mul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
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
		if c == nil || a == nil || scalar == nil {
			return
		}

		c.Mul(a, scalar)
	})
}

func Fuzz_Nosy_twistPoint_Neg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Neg(a)
	})
}

func Fuzz_Nosy_twistPoint_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *twistPoint
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		c.Set(a)
	})
}

func Fuzz_Nosy_twistPoint_SetInfinity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetInfinity()
	})
}

func Fuzz_Nosy_twistPoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *twistPoint
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

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

func Fuzz_Nosy_gfpAdd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		gfpAdd(c, a, b)
	})
}

func Fuzz_Nosy_gfpMul__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		gfpMul(c, a, b)
	})
}

func Fuzz_Nosy_gfpNeg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		gfpNeg(c, a)
	})
}

func Fuzz_Nosy_gfpSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b *gfP
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil || b == nil {
			return
		}

		gfpSub(c, a, b)
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
		if r == nil || p == nil || q == nil || r2 == nil {
			return
		}

		lineFunctionAdd(r, p, q, r2)
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
		if r == nil || q == nil {
			return
		}

		lineFunctionDouble(r, q)
	})
}

func Fuzz_Nosy_montDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		montDecode(c, a)
	})
}

func Fuzz_Nosy_montEncode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *gfP
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var a *gfP
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if c == nil || a == nil {
			return
		}

		montEncode(c, a)
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
		if ret == nil || a == nil || b == nil || c == nil {
			return
		}

		mulLine(ret, a, b, c)
	})
}

func Fuzz_Nosy_round__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var denom *big.Int
		fill_err = tp.Fill(&denom)
		if fill_err != nil {
			return
		}
		if num == nil || denom == nil {
			return
		}

		round(num, denom)
	})
}
