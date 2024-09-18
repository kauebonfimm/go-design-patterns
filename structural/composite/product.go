package composite

type Product struct {
	name  string
	price int
}

func NewProduct(name string, price int) Icomposite {
	return &Product{name, price}
}

func (p *Product) Count() int {
	return p.price
}
