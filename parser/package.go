package parser

import (
	"fmt"
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseProgram() {
	m := p.parseAttributes()
	n := p.parseNamespace()
	program := p.findPackage(n)
	if len(m) > 0 {
		program.Custom = append(program.Custom, m...)
	}

	if p.token == token.Using {
		if len(m) > 0 {
			p.error(m[0].Position, "import should not contain attributes")
		}
		p.parseUsing()
	}

	for p.token != token.EOF {
		modifier := p.parseModifier()
		m = p.parseAttributes()
		switch p.token {
		case token.Const, token.Var:
			v := p.parseVariable(modifier, m)
			name := v.Name.Name
			if _, ok := program.Variables[name]; ok {
				p.error(v.Name.Position, fmt.Sprintf("variable %s redeclared", name))
			}
			program.Variables[name] = v

		case token.Function:
			f := p.parseFunction(modifier, m, nil)
			name := f.Name.Name
			if _, ok := program.Functions[name]; ok {
				p.error(f.Name.Position, fmt.Sprintf("function %s redeclared", name))
			}
			program.Functions[name] = f

		case token.Enum:
			e := p.parseEnum(modifier, m)
			name := e.Name.Name
			if _, ok := program.Enums[name]; ok {
				p.error(e.Name.Position, fmt.Sprintf("enum %s redeclared", name))
			}
			program.Enums[name] = e

		case token.Interface:
			i := p.parseInterface(modifier, m)
			name := i.Name.Name
			if _, ok := program.Interfaces[name]; ok {
				p.error(i.Name.Position, fmt.Sprintf("interface %s redeclared", name))
			}
			program.Interfaces[name] = i

		case token.Class:
			c := p.parseClass(modifier, m)
			name := c.Name.Name
			if existing, ok := program.Classes[name]; ok {
				if !c.Modifier.Equal(existing.Modifier) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's modifier are different", name))
				}
				if !c.TypeParameters.Equal(existing.TypeParameters) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's type parameters are different", name))
				}
				if len(c.Parents) != len(existing.Parents) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's base type are different", name))
				}
				for i, parent := range existing.Parents {
					if !parent.Equal(existing.Parents[i]) {
						p.error(c.Name.Position, fmt.Sprintf("partial class %s's base type are different", name))
					}
				}
				for n, v := range existing.Variables {
					if _, ok := c.Variables[n]; ok {
						p.error(v.Name.Position, fmt.Sprintf("class member %s's is redeclared", n))
					}
					c.Variables[n] = v
				}
				for n, f := range existing.Functions {
					if _, ok := c.Functions[n]; ok {
						p.error(f.Name.Position, fmt.Sprintf("class member %s's is redeclared", n))
					}
					c.Functions[n] = f
				}
				for n, e := range existing.Enums {
					if _, ok := c.Enums[n]; ok {
						p.error(e.Name.Position, fmt.Sprintf("class member %s's is redeclared", n))
					}
					c.Enums[n] = e
				}
				for n, i := range existing.Interfaces {
					if _, ok := c.Interfaces[n]; ok {
						p.error(i.Name.Position, fmt.Sprintf("class member %s's is redeclared", n))
					}
					c.Interfaces[n] = i
				}
				for n, cc := range existing.Classes {
					if _, ok := c.Classes[n]; ok {
						p.error(cc.Name.Position, fmt.Sprintf("class member %s's is redeclared", n))
					}
					c.Classes[n] = cc
				}
			}
			program.Classes[name] = c

		default:
			p.expectedError(p.position, "declaration")
		}
	}
}

func (p *Parser) parseNamespace() []string {
	p.expect(token.Namespace)
	if p.token == token.Semi {
		p.next()
		return nil
	}
	namespace := p.parseQualifiedName("")
	p.expect(token.Semi)
	return namespace
}

func (p *Parser) parseUsing() {
	for p.token == token.Using {
		f := p.scanner.GetFileName()
		p.expect(token.Using)
		i := &usingNamespace{}
		name := p.parseIdentifier()
		if p.token == token.Assign {
			i.alias = name
			p.next()
			name = p.parseIdentifier()
		}
		i.path = p.parseQualifiedName(name.Name)
		p.expect(token.Semi)
		p.imports[f] = append(p.imports[f], i)
	}
}

func (p *Parser) findPackage(namespace []string) *ast.Package {
	if len(namespace) == 0 {
		return p.program
	}
	program := p.program
	for len(namespace) > 0 {
		name := namespace[0]
		if _, ok := program.Children[name]; !ok {
			program.Children[name] = ast.NewPackage(name, program)
		}
		program = program.Children[name]
		namespace = namespace[1:]
	}
	return program
}

/*
func (p *Parser) validateProgram(program *ast.Package) {
	for _, v := range program.Variables {
		fmt.Println(v.Name.Name)
	}
	for _, f := range program.Functions {
		fmt.Println(f.Name.Name)
	}
	for _, e := range program.Enums {
		fmt.Println(e.Name.Name)
	}
	for _, i := range program.Interfaces {
		fmt.Println(i.Name.Name)
	}
	for _, c := range program.Classes {
		fmt.Println(c.Name.Name)
	}
}*/
