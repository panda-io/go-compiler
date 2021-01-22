package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseType() ast.Type {
	if p.token.IsScalar() {
		t := &ast.BuitinType{}
		t.Position = p.position
		t.Token = p.token
		p.next()
		return t
	}
	if p.token == token.Function {
		p.next()
		return p.parseFunctionType()
	}
	return p.parseTypeName()
}

func (p *Parser) parseTypeName() *ast.TypeName {
	t := &ast.TypeName{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	if p.token == token.Dot {
		p.next()
		t.Selector = t.Name
		t.Name = p.parseIdentifier().Name
	}
	if p.token == token.Less {
		t.TypeArguments = p.parseTypeArguments()
	}
	return t
}

func (p *Parser) parseTypeArguments() *ast.TypeArguments {
	t := &ast.TypeArguments{Ellipsis: -1}
	t.Position = p.position
	t.Ellipsis = -1
	p.next() // skip <
	t.Arguments = append(t.Arguments, p.parseType())
	if p.token == token.Ellipsis {
		t.Ellipsis = 0
		p.next()
	}
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseType())
		if p.token == token.Ellipsis {
			if t.Ellipsis > -1 {
				p.error(p.position, "dupicate ellipsis")
			}
			t.Ellipsis = len(t.Arguments) - 1
			p.next()
		}
	}
	p.expect(token.Greater)
	return t
}

func (p *Parser) parseTypeParameters() *ast.TypeParameters {
	t := &ast.TypeParameters{}
	t.Position = p.position
	p.next() // skip <
	t.Parameters = append(t.Parameters, p.parseTypeParameter())
	if p.token == token.Ellipsis {
		t.Ellipsis = true
		p.next()
	}
	for p.token == token.Comma {
		if t.Ellipsis {
			p.error(p.position, "ellipsis must be in the last position")
		}
		p.next()
		t.Parameters = append(t.Parameters, p.parseTypeParameter())
		if p.token == token.Ellipsis {
			t.Ellipsis = true
			p.next()
		}
	}
	p.expect(token.Greater)
	return t
}

func (p *Parser) parseTypeParameter() *ast.TypeParameter {
	t := &ast.TypeParameter{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	if p.token == token.Colon {
		p.next()
		t.Type = p.parseType()
	}
	return t
}

func (p *Parser) parseTypeNames() []*ast.TypeName {
	p.next() // skip :
	result := []*ast.TypeName{p.parseTypeName()}
	for p.token == token.Comma {
		p.next()
		result = append(result, p.parseTypeName())
	}
	return result
}

func (p *Parser) parseParameters() *ast.Parameters {
	t := &ast.Parameters{}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return nil
	}
	t.Parameters = append(t.Parameters, p.parseParameter())
	if p.token == token.Ellipsis {
		t.Ellipsis = true
		p.next()
	}
	for p.token == token.Comma {
		if t.Ellipsis {
			p.error(p.position, "ellipsis must be in the last position")
		}
		p.next()
		t.Parameters = append(t.Parameters, p.parseParameter())
	}
	p.expect(token.RightParen)
	return t
}

func (p *Parser) parseParameter() *ast.Parameter {
	t := &ast.Parameter{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	t.Type = p.parseType()
	return t
}

func (p *Parser) parseArguments() *ast.Arguments {
	t := &ast.Arguments{Ellipsis: -1}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return t
	}
	t.Arguments = append(t.Arguments, p.parseExpression())
	if p.token == token.Ellipsis {
		t.Ellipsis = 0
		p.next()
	}
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseExpression())
		if p.token == token.Ellipsis {
			if t.Ellipsis > -1 {
				p.error(p.position, "dupicate ellipsis")
			}
			t.Ellipsis = len(t.Arguments) - 1
			p.next()
		}
	}
	p.expect(token.RightParen)
	return t
}

func (p *Parser) parseFunctionType() *ast.TypeFunction {
	t := &ast.TypeFunction{}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return t
	}
	t.Parameters = append(t.Parameters, p.parseType())
	for p.token == token.Comma {
		p.next()
		t.Parameters = append(t.Parameters, p.parseType())
	}
	p.expect(token.RightParen)
	if p.token != token.Semi && p.token != token.Assign {
		t.ReturnType = p.parseType()
	}
	return t
}
