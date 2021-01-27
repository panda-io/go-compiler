package ir

import (
	"fmt"
	"strings"
)

// === [ Global variables ] ====================================================

// Global is a global variable declaration or definition.
type Global struct {
	// Global variable name (without '@' prefix).
	GlobalIdent
	// Immutability of global variable (constant or global).
	Immutable bool
	// Content type.
	ContentType Type
	// Initial value; or nil if declaration.
	Init Constant
	// Pointer type to global variable, including an optional address space. If
	// Typ is nil, the first invocation of Type stores a pointer type with
	// ContentType as element.
	Typ *PointerType
}

// NewGlobal returns a new global variable declaration based on the given global
// variable name and content type.
func NewGlobal(name string, contentType Type) *Global {
	global := &Global{ContentType: contentType}
	global.SetName(name)
	// Compute type.
	global.Type()
	return global
}

// NewGlobalDef returns a new global variable definition based on the given
// global variable name and initial value.
func NewGlobalDef(name string, init Constant) *Global {
	global := &Global{ContentType: init.Type(), Init: init}
	global.SetName(name)
	// Compute type.
	global.Type()
	return global
}

// String returns the LLVM syntax representation of the global variable as a
// type-value pair.
func (g *Global) String() string {
	return fmt.Sprintf("%s %s", g.Type(), g.Ident())
}

// Type returns the type of the global variable.
func (g *Global) Type() Type {
	// Cache type if not present.
	if g.Typ == nil {
		g.Typ = NewPointerType(g.ContentType)
	}
	return g.Typ
}

// LLString returns the LLVM syntax representation of the global variable
// definition or declaration.
//
// Global declaration.
func (g *Global) LLString() string {
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s =", g.Ident())
	if g.Immutable {
		buf.WriteString(" constant")
	} else {
		buf.WriteString(" global")
	}
	fmt.Fprintf(buf, " %s", g.ContentType)
	if g.Init != nil {
		// Global definition.
		fmt.Fprintf(buf, " %s", g.Init.Ident())
	}
	return buf.String()
}
