package ast

// Metadata type
type Metadata struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*Literal
}

// Modifier type
type Modifier struct {
	Position int
	Public   bool
	Static   bool
	Async    bool
	Inline   bool
}

// NewProgram to create new program
func NewProgram(packageName string, parent *Program) *Program {
	return &Program{
		Package: packageName,

		Variables:  make(map[string]*Variable),
		Functions:  make(map[string]*Function),
		Enums:      make(map[string]*Enum),
		Interfaces: make(map[string]*Interface),
		Classes:    make(map[string]*Class),

		Parent:   parent,
		Children: make(map[string]*Program),
	}
}

// Program type
type Program struct {
	Package string

	Variables  map[string]*Variable
	Functions  map[string]*Function
	Enums      map[string]*Enum
	Interfaces map[string]*Interface
	Classes    map[string]*Class

	Parent   *Program
	Children map[string]*Program

	Document *Metadata
}
