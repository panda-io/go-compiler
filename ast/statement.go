package ast

import "github.com/panda-foundation/go-compiler/token"

// Statement interface for all statements
type Statement interface {
	statement()
}

func (*BlockStatement) statement()       {}
func (*ReturnStatement) statement()      {}
func (*BranchStatement) statement()      {}
func (*RawStatement) statement()         {}
func (*TryStatement) statement()         {}
func (*ThrowStatement) statement()       {}
func (*DeclarationStatement) statement() {}
func (*IfStatement) statement()          {}

// BlockStatement is statement list
type BlockStatement struct {
	Position   int
	Statements []Statement
}

// RawStatement contain raw souce code for target language
type RawStatement struct {
	Position int
	Source   string
}

// ReturnStatement return statement
type ReturnStatement struct {
	Position   int
	Expression Expression
}

// ThrowStatement return statement
type ThrowStatement struct {
	Position   int
	Expression Expression
}

// BranchStatement branch statement
type BranchStatement struct {
	Position int
	Token    token.Token
}

// TryStatement try statement
type TryStatement struct {
	Position         int
	TryStatement     *BlockStatement
	Catch            *Variable
	CatchStatement   *BlockStatement
	FinallyStatement *BlockStatement
}

// DeclarationStatement local declaration statement
type DeclarationStatement struct {
	Position int
	Name     *Identifier
	Type     Type
	Value    Expression
}

// IfStatement represents an if statement.
type IfStatement struct {
	Position  int
	Condition Expression // condition
	Body      *BlockStatement
	Else      Statement
}
