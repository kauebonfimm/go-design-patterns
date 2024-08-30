package abstractfactory

import "fmt"

type ICloudProvider interface {
	SetCredentials(jsonCredentials map[string]string) error
	// Create abstract multi cloud methods
	// GetStorage() IStorage
	// GetInstance() IComputerIntances
}

func NewCloudProvider(kind string) (ICloudProvider, error) {
	switch kind {
	case "aws":
		return NewAwsProvider(), nil
	default:
		return nil, fmt.Errorf("provider type %s not found", kind)
	}
}
