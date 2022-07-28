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
	//for i := 115; i < 150; i++ {
	//	err = control.Put([]byte(fmt.Sprintf("key%d", i)), []byte(fmt.Sprintf("Val%d", i)))
	//	if err != nil {
	//		fmt.Println("eereeer", err)
	//		return
	//	}
	//}

	//===============Get=====================
	for i := 0; i < 150; i++ {
		key := fmt.Sprintf("key%d", i)
		val, err := control.Get([]byte(key))
		if err != nil {
			fmt.Println("Get Error")
			return
		}
		fmt.Printf("Get KV==%s:%s\n", key, val)
	}

	// ===============Del=====================
	//for i := 50; i < 75; i++ {
	//	key := fmt.Sprintf("key%d", i)
	//	err := control.Del([]byte(key))
	//	if err != nil {
	//		fmt.Println()
	//	}
	//
	//}
	select {}

}
