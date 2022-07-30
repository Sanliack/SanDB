package main

import (
	"SanDB/sanmodel"
	"fmt"
)

// 全局锁，偏移值
// 多conn写入失败原因：offset不会随其他conn写入而改变

func main() {
	client := sanmodel.NewClientModel()
	con1, err := client.Connect("127.0.0.1:3355", "con3")
	if err != nil {
		fmt.Println("1+", err)
		return
	}
	fmt.Println(con1)
	con2, err := client.Connect("127.0.0.1:3355", "con3")
	if err != nil {
		fmt.Println("2+", err)
		return
	}

	//go tools.StartPut(con1, 0, 100)
	//go tools.StartPut(con2, 100, 200)
	fmt.Println(con1, con2)
	select {}

}