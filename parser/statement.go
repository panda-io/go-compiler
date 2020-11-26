package parser

// ----------------------------------------------------------------------------
// Statements

// parseSimpleStmt returns true as 2nd result if it parsed the assignment
// of a range clause (with mode == rangeOk). The returned statement is an
// assignment with a right-hand side that is a single unary expression of
// the form "range x". No guarantees are given for the left-hand side.

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

func (p *Parser) parseReturnStmt() *ReturnStmt {
	pos := p.pos
	p.expect(Return)
	var result Expr
	if p.tok != Semi {
		result = p.parseExpr(false)
	}
	p.expect(Semi)
	return &ReturnStmt{Start: pos, Result: result}
}

func (p *Parser) parseBranchStmt(tok Token) *BranchStmt {
	pos := p.expect(tok)
	p.expect(Semi)
	return &BranchStmt{Start: pos, Tok: tok}
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

func (p *Parser) parseIfStmt() *IfStmt {
	pos := p.expect(If)

	cond := p.parseExpr(true)
	body := p.parseBlockStmt()

	var else_ Stmt
	if p.tok == Else {
		p.next()
		switch p.tok {
		case If:
			else_ = p.parseIfStmt()
		case LeftBrace:
			else_ = p.parseBlockStmt()
			p.expect(Semi)
		default:
			p.errorExpected(p.pos, "if statement or block")
			else_ = &BadStmt{Start: p.pos}
		}
	} else {
		p.expect(Semi)
	}

	return &IfStmt{Start: pos, Condition: cond, Body: body, Else: else_}
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
}

func (p *Parser) parseSwitchStmt() Stmt {
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

func (p *Parser) parseStmt() (s Stmt) {
	switch p.tok {
	case Const, Var:
		m := p.parseModifier()
		s = &DeclStmt{Decl: p.parseValueDecl(m)}
	case
		// tokens that may start an expression
		IDENT, INT, FLOAT, CHAR, STRING,
		LeftParen, LeftBracket,
		Plus, Minus, Star, And, Caret, Not: // unary operators
		s = p.parseSimpleStmt()
		p.expect(Semi)
	case Return:
		s = p.parseReturnStmt()
	case Break, Continue:
		s = p.parseBranchStmt(p.tok)
	case LeftBrace:
		s = p.parseBlockStmt()
		p.expect(Semi)
	case If:
		s = p.parseIfStmt()
	case Switch:
		s = p.parseSwitchStmt()
	case For:
		s = p.parseForStmt()
	case RightBrace:
		// a semicolon may be omitted before a closing "}"
		s = &EmptyStmt{Start: p.pos}
	default:
		// no statement found
		pos := p.pos
		p.errorExpected(pos, "statement")
		s = &BadStmt{Start: pos}
	}

	return
}

// ----------------------------------------------------------------------------
// Blocks
func (p *Parser) parseStmtList() (list []Stmt) {
	for {
		if len(p.cpp) > 0 {
			for _, e := range p.cpp {
				list = append(list, &EmitStmt{
					Start:   e.Pos(),
					Content: e.Text,
				})
			}
			p.cpp = p.cpp[:0]
		}
		if p.tok == RightBrace || p.tok == EOF {
			return
		}
		list = append(list, p.parseStmt())
	}
	return
}

func (p *Parser) parseBody() *BlockStmt {
	start := p.expect(LeftBrace)
	list := p.parseStmtList()
	p.expect(RightBrace)
	return &BlockStmt{Start: start, Stmts: list}
}

func (p *Parser) parseBlockStmt() *BlockStmt {
	start := p.expect(LeftBrace)
	list := p.parseStmtList()
	return &BlockStmt{Start: start, Stmts: list}
}
*/
