package ast

type Declaration interface {
	Node
	Identifier() string
	HasAttribute(attribute string) bool
	Qualified(namespace string) string
}

type Modifier struct {
	Public bool
	Weak   bool
}

func (m *Modifier) Equal(target *Modifier) bool {
	return m.Public == target.Public && m.Weak == target.Weak
}

type Attribute struct {
	Position int
	Name     string
	Text     string
	Values   map[string]*Literal
}

type DeclarationBase struct {
	NodeBase
	Attributes []*Attribute
	Modifier   *Modifier
	Name       *Identifier
	ObjectName string // parent object (class|interface|enum)
}

func (b *DeclarationBase) HasAttribute(attribute string) bool {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return true
		}
	}
	return false
}

func (b *DeclarationBase) GetAttributeValue(attribute string, value string) *Literal {
	for _, a := range b.Attributes {
		if a.Name == attribute {
			return a.Values[value]
		}
	}
	return nil
}

func (b *DeclarationBase) Identifier() string {
	return b.Name.Name
}

func (b *DeclarationBase) Qualified(namespace string) string {
	name := b.Name.Name
	if b.HasAttribute(Extern) {
		if b.ObjectName != "" {
			panic("object member cannot be external")
		}
	} else {
		if b.ObjectName != "" {
			name = b.ObjectName + "." + name
		}
		if !(namespace == Global && b.Name.Name == ProgramEntry) {
			name = namespace + "." + name
		}
	}
	return name
}
