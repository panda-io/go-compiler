package native

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/panda-foundation/go-compiler/ast"
	"github.com/panda-foundation/go-compiler/ast/declaration"
	"github.com/panda-foundation/go-compiler/token"
)

const (
	attributeName = "cpp"
	tabSize       = 4
)

var (
	indents = []byte("        ")
)

//TO-DO use global position

type cppAttribute struct {
	replace  string
	operator bool
}

type writer struct {
	fileset       *token.FileSet
	buffer        *bytes.Buffer
	cppAttributes map[string]*cppAttribute
}

func Write(program *ast.Program, fileset *token.FileSet, file string) {
	w := &writer{
		fileset:       fileset,
		buffer:        bytes.NewBuffer(nil),
		cppAttributes: make(map[string]*cppAttribute),
	}

	writeIncludes(program, w)
	writeForwardDeclaration(program, w)
	//p.PrintForwardDeclaration(buffer)
	//buffer.WriteString("\n")
	//p.PrintDeclaration(buffer)
	//buffer.WriteString("\n")
	//p.PrintImplementation(buffer)

	//TO-DO print main at the last
	err := ioutil.WriteFile(file, w.buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func writeIncludes(program *ast.Program, w *writer) {
	includes := []string{"<cinttypes>", "<cuchar>"}
	includes = append(includes, collectPackageIncludes(program.Global, w)...)
	for _, pkg := range program.Packages {
		includes = append(includes, collectPackageIncludes(pkg, w)...)
	}
	unique := make(map[string]bool)
	for _, include := range includes {
		unique[include] = true
	}
	for include := range unique {
		w.buffer.WriteString(fmt.Sprintf("#include %s\n", include))
	}
	w.buffer.WriteString("\n")
}

func collectPackageIncludes(p *ast.Package, w *writer) []string {
	includes := []string{}
	for i := len(p.Attributes) - 1; i >= 0; i-- {
		attr := p.Attributes[i]
		if attr.Name == attributeName {
			if name, ok := attr.Values["include"]; ok {
				if len(attr.Values) > 1 {
					error(w.fileset.Position(attr.Position), "extra data except include")
				} else {
					if name.Type == token.STRING {
						includes = append(includes, name.Value[1:len(name.Value)-1])
					} else {
						error(w.fileset.Position(attr.Position), "include path must be string")
					}
				}
				p.Attributes[i] = p.Attributes[len(p.Attributes)-1]
				p.Attributes = p.Attributes[:len(p.Attributes)-1]
			} else {
				error(w.fileset.Position(attr.Position), "invalid cpp attribute for package")
			}
		}
	}
	return includes
}

func writeForwardDeclaration(program *ast.Program, w *writer) {
	writePackageForwardDeclaration(program.Global, w)
	for _, pkg := range program.Packages {
		writePackageForwardDeclaration(pkg, w)
	}
}

func writePackageForwardDeclaration(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != "" {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Enum:
			w.buffer.WriteString("enum class " + t.Name.Name + ";\n")
		case *declaration.Interface:
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
		case *declaration.Class:
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
}

/*
func (p *Program) PrintDeclaration(buffer *bytes.Buffer) {
	if p.PackageName != "" {
		buffer.WriteString("namespace " + p.PackageName + "\n{\n")
	}

	for _, v := range p.Children {
		v.PrintDeclaration(buffer)
	}

	//to-do sort class declaration by inheiritance
	// get max inheiritance level, then print by level. (later check level and save it)
	for _, v := range p.Functions {
		v.PrintDeclaration(buffer, 0)
		buffer.WriteString("\n")
	}

	for _, v := range p.Enums {
		v.Print(buffer, 0)
		buffer.WriteString("\n")
	}

	for _, v := range p.Interfaces {
		v.Print(buffer, 0)
		buffer.WriteString("\n")
	}

	for _, v := range p.Classes {
		v.PrintDeclaration(buffer, 0)
		buffer.WriteString("\n")
	}

	if p.PackageName != "" {
		buffer.WriteString("\n}\n")
	}
}

func (p *Program) PrintImplementation(buffer *bytes.Buffer) {
	if p.PackageName != "" {
		buffer.WriteString("namespace " + p.PackageName + "\n{\n")
	}

	for _, v := range p.Children {
		v.PrintImplementation(buffer)
	}

	for _, v := range p.Values {
		v.Print(buffer, 0)
		buffer.WriteString("\n")
	}

	for _, v := range p.Functions {
		v.PrintImplementation(buffer, 0)
		buffer.WriteString("\n")
	}

	for _, v := range p.Classes {
		v.PrintImplementation(buffer, 0)
		buffer.WriteString("\n")
	}

	if p.PackageName != "" {
		buffer.WriteString("\n}\n")
	}
}
*/

func error(position *token.Position, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", position.String(), message))
}

func writeIndent(indent int, w *writer) {
	if indent > len(indents) {
		for i := indent - len(indents); i > -1; i-- {
			indents = append(indents, byte(' '))
		}
	}
	w.buffer.Write(indents[:indent])
}
