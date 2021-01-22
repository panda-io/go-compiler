package ast

type Block struct {
	StatementBase
	Statements []Statement
}

func (b *Block) GenerateIR(c *Context) {
	for _, stmt := range b.Statements {
		stmt.GenerateIR(c)
	}
}
