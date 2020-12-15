package ir

import (
	"fmt"
)

// --- [ Undefined values ] ----------------------------------------------------

// Undef is an LLVM IR undefined value.
type Undef struct {
	// Undefined value type.
	Typ Type
}

// NewUndef returns a new undefined value based on the given type.
func NewUndef(typ Type) *Undef {
	return &Undef{Typ: typ}
}

// String returns the LLVM syntax representation of the constant as a type-value
// pair.
func (c *Undef) String() string {
	return fmt.Sprintf("%s %s", c.Type(), c.Ident())
}

// Type returns the type of the constant.
func (c *Undef) Type() Type {
	return c.Typ
}

// Ident returns the identifier associated with the constant.
func (*Undef) Ident() string {
	// 'undef'
	return "undef"
}
