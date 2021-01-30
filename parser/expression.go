package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseExpression() ast.Expression {
	return p.parseBinaryExpression(0)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	e := &ast.Identifier{}
	e.Position = p.position
	if p.token == token.IDENT {
		e.Name = p.literal
		p.next()
	} else {
		p.expect(token.IDENT)
	}
	return e
}

func (p *Parser) parseOperand() ast.Expression {
	switch p.token {
	case token.IDENT:
		return p.parseIdentifier()

	case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void:
		e := &ast.Literal{}
		e.Position = p.position
		e.Typ = p.token
		e.Value = p.literal
		p.next()
		return e

	case token.This:
		e := &ast.This{}
		e.Position = p.position
		p.next()
		return e

	case token.Base:
		e := &ast.Base{}
		e.Position = p.position
		p.next()
		return e

	case token.LeftParen:
		e := &ast.Parentheses{}
		e.Position = p.position
		p.next()
		e.Expression = p.parseExpression()
		p.expect(token.RightParen)
		return e

	default:
		p.error(p.position, "unexpected "+p.token.String())
		return nil
	}
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	x := p.parseOperand()
	for {
		switch p.token {
		case token.Dot:
			e := &ast.MemberAccess{}
			e.Position = p.position
			p.next()
			e.Parent = x
			e.Member = p.parseIdentifier()
			x = e

		case token.LeftBracket:
			e := &ast.Subscripting{}
			e.Position = p.position
			p.next()
			e.Element = p.parseExpression()
			p.expect(token.RightBracket)
			x = e

		case token.LeftParen:
			e := &ast.Invocation{}
			e.Position = p.position
			e.Function = x
			e.Arguments = p.parseArguments()
			x = e

		case token.PlusPlus:
			e := &ast.Increment{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		case token.MinusMinus:
			e := &ast.Decrement{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		default:
			return x
		}
	}
}

func (p *Parser) parseUnaryExpression() ast.Expression {
	switch p.token {
	case token.Plus, token.Minus, token.Not, token.Complement:
		e := &ast.Unary{}
		e.Position = p.position
		e.Operator = p.token
		p.next()
		e.Expression = p.parseUnaryExpression()
		return e

	case token.New:
		e := &ast.New{}
		e.Position = p.position
		p.next()
		e.Typ = p.parseType()
		e.Arguments = p.parseArguments()
		return e

	default:
		return p.parsePrimaryExpression()
	}
}

func (p *Parser) parseBinaryExpression(precedence int) ast.Expression {
	x := p.parseUnaryExpression()
	for {
		if p.token == token.Semi {
			return x
		}
		op := p.token
		opPrec := p.token.Precedence()
		if opPrec <= precedence {
			return x
		}
		p.next()
		y := p.parseBinaryExpression(opPrec)
		x = &ast.Binary{
			Left:     x,
			Operator: op,
			Right:    y,
		}
	}
}
