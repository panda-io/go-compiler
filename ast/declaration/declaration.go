package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Declaration interface {
	node.Node
	Identifier() string
	Qualified(namespace string) string
}

type Modifier struct {
	Public bool
	//TO-DO Inline
}

func (m0 *Modifier) Equal(m1 *Modifier) bool {
	return m0.Public == m1.Public
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*expression.Literal
}

type Base struct {
	node.Base
	Attributes []*Attribute
	Modifier   *Modifier
	Name       *expression.Identifier
	ObjectName string // parent object (class|interface|enum)
}

func (b *Base) External() bool {
	for _, a := range b.Attributes {
		if a.Name == node.Extern {
			return true
		}
	}
	return false
}

func (b *Base) Identifier() string {
	return b.Name.Name
}

func (b *Base) Qualified(namespace string) string {
	name := b.Name.Name
	if b.External() {
		if b.ObjectName != "" {
			panic("object member cannot be external")
		}
	} else {
		if b.ObjectName != "" {
			name = b.ObjectName + "." + name
		}
		if !(namespace == node.Global && b.Name.Name == node.Entry) {
			name = namespace + "." + name
		}
	}
	return name
}

func FindDeclaration(c *node.Context, declarations map[string]Declaration, t *types.TypeName) (string, Declaration) {
	if t.Selector == "" {
		// search current package
		if c.Namespace != node.Global {
			qualified := c.Namespace + "." + t.Name
			if declarations != nil {
				return qualified, declarations[qualified]
			}
		}
		// search global
		qualified := node.Global + "." + t.Name
		return qualified, declarations[qualified]
	} else {
		// search imports
		for _, i := range c.Imports {
			if i.Alias == t.Selector {
				qualified := i.Namespace + "." + t.Name
				return qualified, declarations[qualified]
			}
		}
		return t.Name, nil
	}
}

func TypeOf(c *node.Context, declarations map[string]Declaration, t types.Type) ir.Type {
	if t == nil {
		return ir.Void
	}

	switch typ := t.(type) {
	case *types.BuitinType:
		switch typ.Token {
		case token.Bool:
			return ir.I1

		case token.Char:
			return ir.I32

		case token.Int8, token.Uint8, token.SByte, token.Byte:
			return ir.I8

		case token.Int16, token.Uint16, token.Short, token.Ushort:
			return ir.I16

		case token.Int32, token.Uint32, token.Int, token.Uint:
			return ir.I32

		case token.Int64, token.Uint64, token.Long, token.Ulong:
			return ir.I64

		case token.Float32, token.Float:
			return ir.Float32

		case token.Float64, token.Double:
			return ir.Float64

		case token.Void:
			return ir.Void

		case token.Pointer:
			return ir.NewPointerType(ir.I8)

		default:
			panic("invalid buitin type: " + typ.Token.String())
		}

	case *types.TypeName:
		_, d := FindDeclaration(c, declarations, typ)
		if d == nil {
			c.Error(t.GetPosition(), "undefined: "+typ.Name)
			return ir.Void
		}
		//TO-DO use shared, weak
		//TO-DO deal with top class
		return ir.Void

	default:
		panic("inviad type define")
	}
}

func CompareMemberFunction(f1 *ir.FuncType, f2 *ir.FuncType) bool {
	if !f1.RetType.Equal(f2.RetType) {
		return false
	}
	if len(f1.Params) != len(f2.Params) {
		return false
	}
	for i := 1; i < len(f1.Params); i++ {
		if !f1.Params[i].Equal(f2.Params[i]) {
			return false
		}
	}
	return f1.Variadic == f2.Variadic
}

/*
func (p *Parameters) GenerateParameters(c *node.Context, structType *ir.StructType) []*ir.Param {
	params := []*ir.Param{}
	if structType != nil {
		param := ir.NewParam(ir.NewPointerType(structType))
		param.LocalName = node.This
		params = append(params, param)
	}
	if p == nil {
		return params
	}
	for _, parameter := range p.Parameters {
		param := ir.NewParam(TypeOf(parameter.Type))
		param.LocalName = parameter.Name
		params = append(params, param)
	}
	return params
}*/
