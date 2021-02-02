package ast

import "github.com/panda-foundation/go-compiler/ir"

type DeclarationStatement struct {
	StatementBase
	Name  *Identifier
	Type  Type
	Value Expression
}

func (d *DeclarationStatement) GenerateIR(c *Context) {
	//TO-DO // zero initialize
	//TO-DO class type conversion with auto pointer
	//Builtin, TypeName
	var alloca *ir.InstAlloca
	switch d.Type.(type) {
	case *BuitinType:
		alloca = ir.NewAlloca(d.Type.Type(c.Program))

	case *TypeName:
		// TO-DO implement
		// alloca.UserData = // param.UserData, _ = p.FindDeclaration(t)

	case *TypeFunction:
		alloca = ir.NewAlloca(d.Type.Type(c.Program))
	}

	if alloca == nil {
		c.Program.Error(d.Position, "invalid declaration")
	} else {
		c.Function.IREntry.InsertAlloca(alloca)
		var store *ir.InstStore
		if d.Value == nil {
			store = ir.NewStore(ir.NewZeroInitializer(d.Type.Type(c.Program)), alloca)
		} else {
			store = ir.NewStore(d.Value.GenerateIR(c, d.Type.Type(c.Program)), alloca)
		}
		c.Block.AddInstruction(store)
		err := c.AddObject(d.Name.Name, alloca)
		if err != nil {
			c.Program.Error(d.Position, err.Error())
		}
	}
}
