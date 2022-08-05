package main

import (
	"SanDB/sanmodel"
	"fmt"
	"sync"
	"time"
)

func main() {
	//client := sanmodel.NewClientModel()
	//c1, err := client.Connect("127.0.0.1:6666", "tttt")
	//if err != nil {
	//	fmt.Println("errrrrr")
	//	return
	//}
	//err = c1.Set().Sadd([]byte("c1"), []byte("valc1"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//
	//}
	//
	//c2, err := client.Connect("127.0.0.1:6666", "tttt")
	//if err != nil {
	//	fmt.Println("errrrrr")
	//	return
	//}
	//err = c2.Set().Sadd([]byte("c2"), []byte("valc2"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//
	//}
	//
	//c3, err := client.Connect("127.0.0.1:6666", "tttt")
	//if err != nil {
	//	fmt.Println("errrrrr")
	//	return
	//}
	//
	//err = c3.Set().Sadd([]byte("c3"), []byte("valc3"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//c4, err := client.Connect("127.0.0.1:6666", "tttt")
	//if err != nil {
	//	fmt.Println("errrrrr")
	//	return
	//}
	//
	//err = c4.Set().Sadd([]byte("c4"), []byte("valc4"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//
	//}
	//select {}
	moregoset()
	//moregostr()
	select {}
}

func moregoset() {
	client := sanmodel.NewClientModel()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(gowg *sync.WaitGroup, i int) {
			dbname := fmt.Sprintf("db1")
			conn, err := client.Connect("127.0.0.1:6666", dbname)
			if err != nil {
				fmt.Println("aaaaa", err)
				return
			}
			for j := 0; j < 100; j++ {
				val := fmt.Sprintf("go%dval%d", i, j)
				err := conn.Set().Sadd([]byte(dbname), []byte(val))
				if err != nil {
					fmt.Println("[Warning]", dbname, err)
				}
			}
			gowg.Done()
		}(&wg, i)
	}
	fmt.Println("start", time.Now().Minute(), time.Now().Second())
	wg.Wait()
	fmt.Println("is end", time.Now().Minute(), time.Now().Second())
	select {}
}

func moregostr() {
	client := sanmodel.NewClientModel()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(gowg *sync.WaitGroup, i int) {
			dbname := fmt.Sprintf("db2")
			conn, err := client.Connect("127.0.0.1:6666", dbname)
			if err != nil {
				fmt.Println("aaaaa", err)
				return
			}
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("go%dkey%d", i, j)
				val := fmt.Sprintf("go%dval%d", i, j)
				err := conn.Str().Put([]byte(key), []byte(val))
				if err != nil {
					fmt.Println("[Warning]", dbname, err)
				}
			}
			gowg.Done()
		}(&wg, i)
	}
	fmt.Println("start", time.Now().Minute(), time.Now().Second())
	wg.Wait()
	fmt.Println("is end", time.Now().Minute(), time.Now().Second())
	select {}
}
