package singleton_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kauebonfimm/go-design-patterns/creational/singleton"
	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	assert := assert.New(t)

	wg := new(sync.WaitGroup)
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		t.Log(i)
		go func(id string) {
			instance := singleton.GetInstanceCache()
			instance.Set(id, "teste")
			t.Log(i)

			wg.Done()

			t.Log(i)

		}(fmt.Sprint(i))
	}

	wg.Wait()

	instance := singleton.GetInstanceCache()

	for i := 0; i <= 4; i++ {
		go func(id string) {
			assert.Equal(instance.Copy()[id], "id")
		}(fmt.Sprint(i))
	}
}
