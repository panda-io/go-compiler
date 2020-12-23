package statement

import "github.com/panda-foundation/go-compiler/ast/node"

type Block struct {
	Base
	Statements []Statement
}

func (b *Block) GenerateIR(c *node.Context) {
	for _, stmt := range b.Statements {
		stmt.GenerateIR(c)
	}
}
