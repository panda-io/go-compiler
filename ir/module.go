package ir

import (
	"io"
)

// === [ Modules ] ===

// Module is an LLVM IR module, which consists of top-level declarations and
// definitions.
type Module struct {
	// Type definitions.
	TypeDefs []Type
	// Global variable declarations and definitions.
	Globals []*Global
	// Function declarations and definitions.
	Funcs []*Func
}

// NewModule returns a new LLVM IR module.
func NewModule() *Module {
	return &Module{}
}

// WriteTo write the string representation of the module in LLVM IR assembly
// syntax to w.
func (m *Module) WriteTo(w io.Writer) (n int64, err error) {
	fw := &FmtWriter{w: w}

	// Type definitions.
	if len(m.TypeDefs) > 0 && fw.size > 0 {
		fw.Fprint("\n")
	}
	for _, t := range m.TypeDefs {
		// Name=LocalIdent '=' 'type' Typ=OpaqueType
		//
		// Name=LocalIdent '=' 'type' Typ=Type
		fw.Fprintf("%s = type %s\n", t, t.LLString())
	}
	// Global declarations and definitions.
	if len(m.Globals) > 0 && fw.size > 0 {
		fw.Fprint("\n")
	}
	for _, g := range m.Globals {
		fw.Fprintln(g.LLString())
	}
	// Function declarations and definitions.
	if len(m.Funcs) > 0 && fw.size > 0 {
		fw.Fprint("\n")
	}
	for i, f := range m.Funcs {
		if i != 0 {
			fw.Fprint("\n")
		}
		fw.Fprintln(f.LLString())
	}
	return fw.size, fw.err
}

// --- [ Functions ] -----------------------------------------------------------

// NewFunc appends a new function to the module based on the given function
// name, return type and function parameters.
//
// The Parent field of the function is set to m.
func (m *Module) NewFunc(name string, retType Type, params ...*Param) *Func {
	f := NewFunc(name, retType, params...)
	m.Funcs = append(m.Funcs, f)
	return f
}

// --- [ Global variables ] ----------------------------------------------------

// NewGlobal appends a new global variable declaration to the module based on
// the given global variable name and content type.
func (m *Module) NewGlobal(name string, contentType Type) *Global {
	g := NewGlobal(name, contentType)
	m.Globals = append(m.Globals, g)
	return g
}

// NewGlobalDef appends a new global variable definition to the module based on
// the given global variable name and initial value.
func (m *Module) NewGlobalDef(name string, init Constant) *Global {
	g := NewGlobalDef(name, init)
	m.Globals = append(m.Globals, g)
	return g
}

// --- [ Type definitions ] ----------------------------------------------------

// NewTypeDef appends a new type definition to the module based on the given
// type name and underlying type.
func (m *Module) NewTypeDef(name string, typ Type) Type {
	typ.SetName(name)
	m.TypeDefs = append(m.TypeDefs, typ)
	return typ
}
