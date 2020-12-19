package ast

const (
	Global = "global"
)

type Program struct {
	Packages map[string]*Package
}

func NewProgram() *Program {
	return &Program{
		Packages: make(map[string]*Package),
	}
}

func (p *Program) AddSource(s *Source) {
	if _, ok := p.Packages[s.Namespace]; !ok {
		p.Packages[s.Namespace] = &Package{
			Namespace: s.Namespace,
		}
	}
	pkg := p.Packages[s.Namespace]
	pkg.Attributes = append(pkg.Attributes, s.Attributes...)
	pkg.Members = append(pkg.Members, s.Members...)
}

func (p *Program) Reset() {
	p.Packages = make(map[string]*Package)
}
