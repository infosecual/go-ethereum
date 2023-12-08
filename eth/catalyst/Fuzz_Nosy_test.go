package catalyst

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
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

func Fuzz_Nosy_ConsensusAPI_ExchangeCapabilities__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var _x2 []string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.ExchangeCapabilities(_x2)
	})
}

func Fuzz_Nosy_ConsensusAPI_ExchangeTransitionConfigurationV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var config engine.TransitionConfigurationV1
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.ExchangeTransitionConfigurationV1(config)
	})
}

func Fuzz_Nosy_ConsensusAPI_ForkchoiceUpdatedV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if eth == nil || payloadAttributes == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.ForkchoiceUpdatedV1(update, payloadAttributes)
	})
}

func Fuzz_Nosy_ConsensusAPI_ForkchoiceUpdatedV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if eth == nil || payloadAttributes == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.ForkchoiceUpdatedV2(update, payloadAttributes)
	})
}

func Fuzz_Nosy_ConsensusAPI_ForkchoiceUpdatedV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if eth == nil || payloadAttributes == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.ForkchoiceUpdatedV3(update, payloadAttributes)
	})
}

func Fuzz_Nosy_ConsensusAPI_GetPayloadBodiesByHashV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.GetPayloadBodiesByHashV1(hashes)
	})
}

func Fuzz_Nosy_ConsensusAPI_GetPayloadBodiesByRangeV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var start hexutil.Uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var count hexutil.Uint64
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.GetPayloadBodiesByRangeV1(start, count)
	})
}

func Fuzz_Nosy_ConsensusAPI_GetPayloadV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.GetPayloadV1(payloadID)
	})
}

func Fuzz_Nosy_ConsensusAPI_GetPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.GetPayloadV2(payloadID)
	})
}

func Fuzz_Nosy_ConsensusAPI_GetPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.GetPayloadV3(payloadID)
	})
}

func Fuzz_Nosy_ConsensusAPI_NewPayloadV1__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.NewPayloadV1(params)
	})
}

func Fuzz_Nosy_ConsensusAPI_NewPayloadV2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.NewPayloadV2(params)
	})
}

func Fuzz_Nosy_ConsensusAPI_NewPayloadV3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var versionedHashes []common.Hash
		fill_err = tp.Fill(&versionedHashes)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if eth == nil || beaconRoot == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.NewPayloadV3(params, versionedHashes, beaconRoot)
	})
}

func Fuzz_Nosy_ConsensusAPI_checkInvalidAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var check common.Hash
		fill_err = tp.Fill(&check)
		if fill_err != nil {
			return
		}
		var head common.Hash
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.checkInvalidAncestor(check, head)
	})
}

func Fuzz_Nosy_ConsensusAPI_delayPayloadImport__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if eth == nil || block == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.delayPayloadImport(block)
	})
}

func Fuzz_Nosy_ConsensusAPI_forkchoiceUpdated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var update engine.ForkchoiceStateV1
		fill_err = tp.Fill(&update)
		if fill_err != nil {
			return
		}
		var payloadAttributes *engine.PayloadAttributes
		fill_err = tp.Fill(&payloadAttributes)
		if fill_err != nil {
			return
		}
		if eth == nil || payloadAttributes == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.forkchoiceUpdated(update, payloadAttributes)
	})
}

func Fuzz_Nosy_ConsensusAPI_getPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var payloadID engine.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		var full bool
		fill_err = tp.Fill(&full)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.getPayload(payloadID, full)
	})
}

func Fuzz_Nosy_ConsensusAPI_heartbeat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.heartbeat()
	})
}

// skipping Fuzz_Nosy_ConsensusAPI_invalid__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ConsensusAPI_newPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var params engine.ExecutableData
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var versionedHashes []common.Hash
		fill_err = tp.Fill(&versionedHashes)
		if fill_err != nil {
			return
		}
		var beaconRoot *common.Hash
		fill_err = tp.Fill(&beaconRoot)
		if fill_err != nil {
			return
		}
		if eth == nil || beaconRoot == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.newPayload(params, versionedHashes, beaconRoot)
	})
}

func Fuzz_Nosy_ConsensusAPI_setInvalidAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var invalid *types.Header
		fill_err = tp.Fill(&invalid)
		if fill_err != nil {
			return
		}
		var origin *types.Header
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		if eth == nil || invalid == nil || origin == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.setInvalidAncestor(invalid, origin)
	})
}

func Fuzz_Nosy_ConsensusAPI_verifyPayloadAttributes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var attr *engine.PayloadAttributes
		fill_err = tp.Fill(&attr)
		if fill_err != nil {
			return
		}
		if eth == nil || attr == nil {
			return
		}

		api := NewConsensusAPI(eth)
		api.verifyPayloadAttributes(attr)
	})
}

