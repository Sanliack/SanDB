package sanmodel

import (
	"SanDB/sanface"
	"fmt"
	"io"
	"net"
)

type ConnModel struct {
	Conn     *net.TCPConn
	Cid      int
	Server   sanface.Server
	StrRoute sanface.StrRouteFace
	SetRoute sanface.SetRouteFace
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
		td := NewWorkerTranData(c, trandata)
		c.Server.AddMsgToMsgQueue(td)
		//err = c.SolveTranData(trandata)
		if err != nil {
			return
		}
	}
}

func (c *ConnModel) SolveTranData(trandata sanface.TranDataFace) error {
	command := trandata.GetCommId()
	switch command {
	case Database:
		database := string(trandata.GetData())
		dm, err := c.Server.GetDataManager(database)

		if err != nil {
			fmt.Println("[error] ConnModel Get datamanager error:", err)
			_ = c.SendErrMsg()
			return err
		}
		sm, err := c.Server.GetSetManager(database)
		if err != nil {
			fmt.Println("[error] ConnModel Get Setmanager error:", err)
			_ = c.SendErrMsg()
			return err
		}
		c.StrRoute = NewStrRouteModel(c, dm)
		c.SetRoute = NewSetRouteModel(c, sm)
		_ = c.SendSucessMsg()
		return nil
	case Str_Get:
		return c.StrRoute.Get(trandata)
	case Str_Put:
		return c.StrRoute.Put(trandata)
	case Str_Del:
		return c.StrRoute.Del(trandata)
	case Str_Clean:
		return c.StrRoute.Clean()
	case Str_Merge:
		return c.StrRoute.Merge()
	case Set_Add:
		return c.SetRoute.Sadd(trandata)
	case Set_Pop:
		return c.SetRoute.Spop(trandata)
	case Set_Card:
		return c.SetRoute.Scard(trandata)
	case Set_Member:
		return c.SetRoute.Smember(trandata)
	case Set_IsMember:
		return c.SetRoute.SIsmember(trandata)
	case Set_DelByKey:
		return c.SetRoute.DelByKey(trandata)
	case Set_Merge:
		return c.SetRoute.MergeFile(trandata)
	case Set_Clean:
		return c.SetRoute.Clean()
	}
	return nil
}

func (c *ConnModel) SendSyntaxMsg() error {
	errmsg := NewTranDataModel(nil, Syntax)
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
