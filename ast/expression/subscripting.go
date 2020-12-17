package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Subscripting struct {
	Base
	Parent  Expression
	Element Expression
}

func (e *Subscripting) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Subscripting) IsConstant() bool {
	//TO-DO search scope
	return false
}
