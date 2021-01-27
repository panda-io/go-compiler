package ir

// === [ Constants ] ===

// Convenience constants.
var (
	// Boolean constants.
	True  = NewInt(I1, 1) // true
	False = NewInt(I1, 0) // false
)

// Constant is an LLVM IR constant; a value that is immutable at runtime, such
// as an integer or floating-point literal, or the address of a function or
// global variable.
//
// A Constant has one of the following underlying types.
//
// Simple constants
//
// https://llvm.org/docs/LangRef.html#simple-constants
//
//    *constant.Int         // https://godoc.org/github.com/llir/llvm/ir/constant#Int
//    *constant.Float       // https://godoc.org/github.com/llir/llvm/ir/constant#Float
//    *constant.Null        // https://godoc.org/github.com/llir/llvm/ir/constant#Null
//
// Complex constants
//
// https://llvm.org/docs/LangRef.html#complex-constants
//
//    *constant.Struct            // https://godoc.org/github.com/llir/llvm/ir/constant#Struct
//    *constant.Array             // https://godoc.org/github.com/llir/llvm/ir/constant#Array
//    *constant.CharArray         // https://godoc.org/github.com/llir/llvm/ir/constant#CharArray
//    *constant.Vector            // https://godoc.org/github.com/llir/llvm/ir/constant#Vector
//
// Global variable and function addresses
//
// https://llvm.org/docs/LangRef.html#global-variable-and-function-addresses
//
//    *ir.Global   // https://godoc.org/github.com/llir/llvm/ir#Global
//    *ir.Func     // https://godoc.org/github.com/llir/llvm/ir#Func
//    *ir.Alias    // https://godoc.org/github.com/llir/llvm/ir#Alias
//    *ir.IFunc    // https://godoc.org/github.com/llir/llvm/ir#IFunc
//
//
// Addresses of basic blocks
//
// https://llvm.org/docs/LangRef.html#addresses-of-basic-blocks
//
//    *constant.BlockAddress   // https://godoc.org/github.com/llir/llvm/ir/constant#BlockAddress
//
// Constant expressions
//
// https://llvm.org/docs/LangRef.html#constant-expressions
//
//    constant.Expression   // https://godoc.org/github.com/llir/llvm/ir/constant#Expression

type Constant interface {
	Value
	isConstant()
}

// === [ constant.Constant ] ===================================================

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Global) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Func) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Float) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Struct) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Array) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*CharArray) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Vector) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Undef) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*BlockAddress) isConstant() {}

// --- [ Unary expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFNeg) isConstant() {}

// --- [ Binary expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAdd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFAdd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSub) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFSub) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprMul) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFMul) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFDiv) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprURem) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSRem) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFRem) isConstant() {}

// --- [ Bitwise expressions ] -------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShl) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprLShr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAShr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAnd) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprOr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprXor) isConstant() {}

// --- [ Vector expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractElement) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertElement) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShuffleVector) isConstant() {}

// --- [ Aggregate expressions ] -----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractValue) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertValue) isConstant() {}

// --- [ Memory expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprGetElementPtr) isConstant() {}

// --- [ Conversion expressions ] ----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprTrunc) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprZExt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSExt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPTrunc) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPExt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToUI) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToSI) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUIToFP) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSIToFP) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprPtrToInt) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprIntToPtr) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprBitCast) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAddrSpaceCast) isConstant() {}

// --- [ Other expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprICmp) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFCmp) isConstant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSelect) isConstant() {}

// --- [ Index ] --------------------------------

func (*Index) isConstant() {}
