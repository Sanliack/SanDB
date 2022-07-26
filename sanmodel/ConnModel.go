package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
	"net"
	"sync"
)

type ConnModel struct {
	Conn      *net.TCPConn
	Cid       int
	SanDBFile sanface.SanDBFileFace
	IndexMap  map[string]int64
	ConnLock  sync.RWMutex
}

func (c *ConnModel) Start() {
	c.InitMap()
	select {}
}

func (c *ConnModel) GetIndexMap() map[string]int64 {
	return c.IndexMap
}

func (c *ConnModel) GetSanDBFile() sanface.SanDBFileFace {
	return c.SanDBFile
}

func (c *ConnModel) Put(key []byte, val []byte) error {
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()

	entry := NewEntryModel(key, val, Put)

	err := c.GetSanDBFile().Write(entry)
	if err != nil {
		fmt.Println("[Error] Conn User SanDBFile Func <Write> appear Error", err)
		return err
	}
	c.IndexMap[string(key)] = c.GetSanDBFile().GetOffset()
	return nil
}

func (c *ConnModel) InitMap() {
	var offset int64
	for {
		entry, err := c.SanDBFile.Read(offset)
		if err != nil {
			fmt.Println("[Warning] Conn init IndexMap appear error")
			continue
		}
		mask := entry.GetMask()
		if mask == Del {
			offset += int64(entry.GetKeySize())
			continue
		}
		key := entry.GetKey()
		c.IndexMap[string(key)] = offset
		offset += int64(entry.GetKeySize())
	}
}

func NewConnModel(conn *net.TCPConn, cid int) *ConnModel {
	sdfile, err := NewSanDBFileModel(conf.ConfigObj.SanDBFilePath)
	if err != nil {
		fmt.Println("[Error] Conn User func <NewSanDBFileModel> appear error", err)
		return nil
	}
	return &ConnModel{
		Conn:      conn,
		Cid:       cid,
		IndexMap:  make(map[string]int64),
		SanDBFile: sdfile,
	}
}
