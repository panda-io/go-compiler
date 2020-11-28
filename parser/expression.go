package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseExpression() ast.Expression {
	return nil
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

// ----------------------------------------------------------------------------
// Expressions

/*
// If lhs is set and the result is an identifier, it is not resolved.
func (p *Parser) parseUnaryExpr(lhs bool) Expr {
	switch p.tok {
	case Plus, Minus, Not, Caret:
		pos, op := p.pos, p.tok
		p.next()
		x := p.parseUnaryExpr(false)
		return &UnaryExpr{Start: pos, Op: op, Expr: x}
	}
	return p.parsePrimaryExpr(lhs)
}

func (p *Parser) tokPrec() (Token, int) {
	tok := p.tok
	if p.inRhs && tok == Assign {
		tok = Equal
	}
	return tok, tok.Precedence()
}

// If lhs is set and the result is an identifier, it is not resolved.
func (p *Parser) parseBinaryExpr(lhs bool, prec1 int) Expr {
	x := p.parseUnaryExpr(lhs)
	for {
		if p.tok == Semi {
			return x
		}

		op, oprec := p.tokPrec()
		if oprec < prec1 {
			return x
		}
		p.expect(op)
		if lhs {
			lhs = false
		}
		y := p.parseBinaryExpr(false, oprec+1)
		if op == Question {
			// TernaryExpr
			p.expect(Colon)
			z := p.parseBinaryExpr(false, oprec+1)
			x = &TernaryExpr{Condition: x, First: y, Second: z}
		} else {
			x = &BinaryExpr{Left: x, Op: op, Right: y}
		}
	}
}

// Continue
// If lhs is set and the result is an identifier, it is not resolved.
// The result may be a type or even a raw type ([...]int). Callers must
// check the result (using checkExpr or checkExprOrType), depending on
// context.
func (p *Parser) parseExpr(lhs bool) Expr {
	return p.parseBinaryExpr(lhs, LowestPrec+1)
}
*/
