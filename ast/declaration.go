package ast

// Declaration interface for all declarations
type Declaration interface {
	declaration()
}

// Class declaration
type Class struct {
}

// Enum declaration
type Enum struct {
}

// Variable declaration
type Variable struct {
}

// Interface declaration
type Interface struct {
}

// Function declaration
type Function struct {
}
