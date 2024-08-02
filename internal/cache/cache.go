package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	c *cache.Cache
}

func NewCache() *Cache {
	return &Cache{c: cache.New(5*time.Minute, 10*time.Minute)}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.c.Get(key)
}

func (c *Cache) Set(key string, value interface{}) {
	c.c.Set(key, value, cache.DefaultExpiration)
}
