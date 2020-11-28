package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseType() ast.Type {
	if p.token.IsScalar() {
		t := &ast.BuitinType{
			Position: p.position,
			Token:    p.token,
		}
		p.next()
		return t
	}
	return p.parseTypeName()
}

func (p *Parser) parseTypeName() *ast.TypeName {
	t := &ast.TypeName{
		QualifiedName: p.parseQualifiedName(nil),
	}
	if p.token == token.Less {
		t.TypeArguments = p.parseTypeArguments()
	}
	return t
}

func (p *Parser) parseTypeArguments() *ast.TypeArguments {
	p.next()
	t := &ast.TypeArguments{
		Position: p.position,
		Ellipsis: -1,
	}
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
	p.next()
	t := &ast.TypeParameters{
		Position: p.position,
	}
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
	t := &ast.TypeParameter{
		Name: p.parseIdentifier(),
	}
	if p.token == token.Colon {
		p.next()
		t.Type = p.parseType()
	}
	return t
}

func (p *Parser) parseBaseTypes() *ast.TypeArguments {
	p.next()
	t := &ast.TypeArguments{
		Position: p.position,
	}
	t.Arguments = append(t.Arguments, p.parseType())
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseType())
	}
	return t
}

func (p *Parser) parseParameters() *ast.Parameters {
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return nil
	}
	t := &ast.Parameters{
		Position: p.position,
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

func (p *Parser) parseParameter() *ast.Variable {
	t := &ast.Variable{
		Name: p.parseIdentifier(),
	}
	t.Type = p.parseType()
	if p.token == token.Assign {
		p.next()
		if !p.token.IsLiteral() || p.token == token.IDENT {
			p.error(p.position, "variable can only be initialized by const value (string, char, float, int)")
		}
		t.Value = &ast.Literal{
			Position: p.position,
			Type:     p.token,
			Value:    p.literal,
		}
		p.next()
	}
	return t
}

func (p *Parser) parseArguments() *ast.Arguments {
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return nil
	}
	t := &ast.Arguments{
		Position: p.position,
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
