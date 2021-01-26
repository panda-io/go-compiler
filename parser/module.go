package parser

import (
	"fmt"
	"strings"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseSourceFile(file *token.File) {
	m := &ast.Module{
		File: file,
	}
	m.Attributes = p.parseAttributes()
	m.Namespace = p.parseNamespace()
	m.Imports = p.parseImports()

	for p.token != token.EOF {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			v := p.parseVariable(modifier, attr, "")
			if p.token == token.Const {
				v.Const = true
				if v.Value == nil {
					p.error(v.Position, "constant declaration must be initalized")
				}
			}
			qualified := m.Namespace + "." + v.Name.Name
			if p.program.Declarations[qualified] != nil {
				p.error(v.Name.Position, fmt.Sprintf("variable %s redeclared", v.Name.Name))
			}
			m.Variables = append(m.Variables, v)
			p.program.Declarations[qualified] = v

		case token.Function:
			f := p.parseFunction(modifier, attr, "")
			qualified := m.Namespace + "." + f.Name.Name
			if p.program.Declarations[qualified] != nil {
				p.error(f.Name.Position, fmt.Sprintf("function %s redeclared", f.Name.Name))
			}
			m.Functions = append(m.Functions, f)
			p.program.Declarations[qualified] = f

		case token.Enum:
			e := p.parseEnum(modifier, attr)
			qualified := m.Namespace + "." + e.Name.Name
			if p.program.Declarations[qualified] != nil {
				p.error(e.Name.Position, fmt.Sprintf("enum %s redeclared", e.Name.Name))
			}
			m.Enums = append(m.Enums, e)
			p.program.Declarations[qualified] = e

		case token.Interface:
			i := p.parseInterface(modifier, attr)
			qualified := m.Namespace + "." + i.Name.Name
			if p.program.Declarations[qualified] != nil {
				p.error(i.Name.Position, fmt.Sprintf("interface %s redeclared", i.Name.Name))
			}
			m.Interfaces = append(m.Interfaces, i)
			p.program.Declarations[qualified] = i

		case token.Class:
			c := p.parseClass(modifier, attr)
			qualified := m.Namespace + "." + c.Name.Name
			if p.program.Declarations[qualified] != nil {
				p.error(c.Name.Position, fmt.Sprintf("class %s redeclared", c.Name.Name))
			}
			m.Classes = append(m.Classes, c)
			p.program.Declarations[qualified] = c

		default:
			p.expectedError(p.position, "declaration")
		}
	}

	p.program.Modules[file.Name] = m
}

func (p *Parser) parseNamespace() string {
	p.expect(token.Namespace)
	if p.token == token.Semi {
		p.next()
		return ast.Global
	}
	namespace := p.parseQualifiedName("")
	p.expect(token.Semi)
	return namespace
}

func (p *Parser) parseImports() []*ast.Import {
	imports := []*ast.Import{}
	for p.token == token.Import {
		p.expect(token.Import)
		u := &ast.Import{}
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
