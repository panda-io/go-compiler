package parser

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseStatementBlock() *ast.BlockStatement {
	block := &ast.BlockStatement{
		Position: p.position,
	}
	p.next()
	for p.token != token.RightBrace {
		block.Statements = append(block.Statements, p.parseStatement())
	}
	return block
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.token {
	case token.IDENT, token.This, token.Base: // unary operators
		//s = p.parseSimpleStmt() // parsePrimaryExpression
		//invocation,[assign, ++, -- (check left must be variable)]
		//p.expect(Semi)

	case token.Var:
		return p.parseDeclarationStatement()

	case token.Return:
		return p.parseReturnStatement()

	case token.Break, token.Continue:
		return p.parseBranchStatement()

	case token.LeftBrace:
		return p.parseStatementBlock()

	case token.If:
		return p.parseIfStatement()

	case token.Switch:
		//s = p.parseSwitchStmt()

	case token.For:
		//s = p.parseForStmt()

	case token.Try:
		return p.parseTryStatement()

	case token.Throw:
		return p.parseThrowStatement()

	case token.Yield:
		// yield statement //TO-DO

	case token.Await:
		// await //TO-DO

	case token.META:
		return p.parseRawStatement()

	default:
		p.expectedError(p.position, "statement")
	}
	return nil
}

func (p *Parser) parseDeclarationStatement() *ast.DeclarationStatement {
	s := &ast.DeclarationStatement{
		Position: p.position,
	}
	p.next()
	s.Name = p.parseIdentifier()
	if p.token != token.Equal && p.token != token.Semi {
		s.Type = p.parseType()
	}
	if p.token == token.Equal {
		s.Value = p.parseExpression()
	}
	p.expect(token.Semi)
	return s
}

func (p *Parser) parseBranchStatement() *ast.BranchStatement {
	s := &ast.BranchStatement{
		Position: p.position,
		Token:    p.token,
	}
	p.next()
	p.expect(token.Semi)
	return s
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	s := &ast.ReturnStatement{
		Position: p.position,
	}
	p.next()
	if p.token != token.Semi {
		s.Expression = p.parseExpression()
	}
	p.expect(token.Semi)
	return s
}

func (p *Parser) parseThrowStatement() *ast.ThrowStatement {
	s := &ast.ThrowStatement{
		Position: p.position,
	}
	p.next()
	s.Expression = p.parseExpression()
	p.expect(token.Semi)
	return s
}

func (p *Parser) parseRawStatement() *ast.RawStatement {
	s := &ast.RawStatement{
		Position: p.position,
	}
	p.next()
	if p.token != token.STRING {
		p.expectedError(p.position, "raw source (string)")
	}
	s.Source = p.literal
	p.next()
	return s
}

func (p *Parser) parseTryStatement() *ast.TryStatement {
	s := &ast.TryStatement{
		Position: p.position,
	}
	p.next()
	s.TryStatement = p.parseStatementBlock()
	p.expect(token.Catch)
	p.expect(token.LeftParen)
	s.Catch = p.parseParameter()
	s.CatchStatement = p.parseStatementBlock()
	if p.token == token.Finally {
		p.next()
		s.FinallyStatement = p.parseStatementBlock()
	}
	return s
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	s := &ast.IfStatement{
		Position: p.position,
	}
	p.next()
	s.Condition = p.parseExpression()
	s.Body = p.parseStatementBlock()

	if p.token == token.Else {
		p.next()
		if p.token == token.If {
			s.Else = p.parseIfStatement()
		} else if p.token == token.LeftBrace {
			s.Else = p.parseStatementBlock()
		} else {
			p.expectedError(p.position, "if statement or block")
		}
	}
	return s
}

/*
func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	pos := p.expect(Switch)

	tag := p.parseSimpleStmt()
	bodyStart := p.expect(LeftBrace)
	var list []Stmt
	for p.tok == Case || p.tok == Default {
		list = append(list, p.parseCaseClause())
	}
	p.expect(RightBrace)
	body := &BlockStmt{Start: bodyStart, Stmts: list}

	return &SwitchStmt{Start: pos, Tag: p.makeExpr(tag, "switch expression"), Body: body}
}

func (p *Parser) parseCaseClause() *CaseClause {
	pos := p.pos
	var expr Expr
	if p.tok == Case {
		p.next()
		expr = p.parseRhs()
	} else {
		p.expect(Default)
	}

	p.expect(Colon)
	body := p.parseStmtList()

	return &CaseClause{Start: pos, Expr: expr, Body: body}
}*/

// ----------------------------------------------------------------------------
// Statements

/*
func (p *Parser) parseSimpleStmt() Stmt {
	x := p.parseExpr(true)

	switch p.tok {
	case
		Assign, PlusAssign,
		MinusAssign, MulAssign, DivAssign,
		ModAssign, AndAssign, OrAssign,
		XorAssign, LeftShiftAssign, RightShiftAssign:
		// assignment statement, possibly part of a range clause
		tok := p.tok
		p.next()
		y := p.parseRhs()
		as := &AssignStmt{Left: x, Tok: tok, Right: y}
		return as
	}

	switch p.tok {
	case PlusPlus, MinusMinus:
		// increment or decrement
		s := &IncDecStmt{Expr: x, Tok: p.tok}
		p.next()
		return s
	}

	// expression
	return &ExprStmt{Expr: x}
}


func (p *Parser) makeExpr(s Stmt, want string) Expr {
	if s == nil {
		return nil
	}
	if es, isExpr := s.(*ExprStmt); isExpr {
		return es.Expr
	}
	found := "simple statement"
	if _, isAss := s.(*AssignStmt); isAss {
		found = "assignment"
	}
	p.error(s.Pos(), fmt.Sprintf("expected %s, found %s (missing parentheses around composite literal?)", want, found))
	return &BadExpr{Start: s.Pos()}
}

func (p *Parser) parseForStmt() Stmt {
	pos := p.expect(For)

	var s1, s2, s3 Stmt
	//var isRange bool
	if p.tok != LeftBrace {
		prevLev := p.exprLev
		p.exprLev = -1
		if p.tok != Semi {
			s2 = p.parseSimpleStmt()
		}
		if p.tok == Semi {
			p.next()
			s1 = s2
			s2 = nil
			if p.tok != Semi {
				s2 = p.parseSimpleStmt()
			}
			p.expect(Semi)
			if p.tok != LeftBrace {
				s3 = p.parseSimpleStmt()
			}
		}
		p.exprLev = prevLev
	}

	body := p.parseBlockStmt()
	p.expect(Semi)

	// regular for statement
	return &ForStmt{
		For:       pos,
		Init:      s1,
		Condition: p.makeExpr(s2, "boolean"),
		Post:      s3,
		Body:      body,
	}
}


*/
