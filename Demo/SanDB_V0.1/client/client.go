package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("Demo/SanDB_V0.1/client/aaaa", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(file.Name())
}
