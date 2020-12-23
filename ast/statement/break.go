package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type Break struct {
	Base
}

func (*Break) GenerateIR(*node.Context) {
	//TO-DO
}
