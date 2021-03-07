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
		qualified, declaration := c.Program.FindDeclaration(t)
		switch declaration.(type) {
		case *Class:
			alloca = ir.NewAlloca(pointerType)
			SetUserData(alloca, qualified)
			if IsBuiltinClass(qualified) {
				c.Function.BuiltinReleasePool = append(c.Function.BuiltinReleasePool, alloca)
			} else {
				c.Function.AutoReleasePool = append(c.Function.AutoReleasePool, alloca)
			}

		case *Enum:
			alloca = ir.NewAlloca(ir.I32)

		case *Interface:
			// TO-DO interface
			//TO-DO need to be some convert
			SetUserData(alloca, qualified)
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
			switch d.Type.(type) {
			case *BuitinType:
				store = ir.NewStore(ir.NewZeroInitializer(d.Type.Type(c.Program)), alloca)
			case *TypeName:
				store = ir.NewStore(ir.NewZeroInitializer(pointerType), alloca)
			}
		} else {
			if n, ok := d.Value.(*New); ok {
				n.HasOwner = true
			}
			instance := d.Value.GenerateIR(c, d.Type.Type(c.Program))
			store = ir.NewStore(instance, alloca)
		}
		c.Block.AddInstruction(store)
		err := c.AddObject(d.Name.Name, alloca)
		if err != nil {
			c.Program.Error(d.Position, err.Error())
		}
	}
}
