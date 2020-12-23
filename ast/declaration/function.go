package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Function struct {
	Base
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Block
	ClassName      string
}

func (f *Function) GenerateIR(c *node.Context) {
	function := c.Module.NewFunc(f.Name.Name, types.TypeOf(f.ReturnType), f.Parameters.GenerateIR(c)...)
	c.Block = function.NewBlock("")
	if f.Body != nil {
		f.Body.GenerateIR(c)
	}
}
