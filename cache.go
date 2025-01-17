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
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if cache.data[key].deadline != nil &&
		cache.data[key].deadline.Before(time.Now()) {
		return "", false
	}

	if _, trueValue := cache.data[key]; !trueValue {
		return "", false
	}

	return cache.data[key].value, true
}

func (cache *Cache) Put(key, value string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.data[key] = Data{value, nil}
}

func (cache *Cache) Keys() []string {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var keys []string
	nowTime := time.Now()

	for key, value := range cache.data {
		if value.deadline != nil && value.deadline.Before(nowTime) {
			continue
		}
		keys = append(keys, key)
	}

	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.data[key] = Data{value, &deadline}
}
