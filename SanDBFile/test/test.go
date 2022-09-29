package main

import (
	"fmt"
	"os"
)

type aa struct {
	file *os.File
}

func main() {

	aa := []int{22, 2, 2, 2, 2, 2, 2, 2}
	for v := range aa {
		fmt.Println(v)
	}

}
