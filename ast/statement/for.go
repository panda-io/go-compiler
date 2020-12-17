package statement

type For struct {
	Base
	Initialization Statement
	Condition      Statement
	Post           Statement
	Body           Statement
}
