package main

import (
	"SanDB/sanface"
	"SanDB/sanmodel"
	"fmt"
	"sync"
	"time"
)

func main() {
	client := sanmodel.NewClientModel()
	c1, err := client.Connect("127.0.0.1:6666", "BefCache")
	if err != nil {
		fmt.Println("start client error")
		return
	}
	StartSet(50, 100, c1)
	select {}
}

func StartSet(i, j int, c sanface.ClientControlFace) {
	for k := i; k <= j; k++ {
		key := fmt.Sprintf("key%d", k)
		val := fmt.Sprintf("val%d", k)
		err := c.Set().Sadd([]byte(key), []byte(val))
		if err != nil {
			fmt.Println("eeeerrrrr")
			break
		}

	}
}

func moregoset() {
	client := sanmodel.NewClientModel()
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(gowg *sync.WaitGroup, i int) {
			dbname := fmt.Sprintf("setTest1")
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
		}(wg, i)
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
