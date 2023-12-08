package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/signer/core"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_confirm__(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		confirm(text)
	})
}

func Fuzz_Nosy_decryptSeed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyjson []byte, auth string) {
		decryptSeed(keyjson, auth)
	})
}

func Fuzz_Nosy_encryptSeed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed []byte, auth []byte, scryptN int, scryptP int) {
		encryptSeed(seed, auth, scryptN, scryptP)
	})
}

func Fuzz_Nosy_initInternalApi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *cli.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		initInternalApi(c)
	})
}

func Fuzz_Nosy_ipcEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ipcPath string, datadir string) {
		ipcEndpoint(ipcPath, datadir)
	})
}

// skipping Fuzz_Nosy_readMasterKey__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/signer/core.UIClientAPI

func Fuzz_Nosy_testExternalUI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *core.SignerAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		testExternalUI(api)
	})
}
