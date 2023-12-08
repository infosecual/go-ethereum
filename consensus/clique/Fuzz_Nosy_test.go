package clique

import (
	"io"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	lru "github.com/ethereum/go-ethereum/common/lru"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
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

func Fuzz_Nosy_API_Discard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
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

		api.Discard(address)
	})
}

func Fuzz_Nosy_API_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var rlpOrBlockNr *blockNumberOrHashOrRLP
		fill_err = tp.Fill(&rlpOrBlockNr)
		if fill_err != nil {
			return
		}
		if api == nil || rlpOrBlockNr == nil {
			return
		}

		api.GetSigner(rlpOrBlockNr)
	})
}

func Fuzz_Nosy_API_GetSigners__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var number *rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil || number == nil {
			return
		}

		api.GetSigners(number)
	})
}

func Fuzz_Nosy_API_GetSignersAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.GetSignersAtHash(hash)
	})
}

func Fuzz_Nosy_API_GetSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var number *rpc.BlockNumber
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if api == nil || number == nil {
			return
		}

		api.GetSnapshot(number)
	})
}

func Fuzz_Nosy_API_GetSnapshotAtHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.GetSnapshotAtHash(hash)
	})
}

func Fuzz_Nosy_API_Proposals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Proposals()
	})
}

func Fuzz_Nosy_API_Propose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var auth bool
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Propose(address, auth)
	})
}

func Fuzz_Nosy_API_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *API
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Status()
	})
}

// skipping Fuzz_Nosy_Clique_APIs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Clique_Author__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Clique
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if c == nil || header == nil {
			return
		}

		c.Author(header)
	})
}

// skipping Fuzz_Nosy_Clique_Authorize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus/clique.SignerFn

// skipping Fuzz_Nosy_Clique_CalcDifficulty__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Clique_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Clique
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

// skipping Fuzz_Nosy_Clique_Finalize__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_FinalizeAndAssemble__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_Prepare__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_Seal__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Clique_SealHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Clique
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if c == nil || header == nil {
			return
		}

		c.SealHash(header)
	})
}

// skipping Fuzz_Nosy_Clique_VerifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_VerifyHeaders__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_VerifyUncles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainReader

// skipping Fuzz_Nosy_Clique_snapshot__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_verifyCascadingFields__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

// skipping Fuzz_Nosy_Clique_verifyHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/consensus.ChainHeaderReader

func Fuzz_Nosy_Clique_verifySeal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Clique
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var snap *Snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var parents []*types.Header
		fill_err = tp.Fill(&parents)
		if fill_err != nil {
			return
		}
		if c == nil || snap == nil || header == nil {
			return
		}

		c.verifySeal(snap, header, parents)
	})
}

func Fuzz_Nosy_Snapshot_apply__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.apply(headers)
	})
}

func Fuzz_Nosy_Snapshot_cast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var authorize bool
		fill_err = tp.Fill(&authorize)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.cast(address, authorize)
	})
}

func Fuzz_Nosy_Snapshot_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.copy()
	})
}

func Fuzz_Nosy_Snapshot_inturn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var n3 uint64
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		var n6 uint64
		fill_err = tp.Fill(&n6)
		if fill_err != nil {
			return
		}
		var signer common.Address
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, n3, hash, signers)
		s.inturn(n6, signer)
	})
}

func Fuzz_Nosy_Snapshot_signers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.signers()
	})
}

// skipping Fuzz_Nosy_Snapshot_store__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_Snapshot_uncast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var authorize bool
		fill_err = tp.Fill(&authorize)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.uncast(address, authorize)
	})
}

func Fuzz_Nosy_Snapshot_validVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *params.CliqueConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var sigcache *lru.Cache[common.Hash, common.Address]
		fill_err = tp.Fill(&sigcache)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var signers []common.Address
		fill_err = tp.Fill(&signers)
		if fill_err != nil {
			return
		}
		var address common.Address
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var authorize bool
		fill_err = tp.Fill(&authorize)
		if fill_err != nil {
			return
		}
		if config == nil || sigcache == nil {
			return
		}

		s := newSnapshot(config, sigcache, number, hash, signers)
		s.validVote(address, authorize)
	})
}

func Fuzz_Nosy_blockNumberOrHashOrRLP_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sb *blockNumberOrHashOrRLP
		fill_err = tp.Fill(&sb)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if sb == nil {
			return
		}

		sb.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_CliqueRLP__(f *testing.F) {
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

		CliqueRLP(header)
	})
}

func Fuzz_Nosy_encodeSigHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
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

		encodeSigHeader(w, header)
	})
}
