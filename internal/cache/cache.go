package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	ch *cache.Cache
}

func NewCache() *Cache {
	return &Cache{ch: cache.New(5*time.Minute, 10*time.Minute)}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.ch.Get(key)
}

func (c *Cache) Set(key string, value interface{}) {
	c.ch.Set(key, value, cache.DefaultExpiration)
}
