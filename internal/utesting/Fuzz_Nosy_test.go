package utesting

import (
	"io"
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

// skipping Fuzz_Nosy_T_Error__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_T_Errorf__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_T_Fail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Fail()
	})
}

func Fuzz_Nosy_T_FailNow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.FailNow()
	})
}

func Fuzz_Nosy_T_Failed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Failed()
	})
}

// skipping Fuzz_Nosy_T_Fatal__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_T_Fatalf__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_T_Helper__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Helper()
	})
}

// skipping Fuzz_Nosy_T_Log__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_T_Logf__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_consoleOutput_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		c := newConsoleOutput(w)
		c.Write(b)
	})
}

func Fuzz_Nosy_consoleOutput_testResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var r Result
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		c := newConsoleOutput(w)
		c.testResult(r)
	})
}

func Fuzz_Nosy_consoleOutput_testStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		c := newConsoleOutput(w)
		c.testStart(name)
	})
}

func Fuzz_Nosy_indentWriter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var indent string
		fill_err = tp.Fill(&indent)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		w := newIndentWriter(indent, out)
		w.Write(b)
	})
}

func Fuzz_Nosy_indentWriter_flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var indent string
		fill_err = tp.Fill(&indent)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		w := newIndentWriter(indent, out)
		w.flush()
	})
}

func Fuzz_Nosy_tapOutput_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var numTests int
		fill_err = tp.Fill(&numTests)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		t1 := newTAP(out, numTests)
		t1.Write(b)
	})
}

func Fuzz_Nosy_tapOutput_testResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var numTests int
		fill_err = tp.Fill(&numTests)
		if fill_err != nil {
			return
		}
		var r Result
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}

		t1 := newTAP(out, numTests)
		t1.testResult(r)
	})
}

func Fuzz_Nosy_tapOutput_testStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var numTests int
		fill_err = tp.Fill(&numTests)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}

		t1 := newTAP(out, numTests)
		t1.testStart(name)
	})
}

// skipping Fuzz_Nosy_testOutput_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/utesting.testOutput

// skipping Fuzz_Nosy_testOutput_testResult__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/utesting.testOutput

// skipping Fuzz_Nosy_testOutput_testStart__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/utesting.testOutput

func Fuzz_Nosy_CountFailures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rr []Result
		fill_err = tp.Fill(&rr)
		if fill_err != nil {
			return
		}

		CountFailures(rr)
	})
}

func Fuzz_Nosy_MatchTests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tests []Test
		fill_err = tp.Fill(&tests)
		if fill_err != nil {
			return
		}
		var expr string
		fill_err = tp.Fill(&expr)
		if fill_err != nil {
			return
		}

		MatchTests(tests, expr)
	})
}

func Fuzz_Nosy_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var test Test
		fill_err = tp.Fill(&test)
		if fill_err != nil {
			return
		}

		Run(test)
	})
}

func Fuzz_Nosy_RunTAP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tests []Test
		fill_err = tp.Fill(&tests)
		if fill_err != nil {
			return
		}
		var report io.Writer
		fill_err = tp.Fill(&report)
		if fill_err != nil {
			return
		}

		RunTAP(tests, report)
	})
}

func Fuzz_Nosy_RunTests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tests []Test
		fill_err = tp.Fill(&tests)
		if fill_err != nil {
			return
		}
		var report io.Writer
		fill_err = tp.Fill(&report)
		if fill_err != nil {
			return
		}

		RunTests(tests, report)
	})
}

// skipping Fuzz_Nosy_run__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/utesting.testOutput

func Fuzz_Nosy_runTest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var test Test
		fill_err = tp.Fill(&test)
		if fill_err != nil {
			return
		}
		var output io.Writer
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}

		runTest(test, output)
	})
}
