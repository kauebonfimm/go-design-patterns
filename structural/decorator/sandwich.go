package decorator

type ISandwich interface {
	Price() float64
	Items() []string
}

type BaseSandWich struct {
}

func (s *BaseSandWich) Price() float64 {
	return 50.80
}

func (s *BaseSandWich) Items() []string {
	return []string{"brad", "meat", "cheese"}
}
