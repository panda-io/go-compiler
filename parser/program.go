package parser

import (
	"fmt"
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseProgram() {
	/*
		program := p.root
		parse metadata
		namespace := p.parseNamespace(metadata)
		if namespace.Path != nil {
			program = p.root.FindPackage(namespace.Path)
		}

		// import
		parse metadata
		p.parseImport(metadata) // metadata is invalid here

		// rest of namespace body
		for p.tok != EOF {
			decl := p.parseDecl(declStart) // parse metadata and modifier
			switch v := decl.(type) {
			case *ValueDecl:
				program.Values = append(program.Values, v)
			case *EnumDecl:
				program.Enums = append(program.Enums, v)
			case *InterfaceDecl:
				program.Interfaces = append(program.Interfaces, v)
			case *ClassDecl:
				program.Classes = append(program.Classes, v)
			case *FuncDecl:
				program.Functions = append(program.Functions, v)
			case *BadDecl:
				fmt.Println("bad decl", v.Pos())
			}
		}*/

	/*
		func (p *Parser) parseDecl(sync map[Token]bool) Decl {
		m := p.parseModifier()
		switch p.tok {
		case Const, Var:
			return p.parseValueDecl(m)

		case Enum:
			return p.parseEnumDecl(m)

		case Interface:
			return p.parseInterfaceDecl(m)

		case Class:
			return p.parseClassDecl(m)

		case Function:
			return p.parseFuncDecl(m)

		default:
			pos := p.pos
			p.errorExpected(pos, "declaration")
			return &BadDecl{Start: pos}
		}
	}*/
}

func (p *Parser) parseModifier() *ast.Modifier {
	m := &ast.Modifier{}
	if p.token == token.Public {
		m.Position = p.position
		m.Public = true
		p.next()
	}
	if p.token == token.Static {
		m.Position = p.position
		m.Static = true
		p.next()
	}
	if p.token == token.Async {
		m.Position = p.position
		m.Async = true
		p.next()
	}
	if p.token == token.Inline {
		m.Position = p.position
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
						case token.INT, token.FLOAT, token.CHAR, token.STRING, token.True, token.False:
							if _, ok := m.Values[name]; ok {
								p.error(p.position, "duplicated meta "+name)
							}
							m.Values[name] = &ast.Literal{
								Position: p.position,
								Type:     p.token,
								Value:    p.literal,
							}
						default:
							p.unexpected(p.position, "basic literal (bool, char, int, float, string)")
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
	if p.token != token.Namespace {
		return nil
	}
	p.next()

	namespace := p.parseQualifiedName("")
	p.expect(token.Semi)
	return namespace
}

//TO-DO currently only skip, later need to be added into scope for checking
func (p *Parser) parseImport() {
	for p.token == token.Import {
		p.expect(token.Import)
		name := p.parseIdentifier().Name
		if p.token == token.Assign {
			//TO-DO alias name here
			p.next()
			name = p.parseIdentifier().Name
		}
		//TO-DO full path
		path := p.parseQualifiedName(name)
		fmt.Println("import:", path)
		p.expect(token.Semi)
		//TO-DO collect imports	// imports = append(imports, importDecl)
	}
}

func (p *Parser) parseQualifiedName(first string) []string {
	if first == "" {
		first = p.parseIdentifier().Name
	}
	qualifiedName := []string{first}
	for p.token == token.Dot {
		p.next()
		qualifiedName = append(qualifiedName, p.parseIdentifier().Name)
	}
	return qualifiedName
}
