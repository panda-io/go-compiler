package ast

import (
	"fmt"
	"strconv"

	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Enum struct {
	DeclarationBase
	Members []*Variable

	IRStruct        *ir.StructType
	IRStructData    *ir.Global
	VariableIndexes map[string]int
}

func (e *Enum) AddVariable(m *Variable) error {
	for _, v := range e.Members {
		if v.Name.Name == m.Name.Name {
			return fmt.Errorf("%s redeclared", m.Name.Name)
		}
	}
	e.Members = append(e.Members, m)
	return nil
}

func (e *Enum) GenerateIR(p *Program) {
	var index int64 = 0
	var types []ir.Type
	var values []ir.Constant
	for _, v := range e.Members {
		if v.Value == nil {
			types = append(types, ir.I32)
			values = append(values, ir.NewInt(ir.I32, index))
			index++
		} else {
			if literal, ok := v.Value.(*Literal); ok {
				if literal.Typ == token.INT {
					if i, _ := strconv.Atoi(literal.Value); int64(i) >= index {
						index = int64(i)
						types = append(types, ir.I32)
						values = append(values, ir.NewInt(ir.I32, index))
						index++
					} else {
						p.Error(v.Position, fmt.Sprintf("enum value here should be greater than %d.", i-1))
					}
				} else {
					p.Error(v.Position, "enum value must be integer.")
				}
			} else {
				p.Error(v.Position, "enum value must be integer.")
			}
		}
	}

	qualified := e.Qualified(p.Module.Namespace)
	e.IRStruct = ir.NewStructType(types...)
	p.IRModule.NewTypeDef(qualified, e.IRStruct)
	data := ir.NewStruct(CreateStruct(qualified), values...)
	e.IRStructData = p.IRModule.NewGlobalDef(qualified+".data", data)
}

func (e *Enum) HasMember(member string) bool {
	_, ok := e.VariableIndexes[member]
	return ok
}

func (e *Enum) GetMember(member string) ir.Constant {
	if index, ok := e.VariableIndexes[member]; ok {
		return ir.NewExprGetElementPtr(e.IRStruct, e.IRStructData, ir.NewInt(ir.I32, 0), ir.NewInt(ir.I32, int64(index)))
	}
	return nil
}
