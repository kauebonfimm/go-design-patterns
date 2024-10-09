package flyweight

type KindOfCar interface {
	BaseCost() string
}

type categoryCar struct {
	Name     string
	baseCost string
}

func (s *categoryCar) BaseCost() string {
	return s.baseCost
}

type Sedan struct {
	categoryCar
}

func NewSedan() *Sedan {
	return &Sedan{categoryCar: categoryCar{
		Name:     "Sedan",
		baseCost: "29.90",
	}}
}

type SUV struct {
	categoryCar
}

func NewSUV() *SUV {
	return &SUV{categoryCar: categoryCar{
		Name:     "SUV",
		baseCost: "59.90",
	}}
}
