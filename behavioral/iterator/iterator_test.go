package iterator_test

import (
	"testing"

	"github.com/kauebonfimm/go-design-patterns/behavioral/iterator"
	"github.com/stretchr/testify/assert"
)

func TestIteratorLinkedList(t *testing.T) {
	linkedList := iterator.NewListLinked(1)
	linkedList2 := iterator.NewListLinked(2)
	linkedList3 := iterator.NewListLinked(3)
	linkedList4 := iterator.NewListLinked(4)
	linkedList5 := iterator.NewListLinked(5)

	linkedList.SetNext(linkedList2)
	linkedList2.SetNext(linkedList3)
	linkedList3.SetNext(linkedList4)
	linkedList4.SetNext(linkedList5)

	current := linkedList
	i := 1
	for current != nil {
		assert.Equal(t, i, current.GetData())
		i++
		current = current.GetNext()
	}
}
