package ir

import (
	"fmt"
	"strings"
)

// --- [ Aggregate expressions ] -----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprExtractValue is an LLVM IR extractvalue expression.
type ExprExtractValue struct {
	// Aggregate value.
	X Constant
	// Element indices.
	Indices []uint64

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewExprExtractValue returns a new extractvalue expression based on the given
// aggregate value and indicies.
func NewExprExtractValue(x Constant, indices ...uint64) *ExprExtractValue {
	e := &ExprExtractValue{X: x, Indices: indices}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprExtractValue) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprExtractValue) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = aggregateElemType(e.X.Type(), e.Indices)
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprExtractValue) Ident() string {
	// 'extractvalue' '(' X=TypeConst Indices=(',' UintLit)* ')'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "extractvalue (%s", e.X)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprExtractValue) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprInsertValue is an LLVM IR insertvalue expression.
type ExprInsertValue struct {
	// Aggregate value.
	X Constant
	// Element to insert.
	Elem Constant
	// Element indices.
	Indices []uint64

	// extra.

	// Type of result produced by the constant expression.
	Typ Type
}

// NewInsertValue returns a new insertvalue expression based on the given
// aggregate value, element and indicies.
func NewExprInsertValue(x, elem Constant, indices ...uint64) *ExprInsertValue {
	e := &ExprInsertValue{X: x, Elem: elem, Indices: indices}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprInsertValue) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprInsertValue) Type() Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprInsertValue) Ident() string {
	// 'insertvalue' '(' X=TypeConst ',' Elem=TypeConst Indices=(',' UintLit)*
	// ')'
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "insertvalue (%s, %s", e.X, e.Elem)
	for _, index := range e.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	buf.WriteString(")")
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprInsertValue) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}
