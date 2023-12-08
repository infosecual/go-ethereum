package main

import (
	"bytes"
	"go/types"
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

func Fuzz_Nosy_Config_process__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.process()
	})
}

func Fuzz_Nosy_buildContext_generate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var typ *types.Named
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var encoder bool
		fill_err = tp.Fill(&encoder)
		if fill_err != nil {
			return
		}
		var decoder bool
		fill_err = tp.Fill(&decoder)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || typ == nil {
			return
		}

		bctx := newBuildContext(packageRLP)
		bctx.generate(typ, encoder, decoder)
	})
}

// skipping Fuzz_Nosy_buildContext_isDecoder__ because parameters include func, chan, or unsupported interface: go/types.Type

// skipping Fuzz_Nosy_buildContext_isEncoder__ because parameters include func, chan, or unsupported interface: go/types.Type

func Fuzz_Nosy_buildContext_makeBasicOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var typ *types.Basic
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || typ == nil {
			return
		}

		_x1 := newBuildContext(packageRLP)
		_x1.makeBasicOp(typ)
	})
}

func Fuzz_Nosy_buildContext_makeByteArrayOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var name *types.Named
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var typ *types.Array
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || name == nil || typ == nil {
			return
		}

		bctx := newBuildContext(packageRLP)
		bctx.makeByteArrayOp(name, typ)
	})
}

func Fuzz_Nosy_buildContext_makeByteSliceOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var typ *types.Slice
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || typ == nil {
			return
		}

		_x1 := newBuildContext(packageRLP)
		_x1.makeByteSliceOp(typ)
	})
}

// skipping Fuzz_Nosy_buildContext_makeOp__ because parameters include func, chan, or unsupported interface: go/types.Type

// skipping Fuzz_Nosy_buildContext_makePtrOp__ because parameters include func, chan, or unsupported interface: go/types.Type

func Fuzz_Nosy_buildContext_makeRawValueOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		if packageRLP == nil {
			return
		}

		bctx := newBuildContext(packageRLP)
		bctx.makeRawValueOp()
	})
}

func Fuzz_Nosy_buildContext_makeSliceOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var typ *types.Slice
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || typ == nil {
			return
		}

		bctx := newBuildContext(packageRLP)
		bctx.makeSliceOp(typ)
	})
}

func Fuzz_Nosy_buildContext_makeStructOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var packageRLP *types.Package
		fill_err = tp.Fill(&packageRLP)
		if fill_err != nil {
			return
		}
		var named *types.Named
		fill_err = tp.Fill(&named)
		if fill_err != nil {
			return
		}
		var typ *types.Struct
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		if packageRLP == nil || named == nil || typ == nil {
			return
		}

		bctx := newBuildContext(packageRLP)
		bctx.makeStructOp(named, typ)
	})
}

// skipping Fuzz_Nosy_buildContext_typeToStructType__ because parameters include func, chan, or unsupported interface: go/types.Type

func Fuzz_Nosy_genContext_addImport__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inPackage *types.Package
		fill_err = tp.Fill(&inPackage)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if inPackage == nil {
			return
		}

		ctx := newGenContext(inPackage)
		ctx.addImport(path)
	})
}

func Fuzz_Nosy_genContext_importsList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inPackage *types.Package
		fill_err = tp.Fill(&inPackage)
		if fill_err != nil {
			return
		}
		if inPackage == nil {
			return
		}

		ctx := newGenContext(inPackage)
		ctx.importsList()
	})
}

func Fuzz_Nosy_genContext_qualify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inPackage *types.Package
		fill_err = tp.Fill(&inPackage)
		if fill_err != nil {
			return
		}
		var pkg *types.Package
		fill_err = tp.Fill(&pkg)
		if fill_err != nil {
			return
		}
		if inPackage == nil || pkg == nil {
			return
		}

		ctx := newGenContext(inPackage)
		ctx.qualify(pkg)
	})
}

func Fuzz_Nosy_genContext_resetTemp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inPackage *types.Package
		fill_err = tp.Fill(&inPackage)
		if fill_err != nil {
			return
		}
		if inPackage == nil {
			return
		}

		ctx := newGenContext(inPackage)
		ctx.resetTemp()
	})
}

func Fuzz_Nosy_genContext_temp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inPackage *types.Package
		fill_err = tp.Fill(&inPackage)
		if fill_err != nil {
			return
		}
		if inPackage == nil {
			return
		}

		ctx := newGenContext(inPackage)
		ctx.temp()
	})
}

func Fuzz_Nosy_basicOp_decodeNeedsConversion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op basicOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}

		op.decodeNeedsConversion()
	})
}

func Fuzz_Nosy_basicOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op basicOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_basicOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op basicOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_basicOp_writeNeedsConversion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op basicOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}

		op.writeNeedsConversion()
	})
}

func Fuzz_Nosy_bigIntOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op bigIntOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_bigIntOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op bigIntOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_byteArrayOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op byteArrayOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_byteArrayOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op byteArrayOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_encoderDecoderOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op encoderDecoderOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_encoderDecoderOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op encoderDecoderOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

// skipping Fuzz_Nosy_op_genDecode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp/rlpgen.op

// skipping Fuzz_Nosy_op_genWrite__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp/rlpgen.op

func Fuzz_Nosy_ptrOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op ptrOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_ptrOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op ptrOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_sliceOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op sliceOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_sliceOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op sliceOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_structOp_decodeOptionalFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op structOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var resultV string
		fill_err = tp.Fill(&resultV)
		if fill_err != nil {
			return
		}
		if b == nil || ctx == nil {
			return
		}

		op.decodeOptionalFields(b, ctx, resultV)
	})
}

func Fuzz_Nosy_structOp_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op structOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_structOp_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op structOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

func Fuzz_Nosy_structOp_writeOptionalFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op structOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var b *bytes.Buffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if b == nil || ctx == nil {
			return
		}

		op.writeOptionalFields(b, ctx, v)
	})
}

func Fuzz_Nosy_uint256Op_genDecode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op uint256Op
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genDecode(ctx)
	})
}

func Fuzz_Nosy_uint256Op_genWrite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op uint256Op
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx *genContext
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		op.genWrite(ctx, v)
	})
}

// skipping Fuzz_Nosy_fatal__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_generateDecoder__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp/rlpgen.op

// skipping Fuzz_Nosy_generateEncoder__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rlp/rlpgen.op

// skipping Fuzz_Nosy_isBigInt__ because parameters include func, chan, or unsupported interface: go/types.Type

// skipping Fuzz_Nosy_isByte__ because parameters include func, chan, or unsupported interface: go/types.Type

// skipping Fuzz_Nosy_isUint256__ because parameters include func, chan, or unsupported interface: go/types.Type

// skipping Fuzz_Nosy_nonZeroCheck__ because parameters include func, chan, or unsupported interface: go/types.Type
