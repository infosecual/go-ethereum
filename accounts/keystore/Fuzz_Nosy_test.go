package keystore

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_AmbiguousAddrError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *AmbiguousAddrError
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

func Fuzz_Nosy_Key_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyjson []byte, auth string) {
		k, err := DecryptKey(keyjson, auth)
		if err != nil {
			return
		}
		k.MarshalJSON()
	})
}

func Fuzz_Nosy_Key_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyjson []byte, auth string, j []byte) {
		k, err := DecryptKey(keyjson, auth)
		if err != nil {
			return
		}
		k.UnmarshalJSON(j)
	})
}

func Fuzz_Nosy_KeyStore_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Accounts()
	})
}

func Fuzz_Nosy_KeyStore_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Delete(a, passphrase)
	})
}

func Fuzz_Nosy_KeyStore_Export__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var newPassphrase string
		fill_err = tp.Fill(&newPassphrase)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Export(a, passphrase, newPassphrase)
	})
}

func Fuzz_Nosy_KeyStore_Find__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Find(a)
	})
}

func Fuzz_Nosy_KeyStore_HasAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.HasAddress(addr)
	})
}

func Fuzz_Nosy_KeyStore_Import__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int, keyJSON []byte, passphrase string, newPassphrase string) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Import(keyJSON, passphrase, newPassphrase)
	})
}

func Fuzz_Nosy_KeyStore_ImportECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var priv *ecdsa.PrivateKey
		fill_err = tp.Fill(&priv)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if priv == nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.ImportECDSA(priv, passphrase)
	})
}

func Fuzz_Nosy_KeyStore_ImportPreSaleKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int, keyJSON []byte, passphrase string) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.ImportPreSaleKey(keyJSON, passphrase)
	})
}

func Fuzz_Nosy_KeyStore_Lock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Lock(addr)
	})
}

func Fuzz_Nosy_KeyStore_NewAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int, passphrase string) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.NewAccount(passphrase)
	})
}

func Fuzz_Nosy_KeyStore_SignHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.SignHash(a, hash)
	})
}

func Fuzz_Nosy_KeyStore_SignHashWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.SignHashWithPassphrase(a, passphrase, hash)
	})
}

func Fuzz_Nosy_KeyStore_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
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

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.SignTx(a, tx, chainID)
	})
}

func Fuzz_Nosy_KeyStore_SignTxWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
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

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.SignTxWithPassphrase(a, passphrase, tx, chainID)
	})
}

// skipping Fuzz_Nosy_KeyStore_Subscribe__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/accounts.WalletEvent

func Fuzz_Nosy_KeyStore_TimedUnlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.TimedUnlock(a, passphrase, timeout)
	})
}

func Fuzz_Nosy_KeyStore_Unlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Unlock(a, passphrase)
	})
}

func Fuzz_Nosy_KeyStore_Update__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var newPassphrase string
		fill_err = tp.Fill(&newPassphrase)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Update(a, passphrase, newPassphrase)
	})
}

func Fuzz_Nosy_KeyStore_Wallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.Wallets()
	})
}

func Fuzz_Nosy_KeyStore_expire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var u *unlocked
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.expire(addr, u, timeout)
	})
}

func Fuzz_Nosy_KeyStore_getDecryptedKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.getDecryptedKey(a, auth)
	})
}

func Fuzz_Nosy_KeyStore_importKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keydir string
		fill_err = tp.Fill(&keydir)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		var key *Key
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.importKey(key, passphrase)
	})
}

func Fuzz_Nosy_KeyStore_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, k1 string, scryptN int, scryptP int, k4 string) {
		ks := NewKeyStore(k1, scryptN, scryptP)
		ks.init(k4)
	})
}

func Fuzz_Nosy_KeyStore_isUpdating__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.isUpdating()
	})
}

func Fuzz_Nosy_KeyStore_refreshWallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.refreshWallets()
	})
}

func Fuzz_Nosy_KeyStore_updater__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string, scryptN int, scryptP int) {
		ks := NewKeyStore(keydir, scryptN, scryptP)
		ks.updater()
	})
}

func Fuzz_Nosy_accountCache_accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.accounts()
	})
}

func Fuzz_Nosy_accountCache_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		var newAccount accounts.Account
		fill_err = tp.Fill(&newAccount)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.add(newAccount)
	})
}

func Fuzz_Nosy_accountCache_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.close()
	})
}

func Fuzz_Nosy_accountCache_delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		var removed accounts.Account
		fill_err = tp.Fill(&removed)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.delete(removed)
	})
}

func Fuzz_Nosy_accountCache_deleteByFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.deleteByFile(path)
	})
}

func Fuzz_Nosy_accountCache_find__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.find(a)
	})
}

func Fuzz_Nosy_accountCache_hasAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.hasAddress(addr)
	})
}

func Fuzz_Nosy_accountCache_maybeReload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.maybeReload()
	})
}

func Fuzz_Nosy_accountCache_scanAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.scanAccounts()
	})
}

func Fuzz_Nosy_accountCache_watcherStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		ac.watcherStarted()
	})
}

func Fuzz_Nosy_fileCache_scan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fc *fileCache
		fill_err = tp.Fill(&fc)
		if fill_err != nil {
			return
		}
		var keyDir string
		fill_err = tp.Fill(&keyDir)
		if fill_err != nil {
			return
		}
		if fc == nil {
			return
		}

		fc.scan(keyDir)
	})
}

