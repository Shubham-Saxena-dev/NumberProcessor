package cache

import (
	"sync"
)

type CacheService interface {
	Get(string) ([]int, bool)
	Set(string, []int)
	Delete(string)
}

type cacheService struct {
	urlCache map[string][]int
	cLock    sync.RWMutex
}

func NewCacheService() CacheService {
	return &cacheService{
		urlCache: make(map[string][]int),
		cLock:    sync.RWMutex{},
	}
}

func (c *cacheService) Get(path string) ([]int, bool) {
	c.cLock.RLock()
	defer c.cLock.RUnlock()
	if cachedNumbers, exists := c.urlCache[path]; exists {
		return cachedNumbers, true
	}
	return nil, false
}

func (c *cacheService) Set(path string, numbers []int) {
	c.cLock.Lock()
	defer c.cLock.Unlock()
	c.urlCache[path] = numbers
}

func (c *cacheService) Delete(path string) {
	_, exists := c.Get(path)
	if exists {
		delete(c.urlCache, path)
	}
}
