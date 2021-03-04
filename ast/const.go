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
	malloc = ir.NewFunc("malloc", pointerType, ir.NewParam(ir.I32))
	free   = ir.NewFunc("free", ir.Void, ir.NewParam(pointerType))
	memcpy = ir.NewFunc("memcpy", pointerType, ir.NewParam(pointerType), ir.NewParam(ir.I32))
	memset = ir.NewFunc("memset", ir.Void, ir.NewParam(ir.I32), ir.NewParam(ir.I32))

	pointerType   = ir.NewPointerType(ir.I8)
	counterType   = ir.NewPointerType(&ir.StructType{TypeName: Counter})
	retainShared  = ir.NewFunc("global.counter.retain_shared", ir.Void, ir.NewParam(pointerType))
	releaseShared = ir.NewFunc("global.counter.release_shared", ir.Void, ir.NewParam(pointerType))
	retainWeak    = ir.NewFunc("global.counter.retain_weak", ir.Void, ir.NewParam(pointerType))
	releaseWeak   = ir.NewFunc("global.counter.release_weak", ir.Void, ir.NewParam(pointerType))
)
