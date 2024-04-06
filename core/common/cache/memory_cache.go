package cache

import (
	"sync"
	"time"
)

type MemoryCache struct {
	data     map[string]interface{}
	mutex    sync.RWMutex
	expire   time.Duration
	cleanup  *time.Timer
	cleanupM sync.Mutex
}

func NewMemoryCache(expire time.Duration) *MemoryCache {
	cache := &MemoryCache{
		data:    make(map[string]interface{}),
		expire:  expire,
		cleanup: time.NewTimer(expire),
	}
	go cache.startCleanup()
	return cache
}

func (c *MemoryCache) startCleanup() {
	for {
		select {
		case <-c.cleanup.C:
			c.cleanupExpired()
			c.cleanup.Reset(c.expire)
		}
	}
}

func (c *MemoryCache) cleanupExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key := range c.data {
		if c.isExpired(key, now) {
			delete(c.data, key)
		}
	}
}

func (c *MemoryCache) isExpired(key string, now time.Time) bool {
	expiration, ok := c.data[key+"_expiration"].(time.Time)
	return ok && expiration.Before(now)
}

func (c *MemoryCache) Set(key string, value string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	return nil
}

func (c *MemoryCache) SetWithExpire(key string, value string, ttl uint64) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.data[key+"_expiration"] = time.Now().Add(time.Duration(ttl) * time.Second)
	return nil
}

func (c *MemoryCache) Get(key string) (string, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.data[key].(string)
	if !ok {
		return "", nil
	}
	return value, nil
}

func (c *MemoryCache) Delete(key string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
	return nil
}

func (c *MemoryCache) Flush() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data = make(map[string]interface{})
	return nil
}
