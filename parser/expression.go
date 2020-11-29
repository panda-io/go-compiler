package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseExpression() ast.Expression {
	return p.parseBinaryExpression(1)
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	position := p.position
	name := ""
	if p.token == token.IDENT {
		name = p.literal
		p.next()
	} else {
		p.expect(token.IDENT)
	}
	return &ast.Identifier{
		Position: position,
		Name:     name,
	}
}

func (p *Parser) parseOperand() ast.Expression {
	switch p.token {
	case token.IDENT:
		return p.parseIdentifier()

	case token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void:
		e := &ast.Literal{
			Position: p.position,
			Type:     p.token,
			Value:    p.literal,
		}
		p.next()
		return e

	case token.This, token.Base:
		e := &ast.ClassReference{
			Position: p.position,
			Type:     p.token,
		}
		p.next()
		return e

	case token.New:
		e := &ast.ClassCreateExpression{}
		p.next()
		e.Type = p.parseType()
		e.Arguments = p.parseArguments()
		return e

	case token.LeftParen:
		e := &ast.ParenExpression{
			Position: p.position,
		}
		p.next()
		e.Expression = p.parseExpression()
		p.expect(token.RightParen)
		return e
	}

	if p.token.IsScalar() {
		e := &ast.BuitinType{
			Position: p.position,
			Token:    p.token,
		}
		p.next()
		return e
	}
	p.error(p.position, "unexpected "+p.token.String())
	return nil
}

func (p *Parser) parsePrimaryExpression() ast.Expression {
	e := p.parseOperand()
	for {
		switch p.token {
		case token.Dot:
			p.next()
			m := &ast.MemberAccessExpression{
				Parent: e,
				Member: p.parseIdentifier(),
			}
			e = m

		case token.LeftBracket:
			p.next()
			m := &ast.ElementAccessExpression{
				Parent:  e,
				Element: p.parseExpression(),
			}
			e = m
			p.expect(token.RightBracket)
			return e

		case token.LeftParen:
			p.next()
			f := &ast.InvocationExpression{
				Function: e,
			}
			f.Arguments = p.parseArguments()
			return f

		default:
			return e
		}
	}
}

func (p *Parser) parseUnaryExpression() ast.Expression {
	switch p.token {
	case token.Plus, token.Minus, token.Not, token.Caret:
		e := &ast.UnaryExpression{
			Position: p.position,
			Operator: p.token,
		}
		p.next()
		e.Expression = p.parseUnaryExpression()
		return e
	}
	return p.parsePrimaryExpression()
}

func (p *Parser) parseBinaryExpression(precedence int) ast.Expression {
	x := p.parseUnaryExpression()
	for {
		op := p.token
		if op == token.Semi {
			return x
		}
		opPrec := p.token.Precedence()
		if opPrec < precedence {
			return x
		}
		p.next()
		y := p.parseBinaryExpression(p.token.Precedence() + 1)
		if op == token.Question {
			p.expect(token.Colon)
			z := p.parseBinaryExpression(opPrec + 1)
			return &ast.TernaryExpression{
				Condition: x,
				First:     y,
				Second:    z,
			}
		}
		return &ast.BinaryExpression{
			Left:     x,
			Operator: op,
			Right:    y,
		}
	}
}
