package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Fuzz_Nosy_ChainConfig_AddFork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		var version []byte
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.AddFork(name, epoch, version)
	})
}

func Fuzz_Nosy_ChainConfig_LoadForks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.LoadForks(path)
	})
}

func Fuzz_Nosy_Fork_computeDomain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Hash
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.computeDomain(genesisValidatorsRoot)
	})
}

func Fuzz_Nosy_Header_Epoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Epoch()
	})
}

func Fuzz_Nosy_Header_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Hash()
	})
}

func Fuzz_Nosy_Header_SyncPeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.SyncPeriod()
	})
}

func Fuzz_Nosy_Header_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_LightClientUpdate_Score__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var update *LightClientUpdate
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		if update == nil {
			return
		}

		update.Score()
	})
}

func Fuzz_Nosy_LightClientUpdate_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var update *LightClientUpdate
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		if update == nil {
			return
		}

		update.Validate()
	})
}

func Fuzz_Nosy_SerializedSyncCommittee_Deserialize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SerializedSyncCommittee
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Deserialize()
	})
}

func Fuzz_Nosy_SerializedSyncCommittee_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SerializedSyncCommittee
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_SerializedSyncCommittee_Root__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SerializedSyncCommittee
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Root()
	})
}

func Fuzz_Nosy_SerializedSyncCommittee_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SerializedSyncCommittee
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_SyncAggregate_SignerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncAggregate
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SignerCount()
	})
}

func Fuzz_Nosy_SyncAggregate_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncAggregate
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_SyncCommittee_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SyncCommittee
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var signingRoot common.Hash
		fill_err = tp.Fill(&signingRoot)
		if fill_err != nil {
			return
		}
		var signature *SyncAggregate
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if sc == nil || signature == nil {
			return
		}

		sc.VerifySignature(signingRoot, signature)
	})
}

func Fuzz_Nosy_UpdateScore_finalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *UpdateScore
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.finalized()
	})
}

func Fuzz_Nosy_Forks_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 Forks
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.Len()
	})
}

func Fuzz_Nosy_Forks_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 Forks
		fill_err = tp.Fill(&f1)
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

		f1.Less(i, j)
	})
}

func Fuzz_Nosy_Forks_SigningRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 Forks
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var header Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}

		f1.SigningRoot(header)
	})
}

func Fuzz_Nosy_Forks_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 Forks
		fill_err = tp.Fill(&f1)
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

		f1.Swap(i, j)
	})
}

func Fuzz_Nosy_Forks_domain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 Forks
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		f1.domain(epoch)
	})
}

func Fuzz_Nosy_Header_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.MarshalJSON()
	})
}

func Fuzz_Nosy_SyncAggregate_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s SyncAggregate
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_UpdateScore_BetterThan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u UpdateScore
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var w UpdateScore
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		u.BetterThan(w)
	})
}

func Fuzz_Nosy_SyncPeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, slot uint64) {
		SyncPeriod(slot)
	})
}

func Fuzz_Nosy_SyncPeriodStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, period uint64) {
		SyncPeriodStart(period)
	})
}
