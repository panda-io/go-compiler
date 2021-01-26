package ast

type Block struct {
	StatementBase
	Statements []Statement
}

func (b *Block) GenerateIR(c *Context) bool {
	for _, stmt := range b.Statements {
		ctx := c
		if _, ok := stmt.(*Block); ok {
			ctx = c.NewContext()
		}
		if stmt.GenerateIR(ctx) {
			return true
		}
	}
	return false
}
