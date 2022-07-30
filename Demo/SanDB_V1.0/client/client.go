package main

import (
	"SanDB/sanmodel"
	"SanDB/tools"
	"fmt"
)

func main() {
	client := sanmodel.NewClientModel()
	control, err := client.Connect("127.0.0.1:3665", "Sanli")
	if err != nil {
		fmt.Println("Get Control error:", err)
		return
	}
	//fmt.Println(control)
	//tools.StartPut(control, 0, 200)
	//tools.StartGet(control, 0, 50)
	//tools.StartDel(control, 50, 150)
	tools.StartMerge(control)
	select {}
}
