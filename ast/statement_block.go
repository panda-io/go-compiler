package ast

type Block struct {
	StatementBase
	Statements []Statement
}

func (b *Block) GenerateIR(c *Context) bool {
	terminated := false
	for _, stmt := range b.Statements {
		ctx := c
		if _, ok := stmt.(*Block); ok {
			ctx = c.NewContext()
		}
		if stmt.GenerateIR(ctx) {
			terminated = true
		}
	}
	return terminated
}
