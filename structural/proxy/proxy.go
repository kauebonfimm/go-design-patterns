package proxy

type Proxy struct {
	resource IResource
	cache    map[string]string
}

func NewProxy(resource IResource) IResource {
	return &Proxy{resource, make(map[string]string)}
}

func (p *Proxy) GetCompositionKey(id string) string {
	res, ok := p.cache[id]

	if !ok {
		res = p.resource.GetCompositionKey(id)
	}

	return res
}
