package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val        []byte
	createTime time.Time
}

func NewCache() *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}

	go DeleteOlderThan5Minutes(c)
	return c

}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:        val,
		createTime: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}

func DeleteOlderThan5Minutes(c *Cache) {
	for {
		time.Sleep(5 * time.Minute)
		for key, entry := range c.cache {
			if time.Now().Sub(entry.createTime) > 5*time.Minute {
				c.Delete(key)
			}
		}
	}
}
