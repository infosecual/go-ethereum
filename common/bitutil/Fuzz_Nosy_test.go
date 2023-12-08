package bitutil

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

func Fuzz_Nosy_ANDBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		ANDBytes(dst, a, b)
	})
}

func Fuzz_Nosy_CompressBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		CompressBytes(d1)
	})
}

func Fuzz_Nosy_DecompressBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, t2 int) {
		DecompressBytes(d1, t2)
	})
}

func Fuzz_Nosy_ORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		ORBytes(dst, a, b)
	})
}

func Fuzz_Nosy_TestBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		TestBytes(p)
	})
}

func Fuzz_Nosy_XORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		XORBytes(dst, a, b)
	})
}

func Fuzz_Nosy_bitsetDecodeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, t2 int) {
		bitsetDecodeBytes(d1, t2)
	})
}

func Fuzz_Nosy_bitsetDecodePartialBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, t2 int) {
		bitsetDecodePartialBytes(d1, t2)
	})
}

func Fuzz_Nosy_bitsetEncodeBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		bitsetEncodeBytes(d1)
	})
}

func Fuzz_Nosy_fastANDBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		fastANDBytes(dst, a, b)
	})
}

func Fuzz_Nosy_fastORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		fastORBytes(dst, a, b)
	})
}

func Fuzz_Nosy_fastTestBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		fastTestBytes(p)
	})
}

func Fuzz_Nosy_fastXORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		fastXORBytes(dst, a, b)
	})
}

func Fuzz_Nosy_safeANDBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		safeANDBytes(dst, a, b)
	})
}

func Fuzz_Nosy_safeORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		safeORBytes(dst, a, b)
	})
}

func Fuzz_Nosy_safeTestBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p []byte) {
		safeTestBytes(p)
	})
}

func Fuzz_Nosy_safeXORBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte, a []byte, b []byte) {
		safeXORBytes(dst, a, b)
	})
}
