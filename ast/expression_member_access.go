package ast

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ir"
)

type MemberAccess struct {
	ExpressionBase
	Parent        Expression
	Member        *Identifier
	FullNamespace string
}

func (m *MemberAccess) Type(c *Context) ir.Type {
	if ident, ok := m.Parent.(*Identifier); ok {
		// resolve here
		_, obj := c.FindSelector(ident.Name, m.Member.Name)
		if obj == nil {
			c.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
			return nil
		}
		return obj.Type()
	} else if _, ok := m.Parent.(*This); ok {
		class := c.Class
		if index, ok := class.VariableIndexes[m.Member.Name]; ok {
			v := ir.NewGetElementPtr(class.IRStruct, c.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
			return v.Type()
		} else if index, ok := class.FunctionIndexes[m.Member.Name]; ok {
			vtable := ir.NewGetElementPtr(class.IRStruct, c.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			f := ir.NewGetElementPtr(class.IRVTable, vtable, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
			return f.Type()
		} else {
			c.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
			return nil
		}
	} /* else {
		// TO-DO
		// generate parent firstly, then check type of parent, then generate ir
	}*/
	return nil
}

func (m *MemberAccess) GenerateParentIR(c *Context) ir.Value {
	if ident, ok := m.Parent.(*Identifier); ok {
		// resolve here
		parent, _ := c.FindSelector(ident.Name, m.Member.Name)
		return parent
	}
	return m.Parent.GenerateIR(c)
}

func (m *MemberAccess) GenerateIR(c *Context) ir.Value {
	if ident, ok := m.Parent.(*Identifier); ok {
		// resolve here
		_, obj := c.FindSelector(ident.Name, m.Member.Name)
		if obj == nil {
			c.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
		}
		return obj
	} else if _, ok := m.Parent.(*This); ok {
		class := c.Class
		if index, ok := class.VariableIndexes[m.Member.Name]; ok {
			v := ir.NewGetElementPtr(class.IRStruct, c.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
			c.Block.AddInstruction(v)
			return v
		} else if index, ok := class.FunctionIndexes[m.Member.Name]; ok {
			vtable := ir.NewGetElementPtr(class.IRStruct, c.FindObject(ClassThis), ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, 0))
			c.Block.AddInstruction(vtable)
			f := ir.NewGetElementPtr(class.IRVTable, vtable, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
			c.Block.AddInstruction(f)
			return f
		} else {
			c.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
			return nil
		}
	} /* else {
		// TO-DO
		// generate parent firstly, then check type of parent, then generate ir
	}*/
	return nil
}
