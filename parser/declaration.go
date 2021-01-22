package parser

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseVariable(modifier *ast.Modifier, attributes []*ast.Attribute, objectName string) *ast.Variable {
	d := &ast.Variable{}
	d.ObjectName = objectName
	d.Modifier = modifier
	d.Attributes = attributes
	d.Token = p.token
	p.next()
	d.Name = p.parseIdentifier()
	d.Type = p.parseType()

	if p.token == token.Assign {
		p.next()
		d.Value = p.parseExpression()
	}
	p.expect(token.Semi)
	return d
}

func (p *Parser) parseFunction(modifier *ast.Modifier, attributes []*ast.Attribute, objectName string) *ast.Function {
	d := &ast.Function{}
	d.ObjectName = objectName
	d.Modifier = modifier
	d.Attributes = attributes
	p.next()
	d.Name = p.parseIdentifier()
	if p.token == token.Less {
		d.TypeParameters = p.parseTypeParameters()
	}
	d.Parameters = p.parseParameters()
	if p.token != token.Semi && p.token != token.LeftBrace {
		d.ReturnType = p.parseType()
	}
	if p.token == token.LeftBrace {
		d.Body = p.parseBlockStatement()
	} else if p.token == token.Semi {
		p.next()
	}
	return d
}

func (p *Parser) parseEnum(modifier *ast.Modifier, attributes []*ast.Attribute) *ast.Enum {
	d := &ast.Enum{}
	d.Modifier = modifier
	d.Attributes = attributes
	p.next()
	d.Name = p.parseIdentifier()
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		v := &ast.Variable{}
		v.Name = p.parseIdentifier()
		v.ObjectName = d.Name.Name
		if p.token == token.Assign {
			p.next()
			v.Value = p.parseExpression()
		}
		if p.redeclared(v.Name.Name, d.Members) {
			p.error(v.Name.Position, fmt.Sprintf("enum %s redeclared", v.Name.Name))
		}
		d.Members = append(d.Members, v)
		if p.token != token.Comma {
			break
		}
		p.next()
	}
	p.expect(token.RightBrace)
	return d
}

func (p *Parser) parseInterface(modifier *ast.Modifier, attributes []*ast.Attribute) *ast.Interface {
	d := &ast.Interface{}
	d.Modifier = modifier
	d.Attributes = attributes
	p.next()
	d.Name = p.parseIdentifier()
	if p.token == token.Less {
		d.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		d.Parents = p.parseTypeNames()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Function:
			m := p.parseFunction(modifier, attr, d.Name.Name)
			if p.redeclared(m.Name.Name, d.Members) {
				p.error(m.Name.Position, fmt.Sprintf("function %s redeclared", m.Name.Name))
			}
			d.Members = append(d.Members, m)
		default:
			p.expectedError(p.position, "declaration")
		}
	}
	p.expect(token.RightBrace)
	return d
}

func (p *Parser) parseClass(modifier *ast.Modifier, attributes []*ast.Attribute) *ast.Class {
	d := &ast.Class{}
	d.Modifier = modifier
	d.Attributes = attributes
	p.next()
	d.Name = p.parseIdentifier()
	if p.token == token.Less {
		d.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		d.Parents = p.parseTypeNames()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			m := p.parseVariable(modifier, attr, d.Name.Name)
			if p.redeclared(m.Name.Name, d.Members) {
				p.error(m.Name.Position, fmt.Sprintf("variable %s redeclared", m.Name.Name))
			}
			d.Members = append(d.Members, m)

		case token.Function:
			m := p.parseFunction(modifier, attr, d.Name.Name)
			if p.redeclared(m.Name.Name, d.Members) {
				p.error(m.Name.Position, fmt.Sprintf("function %s redeclared", m.Name.Name))
			}
			d.Members = append(d.Members, m)

		default:
			p.expectedError(p.position, "declaration")
		}
	}
	p.expect(token.RightBrace)
	return d
}

func (p *Parser) parseModifier() *ast.Modifier {
	m := &ast.Modifier{}
	if p.token == token.Public {
		m.Public = true
		p.next()
	}
	if p.token == token.Weak {
		m.Weak = true
		p.next()
	}
	return m
}

func (p *Parser) parseAttributes() []*ast.Attribute {
	if p.token != token.META {
		return nil
	}
	var attr []*ast.Attribute
	for p.token == token.META {
		p.next()
		if p.token != token.IDENT {
			p.expect(token.IDENT)
		}
		m := &ast.Attribute{Position: p.position}
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
								p.error(p.position, "duplicated attribute "+name)
							}
							m.Values[name] = &ast.Literal{
								Typ:   p.token,
								Value: p.literal,
							}
							m.Values[name].Position = p.position
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
		attr = append(attr, m)
	}
	return attr
}
