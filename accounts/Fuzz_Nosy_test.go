package accounts

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

func Fuzz_Nosy_AuthNeededError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *AuthNeededError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_DerivationPath_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, b []byte) {
		p1, err := ParseDerivationPath(path)
		if err != nil {
			return
		}
		p1.UnmarshalJSON(b)
	})
}

func Fuzz_Nosy_Manager_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Accounts()
	})
}

// skipping Fuzz_Nosy_Manager_AddBackend__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Backend

// skipping Fuzz_Nosy_Manager_Backends__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_Manager_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Close()
	})
}

func Fuzz_Nosy_Manager_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Config()
	})
}

func Fuzz_Nosy_Manager_Find__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		var account Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Find(account)
	})
}

// skipping Fuzz_Nosy_Manager_Subscribe__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/accounts.WalletEvent

func Fuzz_Nosy_Manager_Wallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Wallet(url)
	})
}

func Fuzz_Nosy_Manager_Wallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.Wallets()
	})
}

func Fuzz_Nosy_Manager_update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.update()
	})
}

func Fuzz_Nosy_Manager_walletsNoLock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am *Manager
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		if am == nil {
			return
		}

		am.walletsNoLock()
	})
}

func Fuzz_Nosy_URL_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, input []byte) {
		u, err := parseURL(url)
		if err != nil {
			return
		}
		u.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_AccountsByURL_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AccountsByURL
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.Len()
	})
}

func Fuzz_Nosy_AccountsByURL_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AccountsByURL
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		a.Less(i, j)
	})
}

func Fuzz_Nosy_AccountsByURL_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AccountsByURL
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		a.Swap(i, j)
	})
}

// skipping Fuzz_Nosy_Backend_Subscribe__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Backend

// skipping Fuzz_Nosy_Backend_Wallets__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Backend

func Fuzz_Nosy_DerivationPath_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		p1, err := ParseDerivationPath(path)
		if err != nil {
			return
		}
		p1.MarshalJSON()
	})
}

func Fuzz_Nosy_DerivationPath_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		p1, err := ParseDerivationPath(path)
		if err != nil {
			return
		}
		p1.String()
	})
}

func Fuzz_Nosy_URL_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u1 string
		fill_err = tp.Fill(&u1)
		if fill_err != nil {
			return
		}
		var u2 URL
		fill_err = tp.Fill(&u2)
		if fill_err != nil {
			return
		}

		u, err := parseURL(u1)
		if err != nil {
			return
		}
		u.Cmp(u2)
	})
}

func Fuzz_Nosy_URL_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		u, err := parseURL(url)
		if err != nil {
			return
		}
		u.MarshalJSON()
	})
}

func Fuzz_Nosy_URL_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		u, err := parseURL(url)
		if err != nil {
			return
		}
		u.String()
	})
}

func Fuzz_Nosy_URL_TerminalString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		u, err := parseURL(url)
		if err != nil {
			return
		}
		u.TerminalString()
	})
}

// skipping Fuzz_Nosy_Wallet_Accounts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_Contains__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_Derive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_Open__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SelfDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignData__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignDataWithPassphrase__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignText__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignTextWithPassphrase__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_SignTxWithPassphrase__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_Status__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_Wallet_URL__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_WalletsByURL_Len__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.WalletsByURL

// skipping Fuzz_Nosy_WalletsByURL_Less__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.WalletsByURL

// skipping Fuzz_Nosy_WalletsByURL_Swap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts.WalletsByURL

func Fuzz_Nosy_DefaultIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var base DerivationPath
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}

		DefaultIterator(base)
	})
}

func Fuzz_Nosy_LedgerLiveIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var base DerivationPath
		fill_err = tp.Fill(&base)
		if fill_err != nil {
			return
		}

		LedgerLiveIterator(base)
	})
}

func Fuzz_Nosy_TextAndHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		TextAndHash(d1)
	})
}

func Fuzz_Nosy_TextHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		TextHash(d1)
	})
}

// skipping Fuzz_Nosy_drop__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/accounts.Wallet

// skipping Fuzz_Nosy_merge__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/accounts.Wallet
