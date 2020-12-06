package parser

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseSourceFile() *ast.SoureFile {
	s := &ast.SoureFile{}
	s.Attributes = p.parseAttributes()
	s.Namespace = p.parseNamespace()
	s.Using = p.parseUsing()

	for p.token != token.EOF {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			m := p.parseVariable(modifier, attr)
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("variable %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		case token.Function:
			m := p.parseFunction(modifier, attr, nil)
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("function %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		case token.Enum:
			m := p.parseEnum(modifier, attr)
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("enum %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		case token.Interface:
			m := p.parseInterface(modifier, attr)
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("interface %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		case token.Class:
			m := p.parseClass(modifier, attr)
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("class %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		default:
			p.expectedError(p.position, "declaration")
		}
	}

	return s
}

func (p *Parser) parseNamespace() string {
	p.expect(token.Namespace)
	if p.token == token.Semi {
		p.next()
		return ""
	}
	namespace := p.parseQualifiedName("")
	p.expect(token.Semi)
	return namespace
}

func (p *Parser) parseUsing() []*ast.Using {
	using := []*ast.Using{}
	for p.token == token.Using {
		p.expect(token.Using)
		u := &ast.Using{}
		name := p.parseIdentifier()
		if p.token == token.Assign {
			u.Alias = name.Name
			p.next()
			name = p.parseIdentifier()
		}
		u.Namespace = p.parseQualifiedName(name.Name)
		p.expect(token.Semi)
		using = append(using, u)
	}
	return using
}
