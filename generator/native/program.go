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
	writeForwardDeclarations(program, w)
	writeDeclarations(program, w)
	writeImplements(program, w)

	//TO-DO print main at the last
	err := ioutil.WriteFile(file, w.buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func writeIncludes(program *ast.Program, w *writer) {
	includes := []string{"<cinttypes>", "<cuchar>", "<string>"}
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

func writeForwardDeclarations(program *ast.Program, w *writer) {
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
	total := 0
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Enum:
			if total > 0 {
				w.buffer.WriteString("\n")
			}
			total++
			w.buffer.WriteString("enum class " + t.Name.Name + ";\n")
		case *declaration.Interface:
			if total > 0 {
				w.buffer.WriteString("\n")
			}
			total++
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
		case *declaration.Class:
			if total > 0 {
				w.buffer.WriteString("\n")
			}
			total++
			w.buffer.WriteString("class " + t.Name.Name + ";\n")
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
	w.buffer.WriteString("\n")
}

func writeDeclarations(program *ast.Program, w *writer) {
	writePackageDeclaration(program.Global, w)
	for _, pkg := range program.Packages {
		writePackageDeclaration(pkg, w)
	}
}

func writePackageDeclaration(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != "" {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	//TO-DO sort class declaration by inheiritance
	// get max inheiritance level, then print by level. (later check level and save it)
	total := 0
	for _, m := range p.Members {
		if _, ok := m.(*declaration.Variable); !ok {
			if total > 0 {
				w.buffer.WriteString("\n")
			}
			total++
			writeDeclaration(m, 0, w)
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
}

func writeImplements(program *ast.Program, w *writer) {
	writePackageImplement(program.Global, w)
	for _, pkg := range program.Packages {
		writePackageImplement(pkg, w)
	}
}

func writePackageImplement(p *ast.Package, w *writer) {
	namespace := []string{}
	if p.Namespace != "" {
		namespace = strings.Split(p.Namespace, ".")
		for _, n := range namespace {
			w.buffer.WriteString("namespace " + n + "\n{\n")
		}
	}
	for _, m := range p.Members {
		switch t := m.(type) {
		case *declaration.Variable:
			writeVariable(t, w)

		case *declaration.Function:
			writeFunction(t, w)

		case *declaration.Class:
			writeClass(t, w)
		}
	}
	if p.Namespace != "" {
		for range namespace {
			w.buffer.WriteString("}\n")
		}
	}
}

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
