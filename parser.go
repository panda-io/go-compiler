package main

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the exported entry points for invoking the parser.

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//TO-DO too many errors

func NewParser(scanComments bool, flags []string) *Parser {
	return &Parser{
		scanComments: scanComments,
		flags:        flags,
		root: &Program{
			Children: make(map[string]*Program),
		},
	}
}

// The parser structure holds the parser's internal state.
type Parser struct {
	scanComments bool
	flags        []string
	scanner      *Scanner

	cpp []*Metadata

	// Next token
	pos int    // token position
	tok Token  // one token look-ahead
	lit string // token literal

	// Non-syntactic parser control
	exprLev int  // < 0: in control clause, >= 0: in expression // TO-DO refactor later
	inRhs   bool // if set, the parser is parsing a rhs expression

	root *Program
}

func (p *Parser) ParseString(content string) {
	file := NewFile("source")
	p.parse(file, []byte(content), p.scanComments, p.flags)
}

func (p *Parser) ParseFile(fileName string) {
	// get source
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	file := NewFile(fileName)
	p.parse(file, text, p.scanComments, p.flags)
}

func (p *Parser) ParseFolder(folder string) {
	fd, err := os.Open(folder)
	if err != nil {
		panic(err)
	}

	list, err := fd.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, d := range list {
		if d.IsDir() {
			p.ParseFolder(filepath.Join(folder, d.Name()))
		} else {
			if strings.HasSuffix(d.Name(), ".pd") {
				filename := filepath.Join(folder, d.Name())
				p.ParseFile(filename)
			}
		}
	}
}

func (p *Parser) GetProgram() *Program {
	return p.root
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()

	//TO-DO remove meta here, to the position call this // process like modifier
	for p.tok == META {
		metaData := p.parseMetadata()
		for _, v := range metaData {
			if v.Name == MetaCpp {
				p.cpp = append(p.cpp, v)
			} else {
				//TO-DO unknow meta //ignore doc currently
			}
		}
	}
}

func (p *Parser) error(pos int, msg string) {
	errPos := p.scanner.file.Position(pos)
	fmt.Println("file:", errPos.FileName)
	fmt.Println("line:", errPos.Line)
	fmt.Println("column:", errPos.Column)
	fmt.Println("error:", msg)

	panic("pare file error")
}

func (p *Parser) errorExpected(pos int, msg string) {
	msg = "expected " + msg
	if pos == p.pos {
		// the error happened at the current position;
		// make the error message more specific
		switch {
		case p.tok == Semi && p.lit == "\n":
			msg += ", found newline"
		case p.tok.IsLiteral():
			msg += ", found " + p.lit
		default:
			msg += ", found '" + p.tok.String() + "'"
		}
	}
	p.error(pos, msg)
}

func (p *Parser) expect(tok Token) int {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, "'"+tok.String()+"'")
	}
	p.next() // make progress
	return pos
}

var stmtStart = map[Token]bool{
	Break:    true,
	Catch:    true,
	Const:    true,
	Continue: true,
	For:      true,
	If:       true,
	Return:   true,
	Switch:   true,
	Try:      true,
	Var:      true,
}

var declStart = map[Token]bool{
	Const:     true,
	Class:     true,
	Enum:      true,
	Import:    true,
	Interface: true,
	Public:    true,
	Static:    true,
	Var:       true,
}

var exprEnd = map[Token]bool{
	Comma:        true,
	Colon:        true,
	Semi:         true,
	RightParen:   true,
	RightBrace:   true,
	RightBracket: true,
}

// ----------------------------------------------------------------------------
// Identifiers

func (p *Parser) parseIdent() *Ident {
	pos := p.pos
	name := ""
	if p.tok == IDENT {
		name = p.lit
		p.next()
	} else {
		p.expect(IDENT) // use expect() error handling
	}
	return &Ident{Start: pos, Name: name}
}

func (p *Parser) parseModifier() *Modifier {
	m := &Modifier{}
	if p.tok == Public {
		m.Public = true
		p.next()
	}
	if p.tok == Static {
		m.Static = true
		p.next()
	}
	if p.tok == Async {
		m.Async = true
		p.next()
	}
	return m
}

