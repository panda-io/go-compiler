package parser

import (
	"fmt"
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Resolver struct {
	parser   *Parser
	global   *ast.Package
	declared map[string]TypeKind

	using          []*ast.Package
	currentPackage *ast.Package
	currentClass   *declaration.Class
	currentScope   *Scope // function, statement
}

func NewResolver(p *Parser) *Resolver {
	return &Resolver{
		parser:   p,
		global:   p.Program(),
		declared: make(map[string]TypeKind),
	}
}

func (r *Resolver) Resolve(p *ast.Package) {
	r.declare(p)
	r.resolve(p)
}

func (r *Resolver) declare(p *ast.Package) {
	ns := p.Namespace()
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Enum:
			r.declared[ns+t.Name.Name] = EnumType
		case *declaration.Interface:
			r.declared[ns+t.Name.Name] = InterfaceType
		case *declaration.Class:
			r.declared[ns+t.Name.Name] = ClassType
		}
	}
	for _, c := range p.Children {
		r.declare(c)
	}
}

func (r *Resolver) resolve(p *ast.Package) {
	r.currentPackage = p

	for _, m := range p.Members {
		r.resolveMember(m)
	}
	for _, c := range p.Children {
		r.resolve(c)
	}
}

func (r *Resolver) resolveMember(d declaration.Declaration) {
	switch m := d.(type) {
	case *declaration.Variable:
		r.resolveType(m.Type)

	case *declaration.Function:
	case *declaration.Enum:
	case *declaration.Interface:
	case *declaration.Class:
	}
}

func (r *Resolver) resolveType(typ types.Type) {
	switch t := typ.(type) {
	case *types.BuitinType:
		// no need to resolve for builtin type

	case *types.TypeName:
		fmt.Println("::", t.QualifiedName)

	}
}

// TO-DO validate conflict in using

// resolveType
// resolveCompoundStatement
// resolveExpression

// use scope when start compound, if, for, switch, case, function

//TO-DO
//invalid recursive type
//ambiguous type (namespace conflict)
