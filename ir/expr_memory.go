package ir

import (
	"fmt"
	"strings"
)

// --- [ Memory expressions ] --------------------------------------------------

// ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprGetElementPtr is an LLVM IR getelementptr expression.
type ExprGetElementPtr struct {
	// Element type.
	ElemType Type
	// Source address.
	Src Constant
	// Element indicies.
	Indices []Constant // *Int, *Vector or *Index

	// extra.

	// Type of result produced by the constant expression.
	Typ Type // *PointerType or *VectorType (with elements of pointer type)
}

// NewGetElementPtr returns a new getelementptr expression based on the given
// element type, source address and element indices.
func NewExprGetElementPtr(elemType Type, src Constant, indices ...Constant) *ExprGetElementPtr {
	e := &ExprGetElementPtr{ElemType: elemType, Src: src, Indices: indices}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprGetElementPtr) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprGetElementPtr) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = gepExprType(e.ElemType, e.Src.Type(), e.Indices)
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprGetElementPtr) Ident() string {
	// 'getelementptr' InBoundsopt '(' ElemType=Type ',' Src=TypeConst
	// Indices=(',' GEPIndex)* ')'
	buf := &strings.Builder{}
	buf.WriteString("getelementptr")
	fmt.Fprintf(buf, " (%s, %s", e.ElemType, e.Src)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %s", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprGetElementPtr) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ___ [ gep indices ] _________________________________________________________

// Index is an index of a getelementptr constant expression.
type Index struct {
	// Element index.
	Index Constant

	// extra.
}

// NewIndex returns a new gep element index.
func NewIndex(index Constant) *Index {
	return &Index{Index: index}
}

// String returns a string representation of the getelementptr index.
func (index *Index) String() string {
	// OptInrange Type Constant
	return index.Index.String()
}

// Ident returns a string representation of the getelementptr index.
func (index *Index) Ident() string {
	// OptInrange Type Constant
	return index.Index.Ident()
}

// Ident returns a string representation of the getelementptr index.
func (index *Index) Type() Type {
	// OptInrange Type Constant
	return index.Index.Type()
}

// ### [ Helper functions ] ####################################################

// gepExprType computes the result type of a getelementptr constant expression.
//
//    getelementptr (ElemType, Src, Indices)
func gepExprType(elemType, src Type, indices []Constant) Type {
	var idxs []GepIndex
	for _, index := range indices {
		idx := getIndex(index)
		// Check if index is of vector type.
		if indexType, ok := index.Type().(*VectorType); ok {
			idx.VectorLen = indexType.Len
		}
		idxs = append(idxs, idx)
	}
	return ResultType(elemType, idxs)
}
