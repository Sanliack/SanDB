package test

import (
	"SanDB/sanmodel"
	"testing"
)

func CacheTest(t *testing.T) {
	cachemanager := sanmodel.NewCacheManagerModel(10, 1000)
	cachemanager.Put("cname", "key1", []byte("val1"))
}
