package abi

import (
	"io"
	"reflect"
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

func Fuzz_Nosy_ABI_ErrorByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var sigdata [4]byte
		fill_err = tp.Fill(&sigdata)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.ErrorByID(sigdata)
	})
}

func Fuzz_Nosy_ABI_EventByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var topic common.Hash
		fill_err = tp.Fill(&topic)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.EventByID(topic)
	})
}

func Fuzz_Nosy_ABI_HasFallback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.HasFallback()
	})
}

func Fuzz_Nosy_ABI_HasReceive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.HasReceive()
	})
}

func Fuzz_Nosy_ABI_MethodById__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var sigdata []byte
		fill_err = tp.Fill(&sigdata)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.MethodById(sigdata)
	})
}

func Fuzz_Nosy_ABI_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_Argument_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var argument *Argument
		fill_err = tp.Fill(&argument)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if argument == nil {
			return
		}

		argument.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_Error_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		e := NewError(name, inputs)
		e.Unpack(d3)
	})
}

// skipping Fuzz_Nosy_ABI_Pack__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_ABI_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.Unpack(name, d3)
	})
}

// skipping Fuzz_Nosy_ABI_UnpackIntoInterface__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_ABI_UnpackIntoMap__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_ABI_getArguments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reader io.Reader
		fill_err = tp.Fill(&reader)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		a1, err := JSON(reader)
		if err != nil {
			return
		}
		a1.getArguments(name, d3)
	})
}

// skipping Fuzz_Nosy_Arguments_Copy__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Arguments_NonIndexed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var arguments Arguments
		fill_err = tp.Fill(&arguments)
		if fill_err != nil {
			return
		}

		arguments.NonIndexed()
	})
}

// skipping Fuzz_Nosy_Arguments_Pack__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Arguments_PackValues__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Arguments_Unpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var arguments Arguments
		fill_err = tp.Fill(&arguments)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		arguments.Unpack(d2)
	})
}

// skipping Fuzz_Nosy_Arguments_UnpackIntoMap__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_Arguments_UnpackValues__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var arguments Arguments
		fill_err = tp.Fill(&arguments)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		arguments.UnpackValues(d2)
	})
}

// skipping Fuzz_Nosy_Arguments_copyAtomic__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Arguments_copyTuple__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Arguments_isTuple__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var arguments Arguments
		fill_err = tp.Fill(&arguments)
		if fill_err != nil {
			return
		}

		arguments.isTuple()
	})
}

func Fuzz_Nosy_Error_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}

		e := NewError(name, inputs)
		e.String()
	})
}

func Fuzz_Nosy_Event_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var rawName string
		fill_err = tp.Fill(&rawName)
		if fill_err != nil {
			return
		}
		var anonymous bool
		fill_err = tp.Fill(&anonymous)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}

		e := NewEvent(name, rawName, anonymous, inputs)
		e.String()
	})
}

func Fuzz_Nosy_Method_IsConstant__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var rawName string
		fill_err = tp.Fill(&rawName)
		if fill_err != nil {
			return
		}
		var funType FunctionType
		fill_err = tp.Fill(&funType)
		if fill_err != nil {
			return
		}
		var mutability string
		fill_err = tp.Fill(&mutability)
		if fill_err != nil {
			return
		}
		var isConst bool
		fill_err = tp.Fill(&isConst)
		if fill_err != nil {
			return
		}
		var isPayable bool
		fill_err = tp.Fill(&isPayable)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}
		var outputs Arguments
		fill_err = tp.Fill(&outputs)
		if fill_err != nil {
			return
		}

		method := NewMethod(name, rawName, funType, mutability, isConst, isPayable, inputs, outputs)
		method.IsConstant()
	})
}

func Fuzz_Nosy_Method_IsPayable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var rawName string
		fill_err = tp.Fill(&rawName)
		if fill_err != nil {
			return
		}
		var funType FunctionType
		fill_err = tp.Fill(&funType)
		if fill_err != nil {
			return
		}
		var mutability string
		fill_err = tp.Fill(&mutability)
		if fill_err != nil {
			return
		}
		var isConst bool
		fill_err = tp.Fill(&isConst)
		if fill_err != nil {
			return
		}
		var isPayable bool
		fill_err = tp.Fill(&isPayable)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}
		var outputs Arguments
		fill_err = tp.Fill(&outputs)
		if fill_err != nil {
			return
		}

		method := NewMethod(name, rawName, funType, mutability, isConst, isPayable, inputs, outputs)
		method.IsPayable()
	})
}

func Fuzz_Nosy_Method_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var rawName string
		fill_err = tp.Fill(&rawName)
		if fill_err != nil {
			return
		}
		var funType FunctionType
		fill_err = tp.Fill(&funType)
		if fill_err != nil {
			return
		}
		var mutability string
		fill_err = tp.Fill(&mutability)
		if fill_err != nil {
			return
		}
		var isConst bool
		fill_err = tp.Fill(&isConst)
		if fill_err != nil {
			return
		}
		var isPayable bool
		fill_err = tp.Fill(&isPayable)
		if fill_err != nil {
			return
		}
		var inputs Arguments
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}
		var outputs Arguments
		fill_err = tp.Fill(&outputs)
		if fill_err != nil {
			return
		}

		method := NewMethod(name, rawName, funType, mutability, isConst, isPayable, inputs, outputs)
		method.String()
	})
}

func Fuzz_Nosy_Type_GetType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 string
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var internalType string
		fill_err = tp.Fill(&internalType)
		if fill_err != nil {
			return
		}
		var components []ArgumentMarshaling
		fill_err = tp.Fill(&components)
		if fill_err != nil {
			return
		}

		t1, err := NewType(t1, internalType, components)
		if err != nil {
			return
		}
		t1.GetType()
	})
}

