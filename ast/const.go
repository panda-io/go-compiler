package ast

import "github.com/panda-foundation/go-compiler/ir"

const (
	Global        = "global"
	FunctionEntry = "entry"
	FunctionBody  = "body"
	FunctionExit  = "exit"
	ProgramEntry  = "main"
	ClassThis     = "this"
	Constructor   = "create"
	Destructor    = "destroy"

	Extern   = "extern"
	Variadic = "variadic"
)

type ValueType int

const (
	TVoid ValueType = iota
	TBool
	TI8
	TI16
	TI32
	TI64
	TUI8
	TUI16
	TUI32
	TUI64
	TF32
	TF64
	TChar
	TString
	TFunction
	TClass
	TPointer
)

var (
	malloc  = ir.NewFunc("malloc", ir.NewPointerType(ir.I8), ir.NewParam(ir.I32))
	free    = ir.NewFunc("free", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	memcpy  = ir.NewFunc("memcpy", ir.NewPointerType(ir.I8), ir.NewParam(ir.NewPointerType(ir.I8)), ir.NewParam(ir.I32))
	memset  = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))
	counter = ir.NewPointerType(&ir.StructType{TypeName: "global.counter"})
)
