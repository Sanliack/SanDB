package sanmodel

import (
	"SanDB/sanface"
	"errors"
	"fmt"
	"net"
)

type StrControlModel struct {
	Conn *net.TCPConn
}

func (c *StrControlModel) Put(key []byte, val []byte) error {
	tranbuf := append(key, []byte(" ")...)
	tranbuf = append(tranbuf, val...)
	trandata := NewTranDataModel(tranbuf, Str_Put)
	buf, _ := trandata.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] write []byte To Server error: ", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := c.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error] Client Put Read ReMessage Appear error:", err)
		return err
	}

	remsgtrandata := DecodeTranData(remsg[:n])
	if remsgtrandata.GetCommId() == Suc {
		return nil
	} else if remsgtrandata.GetCommId() == Syntax {
		return errors.New("syntax Error")
	} else {
		return errors.New("something happen wrong")
	}

}

func (c *StrControlModel) Get(key []byte) ([]byte, error) {
	trandata := NewTranDataModel(key, Str_Get)
	buf, _ := trandata.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] write []byte To Server error: ", err)
		return nil, err
	}

	remsg := make([]byte, 2048)
	n, err := c.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error] Client Get Read ReMessage Appear error:", err)
		return nil, err
	}
	remsgtrandata := DecodeTranData(remsg[:n])
	if remsgtrandata.GetCommId() == Suc {
		return remsgtrandata.GetData(), nil
	} else {
		fmt.Println("[Warning] Get Accept Command ID no is msg")
		return nil, errors.New("accept Command ID no is msg")
	}
}

func (c *StrControlModel) Merge() error {
	trandata := NewTranDataModel(nil, Str_Merge)
	buf, _ := trandata.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] write Merge Msg To Server error: ")
		return err
	}
	remsg := make([]byte, 2048)
	n, err := c.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error] Client Merge Read ReMessage Appear error:", err)
		return err
	}
	remsgtrandata := DecodeTranData(remsg[:n])
	if remsgtrandata.GetCommId() == Suc {
		return nil
	} else {
		fmt.Println("[Warning] Mer Accept Command ID no is msg")
		return errors.New("accept Command ID no is msg")
	}
}

func (c *StrControlModel) Del(key []byte) error {
	trandata := NewTranDataModel(key, Str_Del)
	buf, _ := trandata.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] write []byte To Server error: ", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := c.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error] Client Del Read ReMessage Appear error:", err)
		return err
	}
	remsgtrandata := DecodeTranData(remsg[:n])
	if remsgtrandata.GetCommId() == Suc {
		return nil
	} else {
		fmt.Println("[Warning] Get Accept Command ID no is msg")
		return errors.New("accept Command ID no is msg")
	}
}

func (c *StrControlModel) Clean() error {
	trandata := NewTranDataModel(nil, Str_Clean)
	buf, _ := trandata.Encode()
	_, err := c.Conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] write []byte To Server error: ", err)
		return err
	}
	remsg := make([]byte, 2048)
	n, err := c.Conn.Read(remsg)
	if err != nil {
		fmt.Println("[Error] Client Get Read ReMessage Appear error:", err)
		return err
	}
	remsgtrandata := DecodeTranData(remsg[:n])
	if remsgtrandata.GetCommId() == Suc {
		return nil
	} else {
		fmt.Println("[Warning] Get Accept Command ID no is msg")
		return errors.New("accept Command ID no is msg")
	}
}

func NewStrControl(conn *net.TCPConn) sanface.StrControlFace {
	return &StrControlModel{
		Conn: conn,
	}
}
