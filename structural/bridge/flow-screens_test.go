package bridge_test

import (
	"strings"
	"testing"

	"github.com/kauebonfimm/go-design-patterns/structural/bridge"
	"github.com/stretchr/testify/assert"
)

const result = `*1
	* name: Cellphone
	* slug: 1
	* price: R$20.75

*2
	* name: Smartphone
	* slug: 2-A
	* price: R$75.8
`

func TestListProducts(t *testing.T) {
	assert := assert.New(t)

	screen := bridge.NewScreen("", &bridge.Shopify{})

	text := screen.RenderProducts("25")

	assert.Equal(text, result)

	screen = bridge.NewScreen("backoffice", &bridge.Shopify{})

	text = screen.RenderProducts("25")

	assert.Equal(text, strings.ToUpper(text))

}
