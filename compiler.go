package main

import (
	"fmt"
	"io/ioutil"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/parser"
	"github.com/panda-foundation/go-compiler/token"
)

type Compiler struct {
	parser  *parser.Parser
	fileset *token.FileSet
	program *ast.Program
}

func NewCompiler(flags []string) *Compiler {
	return &Compiler{
		parser:  parser.NewParser(flags),
		fileset: &token.FileSet{},
		program: ast.NewProgram(),
	}
}

//TO-DO add project config or folder as project

func (c *Compiler) ParseFile(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	f := c.fileset.AddFile(file, len(b))
	m := c.parser.ParseFile(f, b)
	c.program.AddModule(file, m)
}

func (c *Compiler) Generate(file string) {
	content := c.program.GenerateIR()
	errors := c.program.Errors()
	if len(errors) > 0 {
		fmt.Println("found compile errors:")
		for _, e := range errors {
			//panic(fmt.Sprintf("error: %s \n %s \n", p.scanner.Position(position).String(), message))
			//TO-DO use global position and fileset
			fmt.Println(e.Message)
		}
		panic("compile failed.")
	}
	err := ioutil.WriteFile(file, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

/*
func (p *Parser) ParseFolder(folder string) {
	folderInfo, err := os.Open(folder)
	if err != nil {
		panic(err)
	}
	list, err := folderInfo.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, f := range list {
		if f.IsDir() {
			p.ParseFolder(filepath.Join(folder, f.Name()))
		} else {
			if strings.HasSuffix(f.Name(), ".pd") {
				filename := filepath.Join(folder, f.Name())
				p.ParseFile(filename)
			}
		}
	}
}*/
