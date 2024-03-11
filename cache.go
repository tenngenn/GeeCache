package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu sync.Locker
	lru *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock() // defer：push the code into stack, run the code after return
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil) // Lazy Initialization
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if value, ok := c.lru.Get(key); ok {
		return value.(ByteView), ok
	}
	return
}