package tools

import (
	"SanDB/sanface"
	"fmt"
)

func StartGet(c sanface.ClientControlFace, l, r int) {
	for i := l; i < r; i++ {
		key := []byte(fmt.Sprintf("key%d", i))
		val, err := c.Get(key)
		if err != nil {
			fmt.Println("put error", err)
		}
		fmt.Printf("Client get: %s:%s\n", key, val)
	}
}

func StartPut(c sanface.ClientControlFace, l, r int) {
	for i := l; i < r; i++ {
		key := []byte(fmt.Sprintf("key%d", i))
		val := []byte(fmt.Sprintf("key%d", i))
		err := c.Put(key, val)
		if err != nil {
			fmt.Println("put error", err)
		}
	}
}

func StartMerge(con sanface.ClientControlFace) {
	err := con.Merge()
	if err != nil {
		fmt.Println("merge error", err)
		return
	}
}

func StartDel(c sanface.ClientControlFace, l, r int) {
	for i := l; i < r; i++ {
		key := []byte(fmt.Sprintf("key%d", i))
		//val := []byte(fmt.Sprintf("key%d", i))
		err := c.Del(key)
		if err != nil {
			fmt.Println("put error", err)
		}
	}
}
