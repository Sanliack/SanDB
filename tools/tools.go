package tools

import (
	"SanDB/sanface"
	"encoding/binary"
	"fmt"
)

type ListNode struct {
	Val  []byte
	Key  string
	Pre  *ListNode
	Next *ListNode
}

func DecodeSetMember(buf []byte) ([][]byte, error) {
	var members = make([][]byte, 0)
	tmp := 0
	for {
		if tmp >= len(buf) {
			break
		}
		datalen := binary.BigEndian.Uint16(buf[tmp : 2+tmp])
		var bytedata = make([]byte, datalen)
		copy(bytedata, buf[2+tmp:2+int(datalen)+tmp])
		members = append(members, bytedata)
		tmp += 2 + int(datalen)
	}
	return members, nil
}

func EncodeSetMember(members [][]byte) []byte {
	var tranbyte = make([]byte, 0)
	for _, buf := range members {
		smallbuf := make([]byte, len(buf)+2)
		binary.BigEndian.PutUint16(smallbuf[:2], uint16(len(buf)))
		copy(smallbuf[2:], buf)
		tranbyte = append(tranbyte, smallbuf...)
	}
	return tranbyte
}

func StartGet(c sanface.ClientControlFace, l, r int) {
	for i := l; i < r; i++ {
		key := []byte(fmt.Sprintf("key%d", i))
		val, err := c.Str().Get(key)
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
		err := c.Str().Put(key, val)
		if err != nil {
			fmt.Println("put error", err)
		}
	}
}

func StartMerge(c sanface.ClientControlFace) {
	err := c.Str().Merge()
	if err != nil {
		fmt.Println("merge error", err)
		return
	}
}

func StartDel(c sanface.ClientControlFace, l, r int) {
	for i := l; i < r; i++ {
		key := []byte(fmt.Sprintf("key%d", i))
		//val := []byte(fmt.Sprintf("key%d", i))
		err := c.Str().Del(key)
		if err != nil {
			fmt.Println("put error", err)
		}
	}
}
