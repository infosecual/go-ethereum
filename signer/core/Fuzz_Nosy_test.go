package core

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
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

func Fuzz_Nosy_AuditLogger_EcRecover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 hexutil.Bytes
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var sig hexutil.Bytes
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.EcRecover(ctx, d3, sig)
	})
}

func Fuzz_Nosy_AuditLogger_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.List(ctx)
	})
}

func Fuzz_Nosy_AuditLogger_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.New(ctx)
	})
}

// skipping Fuzz_Nosy_AuditLogger_SignData__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_AuditLogger_SignGnosisSafeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.MixedcaseAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var gnosisTx GnosisSafeTx
		fill_err = tp.Fill(&gnosisTx)
		if fill_err != nil {
			return
		}
		var methodSelector *string
		fill_err = tp.Fill(&methodSelector)
		if fill_err != nil {
			return
		}
		if l == nil || methodSelector == nil {
			return
		}

		l.SignGnosisSafeTx(ctx, addr, gnosisTx, methodSelector)
	})
}

func Fuzz_Nosy_AuditLogger_SignTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args apitypes.SendTxArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var methodSelector *string
		fill_err = tp.Fill(&methodSelector)
		if fill_err != nil {
			return
		}
		if l == nil || methodSelector == nil {
			return
		}

		l.SignTransaction(ctx, args, methodSelector)
	})
}

func Fuzz_Nosy_AuditLogger_SignTypedData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.MixedcaseAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var d4 apitypes.TypedData
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.SignTypedData(ctx, addr, d4)
	})
}

func Fuzz_Nosy_AuditLogger_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *AuditLogger
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.Version(ctx)
	})
}

func Fuzz_Nosy_CommandlineUI_ApproveListing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *ListRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewCommandlineUI()
		ui.ApproveListing(request)
	})
}

func Fuzz_Nosy_CommandlineUI_ApproveNewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *NewAccountRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewCommandlineUI()
		ui.ApproveNewAccount(request)
	})
}

func Fuzz_Nosy_CommandlineUI_ApproveSignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *SignDataRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewCommandlineUI()
		ui.ApproveSignData(request)
	})
}

func Fuzz_Nosy_CommandlineUI_ApproveTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *SignTxRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewCommandlineUI()
		ui.ApproveTx(request)
	})
}

func Fuzz_Nosy_CommandlineUI_OnApprovedTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx ethapi.SignTransactionResult
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}

		ui := NewCommandlineUI()
		ui.OnApprovedTx(tx)
	})
}

func Fuzz_Nosy_CommandlineUI_OnInputRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info UserInputRequest
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		ui := NewCommandlineUI()
		ui.OnInputRequired(info)
	})
}

func Fuzz_Nosy_CommandlineUI_OnSignerStartup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info StartupInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		ui := NewCommandlineUI()
		ui.OnSignerStartup(info)
	})
}

func Fuzz_Nosy_CommandlineUI_RegisterUIServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *UIServerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		ui := NewCommandlineUI()
		ui.RegisterUIServer(api)
	})
}

func Fuzz_Nosy_CommandlineUI_ShowError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message string) {
		ui := NewCommandlineUI()
		ui.ShowError(message)
	})
}

func Fuzz_Nosy_CommandlineUI_ShowInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message string) {
		ui := NewCommandlineUI()
		ui.ShowInfo(message)
	})
}

func Fuzz_Nosy_GnosisSafeTx_ArgsForValidation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *GnosisSafeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.ArgsForValidation()
	})
}

func Fuzz_Nosy_GnosisSafeTx_ToTypedData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *GnosisSafeTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		tx.ToTypedData()
	})
}

func Fuzz_Nosy_SignerAPI_EcRecover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var d3 hexutil.Bytes
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var sig hexutil.Bytes
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.EcRecover(ctx, d3, sig)
	})
}

func Fuzz_Nosy_SignerAPI_List__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.List(ctx)
	})
}

func Fuzz_Nosy_SignerAPI_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.New(ctx)
	})
}

