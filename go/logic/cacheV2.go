package logic

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	data      map[string]interface{}
	expiry    time.Duration
	evictPool chan struct{}
	mutex     sync.RWMutex
}

func NewCache(expiry time.Duration, evictPoolSize int) *Cache {
	return &Cache{
		data:      make(map[string]interface{}),
		expiry:    expiry,
		evictPool: make(chan struct{}, evictPoolSize),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	go c.scheduleEviction(key)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

func (c *Cache) scheduleEviction(key string) {
	select {
	case c.evictPool <- struct{}{}:
		// Goroutine acquired from pool, proceed with eviction
	case <-time.After(c.expiry):
		// Timed out waiting for goroutine from pool, evicting from current goroutine
		c.evict(key)
		return
	}

	go func() {
		c.evict(key)
		<-c.evictPool
	}()
}

func (c *Cache) evict(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.data[key]; ok {
		delete(c.data, key)
		fmt.Printf("Evicted key %s\n", key)
	}
}
