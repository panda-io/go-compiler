package ast

// Declaration interface for all declarations
type Declaration interface {
	declaration()
}

func (*Variable) declaration()  {}
func (*Function) declaration()  {}
func (*Enum) declaration()      {}
func (*Interface) declaration() {}
func (*Class) declaration()     {}

// Variable declaration
type Variable struct {
	Attributes
	Position int
	Name     *Identifier
	Type     Type
	Value    *Literal
}

// Function declaration
type Function struct {
	Attributes
	Position int
	Name     *Identifier
}

// Enum declaration
type Enum struct {
	Attributes
	Position int
	Name     *Identifier
}

// Interface declaration
type Interface struct {
	Attributes
	Position int
	Name     *Identifier
}

// Class declaration
type Class struct {
	Attributes
	Position int
	Name     *Identifier
}
