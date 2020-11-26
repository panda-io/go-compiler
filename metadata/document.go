package metadata

import (
	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/scanner"
)

const documentTag = "doc"

// Document resolver for metadata
type Document struct {
	errorHandler scanner.ErrorHandler
}

// Initialize document metadata resolver
func (d *Document) Initialize(handler scanner.ErrorHandler) {
	d.errorHandler = handler
}

// PackageMetadata resolve package document
func (d *Document) PackageMetadata(metadata []*ast.Metadata, node *ast.Program) []*ast.Metadata {
	documents := filter(metadata, documentTag)
	if len(documents) > 1 {
		d.errorHandler(documents[0].Position, "dupicated package document")
	} else if len(documents) == 1 {
		if node.Document != nil {
			d.errorHandler(documents[0].Position, "package document redeclared")
		}
		node.Document = documents[0]
	}
	return remove(metadata, documentTag)
}

// PackageVariable resolve package variable document
func (*Document) PackageVariable(metadata []*ast.Metadata, node *ast.Variable) []*ast.Metadata {
	return metadata
}

// PackageFunction resolve package function document
func (*Document) PackageFunction(metadata []*ast.Metadata, node *ast.Function) []*ast.Metadata {
	return metadata
}

// Enum resolve enum document
func (*Document) Enum(metadata []*ast.Metadata, node *ast.Enum) []*ast.Metadata { return metadata }

// EnumMember resolve enum member document
func (*Document) EnumMember(metadata []*ast.Metadata, node *ast.EnumMemberStatement) []*ast.Metadata {
	return metadata
}

// Interface resolve interface document
func (*Document) Interface(metadata []*ast.Metadata, node *ast.Interface) []*ast.Metadata {
	return metadata
}

// InterfaceVariable resolve interface variable document
func (*Document) InterfaceVariable(metadata []*ast.Metadata, node *ast.Variable) []*ast.Metadata {
	return metadata
}

// InterfaceFunction resolve interface function document
func (*Document) InterfaceFunction(metadata []*ast.Metadata, node *ast.Function) []*ast.Metadata {
	return metadata
}

// Class resolve class document
func (*Document) Class(metadata []*ast.Metadata, node *ast.Class) []*ast.Metadata {
	return metadata
}

// ClassVariable resolve class variable document
func (*Document) ClassVariable(metadata []*ast.Metadata, node *ast.Variable) []*ast.Metadata {
	return metadata
}

// ClassFunction resolve class function document
func (*Document) ClassFunction(metadata []*ast.Metadata, node *ast.Function) []*ast.Metadata {
	return metadata
}