func Fuzz_Nosy_FullSyncTester_Start__(f *testing.F) {
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
		var backend *eth.Ethereum
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		var t3 common.Hash
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if stack == nil || backend == nil {
			return
		}

		tester, err := RegisterFullSyncTester(stack, backend, t3)
		if err != nil {
			return
		}
		tester.Start()
	})
}

func Fuzz_Nosy_FullSyncTester_Stop__(f *testing.F) {
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
		var backend *eth.Ethereum
		fill_err = tp.Fill(&backend)
		if fill_err != nil {
			return
		}
		var t3 common.Hash
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if stack == nil || backend == nil {
			return
		}

		tester, err := RegisterFullSyncTester(stack, backend, t3)
		if err != nil {
			return
		}
		tester.Stop()
	})
}

func Fuzz_Nosy_SimulatedBeacon_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.Start()
	})
}

func Fuzz_Nosy_SimulatedBeacon_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.Stop()
	})
}

func Fuzz_Nosy_SimulatedBeacon_finalizedBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.finalizedBlockHash(number)
	})
}

func Fuzz_Nosy_SimulatedBeacon_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.loop()
	})
}

func Fuzz_Nosy_SimulatedBeacon_loopOnDemand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.loopOnDemand()
	})
}

func Fuzz_Nosy_SimulatedBeacon_sealBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var withdrawals []*types.Withdrawal
		fill_err = tp.Fill(&withdrawals)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.sealBlock(withdrawals)
	})
}

func Fuzz_Nosy_SimulatedBeacon_setCurrentState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var headHash common.Hash
		fill_err = tp.Fill(&headHash)
		if fill_err != nil {
			return
		}
		var finalizedHash common.Hash
		fill_err = tp.Fill(&finalizedHash)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.setCurrentState(headHash, finalizedHash)
	})
}

func Fuzz_Nosy_SimulatedBeacon_setFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var eth *eth.Ethereum
		fill_err = tp.Fill(&eth)
		if fill_err != nil {
			return
		}
		var feeRecipient common.Address
		fill_err = tp.Fill(&feeRecipient)
		if fill_err != nil {
			return
		}
		if eth == nil {
			return
		}

		c, err := NewSimulatedBeacon(period, eth)
		if err != nil {
			return
		}
		c.setFeeRecipient(feeRecipient)
	})
}

func Fuzz_Nosy_api_AddWithdrawal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *api
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var withdrawal *types.Withdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if a == nil || withdrawal == nil {
			return
		}

		a.AddWithdrawal(ctx, withdrawal)
	})
}

func Fuzz_Nosy_api_SetFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *api
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var feeRecipient common.Address
		fill_err = tp.Fill(&feeRecipient)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.SetFeeRecipient(ctx, feeRecipient)
	})
}

func Fuzz_Nosy_headerQueue_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		q := newHeaderQueue()
		q.get(hash)
	})
}

func Fuzz_Nosy_headerQueue_put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var d2 *types.Header
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if d2 == nil {
			return
		}

		q := newHeaderQueue()
		q.put(hash, d2)
	})
}

func Fuzz_Nosy_payloadQueue_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var full bool
		fill_err = tp.Fill(&full)
		if fill_err != nil {
			return
		}

		q := newPayloadQueue()
		q.get(id, full)
	})
}

func Fuzz_Nosy_payloadQueue_has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		q := newPayloadQueue()
		q.has(id)
	})
}

func Fuzz_Nosy_payloadQueue_put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id engine.PayloadID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var payload *miner.Payload
		fill_err = tp.Fill(&payload)
		if fill_err != nil {
			return
		}
		if payload == nil {
			return
		}

		q := newPayloadQueue()
		q.put(id, payload)
	})
}

func Fuzz_Nosy_withdrawalQueue_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *withdrawalQueue
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var withdrawal *types.Withdrawal
		fill_err = tp.Fill(&withdrawal)
		if fill_err != nil {
			return
		}
		if w == nil || withdrawal == nil {
			return
		}

		w.add(withdrawal)
	})
}

func Fuzz_Nosy_withdrawalQueue_gatherPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *withdrawalQueue
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var maxCount int
		fill_err = tp.Fill(&maxCount)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.gatherPending(maxCount)
	})
}

func Fuzz_Nosy_RegisterSimulatedBeaconAPIs__(f *testing.F) {
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
		var sim *SimulatedBeacon
		fill_err = tp.Fill(&sim)
		if fill_err != nil {
			return
		}
		if stack == nil || sim == nil {
			return
		}

		RegisterSimulatedBeaconAPIs(stack, sim)
	})
}
