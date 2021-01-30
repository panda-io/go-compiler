package ast

type ExpressionStatement struct {
	StatementBase
	Expression Expression
}

func (e *ExpressionStatement) GenerateIR(c *Context) {
	if e.Expression.IsConstant(c.Program) {
		e.Expression.GenerateConstIR(c.Program, nil)
	} else {
		e.Expression.GenerateIR(c, nil)
	}
}
