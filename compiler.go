package main

import (
	"io/ioutil"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/parser"
	"github.com/panda-foundation/go-compiler/resolver"
	"github.com/panda-foundation/go-compiler/token"
)

type Compiler struct {
	parser   *parser.Parser
	resolver *resolver.Resolver
	fileset  *token.FileSet
	sources  map[string]*ast.SoureFile
}

func NewCompiler(flags []string) *Compiler {
	return &Compiler{
		parser:   parser.NewParser(flags),
		resolver: resolver.NewResolver(),
		fileset:  &token.FileSet{},
		sources:  make(map[string]*ast.SoureFile),
	}
}

//TO-DO add project config or folder as project

func (c *Compiler) ParseFile(file string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	f := c.fileset.AddFile(file, len(b))
	c.sources[f.Name] = c.parser.ParseFile(f, b)
}

func (c *Compiler) Generate() {
	c.fileset.Walk(c.declare)
	c.fileset.Walk(c.resolve)
}

func (c *Compiler) declare(f *token.File) {
	c.resolver.Declare(f, c.sources[f.Name])
}

func (c *Compiler) resolve(f *token.File) {
	c.resolver.Resolve(f, c.sources[f.Name])
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
