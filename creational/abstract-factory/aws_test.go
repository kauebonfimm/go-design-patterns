package abstractfactory_test

import (
	"testing"

	absfac "github.com/kauebonfimm/go-design-patterns/creational/abstract-factory"
)

func TestAwsProvider(t *testing.T) {
	type tests struct {
		name        string
		credentials map[string]string
	}

	tt := []tests{
		{
			name:        "instance aws provider",
			credentials: map[string]string{"region": "us-east-1"},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			provider := absfac.NewAwsProvider()

			if err := provider.SetCredentials(test.credentials); err != nil {
				t.Error("Error to set credentials")
			}
		})
	}
}
