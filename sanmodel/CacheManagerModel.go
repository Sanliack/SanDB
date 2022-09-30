package sanmodel

import (
	"SanDB/sanface"
	"fmt"
)

type CacheManagerModel struct {
	CacheMap       map[string]sanface.OneCacheFace
	MaxDBCache     int
	MaxCacheLength int
}

func (c *CacheManagerModel) GetLen() int {
	return len(c.CacheMap)
}

func (c *CacheManagerModel) Put(name, k string, v []byte) {
	onecache, ok := c.CacheMap[name]
	if !ok {
		if c.GetLen() < c.MaxDBCache {
			onecache = NewOneCache(name, c.MaxCacheLength)
			c.CacheMap[name] = onecache
		} else {
			fmt.Println("开启对应database缓存到达上限")
			return
		}
	}
	onecache.Put(k, v)
}

func (c *CacheManagerModel) Del(name, key string) {
	onecache, ok := c.CacheMap[name]
	if !ok {
		return
	}
	onecache.Del(key)
}

func (c *CacheManagerModel) Clean(name string) {
	onecache, ok := c.CacheMap[name]
	if !ok {
		return
	}
	onecache.Clean()
}

func (c *CacheManagerModel) Get(name, key string) ([]byte, bool) {
	if cache, ok := c.CacheMap[name]; ok {
		return cache.Get(key)
	}
	return nil, false
}

func (c *CacheManagerModel) TestDuBug() {
	fmt.Println("最大数据库缓存数量:", c.MaxDBCache, "最大cache缓存数量", c.MaxCacheLength)
	for _, v := range c.CacheMap {
		v.TestDebug()
	}
	fmt.Println("cachemanager test finish")
}

func NewCacheManagerModel(maxnums, cachelength int) sanface.CacheFace {
	return &CacheManagerModel{
		CacheMap:       make(map[string]sanface.OneCacheFace),
		MaxDBCache:     maxnums,
		MaxCacheLength: cachelength,
	}
}
