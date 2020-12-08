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

func (r *Resolver) PrintErrors() {
	for _, e := range r.errors {
		fmt.Printf("error: %s \n %s \n", e.Position.String(), e.Message)
	}
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
		r.resolveMember(f, m, s, nil)
	}
}

func (r *Resolver) resolveMember(f *token.File, d declaration.Declaration, s *ast.SoureFile, t *types.TypeParameters) {
	switch m := d.(type) {
	case *declaration.Variable:
		r.resolveType(f, m.Type, s, t)
		// TO-DO validate const expr
	case *declaration.Function:
		if m.ClassName == "" {
			if s.Namespace != "" {
				m.QualifinedName = s.Namespace + "." + m.Name.Name
			}
		} else {
			if s.Namespace == "" {
				m.QualifinedName = m.ClassName + "." + m.Name.Name
			} else {
				m.QualifinedName = s.Namespace + "." + m.ClassName + "." + m.Name.Name
			}
		}
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s, t)
		}
		if m.Parameters != nil {
			//TO-DO merge class and function type parameters
			tp := t
			if tp == nil {
				if m.TypeParameters != nil {
					tp = m.TypeParameters
				}
			} else if m.TypeParameters != nil {
				tp = &types.TypeParameters{}
				tp.Parameters = t.Parameters
				tp.Parameters = append(tp.Parameters, m.TypeParameters.Parameters...)
			}
			r.resolveParameters(f, m.Parameters, s, tp)
		}
		if m.ReturnType != nil {
			r.resolveType(f, m.ReturnType, s, t)
		}
		// TO-DO resolve statements // local declaration //call, primary :: implement it with scope ?
	case *declaration.Enum:
		// TO-DO validate const expr
	case *declaration.Interface:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s, t)
		}
		r.resolveParents(f, m.Parents, s, t)
		for _, mm := range m.Members {
			r.resolveMember(f, mm, s, m.TypeParameters)
		}
	case *declaration.Class:
		if m.TypeParameters != nil {
			r.resolveTypeParameters(f, m.TypeParameters, s, t)
		}
		r.resolveParents(f, m.Parents, s, t)
		for _, mm := range m.Members {
			r.resolveMember(f, mm, s, m.TypeParameters)
		}
	}
}

func (r *Resolver) resolveParents(f *token.File, parents []*types.TypeName, s *ast.SoureFile, t *types.TypeParameters) {
	for _, p := range parents {
		r.resolveTypeName(f, p, s, t)
	}
}

func (r *Resolver) resolveParameters(f *token.File, tn *types.Parameters, s *ast.SoureFile, t *types.TypeParameters) {
	for _, arg := range tn.Parameters {
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			r.resolveTypeName(f, typeParameter, s, t)
		}
	}
}

func (r *Resolver) resolveTypeParameters(f *token.File, t *types.TypeParameters, s *ast.SoureFile, parent *types.TypeParameters) {
	for _, arg := range t.Parameters {
		if typeParameter, ok := arg.Type.(*types.TypeName); ok {
			if parent != nil {
				for _, p := range parent.Parameters {
					if p.Name == typeParameter.Name {
						r.error(f, typeParameter.GetPosition(), fmt.Sprintf("generic type %s conflict with class template", typeParameter.Name))
					}
				}
			}
			r.resolveTypeName(f, typeParameter, s, nil)
		}
	}
}

func (r *Resolver) resolveType(f *token.File, typ types.Type, s *ast.SoureFile, t *types.TypeParameters) {
	if tn, ok := typ.(*types.TypeName); ok {
		r.resolveTypeName(f, tn, s, t)
		if tn.TypeArguments != nil {
			r.resolveTypeArguments(f, tn.TypeArguments, s, t)
		}
	}
}

func (r *Resolver) resolveTypeArguments(f *token.File, args *types.TypeArguments, s *ast.SoureFile, t *types.TypeParameters) {
	for _, arg := range args.Arguments {
		if typeArgument, ok := arg.(*types.TypeName); ok {
			r.resolveTypeName(f, typeArgument, s, t)
		}
	}
}

func (r *Resolver) resolveTypeName(f *token.File, tn *types.TypeName, s *ast.SoureFile, t *types.TypeParameters) {
	names := r.findQualifiedName(f, tn, s, t)
	if len(names) == 0 {
		r.error(f, t.GetPosition(), fmt.Sprintf("%s undefined", tn.Name))
	} else if len(names) > 1 {
		r.error(f, t.GetPosition(), fmt.Sprintf("ambiguous type %s", tn.Name))
	} else {
		tn.QualifiedName = names[0]
	}
}

func (r *Resolver) findQualifiedName(f *token.File, name *types.TypeName, s *ast.SoureFile, t *types.TypeParameters) []string {
	names := []string{}
	// search in type parameters
	if t != nil {
		for _, p := range t.Parameters {
			if name.Name == p.Name {
				names = append(names, name.Name)
			}
		}
	}
	// search as qualified name directly
	if _, ok := r.declarations[name.Name]; ok {
		names = append(names, name.Name)
	}
	// search by using
	for _, u := range s.Using {
		if u.Alias == "" {
			n := u.Namespace + "." + name.Name
			if d, ok := r.declarations[n]; ok {
				if d.Kind == ClassObject || d.Kind == EnumObject || d.Kind == InterfaceObject {
					names = r.addQualifiedName(names, n)
				} else {
					r.error(f, name.GetPosition(), fmt.Sprintf("%s can not be a type", name.Name))
				}
			}
		} else {
			if strings.HasPrefix(name.Name, u.Alias+".") {
				n := strings.Replace(name.Name, u.Alias, u.Namespace, 1)
				if d, ok := r.declarations[n]; ok {
					if d.Kind == ClassObject || d.Kind == EnumObject || d.Kind == InterfaceObject {
						names = r.addQualifiedName(names, n)
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
					names = r.addQualifiedName(names, n)
				}
			}
		}
	}
	return names
}

func (r *Resolver) addQualifiedName(names []string, name string) []string {
	found := false
	for _, n := range names {
		if n == name {
			found = true
			break
		}
	}
	if !found {
		names = append(names, name)
	}
	return names
}
