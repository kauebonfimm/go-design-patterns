package bridge

type Product struct {
	Id    string
	Name  string
	Slug  string
	Price float64
}

type IEcommerce interface {
	ListProductsByCategory(categoryId string) []Product
}

type Shopify struct {
}

func (*Shopify) ListProductsByCategory(id string) []Product {
	return []Product{
		{
			"1",
			"Cellphone",
			"1",
			20.75,
		},
		{
			"2",
			"Smartphone",
			"2-A",
			75.80,
		},
	}
}
