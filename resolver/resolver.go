package resolver

import (
	"fmt"
	"strings"

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
	declarations  map[string]*Declaration
	globalScope   *Scope
	packageScopes map[string]*Scope

	//TO-DO scope
	//global->package->class->function->statement

	errors []*Error
}

func NewResolver() *Resolver {
	return &Resolver{
		declarations:  make(map[string]*Declaration),
		globalScope:   &Scope{},
		packageScopes: make(map[string]*Scope),
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
			continue
		case *declaration.Function:
			continue
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
		r.resolveMember(f, m, s)
	}
}

func (r *Resolver) resolveMember(f *token.File, d declaration.Declaration, s *ast.SoureFile) {
	switch m := d.(type) {
	case *declaration.Variable:
		r.resolveType(f, m.Type, s)
		// TO-DO validate const expr
	case *declaration.Function:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s)
		}
		if m.Parameters != nil {
			r.resolveParameters(f, m.Parameters, s)
		}
		if m.ReturnType != nil {
			r.resolveType(f, m.ReturnType, s)
		}
		// TO-DO resolve statements // local declaration //call, primary :: implement it with scope ?
	case *declaration.Enum:
		// TO-DO validate const expr
	case *declaration.Interface:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s)
		}
		r.resolveParents(f, m.Parents, s)
		for _, mm := range m.Members {
			r.resolveMember(f, mm, s)
		}
	case *declaration.Class:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s)
		}
		r.resolveParents(f, m.Parents, s)
		for _, mm := range m.Members {
			r.resolveMember(f, mm, s)
		}
	}
}

func (r *Resolver) resolveParents(f *token.File, ts []*types.TypeName, s *ast.SoureFile) {
	for _, t := range ts {
		r.resolveTypeName(f, t, s)
	}
}

func (r *Resolver) resolveParameters(f *token.File, t *types.Parameters, s *ast.SoureFile) {
	for _, arg := range t.Parameters {
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			r.resolveTypeName(f, typeParameter, s)
		}
	}
}

func (r *Resolver) resolveTypeParameters(f *token.File, t *types.TypeParameters, s *ast.SoureFile) {
	for _, arg := range t.Parameters {
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			r.resolveTypeName(f, typeParameter, s)
		}
	}
}

func (r *Resolver) resolveType(f *token.File, typ types.Type, s *ast.SoureFile) {
	if t, ok := typ.(*types.TypeName); ok {
		r.resolveTypeName(f, t, s)
		if t.TypeArguments != nil {
			r.resolveTypeArguments(f, t.TypeArguments, s)
		}
	}
}

func (r *Resolver) resolveTypeArguments(f *token.File, t *types.TypeArguments, s *ast.SoureFile) {
	for _, arg := range t.Arguments {
		if typeArgument, ok := arg.(*types.TypeName); ok {
			r.resolveTypeName(f, typeArgument, s)
		}
	}
}

func (r *Resolver) resolveTypeName(f *token.File, t *types.TypeName, s *ast.SoureFile) {
	names := r.findQualifiedName(f, t, s)
	if len(names) == 0 {
		r.error(f, t.GetPosition(), fmt.Sprintf("%s undefined", t.Name))
	} else if len(names) > 1 {
		r.error(f, t.GetPosition(), fmt.Sprintf("ambiguous type %s", t.Name))
	} else {
		t.QualifiedName = names[0]
		fmt.Println("qualified:", names[0])
	}
}

func (r *Resolver) findQualifiedName(f *token.File, name *types.TypeName, s *ast.SoureFile) []string {
	names := []string{}
	// search qualified name directly
	if _, ok := r.declarations[name.Name]; ok {
		names = append(names, name.Name)
	}
	// search by using
	for _, u := range s.Using {
		if u.Alias == "" {
			n := u.Namespace + "." + name.Name
			if d, ok := r.declarations[n]; ok {
				if d.Kind == ClassObject || d.Kind == EnumObject || d.Kind == InterfaceObject {
					names = append(names, n)
				} else {
					r.error(f, name.GetPosition(), fmt.Sprintf("%s can not be a type", name.Name))
				}
			}
		} else {
			if strings.HasPrefix(name.Name, u.Alias+".") {
				n := strings.Replace(name.Name, u.Alias, u.Namespace, 1)
				if d, ok := r.declarations[n]; ok {
					if d.Kind == ClassObject || d.Kind == EnumObject || d.Kind == InterfaceObject {
						names = append(names, n)
					} else {
						r.error(f, name.GetPosition(), fmt.Sprintf("%s can not be a type", name.Name))
					}
				}
			}
		}
	}
	// search same package
	if !strings.Contains(name.Name, ".") {
		for _, m := range s.Members {
			if m.Identifier() == name.Name {
				n := name.Name
				if s.Namespace != "" {
					n = s.Namespace + "." + n
				}
				if _, ok := r.declarations[n]; ok {
					names = append(names, n)
				}
			}
		}
	}
	return names
}
