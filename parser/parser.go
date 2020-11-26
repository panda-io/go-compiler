package parser

import (
	"github.com/panda-foundation/go-compiler/scanner"
	"github.com/panda-foundation/go-compiler/token"
)

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
	flags   []string
	files   *token.FileSet
	scanner *scanner.Scanner
	root    *Program
}
