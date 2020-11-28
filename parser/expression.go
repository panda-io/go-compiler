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

// ----------------------------------------------------------------------------
// Expressions

// parseOperand may return an expression or a raw type.
// Callers must verify the result.
// If lhs is set and the result is an identifier, it is not resolved.
//
/*
func (p *Parser) parseOperand(lhs bool) Expr {
	switch p.tok {
	case IDENT:
		x := p.parseIdent()
		return x

	case INT, FLOAT, CHAR, STRING, True, False, Void, Null:
		x := &BasicLit{Start: p.pos, Kind: p.tok, Value: p.lit}
		p.next()
		return x

	case LeftParen:
		start := p.pos
		p.next()
		p.exprLev++
		x := p.parseRhs() // types may be parenthesized: (some type)
		p.exprLev--
		p.expect(RightParen)
		return &ParenExpr{Start: start, Expr: x}
	}

	if typ := p.tryType(); typ != nil {
		// could be type for composite literal or conversion
		_, isIdent := typ.(*Ident)
		if isIdent {
			p.error(p.pos, "type cannot be identifier")
		}
		return typ
	}

	// we have an error
	pos := p.pos
	p.errorExpected(pos, "operand")
	return &BadExpr{Start: pos}
}

func (p *Parser) parseSelector(x Expr) Expr {
	return &SelectorExpr{Expr: x, Selector: p.parseIdent()}
}

func (p *Parser) parseIndex(x Expr) Expr {
	p.expect(LeftBracket)
	p.exprLev++
	var index Expr
	index = p.parseRhs()
	p.exprLev--
	p.expect(RightBracket)

	return &IndexExpr{Expr: x, Index: index}
}

func (p *Parser) parseCall(fun Expr) *CallExpr {
	p.expect(LeftParen)
	p.exprLev++
	var list []Expr
	var ellipsis = -1
	for p.tok != RightParen && p.tok != EOF {
		list = append(list, p.parseRhs()) // builtins may expect a type: make(some type, ...)
		if p.tok == Ellipsis {
			if ellipsis != -1 {
				p.errorExpected(p.pos, "only one ... can be used in a function call.")
			}
			ellipsis = len(list) - 1
			p.next()
		}
		if p.tok == Comma {
			p.next()
		}
	}
	p.exprLev--
	p.expect(RightParen)

	return &CallExpr{Func: fun, Args: list, Ellipsis: ellipsis}
}

func (p *Parser) parseValue(lhs bool) Expr {
	if p.tok == LeftBrace {
		return p.parseLiteralValue(nil)
	}
	x := p.parseExpr(lhs)
	return x
}

func (p *Parser) parseElement() Expr {
	x := p.parseValue(true)
	if p.tok == Colon {
		p.next()
		x = &KeyValueExpr{Key: x, Value: p.parseValue(false)}
	}
	return x
}

func (p *Parser) parseElementList() (list []Expr) {
	for p.tok != RightBrace && p.tok != EOF {
		list = append(list, p.parseElement())
		p.next()
	}
	return
}

func (p *Parser) parseLiteralValue(typ Expr) Expr {
	start := p.expect(LeftBrace)
	var elts []Expr
	p.exprLev++
	if p.tok != RightBrace {
		elts = p.parseElementList()
	}
	p.exprLev--
	p.expect(RightBrace)
	return &CompositeLit{Type: typ, Start: start, Values: elts}
}

// isTypeName reports whether x is a (qualified) TypeName.
func isTypeName(x Expr) bool {
	switch t := x.(type) {
	case *BadExpr:
	case *Ident:
	case *SelectorExpr:
		_, isIdent := t.Expr.(*Ident)
		return isIdent
	default:
		return false // all other nodes are not type names
	}
	return true
}

// isLiteralType reports whether x is a legal composite literal type.
func isLiteralType(x Expr) bool {
	switch t := x.(type) {
	case *BadExpr:
	case *Ident:
	case *SelectorExpr:
		_, isIdent := t.Expr.(*Ident)
		return isIdent
	default:
		return false // all other nodes are not legal composite literal types
	}
	return true
}

// If lhs is set and the result is an identifier, it is not resolved.
func (p *Parser) parsePrimaryExpr(lhs bool) Expr {
	x := p.parseOperand(lhs)
	for {
		switch p.tok {
		case Dot:
			p.next()
			switch p.tok {
			case IDENT:
				x = p.parseSelector(x)
			default:
				pos := p.pos
				p.errorExpected(pos, "selector")
				p.next() // make progress
				sel := &Ident{Start: pos}
				x = &SelectorExpr{Expr: x, Selector: sel}
			}
		case LeftBracket:
			x = p.parseIndex(x)
		case LeftParen:
			x = p.parseCall(x)
		case LeftBrace:
			if isLiteralType(x) && (p.exprLev >= 0 || !isTypeName(x)) {
				x = p.parseLiteralValue(x)
			} else {
				return x
			}
		default:
			return x
		}
		lhs = false // no need to try to resolve again
	}
	return x
}

// If lhs is set and the result is an identifier, it is not resolved.
func (p *Parser) parseUnaryExpr(lhs bool) Expr {
	switch p.tok {
	case Plus, Minus, Not, Caret, And:
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

func (p *Parser) parseRhs() Expr {
	old := p.inRhs
	p.inRhs = true
	x := p.parseExpr(false)
	p.inRhs = old
	return x
}
*/
