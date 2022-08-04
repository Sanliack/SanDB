package sanmodel

import (
	"SanDB/sanface"
	"fmt"
	"net"
)

type ServerModel struct {
	Name           string
	ConnNO         int
	ConnNums       int
	Version        string
	Listen         *net.TCPListener
	DataManagerMap map[string]sanface.DataManagerFace
	SetManagerMap  map[string]sanface.SetManagerFace
}

func (s *ServerModel) Start() {
	fmt.Printf("SanDB Server:%s Version:%s 启动成功,开始监听:%s\n", s.Name, s.Version, s.Listen.Addr().String())
	for {
		conn, err := s.Listen.AcceptTCP()
		if err != nil {
			fmt.Println("服务器获取client的Conn Error:", err)
			continue
		}
		newconn := NewConnModel(conn, s.ConnNO, s)
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

func (s *ServerModel) GetDataManager(database string) (sanface.DataManagerFace, error) {
	dm, ok := s.DataManagerMap[database]
	if !ok {
		newdm, err := NewDataManagerModel(database)
		if err != nil {
			fmt.Println("[Error] Server user func <NewDataManagerModel> appear error", err)
			return nil, err
		}
		s.DataManagerMap[database] = newdm
		return newdm, nil
	}
	return dm, nil
}

func (s *ServerModel) GetSetManager(database string) (sanface.SetManagerFace, error) {
	dm, ok := s.SetManagerMap[database]
	if !ok {
		newdm, err := NewSetManagerModel(database)
		if err != nil {
			fmt.Println("[Error] Server user func <NewDataManagerModel> appear error", err)
			return nil, err
		}
		s.SetManagerMap[database] = newdm
		return newdm, nil
	}
	return dm, nil
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
		Name:           name,
		Listen:         listen,
		ConnNO:         0,
		ConnNums:       0,
		Version:        "SanDB_V1.0",
		DataManagerMap: make(map[string]sanface.DataManagerFace),
		SetManagerMap:  make(map[string]sanface.SetManagerFace),
	}
}