func (p *Parser) parseGeneric() *GenericLit {
	if p.tok == Less {
		g := &GenericLit{
			Start: p.pos,
		}
		p.next()
		g.Types = append(g.Types, p.tryType())

		for p.tok == Comma {
			p.next()
			g.Types = append(g.Types, p.tryType())
		}
		p.expect(Greater)
		return g

	}
	return nil
}

func (p *Parser) parseMetadata() []*Metadata {
	if p.tok != META {
		return nil
	}
	var meta []*Metadata
	for p.tok == META {
		p.next()

		if p.tok != IDENT {
			p.expect(IDENT)
		}
		m := &Metadata{Start: p.pos}
		m.Name = p.lit
		p.next()

		if p.tok == LeftParen {
			p.next()
			if p.tok == STRING {
				m.Text = p.lit
				p.next()
			} else {
				m.Values = make(map[string]*BasicLit)
				for {
					if p.tok == IDENT {
						name := p.lit
						p.next()

						p.expect(Assign)
						switch p.tok {
						case INT, FLOAT, CHAR, STRING, True, False:
							//TO-DO check if duplicated
							m.Values[name] = &BasicLit{
								Start: p.pos,
								Kind:  p.tok,
								Value: p.lit,
							}
						default:
							p.errorExpected(p.pos, "basic literal (bool, char, int, float, string)")
						}
						p.next()

						if p.tok == RightParen {
							break
						}
						p.expect(Comma)
					} else {
						p.expect(IDENT)
					}
				}
			}
			p.expect(RightParen)
		}
		meta = append(meta, m)
	}
	return meta
}

// ----------------------------------------------------------------------------
// Types

func (p *Parser) parseType() Expr {
	typ := p.tryType()

	if typ == nil {
		pos := p.pos
		p.errorExpected(pos, "type")
		return &BadExpr{Start: pos}
	}

	return typ
}

// If the result is an identifier, it is not resolved.
func (p *Parser) parseTypeName() Expr {
	ident := p.parseIdent()
	// don't resolve ident yet - it may be a parameter or field name
	if p.tok == Dot {
		// ident is a package name
		p.next()
		sel := p.parseIdent()
		return &SelectorExpr{Expr: ident, Selector: sel}
	}
	return ident
}

// If the result is an identifier, it is not resolved.
func (p *Parser) tryVarType(isParam bool) Expr {
	if isParam && p.tok == Ellipsis {
		pos := p.pos
		p.next()
		typ := p.tryType() // don't use parseType so we can provide better error message
		if typ == nil {
			p.error(pos, "'...' parameter is missing type")
			typ = &BadExpr{Start: pos}
		}
		return &EllipsisLit{Start: pos, Expr: typ}
	}
	return p.tryType()
}

// If the result is an identifier, it is not resolved.
func (p *Parser) parseVarType(isParam bool) Expr {
	typ := p.tryVarType(isParam)
	if typ == nil {
		pos := p.pos
		p.errorExpected(pos, "type")
		p.next() // make progress
		typ = &BadExpr{Start: pos}
	}
	return typ
}

func (p *Parser) parseParameterList() (params []*Field) {
	for p.tok != RightParen {
		field := &Field{}
		typ := p.parseVarType(true)
		if p.tok == Comma {
			field.Name = nil
			field.Type = typ
			p.expect(Comma)
			params = append(params, field)
			continue
		}
		ok := false
		field.Name, ok = typ.(*Ident)
		if !ok {
			p.expect(IDENT)
		}
		field.Type = p.parseVarType(true)
		if p.tok == Equal {
			p.next()
			field.Default = p.parseExpr(false)
		}
		params = append(params, field)
		if p.tok != RightParen {
			p.expect(Comma)
		}
	}
	//TP-DO ...
	//TO-DO check default
	return
}

func (p *Parser) parseParameters() *FieldList {
	var params []*Field
	start := p.expect(LeftParen)
	if p.tok != RightParen {
		params = p.parseParameterList()
	}
	p.expect(RightParen)

	return &FieldList{Start: start, Fields: params}
}

func (p *Parser) parseResult() *Field {
	typ := p.tryType()
	if typ != nil {
		return &Field{Type: typ}
	}
	return &Field{Type: &Scalar{Token: Void}}
}

