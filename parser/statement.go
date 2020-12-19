package parser

import (
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/token"
)

func (p *Parser) parseStatement() statement.Statement {
	switch p.token {
	case token.Break:
		s := &statement.Break{}
		s.Position = p.position
		p.next()
		p.expect(token.Semi)
		return s

	case token.Continue:
		s := &statement.Continue{}
		s.Position = p.position
		p.next()
		p.expect(token.Semi)
		return s

	case token.Return:
		s := &statement.Return{}
		s.Position = p.position
		p.next()
		if p.token != token.Semi {
			s.Expression = p.parseExpression()
		}
		p.expect(token.Semi)
		return s

	case token.Throw:
		s := &statement.Throw{}
		s.Position = p.position
		p.next()
		s.Expression = p.parseExpression()
		p.expect(token.Semi)
		return s

	case token.LeftBrace:
		return p.parseBlockStatement()

	case token.If:
		return p.parseIfStatement()

	case token.Switch:
		return p.parseSwitchStatement()

	case token.For:
		return p.parseForStatement()

	case token.Try:
		return p.parseTryStatement()

	default:
		return p.parseSimpleStatement(true)
	}
}

func (p *Parser) parseSimpleStatement(consumeSemi bool) statement.Statement {
	switch p.token {
	case token.Semi:
		s := &statement.Empty{}
		s.Position = p.position
		if consumeSemi {
			p.expect(token.Semi)
		}
		return s

	case token.Var:
		return p.parseDeclarationStatement(consumeSemi)

	case token.IDENT, token.This, token.Base, token.New,
		token.INT, token.FLOAT, token.CHAR, token.STRING, token.BOOL, token.NULL, token.Void,
		token.LeftParen, token.LeftBracket,
		token.Plus, token.Minus, token.Not, token.BitXor:
		position := p.position
		x := p.parseExpression()
		if consumeSemi {
			p.expect(token.Semi)
		}
		s := &statement.Expression{}
		s.Position = position
		s.Expression = x
		return s

	default:
		p.expectedError(p.position, "statement")
		return nil
	}
}

func (p *Parser) parseDeclarationStatement(consumeSemi bool) *statement.Declaration {
	s := &statement.Declaration{}
	s.Position = p.position
	p.next()
	s.Name = p.parseIdentifier()
	if p.token != token.Assign && p.token != token.Semi && p.token != token.Colon {
		s.Type = p.parseType()
	}
	if p.token == token.Assign {
		p.next()
		s.Value = p.parseExpression()
	}
	if consumeSemi {
		p.expect(token.Semi)
	}
	return s
}

func (p *Parser) parseBlockStatement() *statement.Block {
	s := &statement.Block{}
	s.Position = p.position
	p.next()
	for p.token != token.RightBrace {
		s.Statements = append(s.Statements, p.parseStatement())
	}
	p.next()
	return s
}

func (p *Parser) parseTryStatement() *statement.Try {
	s := &statement.Try{}
	s.Position = p.position
	p.next()
	s.Try = p.parseStatement()
	p.expect(token.Catch)
	s.Operand = p.parseParameters()
	s.Catch = p.parseStatement()
	if p.token == token.Finally {
		p.next()
		s.Finally = p.parseStatement()
	}
	return s
}

func (p *Parser) parseIfStatement() *statement.If {
	s := &statement.If{}
	p.next()
	p.expect(token.LeftParen)
	first := p.parseSimpleStatement(false)
	if p.token == token.Semi {
		p.next()
		s.Initialization = first
		s.Condition = p.parseSimpleStatement(false)
	} else {
		s.Condition = first
	}
	p.expect(token.RightParen)
	s.Body = p.parseStatement()
	if p.token == token.Else {
		p.next()
		if p.token == token.If {
			s.Else = p.parseIfStatement()
		} else {
			s.Else = p.parseStatement()
		}
	}
	return s
}

func (p *Parser) parseSwitchStatement() *statement.Switch {
	s := &statement.Switch{}
	s.Position = p.position
	p.next()
	p.expect(token.LeftParen)
	first := p.parseSimpleStatement(false)
	if p.token == token.Semi {
		p.next()
		s.Initialization = first
		s.Operand = p.parseSimpleStatement(false)
	} else {
		s.Operand = first
	}
	p.expect(token.RightParen)
	p.expect(token.LeftBrace)
	for p.token == token.Case || p.token == token.Default {
		s.Body = append(s.Body, p.parseCaseStatement())
	}
	p.expect(token.RightBrace)
	return s
}

func (p *Parser) parseCaseStatement() *statement.Case {
	s := &statement.Case{}
	s.Position = p.position
	s.Token = p.token
	if p.token == token.Case {
		p.next()
		s.Case = p.parseExpression()
	} else {
		p.expect(token.Default)
	}
	p.expect(token.Colon)
	s.Body = p.parseStatement()
	return s
}

func (p *Parser) parseForStatement() statement.Statement {
	position := p.position
	p.next()
	if p.token != token.LeftParen {
		s := &statement.For{}
		s.Position = position
		s.Body = p.parseStatement()
		return s
	} else {
		p.next()
		first := p.parseSimpleStatement(false)
		if p.token == token.RightParen {
			p.next()
			s := &statement.For{}
			s.Position = position
			s.Condition = first
			s.Body = p.parseStatement()
			return s
		} else if p.token == token.Colon {
			p.next()
			s := &statement.Foreach{}
			s.Position = position
			s.Item = first
			s.Iterator = p.parseExpression()
			p.expect(token.RightParen)
			s.Body = p.parseStatement()
			return s
		} else {
			p.expect(token.Semi)
			second := p.parseSimpleStatement(false)
			if p.token == token.Colon {
				p.next()
				s := &statement.Foreach{}
				s.Position = position
				s.Key = first
				s.Item = second
				s.Iterator = p.parseExpression()
				p.expect(token.RightParen)
				s.Body = p.parseStatement()
				return s
			} else if p.token == token.RightParen {
				p.next()
				s := &statement.For{}
				s.Position = position
				s.Initialization = first
				s.Condition = second
				s.Body = p.parseStatement()
				return s
			} else if p.token == token.Semi {
				p.next()
				s := &statement.For{}
				s.Position = position
				s.Initialization = first
				s.Condition = second
				if p.token != token.RightParen {
					s.Post = p.parseSimpleStatement(false)
				} else {
					e := &statement.Empty{}
					e.Position = p.position
					s.Post = e
				}
				p.expect(token.RightParen)
				s.Body = p.parseStatement()
				return s
			} else {
				p.error(position, "unknow token: "+p.token.String())
				return nil
			}
		}
	}
}
