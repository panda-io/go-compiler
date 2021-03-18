package ast

import "github.com/panda-foundation/go-compiler/ir"

func init() {
	/*
		function read_memory<type>(source pointer, offset int) type;
	    function write_memory<type>(source pointer, value type);*/
	RegisterComplierFunction("type", "to_pointer", CompilerReadMemory)
	RegisterComplierFunction("type", "from_pointer", CompilerWriteMemory)
}

/*
@compiler
function is<type>(value type) bool;
*/

func CompilerReadMemory(c *Context, invocation *Invocation) ir.Value {
	args := invocation.Arguments
	if args != nil && len(args.Arguments) == 1 {
		arg := args.Arguments[0]
		if ir.IsStruct(arg.Type(c, nil)) {
			cast := ir.NewBitCast(arg.GenerateIR(c, nil), pointerType)
			c.Block.AddInstruction(cast)
			return cast
		}
	}
	c.Program.Error(args.Position, "cannot convert target to pointer")
	return nil
}

func CompilerWriteMemory(c *Context, invocation *Invocation) ir.Value {
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
