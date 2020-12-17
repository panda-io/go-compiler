package expression

import (
	"github.com/panda-foundation/go-compiler/ast/types"
	"github.com/panda-foundation/go-compiler/ir"
)

type This struct {
	Base
}

func (e *This) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *This) IsConstant() bool {
	//TO-DO search scope
	return false
}

type Super struct {
	Base
}

func (e *Super) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *Super) IsConstant() bool {
	//TO-DO search scope
	return false
}

type New struct {
	Base
	Type      types.Type
	Arguments *types.Arguments
}

func (e *New) GenerateIR() ir.Value {
	//TO-DO
	return nil
}

func (e *New) IsConstant() bool {
	return false
}
