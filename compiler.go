package main

import (
	"crypto/md5"
	"fmt"
	"go/token"
	"io/ioutil"

	"github.com/panda-foundation/go-compiler/ast/expression"
	"github.com/panda-foundation/go-compiler/ast/statement"
)

func (p *Parser) ParseExpression(file token.File, source []byte) expression.Expression {
	file := p.files.AddFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.scanner.SetFile(file, source)
	p.next()
	return p.parseExpression()
}

func (p *Parser) ParseCompoundStatement(source []byte) statement.Statement {
	file := p.files.AddFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.scanner.SetFile(file, source)
	p.next()
	return p.parseCompoundStatement()
}

func (p *Parser) ParseBytes(source []byte) {
	file := p.files.AddFile("<input>"+fmt.Sprintf("%x", md5.Sum(source)), len(source))
	p.parse(file, source)
}

func (p *Parser) ParseFile(fileName string) {
	source, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	file := p.files.AddFile(fileName, len(source))
	p.parse(file, source)
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
