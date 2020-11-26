package ast

// Statement interface for all statements
type Statement interface {
	statement()
}

func (*EnumMemberStatement) statement() {}

// EnumMemberStatement to declare eum member
type EnumMemberStatement struct {
	Position int
	Name     *Identifier
	Value    *Literal
}
