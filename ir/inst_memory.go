package ir

import (
	"fmt"
	"strings"
)

// --- [ Memory instructions ] -------------------------------------------------

// ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAlloca is an LLVM IR alloca instruction.
type InstAlloca struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Element type.
	ElemType Type
	// Type of result produced by the instruction, including an optional address
	// space.
	Typ *PointerType
	// Save reference type for alloca type
	UserData string
}

// NewAlloca returns a new alloca instruction based on the given element type.
func NewAlloca(elemType Type) *InstAlloca {
	inst := &InstAlloca{ElemType: elemType}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAlloca) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAlloca) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = NewPointerType(inst.ElemType)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'alloca' InAllocaopt SwiftErroropt ElemType=Type NElems=(',' TypeValue)? (',' Align)? (',' AddrSpace)? Metadata=(',' MetadataAttachment)+?
func (inst *InstAlloca) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("alloca")
	fmt.Fprintf(buf, " %s", inst.ElemType)
	return buf.String()
}

// ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstLoad is an LLVM IR load instruction.
type InstLoad struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Element type of src.
	ElemType Type
	// Source address.
	Src Value

	// extra.
}

// NewLoad returns a new load instruction based on the given element type and
// source address.
func NewLoad(elemType Type, src Value) *InstLoad {
	inst := &InstLoad{ElemType: elemType, Src: src}
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstLoad) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstLoad) Type() Type {
	return inst.ElemType
}

// LLString returns the LLVM syntax representation of the instruction.
//
// Load instruction.
//
//    'load' Volatileopt ElemType=Type ',' Src=TypeValue (',' Align)? Metadata=(',' MetadataAttachment)+?
//
// Atomic load instruction.
//
//    'load' Atomic Volatileopt ElemType=Type ',' Src=TypeValue SyncScopeopt Ordering=AtomicOrdering (',' Align)? Metadata=(',' MetadataAttachment)+?
func (inst *InstLoad) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("load")
	fmt.Fprintf(buf, " %s, %s", inst.ElemType, inst.Src)
	return buf.String()
}

// ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstStore is an LLVM IR store instruction.
type InstStore struct {
	// Source
	Src Value
	// Destination address.
	Dst Value

	// extra.
}

// NewStore returns a new store instruction based on the given source value and
// destination address.
func NewStore(src, dst Value) *InstStore {
	// Type-check operands.
	dstPtrType, ok := dst.Type().(*PointerType)
	if !ok {
		panic(fmt.Errorf("invalid store dst operand type; expected *Pointer, got %T", dst.Type()))
	}
	if !src.Type().Equal(dstPtrType.ElemType) {
		panic(fmt.Errorf("store operands are not compatible: src=%v; dst=%v", src.Type(), dst.Type()))
	}
	return &InstStore{Src: src, Dst: dst}
}

// LLString returns the LLVM syntax representation of the instruction.
//
// Store instruction.
//
//    'store' Volatileopt Src=TypeValue ',' Dst=TypeValue (',' Align)? Metadata=(',' MetadataAttachment)+?
//
// Atomic store instruction.
//
//    'store' Atomic Volatileopt Src=TypeValue ',' Dst=TypeValue SyncScopeopt Ordering=AtomicOrdering (',' Align)? Metadata=(',' MetadataAttachment)+?
func (inst *InstStore) LLString() string {
	buf := &strings.Builder{}
	buf.WriteString("store")
	fmt.Fprintf(buf, " %s, %s", inst.Src, inst.Dst)
	return buf.String()
}

// ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstFence is an LLVM IR fence instruction.
type InstFence struct {
	// Atomic memory ordering constraints.
	Ordering AtomicOrdering

	// extra.
}

// NewFence returns a new fence instruction based on the given atomic ordering.
func NewFence(ordering AtomicOrdering) *InstFence {
	return &InstFence{Ordering: ordering}
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'fence' SyncScopeopt Ordering=AtomicOrdering Metadata=(',' MetadataAttachment)+?
func (inst *InstFence) LLString() string {
	buf := &strings.Builder{}
	buf.WriteString("fence")
	fmt.Fprintf(buf, " %d", inst.Ordering)
	return buf.String()
}

// ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstCmpXchg is an LLVM IR cmpxchg instruction.
type InstCmpXchg struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Address to read from, compare against and store to.
	Ptr Value
	// Value to compare against.
	Cmp Value
	// New value to store.
	New Value
	// Atomic memory ordering constraints on success.
	SuccessOrdering AtomicOrdering
	// Atomic memory ordering constraints on failure.
	FailureOrdering AtomicOrdering

	// extra.

	// Type of result produced by the instruction; the first field of the struct
	// holds the old value, and the second field indicates success.
	Typ *StructType
}

// NewCmpXchg returns a new cmpxchg instruction based on the given address,
// value to compare against, new value to store, and atomic orderings for
// success and failure.
func NewCmpXchg(ptr, cmp, new Value, successOrdering, failureOrdering AtomicOrdering) *InstCmpXchg {
	inst := &InstCmpXchg{Ptr: ptr, Cmp: cmp, New: new, SuccessOrdering: successOrdering, FailureOrdering: failureOrdering}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstCmpXchg) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction. The result type is a struct type
// with two fields, the first field has the type of the old value and the second
// field has boolean type.
func (inst *InstCmpXchg) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		oldType := inst.New.Type()
		inst.Typ = NewStructType(oldType, I1)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'cmpxchg' Weakopt Volatileopt Ptr=TypeValue ',' Cmp=TypeValue ',' New=TypeValue SyncScopeopt SuccessOrdering=AtomicOrdering FailureOrdering=AtomicOrdering Metadata=(',' MetadataAttachment)+?
