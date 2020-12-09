package parser

import (
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseType() types.Type {
	if p.token.IsScalar() {
		t := &types.BuitinType{}
		t.Position = p.position
		t.Token = p.token
		p.next()
		return t
	}
	return p.parseTypeName()
}

func (p *Parser) parseTypeName() *types.TypeName {
	t := &types.TypeName{}
	t.Position = p.position
	t.Name = p.parseName("")
	if p.token == token.Less {
		t.TypeArguments = p.parseTypeArguments()
	}
	return t
}

func (p *Parser) parseTypeArguments() *types.TypeArguments {
	t := &types.TypeArguments{Ellipsis: -1}
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

func (p *Parser) parseTypeParameters() *types.TypeParameters {
	t := &types.TypeParameters{}
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

func (p *Parser) parseTypeParameter() *types.TypeParameter {
	t := &types.TypeParameter{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	if p.token == token.Colon {
		p.next()
		t.Type = p.parseType()
	}
	return t
}

func (p *Parser) parseIneritanceTypes() []*types.TypeName {
	p.next() // skip :
	result := []*types.TypeName{p.parseTypeName()}
	for p.token == token.Comma {
		p.next()
		result = append(result, p.parseTypeName())
	}
	return result
}

func (p *Parser) parseParameters() *types.Parameters {
	t := &types.Parameters{}
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

func (p *Parser) parseParameter() *types.Parameter {
	t := &types.Parameter{}
	t.Position = p.position
	t.Name = p.parseIdentifier().Name
	t.Type = p.parseType()
	return t
}

func (p *Parser) parseArguments() *types.Arguments {
	t := &types.Arguments{Ellipsis: -1}
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

func (p *Parser) parseName(identifier string) string {
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