// If the result is an identifier, it is not resolved.
func (p *Parser) tryType() Expr {
	if p.tok.IsScalar() {
		scalar := &Scalar{
			Start: p.pos,
			Token: p.tok,
		}
		p.next()
		return scalar

	} else if p.tok == IDENT {
		typ := p.parseTypeName()
		return typ
	}
	return nil
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

// ----------------------------------------------------------------------------
// Expressions

// parseOperand may return an expression or a raw type.
// Callers must verify the result.
// If lhs is set and the result is an identifier, it is not resolved.
//
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

func (p *Parser) parseIdentOrSelector(existing Expr) Expr {
	if existing == nil {
		existing = p.parseIdent()
	}
	full := existing
	for p.tok == Dot {
		p.next()
		full = p.parseSelector(full)
	}
	return full
}

// ----------------------------------------------------------------------------
// Statements

// parseSimpleStmt returns true as 2nd result if it parsed the assignment
// of a range clause (with mode == rangeOk). The returned statement is an
// assignment with a right-hand side that is a single unary expression of
// the form "range x". No guarantees are given for the left-hand side.
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
		fmt.Println("token:", p.tok.String(), "literal:", p.lit)
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
// Declarations

func (p *Parser) parseNamespaceDecl() *NamespaceDecl {
	if p.tok != Namespace {
		return &NamespaceDecl{}
	}

	p.expect(Namespace)

	path := p.parseIdentOrSelector(nil)
	p.expect(Semi)

	spec := &NamespaceDecl{
		Path: path,
	}

	return spec
}

func (p *Parser) parseImportDecl() []*ImportDecl {
	var imports []*ImportDecl
	for p.tok == Import {
		p.expect(Import)

		var importDecl = &ImportDecl{}
		importDecl.Path = p.parseIdent()

		if p.tok == Assign {
			importDecl.Name = importDecl.Path.(*Ident)
			p.next()
			importDecl.Path = p.parseIdent()
		}

		importDecl.Path = p.parseIdentOrSelector(importDecl.Path)
		p.expect(Semi)

		// collect imports
		imports = append(imports, importDecl)
	}
	return imports
}

func (p *Parser) parseValueDecl(m *Modifier) *ValueDecl {
	p.next()
	name := p.parseIdent()
	typ := p.tryType()

	decl := &ValueDecl{
		Modifier: m,
		Name:     name,
		Type:     typ,
	}

	pos := p.pos
	// always permit optional initialization for more tolerant parsing
	if p.tok == Assign {
		p.next()
		decl.Value = p.parseRhs()
	}
	p.expect(Semi) // call before accessing p.linecomment

	if decl.Value == nil && decl.Type == nil {
		p.error(pos, "missing type or initialization")
		//TO-DO if type is nil, parse type from value

	}
	return decl
}

func (p *Parser) parseEnumDecl(m *Modifier) *EnumDecl {
	p.next()
	name := p.parseIdent()

	decl := &EnumDecl{
		Modifier: m,
		Name:     name,
	}

	p.expect(LeftBrace)
	var list []*EnumStmt
	for p.tok != RightBrace {
		member := &EnumStmt{
			Name: p.parseIdent(),
		}
		if p.tok == Assign {
			p.next()
			//later check there are same values
			if p.tok == INT {
				member.Value = &BasicLit{
					Start: p.pos,
					Kind:  INT,
					Value: p.lit,
				}
			} else {
				p.error(p.pos, "only int value can assign to enum")
			}
			p.next()
		}
		list = append(list, member)
		if p.tok != RightBrace {
			p.expect(Comma)
		}
	}
	p.expect(RightBrace)
	decl.List = list
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseInterfaceDecl(m *Modifier) *InterfaceDecl {
	p.next()
	name := p.parseIdent()

	decl := &InterfaceDecl{
		Modifier: m,
		Name:     name,
	}

	//TO-DO generic

	p.expect(LeftBrace)
	for p.tok != RightBrace {
		switch p.tok {
		case Const, Var:
			decl.Values = append(decl.Values, p.parseValueDecl(nil))

		case Function:
			f := p.parseFuncDecl(m, true)
			f.IsMember = true
			f.InterfaceMember = true
			if f.Name.Name == name.Name {
				p.error(f.Name.Pos(), "interface has no contructor")
			}
			if f.Name.Name[0] == '~' {
				p.error(f.Name.Pos(), "interface has no destructor")
			}
			p.expect(Semi)
			decl.Functions = append(decl.Functions, f)
		default:
			pos := p.pos
			p.errorExpected(pos, "declaration")
			//p.advance(sync)
			//TO-DO advance wrong part
		}
	}
	p.expect(RightBrace)
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseClassDecl(m *Modifier) *ClassDecl {
	p.next()
	name := p.parseIdent()

	decl := &ClassDecl{
		Modifier: m,
		Name:     name,
	}

	//TO-DO generic
	if p.tok == Colon {
		for p.tok != LeftBrace {
			decl.Parents = append(decl.Parents, p.parseIdentOrSelector(nil))
			if p.tok == Comma {
				p.next()
			}
		}
	}

	p.expect(LeftBrace)
	for p.tok != RightBrace {
		m = p.parseModifier()
		switch p.tok {
		case Const, Var:
			decl.Values = append(decl.Values, p.parseValueDecl(m))

		case Function:
			f := p.parseFuncDecl(m, false)
			f.IsMember = true
			f.ClassName = name.Name
			if f.Name.Name == name.Name {
				f.IsConstructor = true
			}
			if f.Name.Name[0] == '~' {
				if f.Name.Name == "~"+name.Name {
					f.IsDestructor = true
				} else {
					p.error(f.Name.Pos(), "invalid destructor name")
				}
				//TO-DO check return type. must be void
			}
			decl.Functions = append(decl.Functions, f)
		default:
			pos := p.pos
			p.errorExpected(pos, "declaration")
			//p.advance(sync)
			//TO-DO advance wrong part
		}
	}
	p.expect(RightBrace)
	//TO-DO check later call.delare ?
	return decl
}

func (p *Parser) parseFuncDecl(m *Modifier, onlyDeclare bool) *FuncDecl {
	p.expect(Function)

	//Tilde
	tilde := false
	if p.tok == Tilde {
		tilde = true
		p.next()
	}
	ident := p.parseIdent()
	if tilde {
		//TO-DO check in class
		ident.Name = "~" + ident.Name
	}
	generic := p.parseGeneric()
	params := p.parseParameters()
	result := p.parseResult()

	decl := &FuncDecl{
		Modifier: m,
		Name:     ident,
		Params:   params,
		Result:   result,
		Generic:  generic,
	}

	if onlyDeclare {
		//TO-DO check later call.delare ?
		return decl
	}

	if p.tok == LeftBrace {
		decl.Body = p.parseBody()
	}

	return decl
}

func (p *Parser) parseDecl(sync map[Token]bool) Decl {
	m := p.parseModifier()
	switch p.tok {
	case Const, Var:
		return p.parseValueDecl(m)

	case Enum:
		return p.parseEnumDecl(m)

	case Interface:
		return p.parseInterfaceDecl(m)

	case Class:
		return p.parseClassDecl(m)

	case Function:
		return p.parseFuncDecl(m, false)

	default:
		pos := p.pos
		p.errorExpected(pos, "declaration")
		return &BadDecl{Start: pos}
	}
}

// ----------------------------------------------------------------------------
// Source files
func (p *Parser) parse(file *File, src []byte, scanComments bool, flags []string) {

	eh := func(pos Position, msg string) { p.error(p.pos, msg) }
	p.scanner = NewScanner(file, src, eh, scanComments, flags)
	p.next()

	program := p.root
	namespace := p.parseNamespaceDecl()
	if namespace.Path != nil {
		program = p.root.FindPackage(namespace.Path)
	}

	// import
	program.Imports = append(program.Imports, p.parseImportDecl()...)

	// rest of namespace body
	for p.tok != EOF {
		decl := p.parseDecl(declStart)
		switch v := decl.(type) {
		case *ValueDecl:
			program.Values = append(program.Values, v)
		case *EnumDecl:
			program.Enums = append(program.Enums, v)
		case *InterfaceDecl:
			program.Interfaces = append(program.Interfaces, v)
		case *ClassDecl:
			program.Classes = append(program.Classes, v)
		case *FuncDecl:
			program.Functions = append(program.Functions, v)
		case *BadDecl:
			fmt.Println("bad decl", v.Pos())
		}
	}
}
