package ast

import "github.com/panda-foundation/go-compiler/ir"

type DeclarationStatement struct {
	StatementBase
	Name  *Identifier
	Type  Type
	Value Expression
}

func (d *DeclarationStatement) GenerateIR(c *Context) {
	var alloca *ir.InstAlloca
	switch t := d.Type.(type) {
	case *BuitinType:
		alloca = ir.NewAlloca(d.Type.Type(c.Program))

	case *TypeName:
		// TO-DO interface
		qualified, declaration := c.Program.FindDeclaration(t)
		switch declaration.(type) {
		case *Class:
			if IsBuiltinType(qualified) {
				alloca = ir.NewAlloca(CreateStructPointer(qualified))
			} else {
				alloca = ir.NewAlloca(counterType)
			}
			alloca.UserData = qualified

		case *Enum:
			alloca = ir.NewAlloca(ir.I32)

		case *Interface:
			//TO-DO need to be some convert
			alloca.UserData = qualified
		}

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
			//TO-DO check counter or instance
			if n, ok := d.Value.(*New); ok {
				n.HasOwner = true
			}
			store = ir.NewStore(d.Value.GenerateIR(c, d.Type.Type(c.Program)), alloca)
		}
		c.Block.AddInstruction(store)
		err := c.AddObject(d.Name.Name, alloca)
		if err != nil {
			c.Program.Error(d.Position, err.Error())
		}
	}
}
