package flyweight

var tCars = make(map[string]KindOfCar, 0)

type Car struct {
	Kind  KindOfCar
	Color string
	Age   uint
}

func NewCar(key string, age uint, color string) *Car {
	var kind KindOfCar

	switch true {
	case key == "sedan":
		if _, ok := tCars[key]; !ok {
			tCars[key] = NewSedan()
		}
		kind = tCars[key]
	case key == "suv":
		if _, ok := tCars[key]; !ok {
			tCars[key] = NewSUV()
		}
		kind = tCars[key]
	}
	return &Car{
		Kind:  kind,
		Color: color,
		Age:   age,
	}
}
