package external

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core/types"
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

// skipping Fuzz_Nosy_ExternalBackend_Subscribe__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/accounts.WalletEvent

func Fuzz_Nosy_ExternalBackend_Wallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		eb, err := NewExternalBackend(endpoint)
		if err != nil {
			return
		}
		eb.Wallets()
	})
}

func Fuzz_Nosy_ExternalSigner_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Accounts()
	})
}

func Fuzz_Nosy_ExternalSigner_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Close()
	})
}

func Fuzz_Nosy_ExternalSigner_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Contains(account)
	})
}

func Fuzz_Nosy_ExternalSigner_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var pin bool
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Derive(path, pin)
	})
}

func Fuzz_Nosy_ExternalSigner_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string, passphrase string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Open(passphrase)
	})
}

// skipping Fuzz_Nosy_ExternalSigner_SelfDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum.ChainStateReader

func Fuzz_Nosy_ExternalSigner_SignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var mimeType string
		fill_err = tp.Fill(&mimeType)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignData(account, mimeType, d4)
	})
}

func Fuzz_Nosy_ExternalSigner_SignDataWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var mimeType string
		fill_err = tp.Fill(&mimeType)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignDataWithPassphrase(account, passphrase, mimeType, d5)
	})
}

func Fuzz_Nosy_ExternalSigner_SignText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignText(account, text)
	})
}

func Fuzz_Nosy_ExternalSigner_SignTextWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignTextWithPassphrase(account, passphrase, text)
	})
}

func Fuzz_Nosy_ExternalSigner_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignTx(account, tx, chainID)
	})
}

func Fuzz_Nosy_ExternalSigner_SignTxWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var tx *types.Transaction
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var chainID *big.Int
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		if tx == nil || chainID == nil {
			return
		}

		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.SignTxWithPassphrase(account, passphrase, tx, chainID)
	})
}

func Fuzz_Nosy_ExternalSigner_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.Status()
	})
}

func Fuzz_Nosy_ExternalSigner_URL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.URL()
	})
}

func Fuzz_Nosy_ExternalSigner_listAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.listAccounts()
	})
}

func Fuzz_Nosy_ExternalSigner_pingVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string) {
		api, err := NewExternalSigner(endpoint)
		if err != nil {
			return
		}
		api.pingVersion()
	})
}
