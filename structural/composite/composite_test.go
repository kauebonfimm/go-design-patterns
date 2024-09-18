package composite_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/structural/composite"
	"github.com/stretchr/testify/assert"
)

func TestCompositePackage(t *testing.T) {
	product1 := composite.NewProduct("Shoes", 29)
	product2 := composite.NewProduct("Hat", 55)
	product3 := composite.NewProduct("Socks", 15)
	product4 := composite.NewProduct("T-shirt", 30)
	product5 := composite.NewProduct("Jeans", 40)
	product6 := composite.NewProduct("Jacket", 75)

	package1 := composite.NewPackage("Promo1", 3)
	package2 := composite.NewPackage("Promo8", 3)
	package3 := composite.NewPackage("Promo5", 3)

	package3.Add(product4)
	package3.Add(product5)
	package3.Add(product6)
	package2.Add(package3)
	package2.Add(product1)
	package2.Add(product2)
	package1.Add(package2)
	package1.Add(product3)

	assert.Equal(t, 253, package1.Count())
}
