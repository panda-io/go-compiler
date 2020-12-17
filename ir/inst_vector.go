package ir

import (
	"fmt"
	"strings"
)

// --- [ Vector instructions ] -------------------------------------------------

// ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractElement is an LLVM IR extractelement instruction.
type InstExtractElement struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vector.
	X Value
	// Element index.
	Index Value

	// extra.

	// Type of result produced by the instruction.
	Typ Type
}

// NewExtractElement returns a new extractelement instruction based on the given
// vector and element index.
func NewExtractElement(x, index Value) *InstExtractElement {
	inst := &InstExtractElement{X: x, Index: index}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstExtractElement) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstExtractElement) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.X.Type().(*VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *VectorType, got %T", inst.X.Type()))
		}
		inst.Typ = t.ElemType
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'extractelement' X=TypeValue ',' Index=TypeValue Metadata=(',' MetadataAttachment)+?
func (inst *InstExtractElement) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "extractelement %s, %s", inst.X, inst.Index)
	return buf.String()
}

// ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertElement is an LLVM IR insertelement instruction.
type InstInsertElement struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vector.
	X Value
	// Element to insert.
	Elem Value
	// Element index.
	Index Value

	// extra.

	// Type of result produced by the instruction.
	Typ *VectorType
}

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and element index.
func NewInsertElement(x, elem, index Value) *InstInsertElement {
	inst := &InstInsertElement{X: x, Elem: elem, Index: index}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstInsertElement) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstInsertElement) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		t, ok := inst.X.Type().(*VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *VectorType, got %T", inst.X.Type()))
		}
		inst.Typ = t
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'insertelement' X=TypeValue ',' Elem=TypeValue ',' Index=TypeValue
// Metadata=(',' MetadataAttachment)+?
func (inst *InstInsertElement) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "insertelement %s, %s, %s", inst.X, inst.Elem, inst.Index)
	return buf.String()
}

// ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstShuffleVector is an LLVM IR shufflevector instruction.
type InstShuffleVector struct {
	// Name of local variable associated with the result.
	LocalIdent
	// Vectors.
	X, Y Value
	// Shuffle mask.
	Mask Value

	// extra.

	// Type of result produced by the instruction.
	Typ *VectorType
}

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask Value) *InstShuffleVector {
	inst := &InstShuffleVector{X: x, Y: y, Mask: mask}
	// Compute type.
	inst.Type()
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstShuffleVector) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstShuffleVector) Type() Type {
	// Cache type if not present.
	if inst.Typ == nil {
		xType, ok := inst.X.Type().(*VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *VectorType, got %T", inst.X.Type()))
		}
		maskType, ok := inst.Mask.Type().(*VectorType)
		if !ok {
			panic(fmt.Errorf("invalid vector type; expected *VectorType, got %T", inst.Mask.Type()))
		}
		inst.Typ = NewVectorType(maskType.Len, xType.ElemType)
	}
	return inst.Typ
}

// LLString returns the LLVM syntax representation of the instruction.
//
// 'shufflevector' X=TypeValue ',' Y=TypeValue ',' Mask=TypeValue
// Metadata=(',' MetadataAttachment)+?
func (inst *InstShuffleVector) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "shufflevector %s, %s, %s", inst.X, inst.Y, inst.Mask)
	return buf.String()
}
