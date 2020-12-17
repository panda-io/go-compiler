package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Identifier struct {
	Base
	Name string
}

func (e *Identifier) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Identifier) IsConstant() bool {
	//TO-DO search scope
	return false
}
