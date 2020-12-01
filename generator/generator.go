package printer

import "github.com/panda-foundation/go-compiler/ast"

type Generator interface {
	//TO-DO add project options
	Generate(program *ast.Package, file string)
}
