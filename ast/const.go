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
	Counter       = "global.counter"

	Extern   = "extern"
	Variadic = "variadic"
)

var (
	malloc  = ir.NewFunc("malloc", ir.NewPointerType(ir.I8), ir.NewParam(ir.I32))
	free    = ir.NewFunc("free", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	memcpy  = ir.NewFunc("memcpy", ir.NewPointerType(ir.I8), ir.NewParam(ir.NewPointerType(ir.I8)), ir.NewParam(ir.I32))
	memset  = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))
	counter = ir.NewPointerType(&ir.StructType{TypeName: Counter})

	releaseShared = ir.NewFunc("global.counter.release_shared", ir.NewPointerType(ir.I8))
	releaseWeak   = ir.NewFunc("global.counter.retain_weak", ir.NewPointerType(ir.I8))
)
