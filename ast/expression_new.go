package ast

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type New struct {
	ExpressionBase
	Typ       *TypeName
	Arguments *Arguments
	HasOwner  bool
}

func (n *New) Type(c *Context, expected ir.Type) ir.Type {
	return n.Typ.Type(c.Program)
}

func (n *New) GenerateIR(ctx *Context, expected ir.Type) ir.Value {
	qualified, d := ctx.Program.FindDeclaration(n.Typ)
	if c, ok := d.(*Class); ok {
		instance := c.CreateInstance(ctx, n.Arguments)
		if IsBuiltinClass(qualified) {
			if !n.HasOwner {
				ctx.Function.BuiltinReleasePool = append(ctx.Function.BuiltinReleasePool, instance)
			}
			return instance
		} else {
			counterClass := ctx.Program.FindQualified(Counter).(*Class)
			counter := counterClass.CreateInstance(ctx, nil)
			if !n.HasOwner {
				ctx.Function.AutoReleasePool = append(ctx.Function.AutoReleasePool, counter)
			}
			// retain shared
			call := ir.NewCall(retainShared, counter)
			ctx.Block.AddInstruction(call)
			// set object
			object := counterClass.GetMember(ctx, counter, "object")
			ctx.Block.AddInstruction(ir.NewStore(instance, object))
			// set destructor
			destructor := counterClass.GetMember(ctx, counter, "destructor")
			ctx.Block.AddInstruction(ir.NewStore(c.IRFunctions[1], destructor))
			return counter
		}
	}
	ctx.Program.Error(n.Position, "invalid type for new operator")
	return nil
}

func (*New) IsConstant(p *Program) bool {
	return false
}

func (*New) GenerateConstIR(p *Program, expected ir.Type) ir.Constant {
	return nil
}
