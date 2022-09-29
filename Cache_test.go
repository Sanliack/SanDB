package test

import (
	"SanDB/sanmodel"
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cachemanager := sanmodel.NewCacheManagerModel(4, 5)
	cachemanager.Put("cname", "key1", []byte("val1"))
	cachemanager.Put("cname", "key2", []byte("val2"))
	cachemanager.Put("cname", "key3", []byte("val3"))
	cachemanager.Put("cname", "key4", []byte("val4"))
	aa, flag := cachemanager.Get("cname", "key1")
	if flag {
		fmt.Println("get val:", string(aa))
	} else {
		fmt.Println("get null")
	}
	cachemanager.Put("cname", "key1", []byte("val5"))
	cachemanager.Put("cname", "key6", []byte("val6"))
	cachemanager.Put("cname", "key7", []byte("val7"))
	cachemanager.Put("cname", "key9", []byte("val9"))
	aa, flag = cachemanager.Get("cname", "key1")
	if flag {
		fmt.Println("get val:", string(aa))
	} else {
		fmt.Println("get null")
	}
	cachemanager.Put("c2", "kk1", []byte("val1"))
	cachemanager.Put("c2", "kk3", []byte("val3"))
	cachemanager.Put("c2", "kk2", []byte("val2"))
	aa, flag = cachemanager.Get("c2", "key1")
	if flag {
		fmt.Println("get val:", string(aa))
	} else {
		fmt.Println("get null")
	}
	cachemanager.Put("d1", "dk", []byte("dd"))
	cachemanager.Put("d2", "dk", []byte("dd"))

	cachemanager.TestDuBug()
}
