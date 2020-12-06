package parser

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/statement"
	"github.com/panda-foundation/go-compiler/scanner"
	"github.com/panda-foundation/go-compiler/token"
)

// NewParser create new parser
func NewParser(flags []string) *Parser {
	p := &Parser{
		scanner: scanner.NewScanner(flags),
	}
	return p
}

type parserState struct {
	position int
	token    token.Token
	literal  string
}

type Parser struct {
	parserState
	scanner *scanner.Scanner
}

func (p *Parser) ParseExpression(source []byte) expression.Expression {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	return p.parseExpression()
}

func (p *Parser) ParseStatements(source []byte) statement.Statement {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	return p.parseCompoundStatement()
}

func (p *Parser) ParseBytes(source []byte) *ast.SoureFile {
	file := token.NewFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.setSource(file, source)
	return p.parseSourceFile()
}

func (p *Parser) ParseFile(file *token.File, path string) *ast.SoureFile {
	source, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	p.setSource(file, source)
	return p.parseSourceFile()
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

func (p *Parser) error(position int, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", p.scanner.Position(position).String(), message))
}

func (p *Parser) setSource(file *token.File, source []byte) {
	p.scanner.SetFile(file, source)
	p.next()
}

func (p *Parser) redeclared(name string, declarations []declaration.Declaration) bool {
	for _, d := range declarations {
		if d.Identifier() == name {
			return true
		}
	}
	return false
}
