package composite

type Package struct {
	name       string
	tax        int
	components []Icomposite
}

func NewPackage(name string, tax int) *Package {
	return &Package{name, tax, make([]Icomposite, 0)}
}

func (p *Package) Count() int {
	total := p.tax

	for _, component := range p.components {
		total += component.Count()
	}

	return total
}

func (p *Package) Add(composite Icomposite) {
	p.components = append(p.components, composite)
}
