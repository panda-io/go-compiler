package ir

import (
	"fmt"
	"strings"
)

// --- [ Binary instructions ] -------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAdd is an LLVM IR add instruction.
type InstAdd struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewAdd returns a new add instruction based on the given operands.
func NewAdd(x, y Value) *InstAdd {
	inst := &InstAdd{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAdd) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'add' OverflowFlags=OverflowFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstAdd) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("add")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFAdd is an LLVM IR fadd instruction.
type InstFAdd struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewFAdd returns a new fadd instruction based on the given operands.
func NewFAdd(x, y Value) *InstFAdd {
	inst := &InstFAdd{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFAdd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFAdd) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fadd' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFAdd) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fadd")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSub is an LLVM IR sub instruction.
type InstSub struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewSub returns a new sub instruction based on the given operands.
func NewSub(x, y Value) *InstSub {
	inst := &InstSub{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSub) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'sub' OverflowFlags=OverflowFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstSub) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("sub")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFSub is an LLVM IR fsub instruction.
type InstFSub struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewFSub returns a new fsub instruction based on the given operands.
func NewFSub(x, y Value) *InstFSub {
	inst := &InstFSub{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFSub) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFSub) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fsub' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFSub) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fsub")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstMul is an LLVM IR mul instruction.
type InstMul struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewMul returns a new mul instruction based on the given operands.
func NewMul(x, y Value) *InstMul {
	inst := &InstMul{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstMul) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'mul' OverflowFlags=OverflowFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstMul) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("mul")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFMul is an LLVM IR fmul instruction.
type InstFMul struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewFMul returns a new fmul instruction based on the given operands.
func NewFMul(x, y Value) *InstFMul {
	inst := &InstFMul{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFMul) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFMul) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fmul' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFMul) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fmul")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUDiv is an LLVM IR udiv instruction.
type InstUDiv struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewUDiv returns a new udiv instruction based on the given operands.
func NewUDiv(x, y Value) *InstUDiv {
	inst := &InstUDiv{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstUDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstUDiv) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'udiv' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstUDiv) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("udiv")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSDiv is an LLVM IR sdiv instruction.
type InstSDiv struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewSDiv returns a new sdiv instruction based on the given operands.
func NewSDiv(x, y Value) *InstSDiv {
	inst := &InstSDiv{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSDiv) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'sdiv' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstSDiv) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("sdiv")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFDiv is an LLVM IR fdiv instruction.
type InstFDiv struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewFDiv returns a new fdiv instruction based on the given operands.
func NewFDiv(x, y Value) *InstFDiv {
	inst := &InstFDiv{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFDiv) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFDiv) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fdiv' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFDiv) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("fdiv")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstURem is an LLVM IR urem instruction.
type InstURem struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewURem returns a new urem instruction based on the given operands.
func NewURem(x, y Value) *InstURem {
	inst := &InstURem{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstURem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstURem) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'urem' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstURem) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "urem %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSRem is an LLVM IR srem instruction.
type InstSRem struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewSRem returns a new srem instruction based on the given operands.
func NewSRem(x, y Value) *InstSRem {
	inst := &InstSRem{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSRem) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'srem' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstSRem) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "srem %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFRem is an LLVM IR frem instruction.
type InstFRem struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // floating-point scalar or floating-point vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewFRem returns a new frem instruction based on the given operands.
func NewFRem(x, y Value) *InstFRem {
	inst := &InstFRem{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFRem) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFRem) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'frem' FastMathFlags=FastMathFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstFRem) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("frem")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}
