package ast

// Declaration interface for all declarations
type Declaration interface {
	declaration()
}

func (*Class) declaration()     {}
func (*Enum) declaration()      {}
func (*Variable) declaration()  {}
func (*Interface) declaration() {}
func (*Function) declaration()  {}

// Class declaration
type Class struct {
	Attributes
	Position int
}

// Enum declaration
type Enum struct {
	Attributes
	Position int
}

// Variable declaration
type Variable struct {
	Attributes
	Position int
}

// Interface declaration
type Interface struct {
	Attributes
	Position int
}

// Function declaration
type Function struct {
	Attributes
	Position int
}
