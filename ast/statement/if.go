package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type If struct {
	Base
	Initialization Statement
	Condition      Statement
	Body           Statement
	Else           Statement
}

func (*If) GenerateIR(*node.Context) {
	//TO-DO
}
