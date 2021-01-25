package ast

type Block struct {
	StatementBase
	Statements []Statement
}

func (b *Block) GenerateIR(c *Context) bool {
	terminated := false
	for _, stmt := range b.Statements {
		terminated = stmt.GenerateIR(c)
		if terminated {
			//TO-DO warning: unreachable codes
			break
		}
	}
	return terminated
}
