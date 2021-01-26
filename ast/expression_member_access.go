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
			c.Program.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
			return nil
		}
		return obj.Type()
	} else if _, ok := m.Parent.(*This); ok {
		t := c.ObjectType(m.Member.Name)
		if t == nil {
			c.Program.Error(m.Position, fmt.Sprintf("%s undefined", m.Member.Name))
		}
		return t
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
			c.Program.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
		}
		return obj
	} else if _, ok := m.Parent.(*This); ok {
		member := c.Function.Class.GetMember(c, c.FindObject(ClassThis), m.Member.Name)
		if member == nil {
			c.Program.Error(m.Position, fmt.Sprintf("%s is undefined", m.Member.Name))
			return nil
		}
		return member
	} /* else {
		// TO-DO
		// generate parent firstly, then check type of parent, then generate ir
	}*/
	return nil
}

func (*MemberAccess) IsConstant() bool {
	//TO-DO
	return false
}

func (*MemberAccess) GenerateConstIR(c Context, exprect ir.Type) ir.Value {
	//TO-DO
	return nil
}
