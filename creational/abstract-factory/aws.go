package abstractfactory

type AWSProvider struct {
	client interface{}
}

func NewAwsProvider() *AWSProvider {
	return &AWSProvider{}
}

func (p *AWSProvider) SetCredentials(credentials map[string]string) error {
	p.client = credentials
	return nil
}
