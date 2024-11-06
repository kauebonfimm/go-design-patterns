package chainresponsabilty_test

import (
	"testing"

	chainresponsabilty "github.com/kauebonfimm/go-design-patterns/behavioral/chain_responsabilty"
	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {

	t.Run("chain success", func(t *testing.T) {
		user := chainresponsabilty.User{Name: "Tester", Number: "2", Email: "test@test.com", Age: 50}

		chain := chainresponsabilty.ProcessorUserSignUpFactory()

		status := chain.Execute(&user)

		assert.EqualValues(t, chainresponsabilty.Status{Type: chainresponsabilty.Success}, status)
	})
}
