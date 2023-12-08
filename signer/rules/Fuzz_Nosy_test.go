package rules

import (
	"testing"

	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/signer/core"
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

func Fuzz_Nosy_rulesetUI_ApproveListing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var request *core.ListRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if r == nil || request == nil {
			return
		}

		r.ApproveListing(request)
	})
}

func Fuzz_Nosy_rulesetUI_ApproveNewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var request *core.NewAccountRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if r == nil || request == nil {
			return
		}

		r.ApproveNewAccount(request)
	})
}

func Fuzz_Nosy_rulesetUI_ApproveSignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var request *core.SignDataRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if r == nil || request == nil {
			return
		}

		r.ApproveSignData(request)
	})
}

func Fuzz_Nosy_rulesetUI_ApproveTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var request *core.SignTxRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if r == nil || request == nil {
			return
		}

		r.ApproveTx(request)
	})
}

func Fuzz_Nosy_rulesetUI_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var javascriptRules string
		fill_err = tp.Fill(&javascriptRules)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Init(javascriptRules)
	})
}

func Fuzz_Nosy_rulesetUI_OnApprovedTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var tx ethapi.SignTransactionResult
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.OnApprovedTx(tx)
	})
}

func Fuzz_Nosy_rulesetUI_OnInputRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var info core.UserInputRequest
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.OnInputRequired(info)
	})
}

func Fuzz_Nosy_rulesetUI_OnSignerStartup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var info core.StartupInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.OnSignerStartup(info)
	})
}

func Fuzz_Nosy_rulesetUI_RegisterUIServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var api *core.UIServerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if r == nil || api == nil {
			return
		}

		r.RegisterUIServer(api)
	})
}

func Fuzz_Nosy_rulesetUI_ShowError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var message string
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ShowError(message)
	})
}

func Fuzz_Nosy_rulesetUI_ShowInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rulesetUI
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var message string
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ShowInfo(message)
	})
}

// skipping Fuzz_Nosy_rulesetUI_checkApproval__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_rulesetUI_execute__ because parameters include func, chan, or unsupported interface: interface{}
