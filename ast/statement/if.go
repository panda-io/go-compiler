package statement

type If struct {
	Base
	Initialization Statement
	Condition      Statement
	Body           Statement
	Else           Statement
}
