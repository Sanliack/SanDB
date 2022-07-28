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

func (s *ClientModel) Connect(tcpAddr string) (sanface.ClientControlFace, error) {
	addr, _ := net.ResolveTCPAddr("tcp", tcpAddr)
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("client net.DialTCP error", err)
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
