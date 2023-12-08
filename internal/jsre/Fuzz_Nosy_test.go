package jsre

import (
	"io"
	"testing"

	"github.com/dop251/goja"
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

func Fuzz_Nosy_JSRE_Compile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var src string
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.Compile(filename, src)
	})
}

func Fuzz_Nosy_JSRE_CompleteKeywords__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var line string
		fill_err = tp.Fill(&line)
		if fill_err != nil {
			return
		}

		j1 := New(assetPath, output)
		j1.CompleteKeywords(line)
	})
}

// skipping Fuzz_Nosy_JSRE_Do__ because parameters include func, chan, or unsupported interface: func(*github.com/dop251/goja.Runtime)

func Fuzz_Nosy_JSRE_Evaluate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var code string
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.Evaluate(code, w)
	})
}

func Fuzz_Nosy_JSRE_Exec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.Exec(file)
	})
}

// skipping Fuzz_Nosy_JSRE_Interrupt__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_JSRE_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var code string
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.Run(code)
	})
}

// skipping Fuzz_Nosy_JSRE_Set__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_JSRE_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var waitForCallbacks bool
		fill_err = tp.Fill(&waitForCallbacks)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.Stop(waitForCallbacks)
	})
}

func Fuzz_Nosy_JSRE_loadScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var call Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.loadScript(call)
	})
}

func Fuzz_Nosy_JSRE_prettyPrintJS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var call goja.FunctionCall
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.prettyPrintJS(call)
	})
}

func Fuzz_Nosy_JSRE_runEventLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var assetPath string
		fill_err = tp.Fill(&assetPath)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}

		re := New(assetPath, output)
		re.runEventLoop()
	})
}

func Fuzz_Nosy_ppctx_fields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx ppctx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var obj *goja.Object
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		if obj == nil {
			return
		}

		ctx.fields(obj)
	})
}

func Fuzz_Nosy_ppctx_indent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx ppctx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var level int
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}

		ctx.indent(level)
	})
}

func Fuzz_Nosy_ppctx_isBigNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx ppctx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var v *goja.Object
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		ctx.isBigNumber(v)
	})
}

func Fuzz_Nosy_ppctx_printObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx ppctx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var obj *goja.Object
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		var level int
		fill_err = tp.Fill(&level)
		if fill_err != nil {
			return
		}
		var inArray bool
		fill_err = tp.Fill(&inArray)
		if fill_err != nil {
			return
		}
		if obj == nil {
			return
		}

		ctx.printObject(obj, level, inArray)
	})
}

// skipping Fuzz_Nosy_ppctx_printValue__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

func Fuzz_Nosy_getCompletions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vm *goja.Runtime
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var line string
		fill_err = tp.Fill(&line)
		if fill_err != nil {
			return
		}
		if vm == nil {
			return
		}

		getCompletions(vm, line)
	})
}

// skipping Fuzz_Nosy_iterOwnAndConstructorKeys__ because parameters include func, chan, or unsupported interface: func(string)

// skipping Fuzz_Nosy_iterOwnKeys__ because parameters include func, chan, or unsupported interface: func(string)

// skipping Fuzz_Nosy_prettyError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_prettyPrint__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

func Fuzz_Nosy_toString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var obj *goja.Object
		fill_err = tp.Fill(&obj)
		if fill_err != nil {
			return
		}
		if obj == nil {
			return
		}

		toString(obj)
	})
}
