package statement

import "github.com/panda-foundation/go-compiler/ir"

type Block struct {
	Base
	Statements []Statement
}

func (e *Block) GenerateIR() ir.Value {
	//block := ir.NewBlock("")

	return nil
}
