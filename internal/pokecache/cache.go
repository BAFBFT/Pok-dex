package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

// return Cache by reference instead of value
// we have a singular cache, every subsequent process
// should be writing to that same cache that was instantiated
func NewCache(interval time.Duration) *Cache {
	newcache := &Cache{entries: make(map[string]cacheEntry), interval: interval}
	// reaper has to run in it's own go routine, since it is an infinite loop
	// return newcache will never be reached if not
	go newcache.reapLoop()
	return newcache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// LOCK EVERYTHING ACCESSING THE CACHE!!
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.entries[key]
	if exists {
		return value.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	// range over a channel means "keep receiving values until the channel is closed."
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.entries {
			if time.Now().After(val.createdAt.Add(c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
