package parser

import (
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseExpression() expression.Expression {
	return p.parseBinaryExpression(0)
}

func (p *Parser) parseIdentifier() *expression.Identifier {
	e := &expression.Identifier{}
	e.Position = p.position
	if p.token == token.IDENT {
		e.Name = p.literal
		p.next()
	} else {
		p.expect(token.IDENT)
	}
	return e
}

func (p *Parser) parseOperand() expression.Expression {
	switch p.token {
	case token.IDENT:
		return p.parseIdentifier()

	case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void:
		e := &expression.Literal{}
		e.Position = p.position
		e.Type = p.token
		e.Value = p.literal
		p.next()
		return e

	case token.This:
		e := &expression.This{}
		e.Position = p.position
		p.next()
		return e

	case token.Base:
		e := &expression.Super{}
		e.Position = p.position
		p.next()
		return e

	case token.LeftParen:
		e := &expression.Parentheses{}
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

func (p *Parser) parsePrimaryExpression() expression.Expression {
	x := p.parseOperand()
	for {
		switch p.token {
		case token.Dot:
			e := &expression.MemberAccess{}
			e.Position = p.position
			p.next()
			e.Parent = x
			e.Member = p.parseIdentifier()
			x = e

		case token.LeftBracket:
			e := &expression.Subscripting{}
			e.Position = p.position
			p.next()
			e.Element = p.parseExpression()
			p.expect(token.RightBracket)
			x = e

		case token.LeftParen:
			e := &expression.Invocation{}
			e.Position = p.position
			e.Function = x
			e.Arguments = p.parseArguments()
			x = e

		case token.PlusPlus:
			e := &expression.Increment{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		case token.MinusMinus:
			e := &expression.Decrement{}
			e.Position = p.position
			e.Expression = x
			p.next()
			return e

		default:
			return x
		}
	}
}

func (p *Parser) parseUnaryExpression() expression.Expression {
	switch p.token {
	case token.Plus, token.Minus, token.Not, token.Complement:
		e := &expression.Unary{}
		e.Position = p.position
		e.Operator = p.token
		p.next()
		e.Expression = p.parseUnaryExpression()
		return e

	case token.New:
		e := &expression.New{}
		e.Position = p.position
		p.next()
		e.Type = p.parseType()
		e.Arguments = p.parseArguments()
		return e

	default:
		return p.parsePrimaryExpression()
	}
}

func (p *Parser) parseBinaryExpression(precedence int) expression.Expression {
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
		if op == token.Question {
			p.expect(token.Colon)
			z := p.parseBinaryExpression(opPrec)
			x = &expression.Conditional{
				Condition: x,
				First:     y,
				Second:    z,
			}
		} else {
			x = &expression.Binary{
				Left:     x,
				Operator: op,
				Right:    y,
			}
		}
	}
}
