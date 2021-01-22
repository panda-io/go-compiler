package ast

import "github.com/panda-foundation/go-compiler/ir"

const (
	Global        = "global"
	FunctionEntry = "entry"
	ClassThis     = "this"
	Constructor   = "create"
	Destructor    = "destroy"

	Extern  = "extern"
	Builtin = "builtin"
)

var (
	malloc  = ir.NewFunc("malloc", ir.NewPointerType(ir.I8), ir.NewParam(ir.I32))
	free    = ir.NewFunc("free", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	memcpy  = ir.NewFunc("memcpy", ir.NewPointerType(ir.I8), ir.NewParam(ir.NewPointerType(ir.I8)), ir.NewParam(ir.I32))
	memset  = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))
	counter = ir.NewPointerType(&ir.StructType{TypeName: "global.counter"})
)
