package sanmodel

import (
	"SanDB/sanface"
	"fmt"
	"net"
)

type ClientModel struct {
}

func (s *ClientModel) Start() {

}

func (s *ClientModel) Connect(tcpAddr string, database string) (sanface.ClientControlFace, error) {
	addr, _ := net.ResolveTCPAddr("tcp", tcpAddr)
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("client net.DialTCP error", err)
		return nil, err
	}

	buf, err := NewTranDataModel([]byte(database), Database).Encode()
	if err != nil {
		fmt.Println("[Error] Connect Pack msg error:", err)
		return nil, err
	}
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println("[Error] Func Connect cannot send msg to Server error:", err)
		return nil, err
	}

	ansbuf := make([]byte, 2048)
	n, err := conn.Read(ansbuf)
	if err != nil {
		fmt.Println("[client] Read Connect msg error", err)
		return nil, err
	}
	entry, err := DecodeEntry(ansbuf[:n])
	if entry.GetMask() == Err {
		fmt.Println("[Info] Sever Send errmsg something is wrong", err)
		return nil, err
	}
	return NewClientContolModel(conn), nil
}

func (s *ClientModel) Server() {
	defer s.Stop()
	s.Start()
	select {}
}

func (s *ClientModel) Stop() {

}

func NewClientModel() sanface.ClientFace {
	return &ClientModel{}
}
