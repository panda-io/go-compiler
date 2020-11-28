package printer

import "github.com/panda-foundation/go-compiler/ast"

// Generator is the interface for all language exporters
type Generator interface {
	//TO-DO add project config as param(library path)
	Generate(program *ast.Program, file string)
}
