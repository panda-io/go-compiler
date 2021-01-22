package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Variable struct {
	DeclarationBase
	Token token.Token
	Type  Type
	Value Expression

	IRVariable *ir.Global
}

func (v *Variable) GenerateIR(c *Context) {
	if v.Value != nil {
		v.IRVariable = c.Program.Module.NewGlobal(v.Qualified(c.Module.Namespace), v.Type.Type(c))
	} else {
		value := v.Value.GenerateIR(c)
		if constValue, ok := value.(ir.Constant); ok {
			v.IRVariable = c.Program.Module.NewGlobalDef(v.Qualified(c.Module.Namespace), constValue)
		} else {
			c.Error(v.Position, "variable initialize value must be constant type")
		}
	}
	err := c.AddObject(v.Name.Name, v.IRVariable)
	if err != nil {
		c.Error(v.Position, err.Error())
	}
}
