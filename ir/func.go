package ir

import (
	"fmt"
	"strings"
)

// === [ Functions ] ===========================================================

// Func is an LLVM IR function. The body of a function definition consists of a
// set of basic blocks, interconnected by terminator control flow instructions.
type Func struct {
	// Function name (without '@' prefix).
	GlobalIdent
	// Function signature.
	Sig *FuncType
	// Function parameters.
	Params []*Param
	// Basic blocks.
	Blocks []*Block // nil if declaration.
	// Pointer type to function, including an optional address space. If Typ is
	// nil, the first invocation of Type stores a pointer type with Sig as
	// element.
	Typ *PointerType
}

// NewFunc returns a new function based on the given function name, return type
// and function parameters.
func NewFunc(name string, retType Type, params ...*Param) *Func {
	paramTypes := make([]Type, len(params))
	for i, param := range params {
		paramTypes[i] = param.Type()
	}
	sig := NewFuncType(retType, paramTypes...)
	f := &Func{Sig: sig, Params: params}
	f.SetName(name)
	// Compute type.
	f.Type()
	return f
}

// NewBlock appends a new basic block to the function based on the given label
// name. An empty label name indicates an unnamed basic block.
//
// The Parent field of the block is set to f.
func (f *Func) NewBlock(name string) *Block {
	block := NewBlock(name)
	f.Blocks = append(f.Blocks, block)
	return block
}

// String returns the LLVM syntax representation of the function as a type-value
// pair.
func (f *Func) String() string {
	return fmt.Sprintf("%s %s", f.Type(), f.Ident())
}

// Type returns the type of the function.
func (f *Func) Type() Type {
	// Cache type if not present.
	if f.Typ == nil {
		f.Typ = NewPointerType(f.Sig)
	}
	return f.Typ
}

// LLString returns the LLVM syntax representation of the function definition or
// declaration.
//
// Function declaration.
//
//    'declare' Metadata=MetadataAttachment* Header=FuncHeader
//
// Function definition.
//
//    'define' Header=FuncHeader Metadata=MetadataAttachment* Body=FuncBody
func (f *Func) LLString() string {
	if err := f.AssignIDs(); err != nil {
		panic(fmt.Errorf("unable to assign IDs of function %q; %v", f.Ident(), err))
	}
	buf := &strings.Builder{}
	if len(f.Blocks) == 0 {
		// Function declaration.
		buf.WriteString("declare")
		buf.WriteString(headerString(f))
		return buf.String()
	} else {
		// Function definition.
		buf.WriteString("define")
		buf.WriteString(headerString(f))
		fmt.Fprintf(buf, " %s", bodyString(f))
		return buf.String()
	}
}

// AssignIDs assigns IDs to unnamed local variables.
func (f *Func) AssignIDs() error {
	id := int64(0)
	for _, param := range f.Params {
		// Assign local IDs to unnamed parameters of function definitions.
		if err := setName(param, &id); err != nil {
			return err
		}
	}
	for _, block := range f.Blocks {
		// Assign local IDs to unnamed basic blocks.
		if err := setName(block, &id); err != nil {
			return err
		}
		for _, inst := range block.Insts {
			n, ok := inst.(Ident)
			if !ok {
				continue
			}
			// Skip void instructions (call with void return).
			if Equal(n.Type(), Void) {
				continue
			}
			// Assign local IDs to unnamed local variables.
			if err := setName(n, &id); err != nil {
				return err
			}
		}
	}
	return nil
}

// ### [ Helper functions ] ####################################################

// headerString returns the string representation of the function header.
func headerString(f *Func) string {
	// (Linkage | ExternLinkage)? Preemptionopt Visibilityopt DLLStorageClassopt CallingConvopt ReturnAttrs=ReturnAttribute* RetType=Type Name=GlobalIdent '(' Params ')' UnnamedAddropt AddrSpaceopt FuncAttrs=FuncAttribute* Sectionopt Partitionopt Comdatopt Alignopt GCopt Prefixopt Prologueopt Personalityopt
	buf := &strings.Builder{}
	fmt.Fprintf(buf, " %s", f.Sig.RetType)
	fmt.Fprintf(buf, " %s(", f.Ident())
	for i, param := range f.Params {
		if i != 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(param.LLString())
	}
	if f.Sig.Variadic {
		if len(f.Params) > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString("...")
	}
	buf.WriteString(")")
	return buf.String()
}

// bodyString returns the string representation of the function body.
func bodyString(body *Func) string {
	// '{' Blocks=Block+ UseListOrders=UseListOrder* '}'
	buf := &strings.Builder{}
	buf.WriteString("{\n")
	for i, block := range body.Blocks {
		if i != 0 {
			buf.WriteString("\n")
		}
		fmt.Fprintf(buf, "%s\n", block.LLString())
	}
	buf.WriteString("}")
	return buf.String()
}

func setName(name Ident, id *int64) error {
	if name.IsUnnamed() {
		if name.ID() != 0 && *id != name.ID() {
			want := *id
			got := name.ID()
			return fmt.Errorf("invalid local ID, expected %s, got %s", LocalID(want), LocalID(got))
		}
		name.SetID(*id)
		(*id)++
	}
	return nil
}

// ___ [ Function parameter ] __________________________________________________

// Param is an LLVM IR function parameter.
type Param struct {
	// (optional) Parameter name (without '%' prefix).
	LocalIdent
	// Parameter type.
	Typ Type
}

// NewParam returns a new function parameter based on the given name and type.
func NewParam(typ Type) *Param {
	return &Param{
		Typ: typ,
	}
}

// String returns the LLVM syntax representation of the function parameter as a
// type-value pair.
func (p *Param) String() string {
	return fmt.Sprintf("%s %s", p.Type(), p.Ident())
}

// Type returns the type of the function parameter.
func (p *Param) Type() Type {
	return p.Typ
}

// LLString returns the LLVM syntax representation of the function parameter.
//
// Typ=Type Attrs=ParamAttribute* Name=LocalIdent?
func (p *Param) LLString() string {
	buf := &strings.Builder{}
	buf.WriteString(p.Typ.String())
	fmt.Fprintf(buf, " %s", p.Ident())
	return buf.String()
}
