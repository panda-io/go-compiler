package resolver

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/token"
)

type Error struct {
	Position *token.Position
	Message  string
}

type Resolver struct {
	declarations map[string]*Object

	packageScopes map[string]*Scope
	classScopes   map[string]*Scope
	importsScopes map[string]*Scope
	currentScope  *Scope

	file   *token.File
	source *ast.Source

	errors []*Error
}

func NewResolver() *Resolver {
	return &Resolver{
		declarations:  make(map[string]*Object),
		packageScopes: map[string]*Scope{ast.Global: NewScope(nil)},
		classScopes:   make(map[string]*Scope),
		importsScopes: make(map[string]*Scope),
	}
}

func (r *Resolver) Declare(f *token.File, s *ast.Source) {
	r.file = f
	r.source = s

	if r.packageScopes[r.source.Namespace] == nil {
		r.packageScopes[r.source.Namespace] = NewScope(nil)
	}
	for _, m := range r.source.Members {
		d := &Object{}
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
		d.Public = m.IsPublic()
		d.Name = m.Identifier()
		d.Declaration = m

		// register declaration
		if d.Kind == EnumObject || d.Kind == InterfaceObject || d.Kind == ClassObject {
			qualifiedName := m.Identifier()
			if r.source.Namespace != ast.Global {
				qualifiedName = r.source.Namespace + "." + qualifiedName
			}
			m.SetQualifiedName(qualifiedName)
			if _, ok := r.declarations[qualifiedName]; ok {
				r.error(m.GetPosition(), fmt.Sprintf("%s redeclared", m.Identifier()))
			} else {
				r.declarations[qualifiedName] = d
			}
		}

		// prepare package scope
		if d.Kind == VariableObject || d.Kind == FunctionObject {
			if r.source.Namespace == "" {
				r.packageScopes[ast.Global].Insert(d)
			} else {
				r.packageScopes[r.source.Namespace].Insert(d)
			}
		}

		// prepare class scope
		if d.Kind == ClassObject {
			//TO-DO parent class scope as outer
			// resolve later and check cycle inheritance
			qualifiedName := m.(*declaration.Class).QualifinedName
			r.classScopes[qualifiedName] = NewScope(nil)
			c := m.(*declaration.Class)
			for _, classMember := range c.Members {
				o := &Object{}
				switch classMember.(type) {
				case *declaration.Variable:
					o.Kind = VariableObject
				case *declaration.Function:
					o.Kind = FunctionObject
				}
				o.Public = classMember.IsPublic()
				o.Name = classMember.Identifier()
				o.Declaration = classMember
				r.classScopes[qualifiedName].Insert(o)
			}
		}
	}
}

func (r *Resolver) Resolve(f *token.File, s *ast.Source) {
	r.file = f
	r.source = s

	// prepare imports scopes
	for _, i := range s.Imports {
		r.importsScopes[i.Alias] = r.packageScopes[i.Namespace]
	}

	for _, m := range r.source.Members {
		r.resolveDeclaration(m)
	}

	for k := range r.importsScopes {
		delete(r.importsScopes, k)
	}
}

func (r *Resolver) Errors() int {
	for _, e := range r.errors {
		fmt.Printf("error: %s \n %s \n", e.Position.String(), e.Message)
	}
	return len(r.errors)
}

func (r *Resolver) error(offset int, message string) {
	r.errors = append(r.errors, &Error{
		Position: r.file.Position(offset),
		Message:  message,
	})
}
