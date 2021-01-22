package ast

import (
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Interface struct {
	DeclarationBase
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration

	ResolvedParents []*Interface
}
