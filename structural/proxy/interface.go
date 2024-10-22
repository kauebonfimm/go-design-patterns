package proxy

type IResource interface {
	GetCompositionKey(string) string
}
