package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type Function struct {
	Base
	TypeParameters *types.TypeParameters
	Parameters     *types.Parameters
	ReturnType     types.Type
	Body           *statement.Block
}

func (f *Function) GenerateIR(c *node.Context) {
	var s *ir.StructType
	if f.ObjectName != "" {
		d := c.FindDelaration(f.ObjectName)
		s = d.(*ir.Global).ContentType.(*ir.StructType)
	}
	function := c.Program.Module.NewFunc(f.Qualified(c.Namespace), types.TypeOf(f.ReturnType), f.Parameters.GenerateIR(c, s)...)
	if f.Body != nil {
		c.Block = function.NewBlock("")
		f.Body.GenerateIR(c)
		if c.Block.Term == nil {
			c.Block.Term = ir.NewRet(ir.NewNull(nil))
		}
	}
}

func (f *Function) GenerateDeclaration(c *node.Context) {
	params := []*ir.Param{}
	if f.ObjectName != "" {
		t := ir.NewStructType()
		t.TypeName = c.Namespace + "." + f.ObjectName
		param := ir.NewParam(ir.NewPointerType(t))
		params = append(params, param)
	}
	if f.Parameters != nil {
		for _, parameter := range f.Parameters.Parameters {
			param := ir.NewParam(types.TypeOf(parameter.Type))
			params = append(params, param)
		}
	}
	d := ir.NewFunc(f.Qualified(c.Namespace), types.TypeOf(f.ReturnType), params...)
	err := c.AddDeclaration(f.Qualified(c.Namespace), d)
	if err != nil {
		c.Error(f.Position, err.Error())
	}
}
