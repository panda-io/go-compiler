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
	return m.GenerateIR(c).Type()
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
		this := c.FindObject(ClassThis)
		fmt.Println(this.Type())
	} /* else {
		// TO-DO
		// generate parent firstly, then check type of parent, then generate ir
	}*/
	return nil
}
