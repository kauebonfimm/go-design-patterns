package abstractfactory_test

import (
	"testing"

	absfac "github.com/kauebonfimm/go-design-patterns/creational/abstract-factory"
)

func TestFactoryProvider(t *testing.T) {
	type tests struct {
		name        string
		provider    string
		credentials map[string]string
	}

	tt := []tests{
		{
			name:        "instance aws provider",
			provider:    "aws",
			credentials: map[string]string{"region": "us-east-1"},
		},
	}

	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			provider, err := absfac.NewCloudProvider(test.provider)

			if err != nil {
				t.Error("Error to instance Provider")
			}

			err = provider.SetCredentials(test.credentials)
			if err != nil {
				t.Error("Error to set credentials")
			}
			t.Logf("Success to instance %s provider", test.provider)
		})
	}
}
