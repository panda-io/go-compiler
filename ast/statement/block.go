package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type Block struct {
	Base
	Statements []Statement
}

func (c *Block) GenerateIR(*node.Context) {
}