// skipping Fuzz_Nosy_SignerAPI_SignData__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_SignerAPI_SignGnosisSafeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var signerAddress common.MixedcaseAddress
		fill_err = tp.Fill(&signerAddress)
		if fill_err != nil {
			return
		}
		var gnosisTx GnosisSafeTx
		fill_err = tp.Fill(&gnosisTx)
		if fill_err != nil {
			return
		}
		var methodSelector *string
		fill_err = tp.Fill(&methodSelector)
		if fill_err != nil {
			return
		}
		if api == nil || methodSelector == nil {
			return
		}

		api.SignGnosisSafeTx(ctx, signerAddress, gnosisTx, methodSelector)
	})
}

func Fuzz_Nosy_SignerAPI_SignTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var args apitypes.SendTxArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		var methodSelector *string
		fill_err = tp.Fill(&methodSelector)
		if fill_err != nil {
			return
		}
		if api == nil || methodSelector == nil {
			return
		}

		api.SignTransaction(ctx, args, methodSelector)
	})
}

func Fuzz_Nosy_SignerAPI_SignTypedData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.MixedcaseAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var typedData apitypes.TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.SignTypedData(ctx, addr, typedData)
	})
}

func Fuzz_Nosy_SignerAPI_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Version(ctx)
	})
}

// skipping Fuzz_Nosy_SignerAPI_derivationLoop__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum/go-ethereum/accounts.WalletEvent

// skipping Fuzz_Nosy_SignerAPI_determineSignatureFormat__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_SignerAPI_lookupOrQueryPassword__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var title string
		fill_err = tp.Fill(&title)
		if fill_err != nil {
			return
		}
		var prompt string
		fill_err = tp.Fill(&prompt)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.lookupOrQueryPassword(address, title, prompt)
	})
}

func Fuzz_Nosy_SignerAPI_lookupPassword__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.lookupPassword(address)
	})
}

func Fuzz_Nosy_SignerAPI_newAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.newAccount()
	})
}

func Fuzz_Nosy_SignerAPI_openTrezor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var url accounts.URL
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.openTrezor(url)
	})
}

func Fuzz_Nosy_SignerAPI_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var req *SignDataRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var legacyV bool
		fill_err = tp.Fill(&legacyV)
		if fill_err != nil {
			return
		}
		if api == nil || req == nil {
			return
		}

		api.sign(req, legacyV)
	})
}

func Fuzz_Nosy_SignerAPI_signTypedData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.MixedcaseAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var typedData apitypes.TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		var validationMessages *apitypes.ValidationMessages
		fill_err = tp.Fill(&validationMessages)
		if fill_err != nil {
			return
		}
		if api == nil || validationMessages == nil {
			return
		}

		api.signTypedData(ctx, addr, typedData, validationMessages)
	})
}

func Fuzz_Nosy_SignerAPI_startUSBListener__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.startUSBListener()
	})
}

func Fuzz_Nosy_StdIOUI_ApproveListing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *ListRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewStdIOUI()
		ui.ApproveListing(request)
	})
}

func Fuzz_Nosy_StdIOUI_ApproveNewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *NewAccountRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewStdIOUI()
		ui.ApproveNewAccount(request)
	})
}

func Fuzz_Nosy_StdIOUI_ApproveSignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *SignDataRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewStdIOUI()
		ui.ApproveSignData(request)
	})
}

func Fuzz_Nosy_StdIOUI_ApproveTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var request *SignTxRequest
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if request == nil {
			return
		}

		ui := NewStdIOUI()
		ui.ApproveTx(request)
	})
}

func Fuzz_Nosy_StdIOUI_OnApprovedTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx ethapi.SignTransactionResult
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}

		ui := NewStdIOUI()
		ui.OnApprovedTx(tx)
	})
}

func Fuzz_Nosy_StdIOUI_OnInputRequired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info UserInputRequest
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		ui := NewStdIOUI()
		ui.OnInputRequired(info)
	})
}

func Fuzz_Nosy_StdIOUI_OnSignerStartup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info StartupInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}

		ui := NewStdIOUI()
		ui.OnSignerStartup(info)
	})
}

func Fuzz_Nosy_StdIOUI_RegisterUIServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *UIServerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		ui := NewStdIOUI()
		ui.RegisterUIServer(api)
	})
}