func Fuzz_Nosy_keystoreWallet_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

// skipping Fuzz_Nosy_keystoreWallet_SelfDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum.ChainStateReader

func Fuzz_Nosy_keystoreWallet_SignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_SignDataWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_SignText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_SignTextWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_SignTxWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_URL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_keystoreWallet_signHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *keystoreWallet
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

func Fuzz_Nosy_watcher_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		w := newWatcher(ac)
		w.close()
	})
}

func Fuzz_Nosy_watcher_enabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		_x1 := newWatcher(ac)
		_x1.enabled()
	})
}

func Fuzz_Nosy_watcher_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		w := newWatcher(ac)
		w.loop()
	})
}

func Fuzz_Nosy_watcher_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ac *accountCache
		fill_err = tp.Fill(&ac)
		if fill_err != nil {
			return
		}
		if ac == nil {
			return
		}

		w := newWatcher(ac)
		w.start()
	})
}

// skipping Fuzz_Nosy_keyStore_GetKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/keystore.keyStore

// skipping Fuzz_Nosy_keyStore_JoinPath__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/keystore.keyStore

// skipping Fuzz_Nosy_keyStore_StoreKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/keystore.keyStore

func Fuzz_Nosy_keyStorePassphrase_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePassphrase
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}

		ks.GetKey(addr, filename, auth)
	})
}

func Fuzz_Nosy_keyStorePassphrase_JoinPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePassphrase
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}

		ks.JoinPath(filename)
	})
}

func Fuzz_Nosy_keyStorePassphrase_StoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePassphrase
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var key *Key
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		ks.StoreKey(filename, key, auth)
	})
}

func Fuzz_Nosy_keyStorePlain_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePlain
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var addr common.Address
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}

		ks.GetKey(addr, filename, auth)
	})
}

func Fuzz_Nosy_keyStorePlain_JoinPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePlain
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}

		ks.JoinPath(filename)
	})
}

func Fuzz_Nosy_keyStorePlain_StoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks keyStorePlain
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var filename string
		fill_err = tp.Fill(&filename)
		if fill_err != nil {
			return
		}
		var key *Key
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		ks.StoreKey(filename, key, auth)
	})
}

func Fuzz_Nosy_DecryptDataV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cryptoJson CryptoJSON
		fill_err = tp.Fill(&cryptoJson)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}

		DecryptDataV3(cryptoJson, auth)
	})
}

func Fuzz_Nosy_EncryptKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key *Key
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		var scryptN int
		fill_err = tp.Fill(&scryptN)
		if fill_err != nil {
			return
		}
		var scryptP int
		fill_err = tp.Fill(&scryptP)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		EncryptKey(key, auth, scryptN, scryptP)
	})
}

func Fuzz_Nosy_aesCBCDecrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, cipherText []byte, iv []byte) {
		aesCBCDecrypt(key, cipherText, iv)
	})
}

func Fuzz_Nosy_aesCTRXOR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, inText []byte, iv []byte) {
		aesCTRXOR(key, inText, iv)
	})
}

func Fuzz_Nosy_byURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a accounts.Account
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b accounts.Account
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		byURL(a, b)
	})
}

func Fuzz_Nosy_decryptKeyV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyProtected *encryptedKeyJSONV1
		fill_err = tp.Fill(&keyProtected)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		if keyProtected == nil {
			return
		}

		decryptKeyV1(keyProtected, auth)
	})
}

func Fuzz_Nosy_decryptKeyV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyProtected *encryptedKeyJSONV3
		fill_err = tp.Fill(&keyProtected)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		if keyProtected == nil {
			return
		}

		decryptKeyV3(keyProtected, auth)
	})
}

// skipping Fuzz_Nosy_ensureInt__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_getKDFKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cryptoJSON CryptoJSON
		fill_err = tp.Fill(&cryptoJSON)
		if fill_err != nil {
			return
		}
		var auth string
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}

		getKDFKey(cryptoJSON, auth)
	})
}

// skipping Fuzz_Nosy_importPreSaleKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/keystore.keyStore

func Fuzz_Nosy_keyFileName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyAddr common.Address
		fill_err = tp.Fill(&keyAddr)
		if fill_err != nil {
			return
		}

		keyFileName(keyAddr)
	})
}

func Fuzz_Nosy_newAccountCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keydir string) {
		newAccountCache(keydir)
	})
}

// skipping Fuzz_Nosy_nonKeyFile__ because parameters include func, chan, or unsupported interface: io/fs.DirEntry

func Fuzz_Nosy_pkcs7Unpad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in []byte) {
		pkcs7Unpad(in)
	})
}

func Fuzz_Nosy_removeAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slice []accounts.Account
		fill_err = tp.Fill(&slice)
		if fill_err != nil {
			return
		}
		var elem accounts.Account
		fill_err = tp.Fill(&elem)
		if fill_err != nil {
			return
		}

		removeAccount(slice, elem)
	})
}

// skipping Fuzz_Nosy_storeNewKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/keystore.keyStore

func Fuzz_Nosy_toISO8601__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 time.Time
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		toISO8601(t1)
	})
}

func Fuzz_Nosy_writeTemporaryKeyFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, content []byte) {
		writeTemporaryKeyFile(file, content)
	})
}

func Fuzz_Nosy_zeroKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *ecdsa.PrivateKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		zeroKey(k)
	})
}
