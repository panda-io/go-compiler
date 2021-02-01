package ast

import "github.com/panda-foundation/go-compiler/ir"

func init() {
	RegisterComplierFunction("type", "to_pointer", CastToPinter)
	RegisterComplierFunction("type", "from_pointer", CastFromPinter)
}

/*
#compiler
function is<type>(value type) bool;
*/

func CastToPinter(c *Context, invocation *Invocation) ir.Value {
	args := invocation.Arguments
	if args != nil && len(args.Arguments) == 1 {
		arg := args.Arguments[0]
		if ir.IsStruct(arg.Type(c, nil)) {
			cast := ir.NewBitCast(arg.GenerateIR(c, nil), ir.NewPointerType(ir.I8))
			c.Block.AddInstruction(cast)
			return cast
		}
	}
	c.Program.Error(args.Position, "cannot convert target to pointer")
	return nil
}

func CastFromPinter(c *Context, invocation *Invocation) ir.Value {
	args := invocation.Arguments
	types := invocation.TypeArguments
	if args != nil && len(args.Arguments) == 1 && types != nil && len(types.Arguments) == 1 {
		arg := args.Arguments[0]
		typ := types.Arguments[0]
		if ir.IsPointer(typ.Type(c.Program)) && ir.IsPointer(arg.Type(c, nil)) && arg.Type(c, nil).(*ir.PointerType).ElemType.Equal(ir.I8) {
			cast := ir.NewBitCast(arg.GenerateIR(c, nil), typ.Type(c.Program))
			c.Block.AddInstruction(cast)
			return cast
		}
	}
	c.Program.Error(args.Position, "cannot convert pointer to target type")
	return nil
}
