package ast

import "github.com/panda-foundation/go-compiler/ir"

var (
	compilerFunctions = map[string]CompilerFunction{}
)

type CompilerFunction = func(c *Context, invocation *Invocation) ir.Value

func RegisterComplierFunction(namespace, name string, f CompilerFunction) {
	compilerFunctions[namespace+"."+name] = f
}

func IsCompilerFunction(c *Context, f Expression) bool {
	return GetCompilerFunctionName(c, f) != ""
}

func InvokeCompilerFunction(c *Context, invocation *Invocation) ir.Value {
	return compilerFunctions[GetCompilerFunctionName(c, invocation.Function)](c, invocation)
}

func GetCompilerFunctionName(c *Context, f Expression) string {
	switch n := f.(type) {
	case *Identifier:
		// search current package
		if c.Program.Module.Namespace != Global {
			qualified := c.Program.Module.Namespace + "." + n.Name
			if compilerFunctions[qualified] != nil {
				return qualified
			}
		}
		// search global
		qualified := Global + "." + n.Name
		if compilerFunctions[qualified] != nil {
			return qualified
		}

	case *MemberAccess:
		if p, ok := n.Parent.(*Identifier); ok {
			// search imports
			for _, i := range c.Program.Module.Imports {
				if i.Alias == p.Name {
					qualified := i.Namespace + "." + n.Member.Name
					if compilerFunctions[qualified] != nil {
						return qualified
					}
				}
			}
		}
	}
	return ""
}
