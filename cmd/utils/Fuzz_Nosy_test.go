package utils

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool/legacypool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
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

// skipping Fuzz_Nosy_ChainDataIterator_Next__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/utils.ChainDataIterator

// skipping Fuzz_Nosy_ChainDataIterator_Release__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/utils.ChainDataIterator

// skipping Fuzz_Nosy_CheckExclusive__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Fatalf__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_GetPassPhrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string, confirmation bool) {
		GetPassPhrase(text, confirmation)
	})
}

func Fuzz_Nosy_GetPassPhraseWithList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var text string
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		var confirmation bool
		fill_err = tp.Fill(&confirmation)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var passwords []string
		fill_err = tp.Fill(&passwords)
		if fill_err != nil {
			return
		}

		GetPassPhraseWithList(text, confirmation, index, passwords)
	})
}

func Fuzz_Nosy_IsNetworkPreset__(f *testing.F) {
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

		IsNetworkPreset(ctx)
	})
}

func Fuzz_Nosy_MakeChain__(f *testing.F) {
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
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		if ctx == nil || stack == nil {
			return
		}

		MakeChain(ctx, stack, readonly)
	})
}

func Fuzz_Nosy_MakeConsolePreloads__(f *testing.F) {
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

		MakeConsolePreloads(ctx)
	})
}

func Fuzz_Nosy_MakeDataDir__(f *testing.F) {
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

		MakeDataDir(ctx)
	})
}

func Fuzz_Nosy_MakeDatabaseHandles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, max int) {
		MakeDatabaseHandles(max)
	})
}

func Fuzz_Nosy_MakePasswordList__(f *testing.F) {
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

		MakePasswordList(ctx)
	})
}

func Fuzz_Nosy_RegisterEthService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if stack == nil || cfg == nil {
			return
		}

		RegisterEthService(stack, cfg)
	})
}

// skipping Fuzz_Nosy_RegisterEthStatsService__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

func Fuzz_Nosy_RegisterFullSyncTester__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stack *node.Node
		fill_err = tp.Fill(&stack)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var t3 common.Hash
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if stack == nil || eth == nil {
			return
		}

		RegisterFullSyncTester(stack, eth, t3)
	})
}

// skipping Fuzz_Nosy_RegisterGraphQLService__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/internal/ethapi.Backend

func Fuzz_Nosy_SetDNSDiscoveryDefaults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var genesis common.Hash
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		SetDNSDiscoveryDefaults(cfg, genesis)
	})
}

func Fuzz_Nosy_SetDataDir__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		SetDataDir(ctx, cfg)
	})
}

func Fuzz_Nosy_SetEthConfig__(f *testing.F) {
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
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || stack == nil || cfg == nil {
			return
		}

		SetEthConfig(ctx, stack, cfg)
	})
}

func Fuzz_Nosy_SetNodeConfig__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		SetNodeConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_SetP2PConfig__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		SetP2PConfig(ctx, cfg)
	})
}

func Fuzz_Nosy_SetupMetrics__(f *testing.F) {
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

		SetupMetrics(ctx)
	})
}

func Fuzz_Nosy_SplitAndTrim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		SplitAndTrim(input)
	})
}

func Fuzz_Nosy_SplitTagsFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, tagsFlag string) {
		SplitTagsFlag(tagsFlag)
	})
}

func Fuzz_Nosy_StartNode__(f *testing.F) {
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
		var isConsole bool
		fill_err = tp.Fill(&isConsole)
		if fill_err != nil {
			return
		}
		if ctx == nil || stack == nil {
			return
		}

		StartNode(ctx, stack, isConsole)
	})
}

func Fuzz_Nosy_getFreeDiskSpace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		getFreeDiskSpace(path)
	})
}

func Fuzz_Nosy_missingBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *core.BlockChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var blocks []*types.Block
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if chain == nil {
			return
		}

		missingBlocks(chain, blocks)
	})
}

// skipping Fuzz_Nosy_monitorFreeDiskSpace__ because parameters include func, chan, or unsupported interface: chan os.Signal

func Fuzz_Nosy_mustParseBootnodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var urls []string
		fill_err = tp.Fill(&urls)
		if fill_err != nil {
			return
		}

		mustParseBootnodes(urls)
	})
}

func Fuzz_Nosy_setBootstrapNodes__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setBootstrapNodes(ctx, cfg)
	})
}

func Fuzz_Nosy_setBootstrapNodesV5__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setBootstrapNodesV5(ctx, cfg)
	})
}

func Fuzz_Nosy_setEtherbase__(f *testing.F) {
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
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setEtherbase(ctx, cfg)
	})
}

func Fuzz_Nosy_setGPO__(f *testing.F) {
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
		var cfg *gasprice.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setGPO(ctx, cfg)
	})
}

func Fuzz_Nosy_setGraphQL__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setGraphQL(ctx, cfg)
	})
}

func Fuzz_Nosy_setHTTP__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setHTTP(ctx, cfg)
	})
}

func Fuzz_Nosy_setIPC__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setIPC(ctx, cfg)
	})
}

func Fuzz_Nosy_setLes__(f *testing.F) {
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
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setLes(ctx, cfg)
	})
}

func Fuzz_Nosy_setListenAddress__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setListenAddress(ctx, cfg)
	})
}

func Fuzz_Nosy_setMiner__(f *testing.F) {
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
		var cfg *miner.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setMiner(ctx, cfg)
	})
}

func Fuzz_Nosy_setNAT__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setNAT(ctx, cfg)
	})
}

func Fuzz_Nosy_setNodeKey__(f *testing.F) {
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
		var cfg *p2p.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setNodeKey(ctx, cfg)
	})
}

func Fuzz_Nosy_setNodeUserIdent__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setNodeUserIdent(ctx, cfg)
	})
}

func Fuzz_Nosy_setRequiredBlocks__(f *testing.F) {
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
		var cfg *ethconfig.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setRequiredBlocks(ctx, cfg)
	})
}

func Fuzz_Nosy_setSmartCard__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setSmartCard(ctx, cfg)
	})
}

func Fuzz_Nosy_setTxPool__(f *testing.F) {
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
		var cfg *legacypool.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setTxPool(ctx, cfg)
	})
}

func Fuzz_Nosy_setWS__(f *testing.F) {
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
		var cfg *node.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if ctx == nil || cfg == nil {
			return
		}

		setWS(ctx, cfg)
	})
}
