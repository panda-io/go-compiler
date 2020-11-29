package ast

import "github.com/panda-foundation/go-compiler/token"

// Statement interface for all statements
type Statement interface {
	statement()
}

func (*EmptyStatement) statement()            {}
func (*BlockStatement) statement()            {}
func (*ReturnStatement) statement()           {}
func (*BranchStatement) statement()           {}
func (*RawStatement) statement()              {}
func (*TryStatement) statement()              {}
func (*ThrowStatement) statement()            {}
func (*DeclarationStatement) statement()      {}
func (*IfStatement) statement()               {}
func (*AssignStatement) statement()           {}
func (*IncreaseDecreaseStatement) statement() {}
func (*SwitchStatement) statement()           {}
func (*CaseStatement) statement()             {}
func (*WhileStatement) statement()            {}
func (*ForStatement) statement()              {}
func (*ForeachStatement) statement()          {}

// EmptyStatement statement
type EmptyStatement struct {
}

// AssignStatement statement
type AssignStatement struct {
	Left  Expression
	Right Expression
}

// IncreaseDecreaseStatement statement
type IncreaseDecreaseStatement struct {
	Expression Expression
	Token      token.Token
}

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
	Condition Expression
	Body      *BlockStatement
	Else      Statement
}

// SwitchStatement statement
type SwitchStatement struct {
	Position int
	Tag      Expression
	Body     []*CaseStatement
}

// CaseStatement statement
type CaseStatement struct {
	Position int
	Case     Expression
	Body     *BlockStatement
}

// WhileStatement statement
type WhileStatement struct {
	Position  int
	Condition Expression
	Body      *BlockStatement
}

// ForStatement statement
type ForStatement struct {
	Position   int
	Initialize Statement
	Condition  Expression
	Post       Statement
	Body       *BlockStatement
}

// ForeachStatement statement
type ForeachStatement struct {
	Position int
	Key      Statement
	Value    Statement
	Iterator Expression
	Body     *BlockStatement
}
