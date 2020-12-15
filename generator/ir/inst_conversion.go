package ir

import (
	"fmt"
	"strings"
)

// --- [ Conversion instructions ] ---------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstTrunc is an LLVM IR trunc instruction.
type InstTrunc struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewTrunc returns a new trunc instruction based on the given source value and
// target type.
func NewTrunc(from Value, to Type) *InstTrunc {
	// Type-check operands.
	fromType := from.Type()
	// Note: intentional alias, so we can use it for checking
	// the integer types within vectors.
	toType := to
	if fromVectorT, ok := fromType.(*VectorType); ok {
		toVectorT, ok := toType.(*VectorType)
		if !ok {
			panic(fmt.Errorf("trunc operands are not compatible: from=%v; to=%v", fromVectorT, to))
		}
		if fromVectorT.Len != toVectorT.Len {
			panic(fmt.Errorf("trunc vector operand length mismatch: from=%v; to=%v", from.Type(), to))
		}
		fromType = fromVectorT.ElemType
		toType = toVectorT.ElemType
	}
	if fromIntT, ok := fromType.(*IntType); ok {
		toIntT, ok := toType.(*IntType)
		if !ok {
			panic(fmt.Errorf("trunc operands are not compatible: from=%v; to=%T", fromIntT, to))
		}
		fromSize := fromIntT.BitSize
		toSize := toIntT.BitSize
		if fromSize < toSize {
			panic(fmt.Errorf("invalid trunc operands: from.BitSize < to.BitSize (%v is smaller than %v)", from.Type(), to))
		}
	}
	return &InstTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstTrunc) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'trunc' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstTrunc) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "trunc %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstZExt is an LLVM IR zext instruction.
type InstZExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewZExt returns a new zext instruction based on the given source value and
// target type.
func NewZExt(from Value, to Type) *InstZExt {
	return &InstZExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstZExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstZExt) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'zext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstZExt) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "zext %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSExt is an LLVM IR sext instruction.
type InstSExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewSExt returns a new sext instruction based on the given source value and
// target type.
func NewSExt(from Value, to Type) *InstSExt {
	return &InstSExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSExt) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'sext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstSExt) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "sext %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPTrunc is an LLVM IR fptrunc instruction.
type InstFPTrunc struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewFPTrunc returns a new fptrunc instruction based on the given source value
// and target type.
func NewFPTrunc(from Value, to Type) *InstFPTrunc {
	return &InstFPTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPTrunc) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPTrunc) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fptrunc' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstFPTrunc) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptrunc %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPExt is an LLVM IR fpext instruction.
type InstFPExt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewFPExt returns a new fpext instruction based on the given source value and
// target type.
func NewFPExt(from Value, to Type) *InstFPExt {
	return &InstFPExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPExt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPExt) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fpext' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstFPExt) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fpext %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToUI is an LLVM IR fptoui instruction.
type InstFPToUI struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewFPToUI returns a new fptoui instruction based on the given source value
// and target type.
func NewFPToUI(from Value, to Type) *InstFPToUI {
	return &InstFPToUI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToUI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToUI) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fptoui' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstFPToUI) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptoui %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFPToSI is an LLVM IR fptosi instruction.
type InstFPToSI struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewFPToSI returns a new fptosi instruction based on the given source value
// and target type.
func NewFPToSI(from Value, to Type) *InstFPToSI {
	return &InstFPToSI{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstFPToSI) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstFPToSI) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fptosi' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstFPToSI) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "fptosi %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstUIToFP is an LLVM IR uitofp instruction.
type InstUIToFP struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewUIToFP returns a new uitofp instruction based on the given source value
// and target type.
func NewUIToFP(from Value, to Type) *InstUIToFP {
	return &InstUIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstUIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstUIToFP) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'uitofp' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstUIToFP) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "uitofp %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstSIToFP is an LLVM IR sitofp instruction.
type InstSIToFP struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewSIToFP returns a new sitofp instruction based on the given source value
// and target type.
func NewSIToFP(from Value, to Type) *InstSIToFP {
	return &InstSIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstSIToFP) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstSIToFP) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'sitofp' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstSIToFP) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "sitofp %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstPtrToInt is an LLVM IR ptrtoint instruction.
type InstPtrToInt struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewPtrToInt returns a new ptrtoint instruction based on the given source
// value and target type.
func NewPtrToInt(from Value, to Type) *InstPtrToInt {
	return &InstPtrToInt{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstPtrToInt) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstPtrToInt) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'ptrtoint' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstPtrToInt) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "ptrtoint %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstIntToPtr is an LLVM IR inttoptr instruction.
type InstIntToPtr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewIntToPtr returns a new inttoptr instruction based on the given source
// value and target type.
func NewIntToPtr(from Value, to Type) *InstIntToPtr {
	return &InstIntToPtr{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstIntToPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstIntToPtr) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'inttoptr' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstIntToPtr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "inttoptr %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstBitCast is an LLVM IR bitcast instruction.
type InstBitCast struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewBitCast returns a new bitcast instruction based on the given source value
// and target type.
func NewBitCast(from Value, to Type) *InstBitCast {
	return &InstBitCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstBitCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstBitCast) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'bitcast' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstBitCast) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "bitcast %s to %s", inst.From, inst.To)
	return buf.String()
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAddrSpaceCast is an LLVM IR addrspacecast instruction.
type InstAddrSpaceCast struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Value before conversion.
	From Value
	// Type after conversion.
	To Type

	// extra.
}

// NewAddrSpaceCast returns a new addrspacecast instruction based on the given
// source value and target type.
func NewAddrSpaceCast(from Value, to Type) *InstAddrSpaceCast {
	return &InstAddrSpaceCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAddrSpaceCast) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAddrSpaceCast) Type() Type {
	return inst.To
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'addrspacecast' From=TypeValue 'to' To=Type Metadata=(',' MetadataAttachment)+?
func (inst *InstAddrSpaceCast) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "addrspacecast %s to %s", inst.From, inst.To)
	return buf.String()
}
