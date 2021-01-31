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
	t := &ast.TypeArguments{}
	t.Position = p.position
	p.next() // skip <
	t.Arguments = append(t.Arguments, p.parseType())
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseType())
	}
	p.expect(token.Greater)
	return t
}

func (p *Parser) parseTypeParameters() *ast.TypeParameters {
	t := &ast.TypeParameters{}
	t.Position = p.position
	p.next() // skip <
	t.Parameters = append(t.Parameters, p.parseTypeParameter())
	for p.token == token.Comma {
		p.next()
		t.Parameters = append(t.Parameters, p.parseTypeParameter())
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
	for p.token == token.Comma {
		p.next()
		t.Parameters = append(t.Parameters, p.parseParameter())
	}
	if p.token == token.Ellipsis {
		t.Ellipsis = true
		p.next()
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
	t := &ast.Arguments{}
	t.Position = p.position
	p.expect(token.LeftParen)
	if p.token == token.RightParen {
		p.next()
		return t
	}
	t.Arguments = append(t.Arguments, p.parseExpression())
	for p.token == token.Comma {
		p.next()
		t.Arguments = append(t.Arguments, p.parseExpression())
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

func (p *Parser) tryParseTypeArguments() (*ast.TypeArguments, bool) {
	t := &ast.TypeArguments{}
	t.Position = p.position
	p.next() // skip <
	arg, success := p.tryParseTypeArgument()
	if !success {
		return nil, false
	}
	t.Arguments = append(t.Arguments, arg)
	for p.token == token.Comma {
		p.next()
		arg, success := p.tryParseTypeArgument()
		if !success {
			return nil, false
		}
		t.Arguments = append(t.Arguments, arg)
	}
	if p.token == token.Greater {
		p.next()
		return t, true
	}
	return nil, false
}

func (p *Parser) tryParseTypeArgument() (ast.Type, bool) {
	if p.token.IsScalar() {
		t := &ast.BuitinType{}
		t.Position = p.position
		t.Token = p.token
		p.next()
		return t, true
	}
	if p.token == token.IDENT {
		t := &ast.TypeName{}
		t.Position = p.position
		t.Name = p.parseIdentifier().Name
		if p.token == token.Dot {
			p.next()
			t.Selector = t.Name
			t.Name = p.parseIdentifier().Name
		}
		if p.token == token.Less {
			var success bool
			t.TypeArguments, success = p.tryParseTypeArguments()
			return t, success
		}
		return t, true
	}
	return nil, false
}
