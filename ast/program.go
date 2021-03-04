package ast

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Program struct {
	Modules  map[string]*Module
	Module   *Module
	IRModule *ir.Module

	Declarations map[string]Declaration
	Strings      map[string]ir.Constant

	Errors []*Error
}

func NewProgram() *Program {
	p := &Program{}
	p.Reset()
	return p
}

func (p *Program) Reset() {
	p.Modules = make(map[string]*Module)
	p.IRModule = ir.NewModule()

	p.Declarations = make(map[string]Declaration)
	p.Strings = make(map[string]ir.Constant)

	p.Errors = p.Errors[:0]
}

func (p *Program) FindSelector(selector, member string) (string, Declaration) {
	if selector == "" || selector == "global" {
		// search current package
		if p.Module.Namespace != Global {
			qualified := p.Module.Namespace + "." + member
			d := p.Declarations[qualified]
			if d != nil {
				return qualified, d
			}
		}
		// search global
		qualified := Global + "." + member
		return qualified, p.Declarations[qualified]
	} else {
		// search imports
		for _, i := range p.Module.Imports {
			if i.Alias == selector {
				qualified := i.Namespace + "." + member
				return qualified, p.Declarations[qualified]
			}
		}
		return "", nil
	}
}

func (p *Program) FindDeclaration(t *TypeName) (string, Declaration) {
	return p.FindSelector(t.Selector, t.Name)
}

func (p *Program) FindQualified(qualified string) Declaration {
	return p.Declarations[qualified]
}

func (p *Program) AddString(value string) ir.Constant {
	bytes := []byte(value)
	bytes = append(bytes, 0)
	hash := fmt.Sprintf("%x", md5.Sum(bytes))
	if v, ok := p.Strings[hash]; ok {
		return v
	}
	s := p.IRModule.NewGlobalDef("string."+hash, ir.NewCharArray(bytes))
	s.Immutable = true
	v := ir.NewExprBitCast(s, pointerType)
	p.Strings[hash] = v
	return v
}

func (p *Program) Error(offset int, message string) {
	p.Errors = append(p.Errors, &Error{
		Position: p.Module.File.Position(offset),
		Message:  message,
	})
}

func (p *Program) GenerateIR() string {
	// zero pass (generate declarations)
	for _, m := range p.Modules {
		// TO-DO check if import is valid // must be valid, cannot import self, cannot duplicated
		p.Module = m

		for _, f := range m.Functions {
			f.GenerateIRDeclaration(p)
		}

		for _, e := range m.Enums {
			e.GenerateIR(p)
		}

		for _, i := range m.Interfaces {
			i.ResolveParents(p)
			//TO-DO something here
		}

		for _, c := range m.Classes {
			c.ResolveParents(p)
			c.PreProcess(p)
			c.GenerateIRDeclaration(p)
		}
	}

	// first pass (resolve oop)
	for _, m := range p.Modules {
		p.Module = m

		for _, c := range m.Classes {
			c.GenerateIRStruct(p)
			c.GenerateIRVTable(p)
		}
	}

	// second pass (generate functions)
	for _, m := range p.Modules {
		p.Module = m

		for _, v := range m.Variables {
			v.GenerateIR(p)
		}

		for _, f := range m.Functions {
			f.GenerateIR(p)
		}

		for _, c := range m.Classes {
			c.GenerateIR(p)
		}
	}

	buf := &strings.Builder{}
	_, err := p.IRModule.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
