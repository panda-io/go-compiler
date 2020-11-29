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
	p.next()
	return block
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.token {
	case token.Semi:
		p.next()
		return &ast.EmptyStatement{}

	case token.Var, token.IDENT, token.This, token.Base:
		return p.parseSimpleStatement()

	case token.Return:
		return p.parseReturnStatement()

	case token.Break, token.Continue:
		return p.parseBranchStatement()

	case token.LeftBrace:
		return p.parseStatementBlock()

	case token.If:
		return p.parseIfStatement()

	case token.Switch:
		return p.parseSwitchStatement()

	case token.While:
		return p.parseWhileStatement()

	case token.For:
		return p.parseForStatement()

	case token.Foreach:
		return p.parseForeachStatement()

	case token.Try:
		return p.parseTryStatement()

	case token.Throw:
		return p.parseThrowStatement()

	case token.Yield:
		// yield statement //TO-DO
		p.next()
		return nil

	case token.Await:
		// await //TO-DO
		p.next()
		return nil

	case token.META:
		return p.parseRawStatement()

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseSimpleStatement() ast.Statement {
	switch p.token {
	case token.IDENT, token.This, token.Base:
		return p.parseAssignStatement(false)

	case token.Var:
		return p.parseDeclarationStatement()

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseAssignStatement(skipSemi bool) ast.Statement {
	e := p.parsePrimaryExpression()
	switch p.token {
	case token.Assign,
		token.PlusAssign, token.MinusAssign, token.MulAssign, token.DivAssign,
		token.ModAssign, token.AndAssign, token.OrAssign,
		token.XorAssign, token.LeftShiftAssign, token.RightShiftAssign:
		s := &ast.AssignStatement{
			Left: e,
		}
		p.next()
		s.Right = p.parseExpression()
		if !skipSemi {
			p.expect(token.Semi)
		}
		return s

	case token.PlusPlus, token.MinusMinus:
		s := &ast.IncreaseDecreaseStatement{
			Expression: e,
			Token:      p.token,
		}
		p.next()
		if !skipSemi {
			p.expect(token.Semi)
		}
		return s

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseDeclarationStatement() *ast.DeclarationStatement {
	s := &ast.DeclarationStatement{
		Position: p.position,
	}
	p.next()
	s.Name = p.parseIdentifier()
	if p.token != token.Assign && p.token != token.Semi && p.token != token.Colon {
		s.Type = p.parseType()
	}
	if p.token == token.Assign {
		p.next()
		s.Value = p.parseExpression()
	}
	if p.token == token.Semi {
		p.next()
	}
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
	p.expect(token.RightParen)
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
	p.expect(token.LeftParen)
	s.Condition = p.parseExpression()
	p.expect(token.RightParen)
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

func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	s := &ast.SwitchStatement{
		Position: p.position,
	}
	p.next()
	s.Tag = p.parseExpression()
	p.expect(token.LeftBrace)
	for p.token == token.Case || p.token == token.Default {
		s.Body = append(s.Body, p.parseCaseStatement())
	}
	p.expect(token.RightBrace)
	return s
}

func (p *Parser) parseCaseStatement() *ast.CaseStatement {
	s := &ast.CaseStatement{
		Position: p.position,
	}
	if p.token == token.Case {
		p.next()
		s.Case = p.parseExpression()
	} else {
		p.expect(token.Default)
	}
	p.expect(token.Colon)
	s.Body = p.parseStatementBlock()
	return s
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	s := &ast.WhileStatement{
		Position: p.position,
	}
	p.next()
	p.expect(token.LeftParen)
	s.Condition = p.parseExpression()
	p.expect(token.RightParen)
	s.Body = p.parseStatementBlock()
	return s
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	s := &ast.ForStatement{
		Position: p.position,
	}
	p.next()
	p.expect(token.LeftParen)
	if p.token != token.Semi {
		s.Initialize = p.parseSimpleStatement()
	} else {
		p.next()
	}
	if p.token != token.Semi {
		s.Condition = p.parseExpression()
	}
	p.expect(token.Semi)
	if p.token != token.RightParen {
		s.Post = p.parseAssignStatement(true)
	}
	p.expect(token.RightParen)
	s.Body = p.parseStatementBlock()
	return s
}

func (p *Parser) parseForeachStatement() *ast.ForeachStatement {
	s := &ast.ForeachStatement{
		Position: p.position,
	}
	p.next()
	p.expect(token.LeftParen)
	s.Value = p.parseDeclarationStatement()
	if p.token != token.Colon {
		s.Key = s.Value
		s.Value = p.parseSimpleStatement()
	}
	p.expect(token.Colon)
	s.Iterator = p.parseExpression()
	p.expect(token.RightParen)
	s.Body = p.parseStatementBlock()
	return s
}