func Fuzz_Nosy_StdIOUI_ShowError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message string) {
		ui := NewStdIOUI()
		ui.ShowError(message)
	})
}

func Fuzz_Nosy_StdIOUI_ShowInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message string) {
		ui := NewStdIOUI()
		ui.ShowInfo(message)
	})
}

// skipping Fuzz_Nosy_StdIOUI_dispatch__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_StdIOUI_notify__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_UIServerAPI_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.ChainId()
	})
}

func Fuzz_Nosy_UIServerAPI_DeriveAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var pin *bool
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}
		if extapi == nil || pin == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.DeriveAccount(url, path, pin)
	})
}

func Fuzz_Nosy_UIServerAPI_Export__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.Export(ctx, addr)
	})
}

func Fuzz_Nosy_UIServerAPI_Import__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var keyJSON json.RawMessage
		fill_err = tp.Fill(&keyJSON)
		if fill_err != nil {
			return
		}
		var oldPassphrase string
		fill_err = tp.Fill(&oldPassphrase)
		if fill_err != nil {
			return
		}
		var newPassphrase string
		fill_err = tp.Fill(&newPassphrase)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		api := NewUIServerAPI(extapi)
		api.Import(ctx, keyJSON, oldPassphrase, newPassphrase)
	})
}

func Fuzz_Nosy_UIServerAPI_ImportRawKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var privkey string
		fill_err = tp.Fill(&privkey)
		if fill_err != nil {
			return
		}
		var password string
		fill_err = tp.Fill(&password)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.ImportRawKey(privkey, password)
	})
}

func Fuzz_Nosy_UIServerAPI_ListAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.ListAccounts(ctx)
	})
}

func Fuzz_Nosy_UIServerAPI_ListWallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.ListWallets()
	})
}

func Fuzz_Nosy_UIServerAPI_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		api := NewUIServerAPI(extapi)
		api.New(ctx)
	})
}

func Fuzz_Nosy_UIServerAPI_OpenWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var passphrase *string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if extapi == nil || passphrase == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.OpenWallet(url, passphrase)
	})
}

func Fuzz_Nosy_UIServerAPI_SetChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var extapi *SignerAPI
		fill_err = tp.Fill(&extapi)
		if fill_err != nil {
			return
		}
		var id math.HexOrDecimal64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if extapi == nil {
			return
		}

		s := NewUIServerAPI(extapi)
		s.SetChainId(id)
	})
}

// skipping Fuzz_Nosy_ExternalAPI_EcRecover__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_List__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_New__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_SignData__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_SignGnosisSafeTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_SignTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_SignTypedData__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

// skipping Fuzz_Nosy_ExternalAPI_Version__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.ExternalAPI

func Fuzz_Nosy_Metadata_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		m := MetadataFromContext(ctx)
		m.String()
	})
}

// skipping Fuzz_Nosy_UIClientAPI_ApproveListing__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_ApproveNewAccount__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_ApproveSignData__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_ApproveTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_OnApprovedTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_OnInputRequired__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_OnSignerStartup__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_RegisterUIServer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_ShowError__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_UIClientAPI_ShowInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

// skipping Fuzz_Nosy_Validator_ValidateTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.Validator

func Fuzz_Nosy_SignTextValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var validatorData apitypes.ValidatorData
		fill_err = tp.Fill(&validatorData)
		if fill_err != nil {
			return
		}

		SignTextValidator(validatorData)
	})
}

func Fuzz_Nosy_cliqueHeaderHashAndRlp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		cliqueHeaderHashAndRlp(header)
	})
}

// skipping Fuzz_Nosy_fromHex__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_logDiff__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var original *SignTxRequest
		fill_err = tp.Fill(&original)
		if fill_err != nil {
			return
		}
		var new *SignTxResponse
		fill_err = tp.Fill(&new)
		if fill_err != nil {
			return
		}
		if original == nil || new == nil {
			return
		}

		logDiff(original, new)
	})
}

func Fuzz_Nosy_sanitize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txt string, limit int) {
		sanitize(txt, limit)
	})
}

func Fuzz_Nosy_showMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var metadata Metadata
		fill_err = tp.Fill(&metadata)
		if fill_err != nil {
			return
		}

		showMetadata(metadata)
	})
}
