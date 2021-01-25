package ast

type ExpressionStatement struct {
	StatementBase
	Expression Expression
}

func (e *ExpressionStatement) GenerateIR(c *Context) bool {
	e.Expression.GenerateIR(c)
	return false
}
