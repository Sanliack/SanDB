package sanmodel

import (
	"net"
)

type ConnModel struct {
	Conn *net.TCPConn
	Cid  int
}

func (c *ConnModel) Start() {

}

func NewConnModel(conn *net.TCPConn, cid int) *ConnModel {
	return &ConnModel{
		Conn: conn,
		Cid:  cid,
	}
}
