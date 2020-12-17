package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Increment struct {
	Base
	Expression Expression
}

func (e *Increment) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Increment) IsConstant() bool {
	return false
}

type Decrement struct {
	Base
	Expression Expression
}

func (e *Decrement) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Decrement) IsConstant() bool {
	return false
}
