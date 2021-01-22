package ast

type ExpressionStatement struct {
	StatementBase
	Expression Expression
}

func (e *ExpressionStatement) GenerateIR(c *Context) {
	e.Expression.GenerateIR(c)
}
