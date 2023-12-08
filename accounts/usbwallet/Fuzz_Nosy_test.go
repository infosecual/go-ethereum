package usbwallet

import (
	"io"
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

// skipping Fuzz_Nosy_Hub_Subscribe__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/accounts.WalletEvent

func Fuzz_Nosy_ledgerDriver_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Close()
	})
}

func Fuzz_Nosy_ledgerDriver_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Derive(path)
	})
}

func Fuzz_Nosy_ledgerDriver_Heartbeat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Heartbeat()
	})
}

func Fuzz_Nosy_ledgerDriver_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var device io.ReadWriter
		fill_err = tp.Fill(&device)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Open(device, passphrase)
	})
}

func Fuzz_Nosy_ledgerDriver_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.SignTx(path, tx, chainID)
	})
}

func Fuzz_Nosy_ledgerDriver_SignTypedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var domainHash []byte
		fill_err = tp.Fill(&domainHash)
		if fill_err != nil {
			return
		}
		var messageHash []byte
		fill_err = tp.Fill(&messageHash)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.SignTypedMessage(path, domainHash, messageHash)
	})
}

func Fuzz_Nosy_ledgerDriver_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Status()
	})
}

func Fuzz_Nosy_ledgerDriver_ledgerDerive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var derivationPath []uint32
		fill_err = tp.Fill(&derivationPath)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.ledgerDerive(derivationPath)
	})
}

func Fuzz_Nosy_ledgerDriver_ledgerExchange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var opcode ledgerOpcode
		fill_err = tp.Fill(&opcode)
		if fill_err != nil {
			return
		}
		var p1 ledgerParam1
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 ledgerParam2
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.ledgerExchange(opcode, p1, p2, d5)
	})
}

func Fuzz_Nosy_ledgerDriver_ledgerSign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var derivationPath []uint32
		fill_err = tp.Fill(&derivationPath)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.ledgerSign(derivationPath, tx, chainID)
	})
}

func Fuzz_Nosy_ledgerDriver_ledgerSignTypedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var derivationPath []uint32
		fill_err = tp.Fill(&derivationPath)
		if fill_err != nil {
			return
		}
		var domainHash []byte
		fill_err = tp.Fill(&domainHash)
		if fill_err != nil {
			return
		}
		var messageHash []byte
		fill_err = tp.Fill(&messageHash)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.ledgerSignTypedMessage(derivationPath, domainHash, messageHash)
	})
}

func Fuzz_Nosy_ledgerDriver_ledgerVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.ledgerVersion()
	})
}

func Fuzz_Nosy_ledgerDriver_offline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *ledgerDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.offline()
	})
}

func Fuzz_Nosy_trezorDriver_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Close()
	})
}

func Fuzz_Nosy_trezorDriver_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Derive(path)
	})
}

func Fuzz_Nosy_trezorDriver_Heartbeat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Heartbeat()
	})
}

func Fuzz_Nosy_trezorDriver_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var device io.ReadWriter
		fill_err = tp.Fill(&device)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Open(device, passphrase)
	})
}

func Fuzz_Nosy_trezorDriver_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.SignTx(path, tx, chainID)
	})
}

func Fuzz_Nosy_trezorDriver_SignTypedMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var domainHash []byte
		fill_err = tp.Fill(&domainHash)
		if fill_err != nil {
			return
		}
		var messageHash []byte
		fill_err = tp.Fill(&messageHash)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.SignTypedMessage(path, domainHash, messageHash)
	})
}

func Fuzz_Nosy_trezorDriver_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Status()
	})
}

func Fuzz_Nosy_trezorDriver_trezorDerive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var derivationPath []uint32
		fill_err = tp.Fill(&derivationPath)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.trezorDerive(derivationPath)
	})
}

// skipping Fuzz_Nosy_trezorDriver_trezorExchange__ because parameters include func, chan, or unsupported interface: google.golang.org/protobuf/runtime/protoiface.MessageV1

func Fuzz_Nosy_trezorDriver_trezorSign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *trezorDriver
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var derivationPath []uint32
		fill_err = tp.Fill(&derivationPath)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.trezorSign(derivationPath, tx, chainID)
	})
}

func Fuzz_Nosy_wallet_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Accounts()
	})
}

func Fuzz_Nosy_wallet_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Close()
	})
}

func Fuzz_Nosy_wallet_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Contains(account)
	})
}

func Fuzz_Nosy_wallet_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.Derive(path, pin)
	})
}

func Fuzz_Nosy_wallet_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Open(passphrase)
	})
}

// skipping Fuzz_Nosy_wallet_SelfDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum.ChainStateReader

func Fuzz_Nosy_wallet_SignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.SignData(account, mimeType, d4)
	})
}

func Fuzz_Nosy_wallet_SignDataWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.SignDataWithPassphrase(account, passphrase, mimeType, d5)
	})
}

func Fuzz_Nosy_wallet_SignText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.SignText(account, text)
	})
}

func Fuzz_Nosy_wallet_SignTextWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil {
			return
		}

		w.SignTextWithPassphrase(account, passphrase, text)
	})
}

func Fuzz_Nosy_wallet_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.SignTx(account, tx, chainID)
	})
}

func Fuzz_Nosy_wallet_SignTxWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
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
		if w == nil || tx == nil || chainID == nil {
			return
		}

		w.SignTxWithPassphrase(account, passphrase, tx, chainID)
	})
}

func Fuzz_Nosy_wallet_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Status()
	})
}

func Fuzz_Nosy_wallet_URL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.URL()
	})
}

func Fuzz_Nosy_wallet_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.close()
	})
}

func Fuzz_Nosy_wallet_heartbeat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.heartbeat()
	})
}

func Fuzz_Nosy_wallet_selfDerive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.selfDerive()
	})
}

func Fuzz_Nosy_wallet_signHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wallet
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.signHash(account, hash)
	})
}

// skipping Fuzz_Nosy_driver_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_Derive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_Heartbeat__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_Open__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_SignTx__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_SignTypedMessage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver

// skipping Fuzz_Nosy_driver_Status__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/usbwallet.driver
