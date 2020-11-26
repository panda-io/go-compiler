package parser

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/scanner"
	"github.com/panda-foundation/go-compiler/token"
)

// NewParser create new parser
func NewParser(flags []string) *Parser {
	p := &Parser{
		files: &token.FileSet{},
		root:  ast.NewProgram("", nil),
	}
	p.scanner = scanner.NewScanner(p.error, flags)

	return p
}

type parserState struct {
	position int
	token    token.Token
	literal  string

	exprLevel   int
	inRightHand bool
}

// Parser to parse panda source
type Parser struct {
	parserState
	files   *token.FileSet
	scanner *scanner.Scanner
	root    *ast.Program
}

// ParseBytes parse string source
func (p *Parser) ParseBytes(source []byte) {
	file := p.files.AddFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.parse(file, source)
}

// ParseFile parse from file
func (p *Parser) ParseFile(fileName string) {
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	file := p.files.AddFile(fileName, len(source))
	p.parse(file, source)
}

// ParseFolder parse files under folder (include sub folders)
func (p *Parser) ParseFolder(folder string) {
	folderInfo, err := os.Open(folder)
	if err != nil {
		panic(err)
	}
	list, err := folderInfo.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, f := range list {
		if f.IsDir() {
			p.ParseFolder(filepath.Join(folder, f.Name()))
		} else {
			if strings.HasSuffix(f.Name(), ".pd") {
				filename := filepath.Join(folder, f.Name())
				p.ParseFile(filename)
			}
		}
	}
}

func (p *Parser) next() {
	p.position, p.token, p.literal = p.scanner.Scan()
}

func (p *Parser) expect(t token.Token) {
	if p.token != t {
		p.unexpected(p.position, fmt.Sprintf("'%s'", t.String()))
	}
}

func (p *Parser) unexpected(position int, expect string) {
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
	panic(fmt.Sprintf("error: %s \n %s \n", p.files.Position(position).String(), message))
}

func (p *Parser) parse(file *token.File, source []byte) {
	p.scanner.SetFile(file, source)
	p.next()
	p.parseProgram()
}
