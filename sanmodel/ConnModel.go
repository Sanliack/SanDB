package sanmodel

import (
	"SanDB/sanface"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type ConnModel struct {
	Conn        *net.TCPConn
	Cid         int
	Server      sanface.Server
	DataManager sanface.DataManagerFace
}

func (c *ConnModel) Start() {
	defer c.Stop()

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
	case Dat:
		database := string(trandata.GetData())
		dm, err := c.Server.GetDataManager(database)
		if err != nil {
			fmt.Println("[error] ConnModel Get datamanager error:", err)
			_ = c.SendErrMsg()
			return err
		}
		c.DataManager = dm
		_ = c.SendSucessMsg()
		return nil
	case Get:
		key := trandata.GetData()
		val, err := c.DataManager.Get(key)
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
		err := c.DataManager.Put([]byte(key), []byte(val))
		if err != nil {
			fmt.Println("[Warning] Conn Slove TranData user Func <conn.Put> appear Error:", err)
			_ = c.SendErrMsg()
			return err
		}
		_ = c.SendSucessMsg()
		return nil
	case Del:
		key := trandata.GetData()
		err := c.DataManager.Del(key)
		if err != nil {
			fmt.Println("[Warning] Conn Slove TranData user Func <conn.Del> appear Error:", err)
			_ = c.SendErrMsg()
			return err
		}
		_ = c.SendSucessMsg()
		return nil
	case Cle:

		err := c.DataManager.Clean()
		if err != nil {
			fmt.Println("[Warning] Conn Slove TranData user Func <conn.Cle> appear Error:", err)
			_ = c.SendErrMsg()
			return err
		}
		_ = c.SendSucessMsg()
		return nil
	case Mer:

		err := c.DataManager.MergeFile()
		if err != nil {
			fmt.Println("[Error] Server User func <MergerFile> Error:", err)
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

func NewConnModel(conn *net.TCPConn, cid int, server sanface.Server) sanface.ConnFace {
	return &ConnModel{
		Conn:   conn,
		Cid:    cid,
		Server: server,
	}
}
