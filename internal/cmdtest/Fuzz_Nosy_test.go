package cmdtest

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

func Fuzz_Nosy_TestCmd_CloseStdin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.CloseStdin()
	})
}

func Fuzz_Nosy_TestCmd_ExitStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.ExitStatus()
	})
}

func Fuzz_Nosy_TestCmd_Expect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var tplsource string
		fill_err = tp.Fill(&tplsource)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.Expect(tplsource)
	})
}

func Fuzz_Nosy_TestCmd_ExpectExit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.ExpectExit()
	})
}

func Fuzz_Nosy_TestCmd_ExpectRegexp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var regex string
		fill_err = tp.Fill(&regex)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.ExpectRegexp(regex)
	})
}

func Fuzz_Nosy_TestCmd_InputLine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.InputLine(s)
	})
}

func Fuzz_Nosy_TestCmd_Interrupt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.Interrupt()
	})
}

func Fuzz_Nosy_TestCmd_Kill__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.Kill()
	})
}

func Fuzz_Nosy_TestCmd_Output__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.Output()
	})
}

func Fuzz_Nosy_TestCmd_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.Run(name, args...)
	})
}

// skipping Fuzz_Nosy_TestCmd_SetTemplateFunc__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_TestCmd_StderrText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.StderrText()
	})
}

func Fuzz_Nosy_TestCmd_WaitExit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.WaitExit()
	})
}

func Fuzz_Nosy_TestCmd_matchExactOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tt *TestCmd
		fill_err = tp.Fill(&tt)
		if fill_err != nil {
			return
		}
		var want []byte
		fill_err = tp.Fill(&want)
		if fill_err != nil {
			return
		}
		if tt == nil {
			return
		}

		tt.matchExactOutput(want)
	})
}

// skipping Fuzz_Nosy_TestCmd_withKillTimeout__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_runeTee_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rtee *runeTee
		fill_err = tp.Fill(&rtee)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if rtee == nil {
			return
		}

		rtee.Read(b)
	})
}

func Fuzz_Nosy_runeTee_ReadByte__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rtee *runeTee
		fill_err = tp.Fill(&rtee)
		if fill_err != nil {
			return
		}
		if rtee == nil {
			return
		}

		rtee.ReadByte()
	})
}

func Fuzz_Nosy_runeTee_ReadRune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rtee *runeTee
		fill_err = tp.Fill(&rtee)
		if fill_err != nil {
			return
		}
		if rtee == nil {
			return
		}

		rtee.ReadRune()
	})
}

func Fuzz_Nosy_testlogger_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tl *testlogger
		fill_err = tp.Fill(&tl)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if tl == nil {
			return
		}

		tl.Write(b)
	})
}
