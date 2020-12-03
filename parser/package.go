package parser

import (
	"fmt"
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseProgram() {
	attr := p.parseAttributes()
	n := p.parseNamespace()
	program := p.findPackage(n)
	if len(attr) > 0 {
		program.Custom = append(program.Custom, attr...)
		attr = attr[:0]
	}

	if p.token == token.Using {
		p.parseUsing()
	}

	for p.token != token.EOF {
		attr = p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			m := p.parseVariable(modifier, attr)
			if p.redeclared(m.Name.Name, program.Members) {
				p.error(m.Name.Position, fmt.Sprintf("variable %s redeclared", m.Name.Name))
			}
			program.Members = append(program.Members, m)

		case token.Function:
			m := p.parseFunction(modifier, attr, nil)
			if p.redeclared(m.Name.Name, program.Members) {
				p.error(m.Name.Position, fmt.Sprintf("function %s redeclared", m.Name.Name))
			}
			program.Members = append(program.Members, m)

		case token.Enum:
			m := p.parseEnum(modifier, attr)
			if p.redeclared(m.Name.Name, program.Members) {
				p.error(m.Name.Position, fmt.Sprintf("enum %s redeclared", m.Name.Name))
			}
			program.Members = append(program.Members, m)

		case token.Interface:
			m := p.parseInterface(modifier, attr)
			if p.redeclared(m.Name.Name, program.Members) {
				p.error(m.Name.Position, fmt.Sprintf("interface %s redeclared", m.Name.Name))
			}
			program.Members = append(program.Members, m)

		case token.Class:
			m := p.parseClass(modifier, attr)
			if p.redeclared(m.Name.Name, program.Members) {
				p.error(m.Name.Position, fmt.Sprintf("class %s redeclared", m.Name.Name))
			}
			program.Members = append(program.Members, m)

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
