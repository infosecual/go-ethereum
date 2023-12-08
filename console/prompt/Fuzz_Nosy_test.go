package prompt

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

func Fuzz_Nosy_terminalPrompter_AppendHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, command string) {
		p := newTerminalPrompter()
		p.AppendHistory(command)
	})
}

func Fuzz_Nosy_terminalPrompter_PromptConfirm__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p1 string) {
		p := newTerminalPrompter()
		p.PromptConfirm(p1)
	})
}

func Fuzz_Nosy_terminalPrompter_PromptInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p1 string) {
		p := newTerminalPrompter()
		p.PromptInput(p1)
	})
}

func Fuzz_Nosy_terminalPrompter_PromptPassword__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p1 string) {
		p := newTerminalPrompter()
		p.PromptPassword(p1)
	})
}

func Fuzz_Nosy_terminalPrompter_SetHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var history []string
		fill_err = tp.Fill(&history)
		if fill_err != nil {
			return
		}

		p := newTerminalPrompter()
		p.SetHistory(history)
	})
}

// skipping Fuzz_Nosy_terminalPrompter_SetWordCompleter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.WordCompleter

// skipping Fuzz_Nosy_UserPrompter_AppendHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_ClearHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_PromptConfirm__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_PromptInput__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_PromptPassword__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_SetHistory__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter

// skipping Fuzz_Nosy_UserPrompter_SetWordCompleter__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/console/prompt.UserPrompter
