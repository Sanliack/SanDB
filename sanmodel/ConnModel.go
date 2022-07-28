package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

type ConnModel struct {
	Conn      *net.TCPConn
	Cid       int
	SanDBFile sanface.FileFace
	IndexMap  map[string]int64
	ConnLock  sync.RWMutex
}

func (c *ConnModel) Start() {
	defer c.Stop()
	c.InitMap()
	c.Listen()
}

func (c *ConnModel) Stop() {
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("[Warning] ConnModel try to close C.Conn appear error:", err)
	}
}

func (c *ConnModel) Listen() {
	for {
		buf := make([]byte, 4068)
		n, err := c.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("[Info] Remote User close the Connect:", err)
				break
			} else {
				fmt.Println("[Warning] c.Conn Read buf from remote appear error:", err)
				return
			}
		}
		trandata := DecodeTranData(buf[:n])
		err = c.SolveTranData(trandata)
		if err != nil {
			return
		}
		fmt.Println("成功处理一条请求:", string(trandata.GetData()))
	}
}

func (c *ConnModel) SolveTranData(trandata sanface.TranDataFace) error {
	command := trandata.GetCommId()
	switch command {
	case Get:
		key := trandata.GetData()
		val, err := c.Get(key)
		remsg := NewTranDataModel(val, Suc)
		buf, err := remsg.Encode()
		if err != nil {
			fmt.Println("[Error] pack Remsg Error：", err)
			_ = c.SendErrMsg()
			return err
		}
		_, err = c.Conn.Write(buf)
		if err != nil {
			fmt.Println("[Error] Conn Write Error：", err)
			_ = c.SendErrMsg()
			return err
		}
		return nil
	case Put:
		keyandval := strings.Split(string(trandata.GetData()), " ")
		if len(keyandval) != 2 {
			fmt.Println("[info] Accept Message syntax Error,pass")
			_ = c.SendSyntaxMsg()
			return errors.New("message syntax Error")
		}
		key := keyandval[0]
		val := keyandval[1]
		err := c.Put([]byte(key), []byte(val))
		if err != nil {
			fmt.Println("[Warning] Conn Slove TranData user Func <conn.Put> appear Error:", err)
			_ = c.SendErrMsg()
			return err
		}
		_ = c.SendSucessMsg()
		return nil
	case Del:
		key := trandata.GetData()
		err := c.Del(key)
		if err != nil {
			fmt.Println("[Warning] Conn Slove TranData user Func <conn.Del> appear Error:", err)
			_ = c.SendErrMsg()
			return err
		}
		_ = c.SendSucessMsg()
		return nil
	}
	return nil
}

func (c *ConnModel) SendSyntaxMsg() error {
	errmsg := NewTranDataModel(nil, Syn)
	buf, _ := errmsg.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] SendErrmsg Error:", err)
		return err
	}
	return nil
}

func (c *ConnModel) SendSucessMsg() error {
	errmsg := NewTranDataModel(nil, Suc)
	buf, _ := errmsg.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] SendErrmsg Error:", err)
		return err
	}
	return nil
}

func (c *ConnModel) SendErrMsg() error {
	errmsg := NewTranDataModel(nil, Err)
	buf, _ := errmsg.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] SendErrmsg Error:", err)
		return err
	}
	return nil
}

func (c *ConnModel) SendNilMsg() error {
	errmsg := NewTranDataModel(nil, Nil)
	buf, _ := errmsg.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] SendErrmsg Error:", err)
		return err
	}
	return nil
}

func (c *ConnModel) GetIndexMap() map[string]int64 {
	return c.IndexMap
}

func (c *ConnModel) GetSanDBFile() sanface.FileFace {
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
		fmt.Println("[Info] key no exist")
		return nil, errors.New("key no exist")
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

	var newfile sanface.FileFace
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

func NewConnModel(conn *net.TCPConn, cid int) sanface.ConnFace {
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
