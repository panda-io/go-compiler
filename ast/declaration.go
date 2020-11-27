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
	Name  *Identifier
	Type  Type
	Value *Literal
}

// Function declaration
type Function struct {
	Attributes
	Name *Identifier

	ReturnType Type
	/*
		{
			"include": "#type_parameter_list"
		},
		{
			"include": "#parenthesized_parameter_list"
		},
		{
			"include": "#type"
		},
		{
			"include": "#block"
		}
	*/
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
