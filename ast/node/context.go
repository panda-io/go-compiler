package node

import (
	"github.com/panda-foundation/go-compiler/ir"
	"github.com/panda-foundation/go-compiler/token"
)

type Error struct {
	Position token.Position
	Message  string
}

type Context struct {
	Block  *ir.Block
	Parent *Context
	Vars   map[string]ir.Value

	//TO-DO all declarations
	//

	Errors []*Error
}

func (c *Context) AddError(p token.Position, message string) {
	if c.Parent != nil {
		c.Parent.AddError(p, message)
	} else {
		c.Errors = append(c.Errors, &Error{
			Position: p,
			Message:  message,
		})
	}
}
