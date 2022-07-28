package sanmodel

import (
	"SanDB/sanface"
	"fmt"
	"net"
	"sync"
)

type ServerModel struct {
	Name     string
	ConnNO   int
	ConnNums int
	Version  string
	Listen   *net.TCPListener
	DBRWLock sync.RWMutex
}

func (s *ServerModel) Start() {
	fmt.Printf("SanDB Server:%s Version:%s 启动成功,开始监听:%s\n", s.Name, s.Version, s.Listen.Addr().String())
	for {
		conn, err := s.Listen.AcceptTCP()
		if err != nil {
			fmt.Println("服务器获取client的Conn Error:", err)
			continue
		}
		newconn := NewConnModel(conn, s.ConnNO)
		s.ConnNO++
		go newconn.Start()
		fmt.Println("SanDB Server Accept Conn Request TCP ADDR:", conn.RemoteAddr())
	}
}

func (s *ServerModel) Stop() {

}

func (s *ServerModel) Server() {
	defer s.Stop()
	s.Start()
}

func (s *ServerModel) GetVersion() string {
	return s.Version
}

func (s *ServerModel) GetConnNums() int {
	return s.ConnNums
}

// ====================================String====================================

func NewServerModel(name string, address string) sanface.Server {
	listenaddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("服务器获取TCPADDR Error:", err)
		return nil
	}
	listen, err := net.ListenTCP("tcp", listenaddr)
	if err != nil {
		fmt.Println("服务器获取Listen Error:", err)
		return nil
	}
	return &ServerModel{
		Name:     name,
		Listen:   listen,
		ConnNO:   0,
		ConnNums: 0,
		Version:  "SanDB_V0.3",
	}
}
