package printer

import "github.com/panda-foundation/go-compiler/ast"

// Printer is the interface for all language exporters
type Printer interface {
	//TO-DO project config (library path)
	PrintSource(program *ast.Program, folder string)
	PrintBinary(program *ast.Program, file string)
}
