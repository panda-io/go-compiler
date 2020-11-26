package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseVariable() *ast.Variable {
	p.next()
	d := &ast.Variable{
		Position: p.position,
	}
	d.Name = p.parseIdentifier()
	d.Type = p.parseType()

	if p.token == token.Assign {
		p.next()
		if !p.token.IsLiteral() || p.token == token.IDENT {
			p.error(p.position, "variable can only be initialized by const value (string, char, float, int)")
		}
		d.Value = &ast.Literal{
			Position: p.position,
			Type:     p.token,
			Value:    p.literal,
		}
		p.next()
	}
	p.expect(token.Semi)
	return d
}

func (p *Parser) parseFunction() *ast.Function {
	return nil
}

func (p *Parser) parseEnum() *ast.Enum {
	// ok
	return nil
}

func (p *Parser) parseInterface() *ast.Interface {
	// ok
	return nil
}

func (p *Parser) parseClass() *ast.Class {
	// ok
	return nil
}

/*
func (p *Parser) parseEnumDecl(m *Modifier) *EnumDecl {
	p.next()
	name := p.parseIdent()

	decl := &EnumDecl{
		Modifier: m,
		Name:     name,
	}

	p.expect(LeftBrace)
	var list []*EnumStmt
	for p.tok != RightBrace {
		member := &EnumStmt{
			Name: p.parseIdent(),
		}
		if p.tok == Assign {
			p.next()
			//later check there are same values
			if p.tok == INT {
				member.Value = &BasicLit{
					Start: p.pos,
					Kind:  INT,
					Value: p.lit,
				}
			} else {
				p.error(p.pos, "only int value can assign to enum")
			}
			p.next()
		}
		list = append(list, member)
		if p.tok != RightBrace {
			p.expect(Comma)
		}
	}
	p.expect(RightBrace)
	decl.List = list
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseInterfaceDecl(m *Modifier) *InterfaceDecl {
	p.next()
	name := p.parseIdent()

	decl := &InterfaceDecl{
		Modifier: m,
		Name:     name,
	}

	//TO-DO generic

	p.expect(LeftBrace)
	for p.tok != RightBrace {
		switch p.tok {
		case Const, Var:
			decl.Values = append(decl.Values, p.parseValueDecl(nil))

		case Function:
			f := p.parseFuncDecl(m, true)
			f.IsMember = true
			f.InterfaceMember = true
			if f.Name.Name == name.Name {
				p.error(f.Name.Pos(), "interface has no contructor")
			}
			if f.Name.Name[0] == '~' {
				p.error(f.Name.Pos(), "interface has no destructor")
			}
			p.expect(Semi)
			decl.Functions = append(decl.Functions, f)
		default:
			pos := p.pos
			p.errorExpected(pos, "declaration")
			//p.advance(sync)
			//TO-DO advance wrong part
		}
	}
	p.expect(RightBrace)
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseClassDecl(m *Modifier) *ClassDecl {
	p.next()
	name := p.parseIdent()

	decl := &ClassDecl{
		Modifier: m,
		Name:     name,
	}

	//TO-DO generic
	if p.tok == Colon {
		for p.tok != LeftBrace {
			decl.Parents = append(decl.Parents, p.parseIdentOrSelector(nil))
			if p.tok == Comma {
				p.next()
			}
		}
	}

	p.expect(LeftBrace)
	for p.tok != RightBrace {
		m = p.parseModifier()
		switch p.tok {
		case Const, Var:
			decl.Values = append(decl.Values, p.parseValueDecl(m))

		case Function:
			f := p.parseFuncDecl(m, false)
			f.IsMember = true
			f.ClassName = name.Name
			if f.Name.Name == name.Name {
				f.IsConstructor = true
			}
			if f.Name.Name[0] == '~' {
				if f.Name.Name == "~"+name.Name {
					f.IsDestructor = true
				} else {
					p.error(f.Name.Pos(), "invalid destructor name")
				}
				//TO-DO check return type. must be void
			}
			decl.Functions = append(decl.Functions, f)
		default:
			pos := p.pos
			p.errorExpected(pos, "declaration")
			//p.advance(sync)
			//TO-DO advance wrong part
		}
	}
	p.expect(RightBrace)
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseFuncDecl(m *Modifier, onlyDeclare bool) *FuncDecl {
	p.expect(Function)

	//Tilde
	tilde := false
	if p.tok == Tilde {
		tilde = true
		p.next()
	}
	ident := p.parseIdent()
	if tilde {
		//TO-DO check in class
		ident.Name = "~" + ident.Name
	}
	generic := p.parseGeneric()
	params := p.parseParameters()
	result := p.parseResult()

	decl := &FuncDecl{
		Modifier: m,
		Name:     ident,
		Params:   params,
		Result:   result,
		Generic:  generic,
	}

	if onlyDeclare {
		//TO-DO check later call.delare ?
		return decl
	}

	if p.tok == LeftBrace {
		decl.Body = p.parseBody()
	}

	return decl
}

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
		return p.parseFuncDecl(m, false)

	default:
		pos := p.pos
		p.errorExpected(pos, "declaration")
		return &BadDecl{Start: pos}
	}
}
*/
