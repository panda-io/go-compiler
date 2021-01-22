package ast

import "github.com/panda-foundation/go-compiler/token"

type Import struct {
	NodeBase
	Alias     string
	Namespace string
}

type Module struct {
	File *token.File

	Namespace string
	Imports   []*Import

	Attributes []*Attribute
	Variables  []*Variable
	Functions  []*Function
	Enums      []*Enum
	Interfaces []*Interface
	Classes    []*Class
}
