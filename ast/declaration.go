package ast

// Declaration interface for all declarations
type Declaration interface {
	declaration()
}

func (*Variable) declaration()  { return }
func (*Function) declaration()  { return }
func (*Enum) declaration()      { return }
func (*Interface) declaration() { return }
func (*Class) declaration()     { return }

// Variable declaration
type Variable struct {
	Attributes
	Name  *Identifier
	Type  Type
	Value *Literal
}

// Function declaration
type Function struct {
	Attributes
	Name           *Identifier
	TypeParameters *TypeParameters
	Parameters     *Parameters
	ReturnType     Type
	Body           *BlockStatement

	Class *Class
}

// Enum declaration
type Enum struct {
	Attributes
	Name    *Identifier
	Members map[string]*Variable
}

// Interface declaration
type Interface struct {
	Attributes
	Name           *Identifier
	TypeParameters *TypeParameters
	Base           *TypeArguments
	Functions      map[string]*Function
}

// Class declaration
type Class struct {
	Attributes
	Name           *Identifier
	TypeParameters *TypeParameters
	Base           *TypeArguments
	Variables      map[string]*Variable
	Functions      map[string]*Function
}
