package declaration

import (
	"github.com/panda-foundation/go-compiler/ast/types"
)

type Interface struct {
	Base
	TypeParameters *types.TypeParameters
	Parents        []*types.TypeName
	Members        []Declaration

	ResolvedParents []*Interface
}