func Fuzz_Nosy_Type_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 string
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var internalType string
		fill_err = tp.Fill(&internalType)
		if fill_err != nil {
			return
		}
		var components []ArgumentMarshaling
		fill_err = tp.Fill(&components)
		if fill_err != nil {
			return
		}

		t1, err := NewType(t1, internalType, components)
		if err != nil {
			return
		}
		t1.String()
	})
}

func Fuzz_Nosy_Type_pack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 string
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var internalType string
		fill_err = tp.Fill(&internalType)
		if fill_err != nil {
			return
		}
		var components []ArgumentMarshaling
		fill_err = tp.Fill(&components)
		if fill_err != nil {
			return
		}
		var v reflect.Value
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		t1, err := NewType(t1, internalType, components)
		if err != nil {
			return
		}
		t1.pack(v)
	})
}

func Fuzz_Nosy_Type_requiresLengthPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 string
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var internalType string
		fill_err = tp.Fill(&internalType)
		if fill_err != nil {
			return
		}
		var components []ArgumentMarshaling
		fill_err = tp.Fill(&components)
		if fill_err != nil {
			return
		}

		t1, err := NewType(t1, internalType, components)
		if err != nil {
			return
		}
		t1.requiresLengthPrefix()
	})
}

// skipping Fuzz_Nosy_ConvertType__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MakeTopics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var query [][]interface{}
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}

		MakeTopics(query...)
	})
}

func Fuzz_Nosy_ReadFixedBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var word []byte
		fill_err = tp.Fill(&word)
		if fill_err != nil {
			return
		}

		ReadFixedBytes(t1, word)
	})
}

func Fuzz_Nosy_ReadInteger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var typ Type
		fill_err = tp.Fill(&typ)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		ReadInteger(typ, b)
	})
}

// skipping Fuzz_Nosy_ResolveNameConflict__ because parameters include func, chan, or unsupported interface: func(string) bool

func Fuzz_Nosy_ToCamelCase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		ToCamelCase(input)
	})
}

func Fuzz_Nosy_UnpackRevert__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		UnpackRevert(d1)
	})
}

// skipping Fuzz_Nosy_assembleArgs__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_forEachUnpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var output []byte
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		var start int
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}

		forEachUnpack(t1, output, start, size)
	})
}

func Fuzz_Nosy_forTupleUnpack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var output []byte
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}

		forTupleUnpack(t1, output)
	})
}

func Fuzz_Nosy_formatSliceString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind reflect.Kind
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var sliceSize int
		fill_err = tp.Fill(&sliceSize)
		if fill_err != nil {
			return
		}

		formatSliceString(kind, sliceSize)
	})
}

func Fuzz_Nosy_genIntType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rule int64, size uint) {
		genIntType(rule, size)
	})
}

func Fuzz_Nosy_getTypeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		getTypeSize(t1)
	})
}

func Fuzz_Nosy_isAlpha__(f *testing.F) {
	f.Fuzz(func(t *testing.T, c byte) {
		isAlpha(c)
	})
}

func Fuzz_Nosy_isDigit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, c byte) {
		isDigit(c)
	})
}

func Fuzz_Nosy_isDynamicType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		isDynamicType(t1)
	})
}

func Fuzz_Nosy_isIdentifierSymbol__(f *testing.F) {
	f.Fuzz(func(t *testing.T, c byte) {
		isIdentifierSymbol(c)
	})
}

func Fuzz_Nosy_isLetter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ch rune) {
		isLetter(ch)
	})
}

func Fuzz_Nosy_isValidFieldName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fieldName string) {
		isValidFieldName(fieldName)
	})
}

func Fuzz_Nosy_lengthPrefixPointsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, index int, output []byte) {
		lengthPrefixPointsTo(index, output)
	})
}

func Fuzz_Nosy_mapArgNamesToStructFields__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var argNames []string
		fill_err = tp.Fill(&argNames)
		if fill_err != nil {
			return
		}
		var value reflect.Value
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		mapArgNamesToStructFields(argNames, value)
	})
}

func Fuzz_Nosy_packBytesSlice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bytes []byte, l int) {
		packBytesSlice(bytes, l)
	})
}

func Fuzz_Nosy_packElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var reflectValue reflect.Value
		fill_err = tp.Fill(&reflectValue)
		if fill_err != nil {
			return
		}

		packElement(t1, reflectValue)
	})
}

func Fuzz_Nosy_packNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var value reflect.Value
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		packNum(value)
	})
}

func Fuzz_Nosy_parseCompositeType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string) {
		parseCompositeType(unescapedSelector)
	})
}

func Fuzz_Nosy_parseElementaryType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string) {
		parseElementaryType(unescapedSelector)
	})
}

func Fuzz_Nosy_parseIdentifier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string) {
		parseIdentifier(unescapedSelector)
	})
}

func Fuzz_Nosy_parseToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string, isIdent bool) {
		parseToken(unescapedSelector, isIdent)
	})
}

func Fuzz_Nosy_parseType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, unescapedSelector string) {
		parseType(unescapedSelector)
	})
}

func Fuzz_Nosy_readBool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, word []byte) {
		readBool(word)
	})
}

func Fuzz_Nosy_readFunctionType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 Type
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var word []byte
		fill_err = tp.Fill(&word)
		if fill_err != nil {
			return
		}

		readFunctionType(t1, word)
	})
}

func Fuzz_Nosy_toGoType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var t2 Type
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var output []byte
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}

		toGoType(index, t2, output)
	})
}

func Fuzz_Nosy_tuplePointsTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, index int, output []byte) {
		tuplePointsTo(index, output)
	})
}
