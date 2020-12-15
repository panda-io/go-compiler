package generator

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
	tabSize = 4
)

type writer struct {
	fileset        *token.FileSet
	buffer         *bytes.Buffer
	replaceClasses map[string]*replaceClass
}

func Write(program *ast.Program, fileset *token.FileSet, file string) {
	w := &writer{
		fileset:        fileset,
		buffer:         bytes.NewBuffer(nil),
		replaceClasses: make(map[string]*replaceClass),
	}

	checkCppAttributes(program, w)

	w.buffer.WriteString("// --------------------------------       includes       --------------------------------\n")
	writeIncludes(program, w)
	w.buffer.WriteString("// -------------------------------- forward declarations --------------------------------\n")
	writeForwardDeclarations(program, w)
	w.buffer.WriteString("// --------------------------------     declarations     --------------------------------\n")
	writeDeclarations(program, w)
	w.buffer.WriteString("// --------------------------------      implements      --------------------------------\n")
	writeImplements(program, w)

	//TO-DO print main at the last
	err := ioutil.WriteFile(file, w.buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func writeForwardDeclarations(program *ast.Program, w *writer) {
	for _, pkg := range program.Packages {
		writePackageForwardDeclaration(pkg, w)
	}
}

func writePackageForwardDeclaration(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != ast.Global {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	first := true
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Enum:
			if !first {
				w.buffer.WriteString("\n")
			}
			w.buffer.WriteString("enum class " + t.Name.Name + ";\n")
			first = false

		case *declaration.Interface:
			if !first {
				w.buffer.WriteString("\n")
			}
			if t.TypeParameters != nil {
				writeType(t.TypeParameters, w)
			}
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
			first = false

		case *declaration.Class:
			if w.replaceClasses[p.Namespace+"."+m.Identifier()] != nil {
				break
			}
			if !first {
				w.buffer.WriteString("\n")
			}
			if t.TypeParameters != nil {
				writeType(t.TypeParameters, w)
			}
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
			first = false
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
	if !first || p.Namespace != "" {
		w.buffer.WriteString("\n")
	}
}

func writeDeclarations(program *ast.Program, w *writer) {
	for _, pkg := range program.Packages {
		writePackageDeclaration(pkg, w)
	}
}

func writePackageDeclaration(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != ast.Global {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	//TO-DO sort class declaration by inheiritance
	// get max inheiritance level, then print by level. (later check level and save it)
	for i, m := range p.Members {
		if i > 0 {
			w.buffer.WriteString("\n")
		}
		writeDeclaration(p, m, 0, w)
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
	w.buffer.WriteString("\n")
}

func writeImplements(program *ast.Program, w *writer) {
	for _, pkg := range program.Packages {
		writePackageImplement(pkg, w)
	}
}

func writePackageImplement(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != ast.Global {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	first := true
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Function:
			if !first {
				w.buffer.WriteString("\n")
			}
			writeFunction(t, nil, w)
			first = false

		case *declaration.Class:
			if w.replaceClasses[p.Namespace+"."+m.Identifier()] != nil {
				break
			}
			if !first {
				w.buffer.WriteString("\n")
			}
			writeClass(t, w)
			first = false
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
	if !first || p.Namespace != "" {
		w.buffer.WriteString("\n")
	}
}

func error(position *token.Position, message string) {
	panic(fmt.Sprintf("error: %s \n %s \n", position.String(), message))
}
