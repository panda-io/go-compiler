package parser

import (
	"crypto/md5"
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/scanner"
	"github.com/panda-foundation/go-compiler/token"
)

// NewParser create new parser
func NewParser(flags []string, program *ast.Program) *Parser {
	p := &Parser{
		program: program,
		scanner: scanner.NewScanner(flags),
	}
	return p
}

type Parser struct {
	position int
	token    token.Token
	literal  string

	program *ast.Program
	scanner *scanner.Scanner
}

func (p *Parser) ParseBytes(source []byte) {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	p.parseSourceFile(file)
}

func (p *Parser) ParseFile(file *token.File, source []byte) {
	p.setSource(file, source)
	p.parseSourceFile(file)
}

func (p *Parser) ParseExpression(source []byte) ast.Expression {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	return p.parseExpression()
}

func (p *Parser) ParseStatements(source []byte) ast.Statement {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	return p.parseBlockStatement()
}

func (p *Parser) next() {
	p.position, p.token, p.literal = p.scanner.Scan()
}

func (p *Parser) expect(t token.Token) {
	if p.token != t {
		p.expectedError(p.position, fmt.Sprintf("'%s'", t.String()))
	}
	p.next()
}

func (p *Parser) expectedError(position int, expect string) {
	expect = "expected " + expect
	if position == p.position {
		switch {
		case p.token == token.Semi && p.literal == "\n":
			expect += ", but found newline"
		case p.token.IsLiteral():
			expect += ", but found " + p.literal
		default:
			expect += ", but found '" + p.token.String() + "'"
		}
	}
	p.error(position, expect)
}

func (p *Parser) setSource(file *token.File, source []byte) {
	p.scanner.SetFile(file, source)
	p.next()
}

func (p *Parser) error(position int, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", p.scanner.Position(position).String(), message))
}
