package ast

import "github.com/panda-foundation/go-compiler/token"

// Statement interface for all statements
type Statement interface {
	statement()
}

func (*EmptyStatement) statement()            { return }
func (*BlockStatement) statement()            { return }
func (*ReturnStatement) statement()           { return }
func (*BranchStatement) statement()           { return }
func (*RawStatement) statement()              { return }
func (*TryStatement) statement()              { return }
func (*ThrowStatement) statement()            { return }
func (*DeclarationStatement) statement()      { return }
func (*IfStatement) statement()               { return }
func (*AssignStatement) statement()           { return }
func (*IncreaseDecreaseStatement) statement() { return }
func (*SwitchStatement) statement()           { return }
func (*CaseStatement) statement()             { return }
func (*WhileStatement) statement()            { return }
func (*ForStatement) statement()              { return }
func (*ForeachStatement) statement()          { return }

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
