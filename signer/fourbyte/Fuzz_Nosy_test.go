package fourbyte

import (
	"testing"

	"github.com/ethereum/go-ethereum/signer/core/apitypes"
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

func Fuzz_Nosy_Database_AddSelector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, selector string, d3 []byte) {
		db, err := NewWithFile(path)
		if err != nil {
			return
		}
		db.AddSelector(selector, d3)
	})
}

func Fuzz_Nosy_Database_Selector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, id []byte) {
		db, err := NewWithFile(path)
		if err != nil {
			return
		}
		db.Selector(id)
	})
}

func Fuzz_Nosy_Database_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		db, err := NewWithFile(path)
		if err != nil {
			return
		}
		db.Size()
	})
}

func Fuzz_Nosy_Database_ValidateCallData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var selector *string
		fill_err = tp.Fill(&selector)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		var messages *apitypes.ValidationMessages
		fill_err = tp.Fill(&messages)
		if fill_err != nil {
			return
		}
		if selector == nil || messages == nil {
			return
		}

		db, err := NewWithFile(path)
		if err != nil {
			return
		}
		db.ValidateCallData(selector, d3, messages)
	})
}

func Fuzz_Nosy_Database_ValidateTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var selector *string
		fill_err = tp.Fill(&selector)
		if fill_err != nil {
			return
		}
		var tx *apitypes.SendTxArgs
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if selector == nil || tx == nil {
			return
		}

		db, err := NewWithFile(path)
		if err != nil {
			return
		}
		db.ValidateTransaction(selector, tx)
	})
}

func Fuzz_Nosy_decodedArgument_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var arg decodedArgument
		fill_err = tp.Fill(&arg)
		if fill_err != nil {
			return
		}

		arg.String()
	})
}

func Fuzz_Nosy_decodedCallData_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, calldata []byte, unescapedAbidata string) {
		cd, err := parseCallData(calldata, unescapedAbidata)
		if err != nil {
			return
		}
		cd.String()
	})
}

func Fuzz_Nosy_parseSelector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string) {
		parseSelector(unescapedSelector)
	})
}
