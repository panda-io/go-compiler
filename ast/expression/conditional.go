package expression

import (
	"github.com/panda-foundation/go-compiler/ir"
)

type Conditional struct {
	Base
	Condition Expression
	First     Expression
	Second    Expression
}

func (e *Conditional) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Conditional) IsConstant() bool {
	return e.Condition.IsConstant() && e.First.IsConstant() && e.Second.IsConstant()
}
