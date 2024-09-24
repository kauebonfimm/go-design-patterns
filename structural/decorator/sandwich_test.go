package decorator_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/structural/decorator"
	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {

	sandwich := decorator.BaseSandWich{}

	assert.Equal(t, 50.80, sandwich.Price())

	xbacon := decorator.AddBacon{
		Sandwich: &sandwich,
	}
	assert.Equal(t, 55.80, xbacon.Price())

	xsalad := decorator.AddTomato{
		Sandwich: &xbacon,
	}
	assert.Equal(t, 56.80, xsalad.Price())

}
