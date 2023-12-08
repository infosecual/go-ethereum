package scwallet

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	pcsc "github.com/gballet/go-libpcsclite"
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

func Fuzz_Nosy_Hub_Wallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, daemonPath string, scheme string, datadir string) {
		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.Wallets()
	})
}

func Fuzz_Nosy_Hub_pairing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var daemonPath string
		fill_err = tp.Fill(&daemonPath)
		if fill_err != nil {
			return
		}
		var scheme string
		fill_err = tp.Fill(&scheme)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var wallet *Wallet
		fill_err = tp.Fill(&wallet)
		if fill_err != nil {
			return
		}
		if wallet == nil {
			return
		}

		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.pairing(wallet)
	})
}

func Fuzz_Nosy_Hub_readPairings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, daemonPath string, scheme string, datadir string) {
		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.readPairings()
	})
}

func Fuzz_Nosy_Hub_refreshWallets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, daemonPath string, scheme string, datadir string) {
		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.refreshWallets()
	})
}

func Fuzz_Nosy_Hub_setPairing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var daemonPath string
		fill_err = tp.Fill(&daemonPath)
		if fill_err != nil {
			return
		}
		var scheme string
		fill_err = tp.Fill(&scheme)
		if fill_err != nil {
			return
		}
		var datadir string
		fill_err = tp.Fill(&datadir)
		if fill_err != nil {
			return
		}
		var wallet *Wallet
		fill_err = tp.Fill(&wallet)
		if fill_err != nil {
			return
		}
		var pairing *smartcardPairing
		fill_err = tp.Fill(&pairing)
		if fill_err != nil {
			return
		}
		if wallet == nil || pairing == nil {
			return
		}

		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.setPairing(wallet, pairing)
	})
}

func Fuzz_Nosy_Hub_updater__(f *testing.F) {
	f.Fuzz(func(t *testing.T, daemonPath string, scheme string, datadir string) {
		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.updater()
	})
}

func Fuzz_Nosy_Hub_writePairings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, daemonPath string, scheme string, datadir string) {
		hub, err := NewHub(daemonPath, scheme, datadir)
		if err != nil {
			return
		}
		hub.writePairings()
	})
}

func Fuzz_Nosy_SecureChannelSession_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.Open()
	})
}

func Fuzz_Nosy_SecureChannelSession_Pair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var pairingPassword []byte
		fill_err = tp.Fill(&pairingPassword)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.Pair(pairingPassword)
	})
}

func Fuzz_Nosy_SecureChannelSession_Unpair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.Unpair()
	})
}

func Fuzz_Nosy_SecureChannelSession_decryptAPDU__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.decryptAPDU(d3)
	})
}

func Fuzz_Nosy_SecureChannelSession_encryptAPDU__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.encryptAPDU(d3)
	})
}

func Fuzz_Nosy_SecureChannelSession_mutuallyAuthenticate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.mutuallyAuthenticate()
	})
}

func Fuzz_Nosy_SecureChannelSession_open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.open()
	})
}

func Fuzz_Nosy_SecureChannelSession_pair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var p1 uint8
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.pair(p1, d4)
	})
}

func Fuzz_Nosy_SecureChannelSession_transmitEncrypted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var cla byte
		fill_err = tp.Fill(&cla)
		if fill_err != nil {
			return
		}
		var ins byte
		fill_err = tp.Fill(&ins)
		if fill_err != nil {
			return
		}
		var p1 byte
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 byte
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var d7 []byte
		fill_err = tp.Fill(&d7)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.transmitEncrypted(cla, ins, p1, p2, d7)
	})
}

func Fuzz_Nosy_SecureChannelSession_updateIV__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var keyData []byte
		fill_err = tp.Fill(&keyData)
		if fill_err != nil {
			return
		}
		var meta []byte
		fill_err = tp.Fill(&meta)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if card == nil {
			return
		}

		s, err := NewSecureChannelSession(card, keyData)
		if err != nil {
			return
		}
		s.updateIV(meta, d4)
	})
}

func Fuzz_Nosy_Session_authenticate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pairing smartcardPairing
		fill_err = tp.Fill(&pairing)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.authenticate(pairing)
	})
}

func Fuzz_Nosy_Session_derivationPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.derivationPath()
	})
}

func Fuzz_Nosy_Session_derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.derive(path)
	})
}

func Fuzz_Nosy_Session_initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var seed []byte
		fill_err = tp.Fill(&seed)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.initialize(seed)
	})
}

func Fuzz_Nosy_Session_pair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var secret []byte
		fill_err = tp.Fill(&secret)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pair(secret)
	})
}

func Fuzz_Nosy_Session_paired__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.paired()
	})
}

