package apitypes

import (
	"testing"

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

func Fuzz_Nosy_NameValueType_Pprint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nvt *NameValueType
		fill_err = tp.Fill(&nvt)
		if fill_err != nil {
			return
		}
		var depth int
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}
		if nvt == nil {
			return
		}

		nvt.Pprint(depth)
	})
}

func Fuzz_Nosy_SendTxArgs_ToTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args *SendTxArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if args == nil {
			return
		}

		args.ToTransaction()
	})
}

func Fuzz_Nosy_Type_isArray__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.isArray()
	})
}

func Fuzz_Nosy_Type_typeName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.typeName()
	})
}

func Fuzz_Nosy_TypedData_Dependencies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		var primaryType string
		fill_err = tp.Fill(&primaryType)
		if fill_err != nil {
			return
		}
		var found []string
		fill_err = tp.Fill(&found)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.Dependencies(primaryType, found)
	})
}

// skipping Fuzz_Nosy_TypedData_EncodeData__ because parameters include func, chan, or unsupported interface: map[string]interface{}

// skipping Fuzz_Nosy_TypedData_EncodePrimitiveValue__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_TypedData_EncodeType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		var primaryType string
		fill_err = tp.Fill(&primaryType)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.EncodeType(primaryType)
	})
}

func Fuzz_Nosy_TypedData_Format__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.Format()
	})
}

// skipping Fuzz_Nosy_TypedData_HashStruct__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_TypedData_Map__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.Map()
	})
}

func Fuzz_Nosy_TypedData_TypeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		var primaryType string
		fill_err = tp.Fill(&primaryType)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.TypeHash(primaryType)
	})
}

// skipping Fuzz_Nosy_TypedData_formatData__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_TypedData_validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData *TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}
		if typedData == nil {
			return
		}

		typedData.validate()
	})
}

func Fuzz_Nosy_TypedDataDomain_Map__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var domain *TypedDataDomain
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		if domain == nil {
			return
		}

		domain.Map()
	})
}

func Fuzz_Nosy_TypedDataDomain_validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var domain *TypedDataDomain
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		if domain == nil {
			return
		}

		domain.validate()
	})
}

func Fuzz_Nosy_ValidationMessages_Crit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs *ValidationMessages
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if vs == nil {
			return
		}

		vs.Crit(msg)
	})
}

func Fuzz_Nosy_ValidationMessages_GetWarnings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidationMessages
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.GetWarnings()
	})
}

func Fuzz_Nosy_ValidationMessages_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs *ValidationMessages
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if vs == nil {
			return
		}

		vs.Info(msg)
	})
}

func Fuzz_Nosy_ValidationMessages_Warn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs *ValidationMessages
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if vs == nil {
			return
		}

		vs.Warn(msg)
	})
}

func Fuzz_Nosy_SendTxArgs_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args SendTxArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		args.String()
	})
}

func Fuzz_Nosy_Types_validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Types
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.validate()
	})
}

func Fuzz_Nosy_TypedDataAndHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typedData TypedData
		fill_err = tp.Fill(&typedData)
		if fill_err != nil {
			return
		}

		TypedDataAndHash(typedData)
	})
}

// skipping Fuzz_Nosy_convertDataToSlice__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_formatPrimitiveValue__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_isPrimitiveTypeValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, primitiveType string) {
		isPrimitiveTypeValid(primitiveType)
	})
}

// skipping Fuzz_Nosy_parseBytes__ because parameters include func, chan, or unsupported interface: interface{}
