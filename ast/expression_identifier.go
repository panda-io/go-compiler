package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type Identifier struct {
	ExpressionBase
	Name string
}

func (i *Identifier) Type(c *Context, expected ir.Type) ir.Type {
	t := c.ObjectType(i.Name)
	if t == nil {
		_, d := c.Program.FindSelector("", i.Name)
		if d == nil {
			c.Program.Error(i.Position, fmt.Sprintf("undefined %s", i.Name))
			return nil
		}
		switch t := d.(type) {
		case *Variable:
			return t.IRVariable.ContentType

		case *Function:
			return t.IRFunction.Sig

		default:
			return nil
		}
	}
	return t
}

func (i *Identifier) GenerateIR(c *Context, expected ir.Type) ir.Value {
	v := c.FindObject(i.Name)
	if v == nil {
		_, d := c.Program.FindSelector("", i.Name)
		if d == nil {
			c.Program.Error(i.Position, fmt.Sprintf("undefined %s", i.Name))
			return nil
		}
		switch t := d.(type) {
		case *Variable:
			return t.IRVariable

		case *Function:
			var call *ir.InstCall
			if t.Class == nil {
				call = ir.NewCall(t.IRFunction)
			} else {
				call = ir.NewCall(t.IRFunction, c.FindObject(ClassThis))
			}
			call.Ref = t.IRFunction
			return call

		default:
			c.Program.Error(i.Position, "invalid type")
			return nil
		}
	}
	return v
}

func (i *Identifier) IsConstant(p *Program) bool {
	_, d := p.FindSelector("", i.Name)
	if d == nil {
		return false
	}
	if v, ok := d.(*Variable); ok {
		return v.Const && v.Value.IsConstant(p)
	}
	if _, ok := d.(*Function); ok {
		return true
	}
	return false
}

func (i *Identifier) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	_, d := p.FindSelector("", i.Name)
	if d == nil {
		p.Error(i.Position, fmt.Sprintf("%s undefined", i.Name))
		return nil
	}
	if v, ok := d.(*Variable); ok {
		if v.Const && v.Value.IsConstant(p) {
			return v.Value.GenerateConstIR(p, nil)
		}
	} else if f, ok := d.(*Function); ok {
		return f.IRFunction
	}
	p.Error(i.Position, "invalid constant declaration")
	return nil
}
