package memcache

import (
	"cinema/common"
	"sync"
	"time"
)

type caching struct {
	store  map[string]interface{}
	locker *sync.RWMutex
}

func NewCaching() *caching {
	return &caching{
		store:  make(map[string]interface{}),
		locker: new(sync.RWMutex),
	}
}

func (c *caching) Set(k string, value interface{}) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value
}

func (c *caching) Get(k string) interface{} {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.store[k]
}

func (c *caching) WriteTTL(k string, value interface{}, exp int) {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.store[k] = value

	go func() {
		defer common.AppRecover()
		<-time.NewTimer(time.Second * time.Duration(exp)).C
		c.Set(k, nil)
	}()
}
