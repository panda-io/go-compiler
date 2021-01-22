package parser

import (
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
	e := &ast.Enum{}
	e.Modifier = modifier
	e.Attributes = attributes
	p.next()
	e.Name = p.parseIdentifier()
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		v := &ast.Variable{}
		v.Name = p.parseIdentifier()
		v.ObjectName = e.Name.Name
		if p.token == token.Assign {
			p.next()
			v.Value = p.parseExpression()
		}
		err := e.AddVariable(v)
		if err != nil {
			p.error(v.Position, err.Error())
		}
		if p.token != token.Comma {
			break
		}
		p.next()
	}
	p.expect(token.RightBrace)
	return e
}

func (p *Parser) parseInterface(modifier *ast.Modifier, attributes []*ast.Attribute) *ast.Interface {
	i := &ast.Interface{}
	i.Modifier = modifier
	i.Attributes = attributes
	p.next()
	i.Name = p.parseIdentifier()
	if p.token == token.Less {
		i.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		i.Parents = p.parseTypeNames()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Function:
			f := p.parseFunction(modifier, attr, i.Name.Name)
			err := i.AddFunction(f)
			if err != nil {
				p.error(f.Position, err.Error())
			}
		default:
			p.expectedError(p.position, "function declaration")
		}
	}
	p.expect(token.RightBrace)
	return i
}

func (p *Parser) parseClass(modifier *ast.Modifier, attributes []*ast.Attribute) *ast.Class {
	c := &ast.Class{}
	c.Modifier = modifier
	c.Attributes = attributes
	p.next()
	c.Name = p.parseIdentifier()
	if p.token == token.Less {
		c.TypeParameters = p.parseTypeParameters()
	}
	if p.token == token.Colon {
		c.Parents = p.parseTypeNames()
	}
	p.expect(token.LeftBrace)
	for p.token != token.RightBrace {
		attr := p.parseAttributes()
		modifier := p.parseModifier()
		switch p.token {
		case token.Const, token.Var:
			v := p.parseVariable(modifier, attr, c.Name.Name)
			err := c.AddVariable(v)
			if err != nil {
				p.error(v.Position, err.Error())
			}

		case token.Function:
			f := p.parseFunction(modifier, attr, c.Name.Name)
			err := c.AddFunction(f)
			if err != nil {
				p.error(f.Position, err.Error())
			}

		default:
			p.expectedError(p.position, "member declaration")
		}
	}
	p.expect(token.RightBrace)
	return c
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
