package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type Continue struct {
	Base
}

func (*Continue) GenerateIR(*node.Context) {
	//TO-DO
}
