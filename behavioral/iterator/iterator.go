package iterator

type ListLinkedIterator struct {
	data int
	next *ListLinkedIterator
}

func NewListLinked(data int) *ListLinkedIterator {
	return &ListLinkedIterator{
		data: data,
	}
}

func (l *ListLinkedIterator) SetNext(next *ListLinkedIterator) {
	l.next = next
}

func (l *ListLinkedIterator) GetNext() *ListLinkedIterator {
	return l.next
}

func (l *ListLinkedIterator) GetData() int {
	return l.data
}
