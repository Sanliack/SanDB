package main

import (
	"SanDB/sanmodel"
	"fmt"
)

func main() {
	client := sanmodel.NewClientModel()
	//client.Server()
	defer client.Stop()
	control, err := client.Connect("127.0.0.1:3665")
	if err != nil {
		fmt.Println("ccc,", err)
		return
	}

	// ===============Put=====================
	for i := 0; i < 50; i++ {
		err = control.Put([]byte(fmt.Sprintf("key%d", i)), []byte(fmt.Sprintf("Val%d", i)))
		if err != nil {
			fmt.Println("eereeer", err)
			return
		}
	}
	fmt.Println(1)

	//===============Get=====================
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := control.Get([]byte(key))
		if err != nil {
			fmt.Println("Get Error")
			return
		}
		fmt.Printf("Get KV==%s:%s\n", key, string(val))
	}
	fmt.Println(2)

	// ===============Del=====================
	for i := 20; i < 40; i++ {
		key := fmt.Sprintf("key%d", i)
		err := control.Del([]byte(key))
		if err != nil {
			fmt.Println()
		}
	}
	fmt.Println(3)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := control.Get([]byte(key))
		if err != nil {
			fmt.Println("Get Error")
			return
		}
		fmt.Printf("Get KV==%s:%s\n", key, val)
	}
	fmt.Println(4)

	// ===============Clear=====================
	err = control.Clean()

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := control.Get([]byte(key))
		if err != nil {
			fmt.Println("Get Error")
			return
		}
		fmt.Printf("Get KV==%s:%s\n", key, val)
	}
	fmt.Println(5)

	for i := 0; i < 50; i++ {
		err = control.Put([]byte(fmt.Sprintf("key%d", i)), []byte(fmt.Sprintf("Val%d", i)))
		if err != nil {
			fmt.Println("eereeer", err)
			return
		}
	}
	fmt.Println(6)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := control.Get([]byte(key))
		if err != nil {
			fmt.Println("Get Error")
			return
		}
		fmt.Printf("Get KV==%s:%s\n", key, val)
	}
	fmt.Println(7)

	if err != nil {
		fmt.Println("clear err", err)
	}
	select {}
}
