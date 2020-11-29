package parser

import (
	"fmt"
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseProgram() {
	m := p.parseMetadata()
	n := p.parseNamespace()
	program := p.findPackage(n)
	if len(m) > 0 {
		program.Custom = append(program.Custom, m...)
	}
	m = p.parseMetadata()

	if p.token == token.Import {
		if len(m) > 0 {
			p.error(m[0].Position, "import should not contain metadata")
		}
		p.parseImport() // ignore currently //TO-DO
		m = p.parseMetadata()
	}

	for p.token != token.EOF {
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			p.next()
			v := p.parseVariable()
			v.Custom = append(v.Custom, m...)
			v.Modifier = modifier
			name := v.Name.Name
			if _, ok := program.Variables[name]; ok {
				p.error(v.Name.Position, fmt.Sprintf("variable %s redeclared", name))
			}
			program.Variables[name] = v

		case token.Function:
			p.next()
			f := p.parseFunction(nil)
			f.Custom = append(f.Custom, m...)
			f.Modifier = modifier
			name := f.Name.Name
			if _, ok := program.Functions[name]; ok {
				p.error(f.Name.Position, fmt.Sprintf("function %s redeclared", name))
			}
			program.Functions[name] = f

		case token.Enum:
			p.next()
			e := p.parseEnum()
			e.Custom = append(e.Custom, m...)
			e.Modifier = modifier
			name := e.Name.Name
			if _, ok := program.Enums[name]; ok {
				p.error(e.Name.Position, fmt.Sprintf("enum %s redeclared", name))
			}
			program.Enums[name] = e

		case token.Interface:
			p.next()
			i := p.parseInterface()
			i.Custom = append(i.Custom, m...)
			i.Modifier = modifier
			name := i.Name.Name
			if _, ok := program.Interfaces[name]; ok {
				p.error(i.Name.Position, fmt.Sprintf("interface %s redeclared", name))
			}
			program.Interfaces[name] = i

		case token.Class:
			p.next()
			c := p.parseClass()
			c.Custom = append(c.Custom, m...)
			c.Modifier = modifier
			name := c.Name.Name
			if existing, ok := program.Classes[name]; ok {
				if !c.Modifier.Equal(existing.Modifier) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's modifier are different", name))
				}
				if !c.TypeParameters.Equal(existing.TypeParameters) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's type parameters are different", name))
				}
				if !c.Base.Equal(existing.Base) {
					p.error(c.Name.Position, fmt.Sprintf("partial class %s's base type are different", name))
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
			}
			program.Classes[name] = c

		default:
			p.expectedError(p.position, "declaration")
		}
		m = p.parseMetadata()
	}
}

func (p *Parser) parseModifier() *ast.Modifier {
	m := &ast.Modifier{}
	if p.token == token.Public {
		m.Public = true
		p.next()
	}
	if p.token == token.Static {
		m.Static = true
		p.next()
	}
	if p.token == token.Async {
		m.Async = true
		p.next()
	}
	if p.token == token.Inline {
		m.Inline = true
		p.next()
	}
	return m
}

func (p *Parser) parseMetadata() []*ast.Metadata {
	if p.token != token.META {
		return nil
	}

	var meta []*ast.Metadata
	for p.token == token.META {
		p.next()
		if p.token != token.IDENT {
			p.expect(token.IDENT)
		}
		m := &ast.Metadata{Position: p.position}
		m.Name = p.literal
		p.next()

		if p.token == token.STRING {
			m.Text = p.literal
			p.next()
		} else if p.token == token.LeftParen {
			p.next()
			if p.token == token.STRING {
				m.Text = p.literal
				p.next()
			} else {
				m.Values = make(map[string]*ast.Literal)
				for {
					if p.token == token.IDENT {
						name := p.literal
						p.next()
						p.expect(token.Assign)
						switch p.token {
						case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL:
							if _, ok := m.Values[name]; ok {
								p.error(p.position, "duplicated meta "+name)
							}
							m.Values[name] = &ast.Literal{
								Position: p.position,
								Type:     p.token,
								Value:    p.literal,
							}
						default:
							p.expectedError(p.position, "basic literal (bool, char, int, float, string)")
						}
						p.next()
						if p.token == token.RightParen {
							break
						}
						p.expect(token.Comma)
					} else {
						p.expect(token.IDENT)
					}
				}
			}
			p.expect(token.RightParen)
		}
		meta = append(meta, m)
	}
	return meta
}

func (p *Parser) parseNamespace() []string {
	p.expect(token.Namespace)
	if p.token == token.Semi {
		return nil
	}
	name := p.parseQualifiedName(nil)
	namespace := []string{}
	for _, n := range name {
		namespace = append(namespace, n.Name)
	}
	p.expect(token.Semi)
	return namespace
}

//TO-DO currently only skip, later need to be added into scope for checking
func (p *Parser) parseImport() {
	for p.token == token.Import {
		p.expect(token.Import)
		name := p.parseIdentifier()
		if p.token == token.Assign {
			//TO-DO alias name here
			p.next()
			name = p.parseIdentifier()
		}
		//TO-DO full path
		path := p.parseQualifiedName(name)
		fmt.Println("import:", path)
		p.expect(token.Semi)
		//TO-DO collect imports	// imports = append(imports, importDecl)
	}
}

func (p *Parser) parseQualifiedName(identifier *ast.Identifier) []*ast.Identifier {
	if identifier == nil {
		identifier = p.parseIdentifier()
	}
	qualifiedName := []*ast.Identifier{identifier}
	for p.token == token.Dot {
		p.next()
		qualifiedName = append(qualifiedName, p.parseIdentifier())
	}
	return qualifiedName
}

func (p *Parser) findPackage(namespace []string) *ast.Program {
	if len(namespace) == 0 {
		return p.root
	}

	program := p.root
	for len(namespace) > 0 {
		name := namespace[0]
		if _, ok := program.Children[name]; !ok {
			program.Children[name] = ast.NewProgram(name, program)
		}
		program = program.Children[name]
		namespace = namespace[1:len(namespace)]
	}
	return program
}
