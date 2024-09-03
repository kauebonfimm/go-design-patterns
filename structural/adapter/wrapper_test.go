package adapter_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/structural/adapter"
	"github.com/stretchr/testify/assert"
)

func TestGetCategories(t *testing.T) {
	assert := assert.New(t)

	service := adapter.NewServiceProxyEcommerce(&adapter.ShopifyAdapter{})

	res, err := service.GetCategories("cars")

	assert.Nil(err)
	assert.Equal("<Category><Name>1</Name><Description>1</Description><Slug>1</Slug></Category><Category><Name>2</Name><Description>2</Description><Slug>2</Slug></Category><Category><Name>3</Name><Description>3</Description><Slug>3</Slug></Category><Category><Name>4</Name><Description>4</Description><Slug>4</Slug></Category><Category><Name>6</Name><Description>6</Description><Slug>6</Slug></Category>", res)
}
