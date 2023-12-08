package compiler

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

func Fuzz_Nosy_ParseCombinedJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, combinedJSON []byte, source string, languageVersion string, compilerVersion string, compilerOptions string) {
		ParseCombinedJSON(combinedJSON, source, languageVersion, compilerVersion, compilerOptions)
	})
}

func Fuzz_Nosy_parseCombinedJSONV8__(f *testing.F) {
	f.Fuzz(func(t *testing.T, combinedJSON []byte, source string, languageVersion string, compilerVersion string, compilerOptions string) {
		parseCombinedJSONV8(combinedJSON, source, languageVersion, compilerVersion, compilerOptions)
	})
}
