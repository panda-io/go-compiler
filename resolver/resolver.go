package resolver

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

/********************************
global -> package -> class -> function -> statement
variable
function
interface
enum
class
*********************************/

type Declaration struct {
	Kind   ObjectKind
	Public bool
}

type Error struct {
	Position *token.Position
	Message  string
}

type Resolver struct {
	declarations map[string]*Declaration

	//global->package->class->function->statement

	errors []*Error
}

func NewResolver() *Resolver {
	return &Resolver{
		declarations: make(map[string]*Declaration),
	}
}

func (r *Resolver) Declare(f *token.File, s *ast.SoureFile) {
	r.declare(f, s)
}

func (r *Resolver) Resolve(f *token.File, s *ast.SoureFile) {
	r.resolve(f, s)
}

func (r *Resolver) error(f *token.File, offset int, message string) {
	r.errors = append(r.errors, &Error{
		Position: f.Position(offset),
		Message:  message,
	})
}

func (r *Resolver) declare(f *token.File, p *ast.SoureFile) {
	for _, m := range p.Members {
		d := &Declaration{}
		d.Public = m.IsPublic()
		switch m.(type) {
		case *declaration.Variable:
			d.Kind = VariableObject
		case *declaration.Function:
			d.Kind = FunctionObject
		case *declaration.Enum:
			d.Kind = EnumObject
		case *declaration.Interface:
			d.Kind = InterfaceObject
		case *declaration.Class:
			d.Kind = ClassObject
		}
		qualifiedName := m.Identifier()
		if p.Namespace != "" {
			qualifiedName = p.Namespace + "." + qualifiedName
		}
		if _, ok := r.declarations[qualifiedName]; ok {
			r.error(f, m.GetPosition(), fmt.Sprintf("%s redeclared", m.Identifier()))
		} else {
			r.declarations[qualifiedName] = d
		}
	}
}

func (r *Resolver) resolve(f *token.File, s *ast.SoureFile) {
	for _, m := range s.Members {
		r.resolveMember(m)
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
		fmt.Println("resolve :", t.QualifiedName)

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
