package proxy_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/structural/proxy"
	"github.com/stretchr/testify/assert"
)

func TestProxySuccess(t *testing.T) {

	data := proxy.NewProxy(proxy.NewResource("car", "HB20X", "/test/"))

	assert.Equal(t, data.GetCompositionKey(""), "car:HB20X")

}
