package factorymethod

type ICustomer interface {
	GetID() string
	GetName() string
	GetCode() string
	GetKind() string
	CalculatePrice(int64) float64
}

type Customer struct {
	id   string
	name string
	code string
	kind string
}

func (c *Customer) GetID() string {
	return c.id
}
func (c *Customer) GetName() string {
	return c.name
}
func (c *Customer) GetCode() string {
	return c.code
}
func (c *Customer) GetKind() string {
	return c.kind
}

const ENTERPRISE string = "enterprise"
const BASE_CUSTOMER string = "base_customer"

type Enterprise struct {
	Customer
}

func (c *Enterprise) CalculatePrice(dolar int64) float64 {
	return float64(dolar) * 10.5
}

type BaseCustomer struct {
	Customer
}

func (c *BaseCustomer) CalculatePrice(dolar int64) float64 {
	return float64(dolar) * 5.5
}

func CustomerFactory(id, name, code, kind string) ICustomer {
	customer := Customer{
		id,
		name,
		code,
		kind,
	}
	if kind == ENTERPRISE {
		return &Enterprise{
			Customer: customer,
		}
	}
	return &BaseCustomer{
		Customer: customer,
	}
}
