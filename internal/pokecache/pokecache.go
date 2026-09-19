// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	c        map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		c:        make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.ReadLoop()
	return cache
}

func (cache *Cache) Get(key string) (val []byte, found bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	entry, found := cache.c[key]
	if found {
		if val := time.Now().Compare(entry.createdAt.Add(cache.interval)); val == 1 || val == 0 {
			return nil, false
		}
	}
	return entry.val, found
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.c[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) ReadLoop() {
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		cache.update()
	}
}

func (cache *Cache) update() {
	cache.mu.Lock()
	for key, value := range cache.c {
		res := time.Now().Compare(value.createdAt.Add(cache.interval))
		if res == 0 || res == 1 {
			delete(cache.c, key)
		}
	}
	cache.mu.Unlock()
}
