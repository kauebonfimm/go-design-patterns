package singleton

import "sync"

type cacheSingleton struct {
	cache map[string]string
}

func (c *cacheSingleton) Set(k, v string) { c.cache[k] = v }

func (c *cacheSingleton) Copy() map[string]string {
	return c.cache
}

var cache *cacheSingleton
var locker = new(sync.RWMutex)

func GetInstanceCache() *cacheSingleton {
	locker.Lock()
	defer locker.Unlock()
	if cache == nil {
		cache = &cacheSingleton{cache: make(map[string]string)}
	}
	return cache
}
