package adapter

import "encoding/xml"

type Category struct {
	Name        string
	Description string
	Slug        string
}

type EcommerceProxy interface {
	GetCategories(string) []Category
}

type ServiceProxyEcommerce struct {
	adapter EcommerceProxy
}

func NewServiceProxyEcommerce(adapter EcommerceProxy) ServiceProxyEcommerce {
	return ServiceProxyEcommerce{adapter}
}

func (s *ServiceProxyEcommerce) GetCategories(id string) (string, error) {
	categories := s.adapter.GetCategories(id)

	res, err := xml.Marshal(categories)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
