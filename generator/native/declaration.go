package native

import (
	"github.com/panda-foundation/go-compiler/ast/declaration"
)

func writeDeclaration(d declaration.Declaration, w *writer) {
	switch t := d.(type) {
	case *declaration.Variable:
		writeType(t.Type, w)
		w.buffer.WriteString(" ")
		w.buffer.WriteString(t.Identifier())
		if t.Value != nil {
			w.buffer.WriteString(" = ")
			writeExpression(t.Value, w)
		}
		w.buffer.WriteString(";\n")

	case *declaration.Function:
		/*
				//TO-DO indent
			if f.Generic != nil {
				WriteIndent(buffer, indent)
				//template <class T, int N>
				buffer.WriteString("template <")
				for i, t := range f.Generic.Types {
					if i > 0 {
						buffer.WriteString(", ")
					}
					buffer.WriteString("class ")
					t.Print(buffer)
				}
				buffer.WriteString(">\n")
			}
			WriteIndent(buffer, indent)
			if f.IsMember && !f.IsConstructor {
				buffer.WriteString("virtual ")
			}
			if !(f.IsDestructor || f.IsConstructor) {
				if f.Result == nil {
					buffer.WriteString("void ")
				} else {
					f.Result.Type.Print(buffer)
					buffer.WriteString(" ")
				}
			}
			f.Name.Print(buffer)
			buffer.WriteString("(")
			f.Params.Print(buffer)
			buffer.WriteString(")")

			if f.InterfaceMember {
				buffer.WriteString(" = 0")
			}
			buffer.WriteString(";\n")
		*/

	case *declaration.Enum:
		/*
					WriteIndent(buffer, indent)
			buffer.WriteString("enum class ")
			e.Name.Print(buffer)
			buffer.WriteString("\n")
			WriteIndent(buffer, indent)
			buffer.WriteString("{\n")
			for _, item := range e.List {
				item.Print(buffer, indent+4)
			}
			buffer.WriteString("};\n")*/

	case *declaration.Interface:
		/*
					WriteIndent(buffer, indent)
			buffer.WriteString("class ")
			i.Name.Print(buffer)
			buffer.WriteString("\n")
			WriteIndent(buffer, indent)
			buffer.WriteString("{\n")
			buffer.WriteString("public:\n")
			for _, v := range i.Values {
				//TO-DO no init
				v.Print(buffer, indent+TabSize)
				buffer.WriteString("\n")
			}
			for _, f := range i.Functions {
				f.PrintDeclaration(buffer, indent+TabSize)
				buffer.WriteString("\n")
			}
			buffer.WriteString("};\n")*/

	case *declaration.Class:
		/*
					WriteIndent(buffer, indent)
			buffer.WriteString("class ")
			c.Name.Print(buffer)
			//TO-DO inheritance

			buffer.WriteString("\n")
			WriteIndent(buffer, indent)
			buffer.WriteString("{\n")
			buffer.WriteString("public:\n")
			for _, v := range c.Values {
				v.Print(buffer, indent+TabSize)
				buffer.WriteString("\n")
			}
			for _, f := range c.Functions {
				f.PrintDeclaration(buffer, indent+TabSize)
				buffer.WriteString("\n")
			}
			buffer.WriteString("};\n")*/
	}
}

func writeImplement() {
	/* for class
	 	for _, f := range c.Functions {
			f.PrintImplementation(buffer, indent)
			buffer.WriteString("\n")
		}*/
}

func writeFunction() {
	/*
	   //TO-DO indent
	   	if f.Generic != nil {
	   		WriteIndent(buffer, indent)
	   		//template <class T, int N>
	   		buffer.WriteString("template <")
	   		for i, t := range f.Generic.Types {
	   			if i > 0 {
	   				buffer.WriteString(", ")
	   			}
	   			buffer.WriteString("class ")
	   			t.Print(buffer)
	   		}
	   		buffer.WriteString(">\n")
	   	}
	   	WriteIndent(buffer, indent)
	   	if !(f.IsDestructor || f.IsConstructor) {
	   		if f.Result == nil {
	   			buffer.WriteString("void ")
	   		} else {
	   			f.Result.Type.Print(buffer)
	   			buffer.WriteString(" ")
	   		}
	   	}
	   	if f.ClassName != "" {
	   		buffer.WriteString(f.ClassName + "::")
	   	}
	   	f.Name.Print(buffer)
	   	buffer.WriteString("(")
	   	f.Params.Print(buffer)
	   	buffer.WriteString(")")

	   	buffer.WriteString("\n")
	   	f.Body.Print(buffer, indent)*/
}
