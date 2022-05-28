package cache

import (
	"sync"
	"time"
)

type Data struct {
	value    string
	deadline *time.Time
}

type Cache struct {
	mu   sync.Mutex
	data map[string]Data
}

func NewCache() Cache {
	return Cache{
		data: map[string]Data{},
	}
}

func (cache *Cache) Get(key string) (string, bool) {
	return "", true
}

func (cache *Cache) Put(key, value string) {
}

func (cache *Cache) Keys() []string {
	var keys []string
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.data[key] = Data{value, &deadline}
}
