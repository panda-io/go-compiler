package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type Empty struct {
	Base
}

func (*Empty) GenerateIR(*node.Context) {
	//TO-DO
}
