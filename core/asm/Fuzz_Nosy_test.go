package asm

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
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

func Fuzz_Nosy_Compiler_Compile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool) {
		c := NewCompiler(debug)
		c.Compile()
	})
}

// skipping Fuzz_Nosy_Compiler_Feed__ because parameters include func, chan, or unsupported interface: <-chan github.com/ethereum/go-ethereum/core/asm.token

func Fuzz_Nosy_Compiler_compileElement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var debug bool
		fill_err = tp.Fill(&debug)
		if fill_err != nil {
			return
		}
		var element token
		fill_err = tp.Fill(&element)
		if fill_err != nil {
			return
		}

		c := NewCompiler(debug)
		c.compileElement(element)
	})
}

func Fuzz_Nosy_Compiler_compileJump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool, jumpType string) {
		c := NewCompiler(debug)
		c.compileJump(jumpType)
	})
}

func Fuzz_Nosy_Compiler_compileLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool) {
		c := NewCompiler(debug)
		c.compileLabel()
	})
}

func Fuzz_Nosy_Compiler_compileLine__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool) {
		c := NewCompiler(debug)
		c.compileLine()
	})
}

func Fuzz_Nosy_Compiler_compilePush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool) {
		c := NewCompiler(debug)
		c.compilePush()
	})
}

func Fuzz_Nosy_Compiler_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool) {
		c := NewCompiler(debug)
		c.next()
	})
}

func Fuzz_Nosy_Compiler_outputBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, debug bool, b []byte) {
		c := NewCompiler(debug)
		c.outputBytes(b)
	})
}

func Fuzz_Nosy_Compiler_outputOpcode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var debug bool
		fill_err = tp.Fill(&debug)
		if fill_err != nil {
			return
		}
		var op vm.OpCode
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}

		c := NewCompiler(debug)
		c.outputOpcode(op)
	})
}

func Fuzz_Nosy_instructionIterator_Arg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, code []byte) {
		it := NewInstructionIterator(code)
		it.Arg()
	})
}

func Fuzz_Nosy_instructionIterator_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, code []byte) {
		it := NewInstructionIterator(code)
		it.Error()
	})
}

func Fuzz_Nosy_instructionIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, code []byte) {
		it := NewInstructionIterator(code)
		it.Next()
	})
}

func Fuzz_Nosy_instructionIterator_Op__(f *testing.F) {
	f.Fuzz(func(t *testing.T, code []byte) {
		it := NewInstructionIterator(code)
		it.Op()
	})
}

func Fuzz_Nosy_instructionIterator_PC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, code []byte) {
		it := NewInstructionIterator(code)
		it.PC()
	})
}

func Fuzz_Nosy_lexer_accept__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var valid string
		fill_err = tp.Fill(&valid)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.accept(valid)
	})
}

func Fuzz_Nosy_lexer_acceptRun__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var valid string
		fill_err = tp.Fill(&valid)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.acceptRun(valid)
	})
}

func Fuzz_Nosy_lexer_acceptRunUntil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var until rune
		fill_err = tp.Fill(&until)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.acceptRunUntil(until)
	})
}

func Fuzz_Nosy_lexer_backup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.backup()
	})
}

func Fuzz_Nosy_lexer_blob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.blob()
	})
}

func Fuzz_Nosy_lexer_emit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var t2 tokenType
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.emit(t2)
	})
}

func Fuzz_Nosy_lexer_ignore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.ignore()
	})
}

func Fuzz_Nosy_lexer_next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.next()
	})
}

func Fuzz_Nosy_lexer_peek__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l *lexer
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if l == nil {
			return
		}

		l.peek()
	})
}

func Fuzz_Nosy_compileError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err compileError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_tokenType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i tokenType
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}

		i.String()
	})
}

func Fuzz_Nosy_Disassemble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, script []byte) {
		Disassemble(script)
	})
}

func Fuzz_Nosy_Lex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, source []byte, debug bool) {
		Lex(source, debug)
	})
}

func Fuzz_Nosy_isJump__(f *testing.F) {
	f.Fuzz(func(t *testing.T, op string) {
		isJump(op)
	})
}

func Fuzz_Nosy_isLetter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, t1 rune) {
		isLetter(t1)
	})
}

func Fuzz_Nosy_isNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, t1 rune) {
		isNumber(t1)
	})
}

func Fuzz_Nosy_isPush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, op string) {
		isPush(op)
	})
}

func Fuzz_Nosy_isSpace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, t1 rune) {
		isSpace(t1)
	})
}

func Fuzz_Nosy_parseNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tok token
		fill_err = tp.Fill(&tok)
		if fill_err != nil {
			return
		}

		parseNumber(tok)
	})
}
