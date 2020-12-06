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
)

//TO-DO use global position

func Write(program *ast.Program, fileset *token.FileSet, file string) {
	buffer := bytes.NewBuffer(nil)

	writeIncludes(program, fileset, buffer)
	writeForwardDeclaration(program, fileset, buffer)
	//p.PrintForwardDeclaration(buffer)
	//buffer.WriteString("\n")
	//p.PrintDeclaration(buffer)
	//buffer.WriteString("\n")
	//p.PrintImplementation(buffer)

	//TO-DO print main at the last
	err := ioutil.WriteFile(file, buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func writeIncludes(program *ast.Program, fileset *token.FileSet, buffer *bytes.Buffer) {
	includes := []string{"<cinttypes>"}
	includes = append(includes, collectPackageIncludes(program.Global, fileset)...)
	for _, pkg := range program.Packages {
		includes = append(includes, collectPackageIncludes(pkg, fileset)...)
	}
	unique := make(map[string]bool)
	for _, include := range includes {
		unique[include] = true
	}
	for include := range unique {
		buffer.WriteString(fmt.Sprintf("#include %s\n", include))
	}
	buffer.WriteString("\n")
}

func collectPackageIncludes(p *ast.Package, fileset *token.FileSet) []string {
	includes := []string{}
	for i := len(p.Attributes) - 1; i >= 0; i-- {
		attr := p.Attributes[i]
		if attr.Name == attributeName {
			if name, ok := attr.Values["include"]; ok {
				if len(attr.Values) > 1 {
					error(fileset, attr.Position, "extra data except include")
				} else {
					if name.Type == token.STRING {
						includes = append(includes, name.Value[1:len(name.Value)-1])
					} else {
						error(fileset, attr.Position, "include path must be string")
					}
				}
				p.Attributes[i] = p.Attributes[len(p.Attributes)-1]
				p.Attributes = p.Attributes[:len(p.Attributes)-1]
			} else {
				error(fileset, attr.Position, "invalid cpp attribute for package")
			}
		}
	}
	return includes
}

func writeForwardDeclaration(program *ast.Program, fileset *token.FileSet, buffer *bytes.Buffer) {
	writePackageForwardDeclaration(program.Global, fileset, buffer)
	for _, pkg := range program.Packages {
		writePackageForwardDeclaration(pkg, fileset, buffer)
	}
}

func writePackageForwardDeclaration(p *ast.Package, fileset *token.FileSet, buffer *bytes.Buffer) {
	namespace := []string{}
	if p.Namespace != "" {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			buffer.WriteString("namespace " + n + "\n{\n")
		}
	}

	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Enum:
			buffer.WriteString("enum class " + t.Name.Name + ";\n")
		case *declaration.Interface:
			buffer.WriteString("class " + t.Name.Name + ";\n")
		case *declaration.Class:
			buffer.WriteString("class " + t.Name.Name + ";\n")
		}
	}

	if p.Namespace != "" {
		for range namespace {
			buffer.WriteString("}\n")
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

func error(fileset *token.FileSet, offset int, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", fileset.Position(offset).String(), message))
}
