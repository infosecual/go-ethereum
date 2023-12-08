package bind

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

// skipping Fuzz_Nosy_BoundContract_Call__ because parameters include func, chan, or unsupported interface: *[]interface{}

func Fuzz_Nosy_BoundContract_FilterLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *FilterOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var query [][]interface{}
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil {
			return
		}

		c.FilterLogs(opts, name, query...)
	})
}

func Fuzz_Nosy_BoundContract_RawTransact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var calldata []byte
		fill_err = tp.Fill(&calldata)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil {
			return
		}

		c.RawTransact(opts, calldata)
	})
}

// skipping Fuzz_Nosy_BoundContract_Transact__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_BoundContract_Transfer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil {
			return
		}

		c.Transfer(opts)
	})
}

// skipping Fuzz_Nosy_BoundContract_UnpackLog__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BoundContract_UnpackLogIntoMap__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_BoundContract_WatchLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *WatchOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var query [][]interface{}
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil {
			return
		}

		c.WatchLogs(opts, name, query...)
	})
}

func Fuzz_Nosy_BoundContract_createDynamicTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var contract *common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil || contract == nil || head == nil {
			return
		}

		c.createDynamicTx(opts, contract, input, head)
	})
}

func Fuzz_Nosy_BoundContract_createLegacyTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var contract *common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil || contract == nil {
			return
		}

		c.createLegacyTx(opts, contract, input)
	})
}

func Fuzz_Nosy_BoundContract_estimateGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var contract *common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var gasPrice *big.Int
		fill_err = tp.Fill(&gasPrice)
		if fill_err != nil {
			return
		}
		var gasTipCap *big.Int
		fill_err = tp.Fill(&gasTipCap)
		if fill_err != nil {
			return
		}
		var gasFeeCap *big.Int
		fill_err = tp.Fill(&gasFeeCap)
		if fill_err != nil {
			return
		}
		var value *big.Int
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil || contract == nil || gasPrice == nil || gasTipCap == nil || gasFeeCap == nil || value == nil {
			return
		}

		c.estimateGasLimit(opts, contract, input, gasPrice, gasTipCap, gasFeeCap, value)
	})
}

func Fuzz_Nosy_BoundContract_getNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil {
			return
		}

		c.getNonce(opts)
	})
}

func Fuzz_Nosy_BoundContract_transact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BoundContract
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var opts *TransactOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var contract *common.Address
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil || opts == nil || contract == nil {
			return
		}

		c.transact(opts, contract, input)
	})
}

func Fuzz_Nosy_MetaData_GetAbi__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MetaData
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAbi()
	})
}

// skipping Fuzz_Nosy_BlockHashContractCaller_CallContractAtHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.BlockHashContractCaller

// skipping Fuzz_Nosy_BlockHashContractCaller_CodeAtHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.BlockHashContractCaller

// skipping Fuzz_Nosy_ContractCaller_CallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractCaller

// skipping Fuzz_Nosy_ContractCaller_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractCaller

// skipping Fuzz_Nosy_ContractFilterer_FilterLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractFilterer

// skipping Fuzz_Nosy_ContractFilterer_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractFilterer

// skipping Fuzz_Nosy_ContractTransactor_EstimateGas__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_PendingCodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_PendingNonceAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_SendTransaction__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_SuggestGasPrice__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_ContractTransactor_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractTransactor

// skipping Fuzz_Nosy_DeployBackend_CodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.DeployBackend

// skipping Fuzz_Nosy_DeployBackend_TransactionReceipt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.DeployBackend

// skipping Fuzz_Nosy_PendingContractCaller_PendingCallContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.PendingContractCaller

// skipping Fuzz_Nosy_PendingContractCaller_PendingCodeAt__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.PendingContractCaller

func Fuzz_Nosy_Bind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var types []string
		fill_err = tp.Fill(&types)
		if fill_err != nil {
			return
		}
		var abis []string
		fill_err = tp.Fill(&abis)
		if fill_err != nil {
			return
		}
		var bytecodes []string
		fill_err = tp.Fill(&bytecodes)
		if fill_err != nil {
			return
		}
		var fsigs []map[string]string
		fill_err = tp.Fill(&fsigs)
		if fill_err != nil {
			return
		}
		var pkg string
		fill_err = tp.Fill(&pkg)
		if fill_err != nil {
			return
		}
		var lang Lang
		fill_err = tp.Fill(&lang)
		if fill_err != nil {
			return
		}
		var libs map[string]string
		fill_err = tp.Fill(&libs)
		if fill_err != nil {
			return
		}
		var aliases map[string]string
		fill_err = tp.Fill(&aliases)
		if fill_err != nil {
			return
		}

		Bind(types, abis, bytecodes, fsigs, pkg, lang, libs, aliases)
	})
}

// skipping Fuzz_Nosy_DeployContract__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/accounts/abi/bind.ContractBackend

func Fuzz_Nosy_alias__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var aliases map[string]string
		fill_err = tp.Fill(&aliases)
		if fill_err != nil {
			return
		}
		var n string
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		alias(aliases, n)
	})
}

func Fuzz_Nosy_bindBasicTypeGo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind abi.Type
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}

		bindBasicTypeGo(kind)
	})
}

func Fuzz_Nosy_bindStructTypeGo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind abi.Type
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var structs map[string]*tmplStruct
		fill_err = tp.Fill(&structs)
		if fill_err != nil {
			return
		}

		bindStructTypeGo(kind, structs)
	})
}

func Fuzz_Nosy_bindTopicTypeGo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind abi.Type
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var structs map[string]*tmplStruct
		fill_err = tp.Fill(&structs)
		if fill_err != nil {
			return
		}

		bindTopicTypeGo(kind, structs)
	})
}

func Fuzz_Nosy_bindTypeGo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind abi.Type
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var structs map[string]*tmplStruct
		fill_err = tp.Fill(&structs)
		if fill_err != nil {
			return
		}

		bindTypeGo(kind, structs)
	})
}

func Fuzz_Nosy_decapitalise__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		decapitalise(input)
	})
}

func Fuzz_Nosy_hasStruct__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 abi.Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		hasStruct(t1)
	})
}

func Fuzz_Nosy_isKeyWord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, arg string) {
		isKeyWord(arg)
	})
}

func Fuzz_Nosy_structured__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args abi.Arguments
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		structured(args)
	})
}
