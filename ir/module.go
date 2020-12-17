package ir

import (
	"fmt"
	"io"
	"strings"
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

// String returns the string representation of the module in LLVM IR assembly
// syntax.
func (m *Module) String() string {
	buf := &strings.Builder{}
	if _, err := m.WriteTo(buf); err != nil {
		panic(fmt.Errorf("unable to write to string buffer; %v", err))
	}
	return buf.String()
}

// WriteTo write the string representation of the module in LLVM IR assembly
// syntax to w.
func (m *Module) WriteTo(w io.Writer) (n int64, err error) {
	fw := &FmtWriter{w: w}
	/*
		// Assign global IDs.
		if err := m.AssignGlobalIDs(); err != nil {
			panic(fmt.Errorf("unable to assign globals IDs of module; %v", err))
		}
		// Assign metadata IDs.
		if err := m.AssignMetadataIDs(); err != nil {
			panic(fmt.Errorf("unable to assign metadata IDs of module; %v", err))
		}*/
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

/*
// ### [ Helper functions ] ####################################################

// AssignGlobalIDs assigns IDs to unnamed global variables.
func (m *Module) AssignGlobalIDs() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := int64(0)
	setName := func(n namedVar) error {
		if n.IsUnnamed() {
			if n.ID() != 0 && id != n.ID() {
				want := id
				got := n.ID()
				return errors.Errorf("invalid global ID, expected %s, got %s", enc.GlobalID(want), enc.GlobalID(got))
			}
			n.SetID(id)
			id++
		}
		return nil
	}
	// Assign global IDs to unnamed global variables.
	for _, n := range m.Globals {
		if err := setName(n); err != nil {
			return errors.WithStack(err)
		}
	}
	// Assign global IDs to unnamed functions.
	for _, n := range m.Funcs {
		if err := setName(n); err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// AssignMetadataIDs assigns metadata IDs to the unnamed metadata definitions of
// the module.
func (m *Module) AssignMetadataIDs() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Index used IDs.
	used := make(map[int64]bool)
	for _, md := range m.MetadataDefs {
		id := md.ID()
		if id != -1 {
			if _, ok := used[id]; ok {
				return errors.Errorf("metadata ID %s already in use", enc.MetadataID(id))
			}
			used[id] = true
		}
	}
	// nextID returns the next unused metdata ID.
	curID := int64(-1)
	nextID := func() int64 {
		for {
			curID++
			if !used[curID] {
				return curID
			}
		}
	}
	// Assign IDs to unnamed metadata definitions.
	for _, md := range m.MetadataDefs {
		id := md.ID()
		if id != -1 {
			// Metadata definition already has ID.
			continue
		}
		newID := nextID()
		md.SetID(newID)
	}
	return nil
}
*/

// --- [ Functions ] -----------------------------------------------------------

// NewFunc appends a new function to the module based on the given function
// name, return type and function parameters.
//
// The Parent field of the function is set to m.
func (m *Module) NewFunc(name string, retType Type, params ...*Param) *Func {
	f := NewFunc(name, retType, params...)
	f.Parent = m
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
