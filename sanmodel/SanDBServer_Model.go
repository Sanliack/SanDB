package sanmodel

import (
	"fmt"
	"net"
	"sync"
)

type SanDBServerModel struct {
	Name      string
	ConnNO    int
	ConnNums  int
	Version   string
	Listen    *net.TCPListener
	StringMap map[string]interface{}
	DBRWLock  sync.RWMutex
}

func (s *SanDBServerModel) Start() {
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

func (s *SanDBServerModel) Stop() {

}

func (s *SanDBServerModel) Server() {
	defer s.Stop()
	s.Start()
}

// ====================================String====================================

func NewSanDBServerModel(name string, address string) *SanDBServerModel {
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
	return &SanDBServerModel{
		Name:      name,
		Listen:    listen,
		ConnNO:    0,
		ConnNums:  0,
		Version:   "SanDB_V0.3",
		StringMap: make(map[string]interface{}, 24),
	}
}
