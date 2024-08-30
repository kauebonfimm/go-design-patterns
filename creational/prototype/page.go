package prototype

type PagePrototype interface {
	Render() string
	Clone(string) PagePrototype
	Title() string
	Set(string)
}

type BasePage struct {
	name string
	html string
}

func NewPage() PagePrototype {
	return &BasePage{html: `<h1>Teste</h1>`, name: "page 1"}
}

func (p *BasePage) Render() string {
	return p.html
}

func (b *BasePage) Clone(name string) PagePrototype {
	if name == "" {
		name = b.name + " clone"
	}
	return &BasePage{name: name, html: b.html}
}

func (b *BasePage) Title() string {
	return b.name
}

func (b *BasePage) Set(html string) {
	b.html = html
}