func (inst *InstCmpXchg) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("cmpxchg")
	fmt.Fprintf(buf, " %s, %s, %s", inst.Ptr, inst.Cmp, inst.New)
	fmt.Fprintf(buf, " %d", inst.SuccessOrdering)
	fmt.Fprintf(buf, " %d", inst.FailureOrdering)
	return buf.String()
}

// ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstAtomicRMW is an LLVM IR atomicrmw instruction.
type InstAtomicRMW struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Atomic operation.
	Op AtomicOp
	// Destination address.
	Dst Value
	// Operand.
	X Value
	// Atomic memory ordering constraints.
	Ordering AtomicOrdering

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewAtomicRMW returns a new atomicrmw instruction based on the given atomic
// operation, destination address, operand and atomic ordering.
func NewAtomicRMW(op AtomicOp, dst, x Value, ordering AtomicOrdering) *InstAtomicRMW {
	inst := &InstAtomicRMW{Op: op, Dst: dst, X: x, Ordering: ordering}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstAtomicRMW) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstAtomicRMW) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.Dst.Type().(*PointerType)
		if !ok {
			panic(fmt.Errorf("invalid destination type; expected *PointerType, got %T", inst.Dst.Type()))
		}
		inst.Typ = t.ElemType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'atomicrmw' Volatileopt Op=AtomicOp Dst=TypeValue ',' X=TypeValue SyncScopeopt Ordering=AtomicOrdering Metadata=(',' MetadataAttachment)+?
func (inst *InstAtomicRMW) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("atomicrmw")
	fmt.Fprintf(buf, " %d %s, %s", inst.Op, inst.Dst, inst.X)
	fmt.Fprintf(buf, " %d", inst.Ordering)
	return buf.String()
}

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstGetElementPtr is an LLVM IR getelementptr instruction.
type InstGetElementPtr struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Element type.
	ElemType Type
	// Source address.
	Src Value
	// Element indicies.
	Indices []Value

	// extra.

	// Type of result produced by the instruction.
	Typ Type // *PointerType or *VectorType (with elements of pointer type)
}

// NewGetElementPtr returns a new getelementptr instruction based on the given
// element type, source address and element indices.
func NewGetElementPtr(elemType Type, src Value, indices ...Value) *InstGetElementPtr {
	inst := &InstGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstGetElementPtr) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = GepInstType(inst.ElemType, inst.Indices)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'getelementptr' InBoundsopt ElemType=Type ',' Src=TypeValue Indices=(',' TypeValue)* Metadata=(',' MetadataAttachment)+?
func (inst *InstGetElementPtr) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	buf.WriteString("getelementptr")
	fmt.Fprintf(buf, " %s, %s", inst.ElemType, inst.Src)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %s", index)
	}
	return buf.String()
}

// ### [ Helper functions ] ####################################################

// gepInstType computes the result type of a getelementptr instruction.
//
//    getelementptr ElemType, Indices
func GepInstType(elemType Type, indices []Value) Type {
	var idxs []GepIndex
	for _, index := range indices {
		var idx GepIndex
		switch index := index.(type) {
		case Constant:
			idx = getIndex(index)
		default:
			idx = GepIndex{HasVal: false}
			// Check if index is of vector type.
			if indexType, ok := index.Type().(*VectorType); ok {
				idx.VectorLen = indexType.Len
			}
		}
		idxs = append(idxs, idx)
	}
	return ResultType(elemType, idxs)
}

// NOTE: keep getIndex in sync with getIndex in:
//
//    * ast/inst_memory.go
//    * ir/inst_memory.go
//    * ir/constant/expr_memory.go
//
// The reference point and source of truth is in ir/constant/expr_memory.go.

// getIndex returns the gep index corresponding to the given constant index.
func getIndex(index Constant) GepIndex {
	// unpack inrange indices.
	if idx, ok := index.(*Index); ok {
		index = idx.Index
	}
	// Use index.Simplify() to simplify the constant expression to a concrete
	// integer constant or vector of integers
	if idx, ok := index.(Expression); ok {
		index = idx.Simplify()
	}
	switch index := index.(type) {
	case *Int:
		val := index.X.Int64()
		return NewGepIndex(val)
	case *ZeroInitializer:
		return NewGepIndex(0)
	case *Vector:
		// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
		//
		// > The getelementptr returns a vector of pointers, instead of a single
		// > address, when one or more of its arguments is a vector. In such
		// > cases, all vector arguments should have the same number of elements,
		// > and every scalar argument will be effectively broadcast into a vector
		// > during address calculation.
		if len(index.Elems) == 0 {
			return GepIndex{HasVal: false}
		}
		// Sanity check. All vector elements must be integers, and must have the
		// same
		var val int64
		for i, elem := range index.Elems {
			switch elem := elem.(type) {
			case *Int:
				x := elem.X.Int64()
				if i == 0 {
					val = x
				} else if x != val {
					// since all elements were not identical, we must conclude that
					// the index vector does not have a concrete
					return GepIndex{
						HasVal:    false,
						VectorLen: uint64(len(index.Elems)),
					}
				}
			default:
				// TODO: remove debug output.
				panic(fmt.Errorf("support for gep index vector element type %T not yet implemented", elem))
				//return gep.Index{HasVal: false}
			}
		}
		return GepIndex{
			HasVal:    true,
			Val:       val,
			VectorLen: uint64(len(index.Elems)),
		}
	case *Undef:
		return GepIndex{HasVal: false}
	case Expression:
		// should already have been simplified to a form we can handle.
		return GepIndex{HasVal: false}
	default:
		// TODO: add support for more constant expressions.
		// TODO: remove debug output.
		panic(fmt.Errorf("support for gep index type %T not yet implemented", index))
	}
}
