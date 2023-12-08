package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/node"
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

func Fuzz_Nosy_preimageIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *preimageIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Next()
	})
}

func Fuzz_Nosy_preimageIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *preimageIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Release()
	})
}

func Fuzz_Nosy_snapshotIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *snapshotIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Next()
	})
}

func Fuzz_Nosy_snapshotIterator_Release__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var iter *snapshotIterator
		fill_err = tp.Fill(&iter)
		if fill_err != nil {
			return
		}
		if iter == nil {
			return
		}

		iter.Release()
	})
}

func Fuzz_Nosy_applyMetricConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var cfg *gethConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		applyMetricConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_confirmAndRemoveDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, database string, kind string) {
		confirmAndRemoveDB(database, kind)
	})
}

func Fuzz_Nosy_deprecated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, field string) {
		deprecated(field)
	})
}

func Fuzz_Nosy_fetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		fetch(url)
	})
}

func Fuzz_Nosy_hashish__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x string) {
		hashish(x)
	})
}

func Fuzz_Nosy_keyID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id [8]byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		keyID(id)
	})
}

func Fuzz_Nosy_makeConfigNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		makeConfigNode(ctx)
	})
}

func Fuzz_Nosy_makeFullNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		makeFullNode(ctx)
	})
}

func Fuzz_Nosy_parseDumpConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if ctx == nil || stack == nil {
			return
		}

		parseDumpConfig(ctx, stack)
	})
}

func Fuzz_Nosy_prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		prepare(ctx)
	})
}

// skipping Fuzz_Nosy_showLeveldbStats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueStater

// skipping Fuzz_Nosy_startNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

func Fuzz_Nosy_unlockAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks *keystore.KeyStore
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var passwords []string
		fill_err = tp.Fill(&passwords)
		if fill_err != nil {
			return
		}
		if ks == nil {
			return
		}

		unlockAccount(ks, address, i, passwords)
	})
}

func Fuzz_Nosy_unlockAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		if ctx == nil || stack == nil {
			return
		}

		unlockAccounts(ctx, stack)
	})
}
