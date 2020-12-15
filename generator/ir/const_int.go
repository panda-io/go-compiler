package ir

import (
	"fmt"
	"math/big"
)

// --- [ Integer constants ] ---------------------------------------------------

// ConstInt is an LLVM IR integer constant.
type Int struct {
	// Integer type.
	Typ *IntType
	// Integer constant.
	X *big.Int
}

// NewInt returns a new integer constant based on the given integer type and
// 64-bit interger value.
func NewInt(typ *IntType, x int64) *Int {
	return &Int{Typ: typ, X: big.NewInt(x)}
}

// NewBool returns a new boolean constant based on the given boolean value.
func NewBool(x bool) *Int {
	if x {
		return True
	}
	return False
}

// NewIntFromString returns a new integer constant based on the given integer
// type and string.
//
// The integer string may be expressed in one of the following forms.
//
//    * boolean literal
//         true | false
//    * integer literal
//         [-]?[0-9]+
//    * hexadecimal integer literal
//         0x[0-9A-Fa-f]+
//    * binary integer literal
//         0b[01]+
//    * oct integer literal
//         0o[0-7]
func NewIntFromString(typ *IntType, s string) (*Int, error) {
	// Boolean literal.
	switch s {
	case "true":
		if !typ.Equal(I1) {
			return nil, fmt.Errorf("invalid boolean type; expected i1, got %T", typ)
		}
		return True, nil
	case "false":
		if !typ.Equal(I1) {
			return nil, fmt.Errorf("invalid boolean type; expected i1, got %T", typ)
		}
		return False, nil
	}
	x, _ := (&big.Int{}).SetString(s, 0)
	if x == nil {
		return nil, fmt.Errorf("unable to parse integer constant %q", s)
	}
	return &Int{Typ: typ, X: x}, nil
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Int) String() string {
	return fmt.Sprintf("%v %v", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Int) Type() Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (c *Int) Ident() string {
	if c.Typ.BitSize == 1 {
		switch c.X.Int64() {
		case 0:
			return "false"
		case 1:
			return "true"
		default:
			panic(fmt.Errorf("invalid integer value of boolean type; expected 0 or 1, got %d", c.X.Int64()))
		}
	}
	return c.X.String()
}
