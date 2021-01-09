package parser

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseSourceFile() *ast.Module {
	s := &ast.Module{}
	s.Attributes = p.parseAttributes()
	s.Namespace = p.parseNamespace()
	s.Imports = p.parseImports()

	for p.token != token.EOF {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			m := p.parseVariable(modifier, attr, "")
			if p.redeclared(m.Name.Name, s.Members) {
				p.error(m.Name.Position, fmt.Sprintf("variable %s redeclared", m.Name.Name))
			}
			s.Members = append(s.Members, m)

		case token.Function:
			m := p.parseFunction(modifier, attr, "")
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
		return node.Global
	}
	namespace := p.parseQualifiedName("")
	p.expect(token.Semi)
	return namespace
}

func (p *Parser) parseImports() []*node.Import {
	imports := []*node.Import{}
	for p.token == token.Import {
		p.expect(token.Import)
		u := &node.Import{}
		name := p.parseIdentifier()
		if p.token == token.Assign {
			u.Alias = name.Name
			p.next()
			name = p.parseIdentifier()
		}
		u.Namespace = p.parseQualifiedName(name.Name)
		if u.Alias == "" {
			names := strings.Split(u.Namespace, ".")
			u.Alias = names[len(names)-1]
		}
		p.expect(token.Semi)
		imports = append(imports, u)
	}
	return imports
}

func (p *Parser) parseQualifiedName(identifier string) string {
	if identifier == "" {
		identifier = p.parseIdentifier().Name
	}
	qualifiedName := identifier
	for p.token == token.Dot {
		p.next()
		qualifiedName += "." + p.parseIdentifier().Name
	}
	return qualifiedName
}
