package proxy

import (
	"strings"
	"time"
)

type Resource struct {
	Kind      string
	Namespace string
	Path      string
	Delay     time.Duration
}

func NewResource(kind, namespace, path string) IResource {
	return &Resource{kind, namespace, path, 10 * time.Second}
}

func (r *Resource) GetCompositionKey(_ string) string {
	time.Sleep(r.Delay)
	return strings.Join([]string{r.Kind, r.Namespace}, ":")
}
