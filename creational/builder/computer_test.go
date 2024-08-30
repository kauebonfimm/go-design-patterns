package builder_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/creational/builder"
	"github.com/stretchr/testify/assert"
)

func TestComputerBuilder(t *testing.T) {
	assert := assert.New(t)
	b := builder.NewBuilder()

	computer := b.Build()

	assert.Equal(computer.Price(), float64(155.20))

	b.WithCPU(builder.CPU{Component: builder.Component{Name: "test", Price: 30.01}})

	computer = b.Build()

	assert.Equal(computer.Price(), float64(135.16))
}
