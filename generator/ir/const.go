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
	Constant()
}

// === [ constant.Constant ] ===================================================

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Int) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Float) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Null) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Struct) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Array) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*CharArray) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Vector) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ZeroInitializer) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*Undef) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*BlockAddress) Constant() {}

// --- [ Unary expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFNeg) Constant() {}

// --- [ Binary expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAdd) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFAdd) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSub) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFSub) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprMul) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFMul) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUDiv) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSDiv) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFDiv) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprURem) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSRem) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFRem) Constant() {}

// --- [ Bitwise expressions ] -------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShl) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprLShr) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAShr) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAnd) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprOr) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprXor) Constant() {}

// --- [ Vector expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractElement) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertElement) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprShuffleVector) Constant() {}

// --- [ Aggregate expressions ] -----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprExtractValue) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprInsertValue) Constant() {}

// --- [ Memory expressions ] --------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprGetElementPtr) Constant() {}

// --- [ Conversion expressions ] ----------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprTrunc) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprZExt) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSExt) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPTrunc) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPExt) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToUI) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFPToSI) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprUIToFP) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSIToFP) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprPtrToInt) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprIntToPtr) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprBitCast) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprAddrSpaceCast) Constant() {}

// --- [ Other expressions ] ---------------------------------------------------

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprICmp) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprFCmp) Constant() {}

// IsConstant ensures that only constants can be assigned to the
// constant.Constant interface.
func (*ExprSelect) Constant() {}

// --- [ Index ] --------------------------------

func (*Index) Constant() {}
