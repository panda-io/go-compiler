package ir

import (
	"fmt"
	"strings"
)

// --- [ Bitwise instructions ] ------------------------------------------------

// ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShl is an LLVM IR shl instruction.
type InstShl struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalar or integer vector

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewShl returns a new shl instruction based on the given operands.
func NewShl(x, y Value) *InstShl {
	inst := &InstShl{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstShl) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstShl) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'shl' OverflowFlags=OverflowFlag* X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstShl) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("shl")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLShr is an LLVM IR lshr instruction.
type InstLShr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewLShr returns a new lshr instruction based on the given operands.
func NewLShr(x, y Value) *InstLShr {
	inst := &InstLShr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLShr) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'lshr' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstLShr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("lshr")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAShr is an LLVM IR ashr instruction.
type InstAShr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewAShr returns a new ashr instruction based on the given operands.
func NewAShr(x, y Value) *InstAShr {
	inst := &InstAShr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAShr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAShr) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'ashr' Exactopt X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstAShr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("ashr")
	fmt.Fprintf(buf, " %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAnd is an LLVM IR and instruction.
type InstAnd struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewAnd returns a new and instruction based on the given operands.
func NewAnd(x, y Value) *InstAnd {
	inst := &InstAnd{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAnd) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAnd) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'and' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstAnd) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "and %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstOr is an LLVM IR or instruction.
type InstOr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewOr returns a new or instruction based on the given operands.
func NewOr(x, y Value) *InstOr {
	inst := &InstOr{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstOr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstOr) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'or' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstOr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "or %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}

// ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstXor is an LLVM IR xor instruction.
type InstXor struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Operands.
	X, Y Value // integer scalars or vectors

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewXor returns a new xor instruction based on the given operands.
func NewXor(x, y Value) *InstXor {
	inst := &InstXor{X: x, Y: y}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstXor) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstXor) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'xor' X=TypeValue ',' Y=Value Metadata=(',' MetadataAttachment)+?
func (inst *InstXor) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "xor %s, %s", inst.X, inst.Y.Ident())
	return buf.String()
}
