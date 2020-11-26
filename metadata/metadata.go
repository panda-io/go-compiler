package metadata

import "github.com/panda-foundation/go-compiler/ast"

// Resolver to resolve interesting metadata
type Resolver interface {
	// before namespace (document, include)
	PackageMetadata([]*ast.Metadata, *ast.Program) []*ast.Metadata
	// before package variable declaration
	PackageVariable([]*ast.Metadata, *ast.Variable) []*ast.Metadata
	// before package function declaration
	PackageFunction([]*ast.Metadata, *ast.Function) []*ast.Metadata
	// before enum declaration
	Enum([]*ast.Metadata, *ast.Enum) []*ast.Metadata
	// before enum member declaration
	EnumMember([]*ast.Metadata, *ast.EnumMemberStatement) []*ast.Metadata
	// before interface declaration
	Interface([]*ast.Metadata, *ast.Interface) []*ast.Metadata
	// before interface variable declaration
	InterfaceVariable([]*ast.Metadata, *ast.Variable) []*ast.Metadata
	// before interface function declaration
	InterfaceFunction([]*ast.Metadata, *ast.Function) []*ast.Metadata
	// before class declaration
	Class([]*ast.Metadata, *ast.Class) []*ast.Metadata
	// before class variable declaration
	ClassVariable([]*ast.Metadata, *ast.Variable) []*ast.Metadata
	// before class function declaration
	ClassFunction([]*ast.Metadata, *ast.Function) []*ast.Metadata
}
