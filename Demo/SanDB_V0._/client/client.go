package main

import "fmt"

func main() {
	var aa = make(map[int][]int, 0)
	aa[1] = []int{1, 3, 6, 4}
	fmt.Println(aa[0])
	fmt.Println(aa[1] == nil)
}
