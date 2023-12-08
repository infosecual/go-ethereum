package params

import (
	"math/big"
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

func Fuzz_Nosy_ChainConfig_BaseFeeChangeDenominator__(f *testing.F) {
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
		if c == nil {
			return
		}

		c.BaseFeeChangeDenominator()
	})
}

func Fuzz_Nosy_ChainConfig_CheckCompatible__(f *testing.F) {
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
		var newcfg *ChainConfig
		fill_err = tp.Fill(&newcfg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil || newcfg == nil {
			return
		}

		c.CheckCompatible(newcfg, height, time)
	})
}

func Fuzz_Nosy_ChainConfig_CheckConfigForkOrder__(f *testing.F) {
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
		if c == nil {
			return
		}

		c.CheckConfigForkOrder()
	})
}

func Fuzz_Nosy_ChainConfig_Description__(f *testing.F) {
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
		if c == nil {
			return
		}

		c.Description()
	})
}

func Fuzz_Nosy_ChainConfig_ElasticityMultiplier__(f *testing.F) {
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
		if c == nil {
			return
		}

		c.ElasticityMultiplier()
	})
}

func Fuzz_Nosy_ChainConfig_IsArrowGlacier__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsArrowGlacier(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsBerlin__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsBerlin(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsByzantium__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsByzantium(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsCancun__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsCancun(num, time)
	})
}

func Fuzz_Nosy_ChainConfig_IsConstantinople__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsConstantinople(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsDAOFork__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsDAOFork(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsEIP150__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsEIP150(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsEIP155__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsEIP155(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsEIP158__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsEIP158(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsGrayGlacier__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsGrayGlacier(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsHomestead__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsHomestead(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsIstanbul__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsIstanbul(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsLondon__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsLondon(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsMuirGlacier__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsMuirGlacier(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsPetersburg__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsPetersburg(num)
	})
}

func Fuzz_Nosy_ChainConfig_IsPrague__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsPrague(num, time)
	})
}

func Fuzz_Nosy_ChainConfig_IsShanghai__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsShanghai(num, time)
	})
}

func Fuzz_Nosy_ChainConfig_IsTerminalPoWBlock__(f *testing.F) {
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
		var parentTotalDiff *big.Int
		fill_err = tp.Fill(&parentTotalDiff)
		if fill_err != nil {
			return
		}
		var totalDiff *big.Int
		fill_err = tp.Fill(&totalDiff)
		if fill_err != nil {
			return
		}
		if c == nil || parentTotalDiff == nil || totalDiff == nil {
			return
		}

		c.IsTerminalPoWBlock(parentTotalDiff, totalDiff)
	})
}

func Fuzz_Nosy_ChainConfig_IsVerkle__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var time uint64
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.IsVerkle(num, time)
	})
}

func Fuzz_Nosy_ChainConfig_Rules__(f *testing.F) {
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
		var num *big.Int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var isMerge bool
		fill_err = tp.Fill(&isMerge)
		if fill_err != nil {
			return
		}
		var timestamp uint64
		fill_err = tp.Fill(&timestamp)
		if fill_err != nil {
			return
		}
		if c == nil || num == nil {
			return
		}

		c.Rules(num, isMerge, timestamp)
	})
}

func Fuzz_Nosy_ChainConfig_checkCompatible__(f *testing.F) {
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
		var newcfg *ChainConfig
		fill_err = tp.Fill(&newcfg)
		if fill_err != nil {
			return
		}
		var headNumber *big.Int
		fill_err = tp.Fill(&headNumber)
		if fill_err != nil {
			return
		}
		var headTimestamp uint64
		fill_err = tp.Fill(&headTimestamp)
		if fill_err != nil {
			return
		}
		if c == nil || newcfg == nil || headNumber == nil {
			return
		}

		c.checkCompatible(newcfg, headNumber, headTimestamp)
	})
}

func Fuzz_Nosy_CliqueConfig_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CliqueConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_ConfigCompatError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var what string
		fill_err = tp.Fill(&what)
		if fill_err != nil {
			return
		}
		var storedblock *big.Int
		fill_err = tp.Fill(&storedblock)
		if fill_err != nil {
			return
		}
		var newblock *big.Int
		fill_err = tp.Fill(&newblock)
		if fill_err != nil {
			return
		}
		if storedblock == nil || newblock == nil {
			return
		}

		err := newBlockCompatError(what, storedblock, newblock)
		err.Error()
	})
}

func Fuzz_Nosy_EthashConfig_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EthashConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_ArchiveVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, gitCommit string) {
		ArchiveVersion(gitCommit)
	})
}

func Fuzz_Nosy_KnownDNSNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var genesis common.Hash
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}

		KnownDNSNetwork(genesis, protocol)
	})
}

func Fuzz_Nosy_VersionWithCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, gitCommit string, gitDate string) {
		VersionWithCommit(gitCommit, gitDate)
	})
}

func Fuzz_Nosy_configBlockEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *big.Int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *big.Int
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		configBlockEqual(x, y)
	})
}

func Fuzz_Nosy_configTimestampEqual__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x *uint64
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var y *uint64
		fill_err = tp.Fill(&y)
		if fill_err != nil {
			return
		}
		if x == nil || y == nil {
			return
		}

		configTimestampEqual(x, y)
	})
}

func Fuzz_Nosy_isBlockForked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *big.Int
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head *big.Int
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s == nil || head == nil {
			return
		}

		isBlockForked(s, head)
	})
}

func Fuzz_Nosy_isForkBlockIncompatible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 *big.Int
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var s2 *big.Int
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		var head *big.Int
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s1 == nil || s2 == nil || head == nil {
			return
		}

		isForkBlockIncompatible(s1, s2, head)
	})
}

func Fuzz_Nosy_isForkTimestampIncompatible__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 *uint64
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var s2 *uint64
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s1 == nil || s2 == nil {
			return
		}

		isForkTimestampIncompatible(s1, s2, head)
	})
}

func Fuzz_Nosy_isTimestampForked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *uint64
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		isTimestampForked(s, head)
	})
}

func Fuzz_Nosy_newUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, val uint64) {
		newUint64(val)
	})
}
