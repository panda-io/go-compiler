package main

import (
	"io/ioutil"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/node"
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/parser"
	"github.com/panda-foundation/go-compiler/token"
)

type Compiler struct {
	parser  *parser.Parser
	fileset *token.FileSet
	sources map[string]*ast.Source
	program *ast.Program
}

func NewCompiler(flags []string) *Compiler {
	return &Compiler{
		parser:  parser.NewParser(flags),
		fileset: &token.FileSet{},
		sources: make(map[string]*ast.Source),
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
	s := c.parser.ParseFile(f, b)
	c.sources[f.Name] = s
	c.program.AddSource(s)
}

func (c *Compiler) Generate(file string) {
	module := ir.NewModule()
	ctx := node.NewContext(c.program.Declarations, module)
	content := c.program.Packages[ast.Global].GenerateIR(ctx)
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
