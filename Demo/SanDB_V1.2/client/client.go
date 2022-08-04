package main

import (
	"SanDB/sanmodel"
	"fmt"
)

func main() {
	client := sanmodel.NewClientModel()
	conn, err := client.Connect("127.0.0.1:6666", "test_set")
	if err != nil {
		fmt.Println("aaaaa", err)
		return
	}
	//_ = conn.Set().Clean()
	//
	for i := 0; i < 20; i++ {
		val := fmt.Sprintf("val%d", i)
		err := conn.Set().Sadd([]byte("key1"), []byte(val))
		if err != nil {
			fmt.Println("sadd error", err)
			return
		}
		fmt.Printf("sucess Sadd kv==%s:%s\n", "AllMembers", val)
	}
	//
	//for i := 50; i < 80; i++ {
	//	val := fmt.Sprintf("val%d", i)
	//	err := conn.Set().Sadd([]byte("key2"), []byte(val))
	//	if err != nil {
	//		fmt.Println("sadd error", err)
	//		return
	//	}
	//	fmt.Printf("sucess Sadd kv==%s:%s\n", "AllMembers", val)
	//}

	//aa, err := conn.Set().Smember([]byte("key2"))
	//for _, v := range aa {
	//	fmt.Println("key2:", string(v))
	//}

	//n, err := conn.Set().Scard([]byte("key2"))
	//fmt.Println("get n== ", n)
	////
	//for i := 55; i < 70; i++ {
	//	val := fmt.Sprintf("val%d", i)
	//	err := conn.Set().Spop([]byte("key2"), []byte(val))
	//	if err != nil {
	//		fmt.Println("sadd error", err)
	//		return
	//	}
	//	fmt.Printf("sucess Sadd kv==%s:%s\n", "AllMembers", val)
	//}
	//
	//n, err = conn.Set().Scard([]byte("key2"))
	//fmt.Println("get n== ", n)
	booll, err := conn.Set().SIsMember([]byte("key1"), []byte("val11"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(booll)
	//_ = conn.Set().MergeFile()
	//_ = conn.Set().Clean()
	select {}
}