func Fuzz_Nosy_Session_publicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.publicKey()
	})
}

func Fuzz_Nosy_Session_release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.release()
	})
}

func Fuzz_Nosy_Session_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.sign(path, hash)
	})
}

func Fuzz_Nosy_Session_unblockPin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pukpin []byte
		fill_err = tp.Fill(&pukpin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.unblockPin(pukpin)
	})
}

func Fuzz_Nosy_Session_unpair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.unpair()
	})
}

func Fuzz_Nosy_Session_verifyPin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var pin []byte
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.verifyPin(pin)
	})
}

func Fuzz_Nosy_Session_walletStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.walletStatus()
	})
}

func Fuzz_Nosy_Wallet_Accounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Accounts()
	})
}

func Fuzz_Nosy_Wallet_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Close()
	})
}

func Fuzz_Nosy_Wallet_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Contains(account)
	})
}

func Fuzz_Nosy_Wallet_Derive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Derive(path, pin)
	})
}

func Fuzz_Nosy_Wallet_Initialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var seed []byte
		fill_err = tp.Fill(&seed)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Initialize(seed)
	})
}

func Fuzz_Nosy_Wallet_Open__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Open(passphrase)
	})
}

// skipping Fuzz_Nosy_Wallet_SelfDerive__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum.ChainStateReader

func Fuzz_Nosy_Wallet_SignData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		var d5 []byte
		fill_err = tp.Fill(&d5)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignData(account, mimeType, d5)
	})
}

func Fuzz_Nosy_Wallet_SignDataWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		var d6 []byte
		fill_err = tp.Fill(&d6)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignDataWithPassphrase(account, passphrase, mimeType, d6)
	})
}

func Fuzz_Nosy_Wallet_SignText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignText(account, text)
	})
}

func Fuzz_Nosy_Wallet_SignTextWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignTextWithPassphrase(account, passphrase, text)
	})
}

func Fuzz_Nosy_Wallet_SignTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil || tx == nil || chainID == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignTx(account, tx, chainID)
	})
}

func Fuzz_Nosy_Wallet_SignTxWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil || tx == nil || chainID == nil {
			return
		}

		w := NewWallet(hub, card)
		w.SignTxWithPassphrase(account, passphrase, tx, chainID)
	})
}

func Fuzz_Nosy_Wallet_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Status()
	})
}

func Fuzz_Nosy_Wallet_URL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.URL()
	})
}

func Fuzz_Nosy_Wallet_Unpair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var pin []byte
		fill_err = tp.Fill(&pin)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.Unpair(pin)
	})
}

func Fuzz_Nosy_Wallet_connect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.connect()
	})
}

func Fuzz_Nosy_Wallet_doselect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.doselect()
	})
}

func Fuzz_Nosy_Wallet_findAccountPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var account accounts.Account
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.findAccountPath(account)
	})
}

func Fuzz_Nosy_Wallet_makeAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var path accounts.DerivationPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.makeAccount(address, path)
	})
}

func Fuzz_Nosy_Wallet_pair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var puk []byte
		fill_err = tp.Fill(&puk)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.pair(puk)
	})
}

func Fuzz_Nosy_Wallet_ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.ping()
	})
}

func Fuzz_Nosy_Wallet_release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.release()
	})
}

func Fuzz_Nosy_Wallet_selfDerive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.selfDerive()
	})
}

func Fuzz_Nosy_Wallet_signHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.signHash(account, hash)
	})
}

func Fuzz_Nosy_Wallet_signHashWithPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hub *Hub
		fill_err = tp.Fill(&hub)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
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
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hub == nil || card == nil {
			return
		}

		w := NewWallet(hub, card)
		w.signHashWithPassphrase(account, passphrase, hash)
	})
}

func Fuzz_Nosy_responseAPDU_deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var card *pcsc.Card
		fill_err = tp.Fill(&card)
		if fill_err != nil {
			return
		}
		var command *commandAPDU
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if card == nil || command == nil {
			return
		}

		ra, err := transmit(card, command)
		if err != nil {
			return
		}
		ra.deserialize(d3)
	})
}

func Fuzz_Nosy_commandAPDU_serialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ca commandAPDU
		fill_err = tp.Fill(&ca)
		if fill_err != nil {
			return
		}

		ca.serialize()
	})
}

func Fuzz_Nosy_makeRecoverableSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, sig []byte, expectedPubkey []byte) {
		makeRecoverableSignature(hash, sig, expectedPubkey)
	})
}

func Fuzz_Nosy_pad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, terminator byte) {
		pad(d1, terminator)
	})
}

func Fuzz_Nosy_unpad__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, terminator byte) {
		unpad(d1, terminator)
	})
}
