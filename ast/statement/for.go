package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type For struct {
	Base
	Initialization Statement
	Condition      Statement
	Post           Statement
	Body           Statement
}

func (*For) GenerateIR(*node.Context) {
	//TO-DO
}
