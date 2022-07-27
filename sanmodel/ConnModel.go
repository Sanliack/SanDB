package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
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
	//err := c.Put([]byte("testKey1"), []byte("Val1"))
	//if err != nil {
	//	fmt.Println(err)
	//}
	kk, _ := c.Get([]byte("testKey1"))
	fmt.Println(string(kk))
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

func (c *ConnModel) Get(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("len of key is 0")
	}
	c.ConnLock.RLock()
	defer c.ConnLock.RUnlock()
	offset, ok := c.IndexMap[string(key)]
	if !ok {
		return nil, nil
	}
	entry, err := c.SanDBFile.Read(offset)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return entry.GetVal(), nil
}

func (c *ConnModel) Del(key []byte) error {
	if len(key) == 0 {
		return errors.New("len of key is 0")
	}
	c.ConnLock.Lock()
	defer c.ConnLock.Unlock()
	delentry := NewEntryModel(key, nil, Del)
	err := c.SanDBFile.Write(delentry)
	if err != nil {
		fmt.Println("[Warning] ConnModel use Func <Del> Appear error:", err)
		return err
	}
	delete(c.IndexMap, string(key))
	return nil
}

func (c *ConnModel) MergeFile() error {
	var offset int64
	var newEntrys []sanface.EntryFace
	for {
		entry, err := c.SanDBFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("[Error] ConnModel MergeFile user Func <SanDBFile.Read> appear Error", err)
				return err
			}
		}
		off, ok := c.IndexMap[string(entry.GetKey())]
		if !(!ok || entry.GetMask() == Del || off != offset) {
			newEntrys = append(newEntrys, entry)
		}
		offset += entry.GetSize()
	}

	var newfile sanface.SanDBFileFace
	var err error

	if len(newEntrys) > 0 {
		newfile, err = NewSanDBFileModel(conf.ConfigObj.SanDBFileMergePath)
		if err != nil {
			fmt.Println("[Error] ConnModel-MergeFile user Func <NewSanDBFileModel> appear error:", err)
			return err
		}
		for _, entry := range newEntrys {
			newoffset := newfile.GetOffset()
			err = newfile.Write(entry)
			if err != nil {

			}
			c.IndexMap[string(entry.GetKey())] = newoffset
		}
	}
	olddatafilename := c.SanDBFile.GetFile().Name()
	err = c.SanDBFile.GetFile().Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close OldDataFile appear Error:", err)
		return err
	}

	newfileneme := newfile.GetFile().Name()
	err = newfile.GetFile().Close()
	if err != nil {
		fmt.Println("[Error] ConnModel Close NewDataFile appear Error:", err)
		return err
	}

	err = os.Rename(olddatafilename, newfileneme)
	if err != nil {
		fmt.Println("[Error] ConnModel Merge NewDataFile Change Name instead OldDataFile appear Error:", err)
		return err
	}

	c.SanDBFile = newfile
	return nil
}

func (c *ConnModel) InitMap() {
	var offset int64
	for {
		entry, err := c.SanDBFile.Read(offset)
		if err != nil {
			//if err == io.EOF {
			//	break
			//}
			fmt.Println("[Warning] Conn init IndexMap appear error", err)
			break
		}
		mask := entry.GetMask()
		if mask == Del {
			offset += entry.GetSize()
			continue
		}
		key := entry.GetKey()
		c.IndexMap[string(key)] = offset
		offset += entry.GetSize()
	}
	fmt.Printf("Conn%d init Map sucess len of IndexMap:%d IndexMap KV:%v \n", c.Cid, len(c.IndexMap), c.IndexMap)
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
