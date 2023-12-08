package console

import (
	"testing"

	"github.com/dop251/goja"
	"github.com/ethereum/go-ethereum/internal/jsre"
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

func Fuzz_Nosy_Console_AutoCompleteInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var line string
		fill_err = tp.Fill(&line)
		if fill_err != nil {
			return
		}
		var pos int
		fill_err = tp.Fill(&pos)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.AutoCompleteInput(line, pos)
	})
}

func Fuzz_Nosy_Console_Evaluate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var statement string
		fill_err = tp.Fill(&statement)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.Evaluate(statement)
	})
}

func Fuzz_Nosy_Console_Interactive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.Interactive()
	})
}

func Fuzz_Nosy_Console_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var graceful bool
		fill_err = tp.Fill(&graceful)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.Stop(graceful)
	})
}

func Fuzz_Nosy_Console_StopInteractive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.StopInteractive()
	})
}

func Fuzz_Nosy_Console_Welcome__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.Welcome()
	})
}

func Fuzz_Nosy_Console_clearHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.clearHistory()
	})
}

func Fuzz_Nosy_Console_clearSignalReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.clearSignalReceived()
	})
}

func Fuzz_Nosy_Console_consoleOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var call goja.FunctionCall
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.consoleOutput(call)
	})
}

func Fuzz_Nosy_Console_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var preload []string
		fill_err = tp.Fill(&preload)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.init(preload)
	})
}

func Fuzz_Nosy_Console_initAdmin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var vm *goja.Runtime
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var bridge *bridge
		fill_err = tp.Fill(&bridge)
		if fill_err != nil {
			return
		}
		if vm == nil || bridge == nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.initAdmin(vm, bridge)
	})
}

func Fuzz_Nosy_Console_initConsoleObject__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.initConsoleObject()
	})
}

func Fuzz_Nosy_Console_initExtensions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.initExtensions()
	})
}

func Fuzz_Nosy_Console_initPersonal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var vm *goja.Runtime
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var bridge *bridge
		fill_err = tp.Fill(&bridge)
		if fill_err != nil {
			return
		}
		if vm == nil || bridge == nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.initPersonal(vm, bridge)
	})
}

func Fuzz_Nosy_Console_initWeb3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var bridge *bridge
		fill_err = tp.Fill(&bridge)
		if fill_err != nil {
			return
		}
		if bridge == nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.initWeb3(bridge)
	})
}

func Fuzz_Nosy_Console_interruptHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.interruptHandler()
	})
}

// skipping Fuzz_Nosy_Console_readLines__ because parameters include func, chan, or unsupported interface: chan<- string

func Fuzz_Nosy_Console_setSignalReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.setSignalReceived()
	})
}

func Fuzz_Nosy_Console_writeHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}

		c, err := New(config)
		if err != nil {
			return
		}
		c.writeHistory()
	})
}

func Fuzz_Nosy_bridge_NewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.NewAccount(call)
	})
}

func Fuzz_Nosy_bridge_OpenWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.OpenWallet(call)
	})
}

func Fuzz_Nosy_bridge_Send__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Send(call)
	})
}

func Fuzz_Nosy_bridge_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Sign(call)
	})
}

func Fuzz_Nosy_bridge_Sleep__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Sleep(call)
	})
}

func Fuzz_Nosy_bridge_SleepBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SleepBlocks(call)
	})
}

func Fuzz_Nosy_bridge_UnlockAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnlockAccount(call)
	})
}

func Fuzz_Nosy_bridge_readPassphraseAndReopenWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.readPassphraseAndReopenWallet(call)
	})
}

func Fuzz_Nosy_bridge_readPinAndReopenWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *bridge
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var call jsre.Call
		fill_err = tp.Fill(&call)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.readPinAndReopenWallet(call)
	})
}

func Fuzz_Nosy_countIndents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		countIndents(input)
	})
}

// skipping Fuzz_Nosy_isNumber__ because parameters include func, chan, or unsupported interface: github.com/dop251/goja.Value

// skipping Fuzz_Nosy_setError__ because parameters include func, chan, or unsupported interface: interface{}
