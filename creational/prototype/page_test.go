package prototype_test

import (
	"fmt"
	"testing"

	"github.com/kauebonfimm/go-design-patterns/creational/prototype"
	"github.com/stretchr/testify/assert"
)

func TestPrototypeSuccess(t *testing.T) {
	assert := assert.New(t)

	page := prototype.NewPage()

	assert.Equal("page 1", page.Title())

	assert.Equal("<h1>Teste</h1>", page.Render())

	clone := page.Clone("")

	assert.Equal("page 1 clone", clone.Title())
	assert.Equal("<h1>Teste</h1>", clone.Render())

	clone.Set(fmt.Sprintf("<article>%s</article>", clone.Render()))

	assert.Equal("<article><h1>Teste</h1></article>", clone.Render())

	assert.NotEqual(page.Render(), clone.Render())

	clone2 := clone.Clone("New page")
	assert.Equal("New page", clone2.Title())
	assert.Equal("<article><h1>Teste</h1></article>", clone2.Render())

}
