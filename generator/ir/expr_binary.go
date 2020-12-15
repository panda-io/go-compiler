package ir

import (
	"fmt"
	"strings"
)

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAdd is an LLVM IR add expression.
type ExprAdd struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewExprAdd returns a new add expression based on the given operands.
func NewExprAdd(x, y Constant) *ExprAdd {
	e := &ExprAdd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAdd) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAdd) Ident() string {
	// 'add' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("add")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAdd) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFAdd is an LLVM IR fadd expression.
type ExprFAdd struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewExprFAdd returns a new fadd expression based on the given operands.
func NewExprFAdd(x, y Constant) *ExprFAdd {
	e := &ExprFAdd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFAdd) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFAdd) Ident() string {
	// 'fadd' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fadd (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFAdd) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSub is an LLVM IR sub expression.
type ExprSub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewExprSub returns a new sub expression based on the given operands.
func NewExprSub(x, y Constant) *ExprSub {
	e := &ExprSub{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSub) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSub) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSub) Ident() string {
	// 'sub' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("sub")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSub) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFSub is an LLVM IR fsub expression.
type ExprFSub struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewExprFSub returns a new fsub expression based on the given operands.
func NewExprFSub(x, y Constant) *ExprFSub {
	e := &ExprFSub{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFSub) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFSub) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFSub) Ident() string {
	// 'fsub' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fsub (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFSub) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprMul is an LLVM IR mul expression.
type ExprMul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewMul returns a new mul expression based on the given operands.
func NewExprMul(x, y Constant) *ExprMul {
	e := &ExprMul{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprMul) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprMul) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprMul) Ident() string {
	// 'mul' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("mul")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprMul) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFMul is an LLVM IR fmul expression.
type ExprFMul struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewFMul returns a new fmul expression based on the given operands.
func NewExprFMul(x, y Constant) *ExprFMul {
	e := &ExprFMul{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFMul) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFMul) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFMul) Ident() string {
	// 'fmul' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fmul (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFMul) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprUDiv is an LLVM IR udiv expression.
type ExprUDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewUDiv returns a new udiv expression based on the given operands.
func NewExprUDiv(x, y Constant) *ExprUDiv {
	e := &ExprUDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprUDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprUDiv) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprUDiv) Ident() string {
	// 'udiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("udiv")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprUDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSDiv is an LLVM IR sdiv expression.
type ExprSDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewSDiv returns a new sdiv expression based on the given operands.
func NewExprSDiv(x, y Constant) *ExprSDiv {
	e := &ExprSDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSDiv) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSDiv) Ident() string {
	// 'sdiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("sdiv")
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFDiv is an LLVM IR fdiv expression.
type ExprFDiv struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewFDiv returns a new fdiv expression based on the given operands.
func NewExprFDiv(x, y Constant) *ExprFDiv {
	e := &ExprFDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFDiv) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFDiv) Ident() string {
	// 'fdiv' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fdiv (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprURem is an LLVM IR urem expression.
type ExprURem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewURem returns a new urem expression based on the given operands.
func NewExprURem(x, y Constant) *ExprURem {
	e := &ExprURem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprURem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprURem) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprURem) Ident() string {
	// 'urem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("urem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprURem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSRem is an LLVM IR srem expression.
type ExprSRem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewSRem returns a new srem expression based on the given operands.
func NewExprSRem(x, y Constant) *ExprSRem {
	e := &ExprSRem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSRem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSRem) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSRem) Ident() string {
	// 'srem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("srem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSRem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFRem is an LLVM IR frem expression.
type ExprFRem struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewFRem returns a new frem expression based on the given operands.
func NewExprFRem(x, y Constant) *ExprFRem {
	e := &ExprFRem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFRem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFRem) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFRem) Ident() string {
	// 'frem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("frem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFRem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}
