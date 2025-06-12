package pokecache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(dur time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: dur,
	}
	go cache.reapLoop()
	return cache

}

func (c *Cache) Add(key string, cVal []byte) error {
	entry := cacheEntry{createdAt: time.Now(), val: cVal}

	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.entries[key]
	if ok {
		return errors.New("value already exists, can not add")
	}

	c.entries[key] = entry
	return nil

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.entries[key]
	if ok {
		return value.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}

}
