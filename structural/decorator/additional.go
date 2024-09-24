package decorator

type AddBacon struct {
	Sandwich ISandwich
}

func (s *AddBacon) Price() float64 {
	return 5 + s.Sandwich.Price()
}

func (s *AddBacon) Items() []string {
	return append(s.Sandwich.Items(), "bacon")
}

type AddTomato struct {
	Sandwich ISandwich
}

func (s *AddTomato) Price() float64 {
	return 1 + s.Sandwich.Price()
}

func (s *AddTomato) Items() []string {
	return append(s.Sandwich.Items(), "tomato")
}
