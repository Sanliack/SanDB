package main

import (
	"fmt"
	"os"
)

type aa struct {
	file *os.File
}

func main() {
	ori, err := os.OpenFile("./SanDBFile/test/123.data", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	var kk = &aa{ori}
	fmt.Println(kk.file.Name())
	ne, err := os.OpenFile("./SanDBFile/test/456.data", os.O_CREATE|os.O_RDWR, 0644)
	fmt.Println(ne.Name())

	olddatafilename := ori.Name()
	err = ori.Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close OldDataFile appear Error:", err)
		return
	}

	newfileneme := ne.Name()
	err = ne.Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close NewDataFile appear Error:", err)
		return
	}
	_ = os.Remove(olddatafilename)
	err = os.Rename(newfileneme, olddatafilename)
	if err != nil {
		fmt.Println("[Error] ConnModel Merge NewDataFile Change Name instead OldDataFile appear Error:", err)
		return
	}
	kk.file = ne
	fmt.Println(kk.file)

}
