package flags

import (
	"flag"
	"testing"

	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_BigFlag_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var set *flag.FlagSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if f1 == nil || set == nil {
			return
		}

		f1.Apply(set)
	})
}

func Fuzz_Nosy_BigFlag_GetCategory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetCategory()
	})
}

func Fuzz_Nosy_BigFlag_GetDefaultText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetDefaultText()
	})
}

func Fuzz_Nosy_BigFlag_GetEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetEnvVars()
	})
}

func Fuzz_Nosy_BigFlag_GetUsage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetUsage()
	})
}

func Fuzz_Nosy_BigFlag_GetValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetValue()
	})
}

func Fuzz_Nosy_BigFlag_IsRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsRequired()
	})
}

func Fuzz_Nosy_BigFlag_IsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsSet()
	})
}

func Fuzz_Nosy_BigFlag_IsVisible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsVisible()
	})
}

func Fuzz_Nosy_BigFlag_Names__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Names()
	})
}

func Fuzz_Nosy_BigFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.String()
	})
}

func Fuzz_Nosy_BigFlag_TakesValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *BigFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.TakesValue()
	})
}

func Fuzz_Nosy_DirectoryFlag_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var set *flag.FlagSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if f1 == nil || set == nil {
			return
		}

		f1.Apply(set)
	})
}

func Fuzz_Nosy_DirectoryFlag_GetCategory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetCategory()
	})
}

func Fuzz_Nosy_DirectoryFlag_GetDefaultText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetDefaultText()
	})
}

func Fuzz_Nosy_DirectoryFlag_GetEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetEnvVars()
	})
}

func Fuzz_Nosy_DirectoryFlag_GetUsage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetUsage()
	})
}

func Fuzz_Nosy_DirectoryFlag_GetValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetValue()
	})
}

func Fuzz_Nosy_DirectoryFlag_IsRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsRequired()
	})
}

func Fuzz_Nosy_DirectoryFlag_IsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsSet()
	})
}

func Fuzz_Nosy_DirectoryFlag_IsVisible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsVisible()
	})
}

func Fuzz_Nosy_DirectoryFlag_Names__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Names()
	})
}

func Fuzz_Nosy_DirectoryFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.String()
	})
}

func Fuzz_Nosy_DirectoryFlag_TakesValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *DirectoryFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.TakesValue()
	})
}

func Fuzz_Nosy_DirectoryString_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DirectoryString
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Set(value)
	})
}

func Fuzz_Nosy_DirectoryString_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DirectoryString
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_Apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var set *flag.FlagSet
		fill_err = tp.Fill(&set)
		if fill_err != nil {
			return
		}
		if f1 == nil || set == nil {
			return
		}

		f1.Apply(set)
	})
}

func Fuzz_Nosy_TextMarshalerFlag_GetCategory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetCategory()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_GetDefaultText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetDefaultText()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_GetEnvVars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetEnvVars()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_GetUsage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetUsage()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_GetValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetValue()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_IsRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsRequired()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_IsSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsSet()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_IsVisible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.IsVisible()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_Names__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Names()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.String()
	})
}

func Fuzz_Nosy_TextMarshalerFlag_TakesValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *TextMarshalerFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.TakesValue()
	})
}

func Fuzz_Nosy_bigValue_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bigValue
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Set(s)
	})
}

func Fuzz_Nosy_bigValue_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bigValue
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_textMarshalerVal_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v textMarshalerVal
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		v.Set(s)
	})
}

func Fuzz_Nosy_textMarshalerVal_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v textMarshalerVal
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.String()
	})
}

// skipping Fuzz_Nosy_AutoEnvVars__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

// skipping Fuzz_Nosy_CheckEnvVars__ because parameters include func, chan, or unsupported interface: []github.com/urfave/cli/v2.Flag

// skipping Fuzz_Nosy_FlagString__ because parameters include func, chan, or unsupported interface: github.com/urfave/cli/v2.Flag

func Fuzz_Nosy_Merge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var groups [][]cli.Flag
		fill_err = tp.Fill(&groups)
		if fill_err != nil {
			return
		}

		Merge(groups...)
	})
}

func Fuzz_Nosy_MigrateGlobalFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		MigrateGlobalFlags(ctx)
	})
}

func Fuzz_Nosy_doMigrateFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		doMigrateFlags(ctx)
	})
}

// skipping Fuzz_Nosy_eachName__ because parameters include func, chan, or unsupported interface: github.com/urfave/cli/v2.Flag

func Fuzz_Nosy_expandPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		expandPath(p)
	})
}

func Fuzz_Nosy_indent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, nspace int) {
		indent(s, nspace)
	})
}

func Fuzz_Nosy_wordWrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, width int) {
		wordWrap(s, width)
	})
}
