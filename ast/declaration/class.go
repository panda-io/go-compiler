package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Class struct {
	Base
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration
}

func (c *Class) GenerateIR(ctx *node.Context) {
	// generate struct
	s := ctx.Program.Structs[c.Qualified(ctx.Namespace)]
	//TO-DO contact with parent struct (check redeclared)
	ctx.Program.Module.NewGlobal(c.Qualified(ctx.Namespace), ir.NewStructType(s.Variables...))

	// generate function

	// generate vtable
}

func (c *Class) GenerateStructDeclaration(*node.Context) *node.Struct {
	// resolve parent later
	s := &node.Struct{
		Indexes: make(map[string]ir.Type),
	}
	for _, m := range c.Members {
		if v, ok := m.(*Variable); ok {
			t := types.TypeOf(v.Type)
			s.Variables = append(s.Variables, t)
			s.Indexes[v.Identifier()] = t
		}
	}
	return s
}
