package generator

import (
	"fmt"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/token"
)

const (
	tabSize = 4
)

func Write(program *ast.Program, fileset *token.FileSet, file string) {

	//TO-DO print main at the last
	/*
		err := ioutil.WriteFile(file, w.buffer.Bytes(), 0644)
		if err != nil {
			panic(err)
		}*/
}

func error(position *token.Position, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", position.String(), message))
}
