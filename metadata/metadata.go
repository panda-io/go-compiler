package metadata

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/scanner"
)

// Resolvers create resolvers
func Resolvers() []Resolver {
	var resolvers []Resolver
	resolvers = append(resolvers, &Document{})

	return resolvers
}

// Resolver to resolve interesting metadata
type Resolver interface {
	// init resolver with error handler
	Initialize(scanner.ErrorHandler)

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

func remove(metadata []*ast.Metadata, tag string) []*ast.Metadata {
	result := []*ast.Metadata{}
	for _, m := range metadata {
		if m.Name != tag {
			result = append(result, m)
		}
	}
	return result
}

func filter(metadata []*ast.Metadata, tag string) []*ast.Metadata {
	result := []*ast.Metadata{}
	for _, m := range metadata {
		if m.Name == tag {
			result = append(result, m)
		}
	}
	return result
}
