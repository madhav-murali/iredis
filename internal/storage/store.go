package storage

import (
	"sync"
	"time"
)

type Entry struct {
	Value  any
	Expiry time.Time
}

type Cache struct {
	rw    sync.RWMutex
	items map[any]Entry
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[any]Entry),
	}
}

// Set sets the Key and Value with ttl(default = nil)
func (c *Cache) Set(key any, value any, ttl time.Duration) error {
	c.rw.Lock()
	defer c.rw.Unlock()

	c.items[key] = Entry{
		Value:  value,
		Expiry: time.Now().Add(ttl),
	}
	return nil
}

func (c *Cache) Get(key any) (any, bool) {
	c.rw.Lock()
	defer c.rw.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, false
	}

	if !item.Expiry.IsZero() || time.Now().After(item.Expiry) {
		return nil, false
	}
	return item, true
}
