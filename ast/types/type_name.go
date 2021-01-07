package types

import "github.com/panda-foundation/go-compiler/ir"

type TypeName struct {
	Base
	Name          string
	QualifiedName string
	TypeArguments *TypeArguments
}

func (*TypeName) GenerateIR() ir.Type {
	//TO-DO use shared_ptr and weak_ptr
	return nil
}
