package factorymethod_test

import (
	"testing"

	factorymethod "github.com/kauebonfimm/go-design-patterns/creational/factory-method"
	"github.com/stretchr/testify/assert"
)

func TestFactoryMethod(t *testing.T) {
	assert := assert.New(t)
	enterprise := factorymethod.CustomerFactory("1", "Ambev", "123", factorymethod.ENTERPRISE)

	assert.Equal(enterprise.GetID(), "1")
	assert.Equal(enterprise.GetName(), "Ambev")
	assert.Equal(enterprise.GetCode(), "123")
	assert.Equal(enterprise.GetKind(), factorymethod.ENTERPRISE)
	assert.Equal(float64(105), enterprise.CalculatePrice(10))

	customerSample := factorymethod.CustomerFactory("10", "Joao mail", "12", factorymethod.BASE_CUSTOMER)

	assert.Equal(customerSample.GetID(), "10")
	assert.Equal(customerSample.GetName(), "Joao mail")
	assert.Equal(customerSample.GetCode(), "12")
	assert.Equal(customerSample.GetKind(), factorymethod.BASE_CUSTOMER)
	assert.Equal(float64(105), enterprise.CalculatePrice(10))

}
