package main

import (
	"SanDB/sanface"
	"SanDB/sanmodel"
	"fmt"
)

func main() {
	client := sanmodel.NewClientModel()
	defer client.Stop()
	control, err := client.Connect("127.0.0.1:3665")
	if err != nil {
		fmt.Println("Get Control error:", err)
		return
	}
	//TestMerge1(control)

	Mergeee(control)
	select {}
}

func Mergeee(con sanface.ClientControlFace) {
	err := con.Merge()
	if err != nil {
		fmt.Println("merge error", err)
		return
	}
}

func TestMerge1(con sanface.ClientControlFace) {
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("kkey%d", i)
		val := fmt.Sprintf("vall%d", i)
		err := con.Put([]byte(key), []byte(val))
		if err != nil {
			fmt.Println("put error", err)
			return
		}
	}

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("kkey%d", i)
		val := fmt.Sprintf("vvvv%d", i)
		err := con.Put([]byte(key), []byte(val))
		if err != nil {
			fmt.Println("put error", err)
			return
		}
	}
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("kkey%d", i)
		val, err := con.Get([]byte(key))
		fmt.Printf("kv===%s:%s", key, val)
		if err != nil {
			fmt.Println("put error", err)
			return
		}
	}
}
