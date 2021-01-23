// Package gep computes the result type of getelementptr instructions and
// constant expressions.
//
// ref: https://llvm.org/docs/GetElementPtr.html
// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
package ir

import (
	"fmt"
	"log"
)

// Index is a gep index.
type GepIndex struct {
	// HasVal specifies whether Val has a valid value. If index is a constant
	// integer or a constant integer vector of which all elements have the same
	// value, then HasVal is set. Note, this is a requirement to index into
	// structure
	HasVal bool
	// Index integer value. Val is only valid if HasVal is set.
	Val int64
	// Length of index vector; or 0 if index is scalar. VectorLen may be non-zero
	// even if HasVal is false.
	VectorLen uint64
}

// NewIndex returns a new constant index with the given value.
func NewGepIndex(val int64) GepIndex {
	return GepIndex{
		HasVal: true,
		Val:    val,
	}
}

// ResultType computes the result type of a getelementptr instruction or
// constant expression.
//
//    getelementptr (ElemType, Src, Indices)
func ResultType(elemType Type, indices []GepIndex) Type {
	// ref: http://llvm.org/docs/GetElementPtr.html#what-effect-do-address-spaces-have-on-geps
	//
	// > the address space qualifier on the second operand pointer type always
	// > matches the address space qualifier on the result type.
	var resultVectorLength uint64

	// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
	//
	// > The first argument is always a type used as the basis for the
	// > calculations.
	e := elemType
	for i, index := range indices {
		// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
		//
		// > The getelementptr returns a vector of pointers, instead of a single
		// > address, when one or more of its arguments is a vector. In such
		// > cases, all vector arguments should have the same number of elements,
		// > and every scalar argument will be effectively broadcast into a vector
		// > during address calculation.
		if index.VectorLen != 0 && resultVectorLength != 0 && index.VectorLen != resultVectorLength {
			panic(fmt.Errorf("vector length mismatch of index vector (%d) and result type vector (%d)", index.VectorLen, resultVectorLength))
		}
		if resultVectorLength == 0 && index.VectorLen != 0 {
			resultVectorLength = index.VectorLen
		}
		// ref: https://llvm.org/docs/GetElementPtr.html#why-is-the-extra-0-index-required
		//
		// > Since the second argument to the GEP instruction must always be a
		// > value of pointer type, the first index steps through that pointer.
		if i == 0 {
			continue
		}
		switch elm := e.(type) {
		case *PointerType:
			panic(fmt.Errorf("cannot index into pointer type at %d:th gep index, only valid at 0:th gep index; see https://llvm.org/docs/GetElementPtr.html#what-is-dereferenced-by-gep", i))
		case *VectorType:
			// ref: https://llvm.org/docs/GetElementPtr.html#can-gep-index-into-vector-elements
			//
			// > This hasn’t always been forcefully disallowed, though it’s not
			// > recommended. It leads to awkward special cases in the optimizers,
			// > and fundamental inconsistency in the IR. In the future, it will
			// > probably be outright disallowed.
			log.Printf("using gep to index into vector types will be disallowed in a future release or llir/llvm; see https://llvm.org/docs/GetElementPtr.html#can-gep-index-into-vector-elements")
			e = elm.ElemType
		case *ArrayType:
			e = elm.ElemType
		case *StructType:
			// ref: https://llvm.org/docs/LangRef.html#getelementptr-instruction
			//
			// > When indexing into a (optionally packed) structure, only i32
			// > integer constants are allowed (when using a vector of indices they
			// > must all be the same i32 integer constant).
			if !index.HasVal {
				panic(fmt.Errorf("unable to index into struct type `%v` using gep with non-constant index", e))
			}
			e = elm.Fields[index.Val]
		default:
			panic(fmt.Errorf("cannot index into type %T using gep", e))
		}
	}
	ptr := NewPointerType(e)
	if resultVectorLength != 0 {
		vec := NewVectorType(resultVectorLength, ptr)
		return vec
	}
	return ptr
}
