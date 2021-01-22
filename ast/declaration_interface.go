package ast

type Interface struct {
	DeclarationBase
	TypeParameters *TypeParameters
	Parents        []*TypeName
	Members        []Declaration

	ResolvedParents []*Interface
}
