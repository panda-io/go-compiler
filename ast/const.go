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
	malloc = ir.NewFunc("malloc", ir.NewPointerType(ir.I8), ir.NewParam(ir.I32))
	free   = ir.NewFunc("free", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	memcpy = ir.NewFunc("memcpy", ir.NewPointerType(ir.I8), ir.NewParam(ir.NewPointerType(ir.I8)), ir.NewParam(ir.I32))
	memset = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))

	counterType   = ir.NewPointerType(&ir.StructType{TypeName: Counter})
	retainShared  = ir.NewFunc("global.counter.retain_shared", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	releaseShared = ir.NewFunc("global.counter.release_shared", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	retainWeak    = ir.NewFunc("global.counter.retain_weak", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
	releaseWeak   = ir.NewFunc("global.counter.release_weak", ir.Void, ir.NewParam(ir.NewPointerType(ir.I8)))
)
