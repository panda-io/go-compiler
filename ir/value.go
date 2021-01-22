// Package ir provides a definition of LLVM IR values.
package ir

import (
	"fmt"
)

// Value is an LLVM IR value, which may be used as an operand of instructions
// and terminators.
//
// A Value has one of the following underlying types.
//
//    constant.Constant   // https://godoc.org/github.com/llir/llvm/ir/constant#Constant
//    value.Named         // https://godoc.org/github.com/llir/llvm/ir/value#Named
type Value interface {
	// String returns the LLVM syntax representation of the value as a type-value
	// pair.
	fmt.Stringer
	// Type returns the type of the value.
	Type() Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

// Named is a named LLVM IR value.
//
// A Named value has one of the following underlying types.
//
//    *ir.Global            // https://godoc.org/github.com/llir/llvm/ir#Global
//    *ir.Func              // https://godoc.org/github.com/llir/llvm/ir#Func
//    *ir.Param             // https://godoc.org/github.com/llir/llvm/ir#Param
//    *ir.Block             // https://godoc.org/github.com/llir/llvm/ir#Block
//    *ir.Instruction       // https://godoc.org/github.com/llir/llvm/ir#Instruction (except store and fence)
//    *ir.TermInvoke        // https://godoc.org/github.com/llir/llvm/ir#TermInvoke
//    *ir.TermCatchSwitch   // https://godoc.org/github.com/llir/llvm/ir#TermCatchSwitch (token result used by catchpad)
type Named interface {
	Value
	// Name returns the name of the value.
	Name() string
	// SetName sets the name of the value.
	SetName(name string)
}

func IsConstant(v Value) bool {
	_, ok := v.(Constant)
	return ok
}
