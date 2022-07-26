package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"net"
)

type ConnModel struct {
	Conn      *net.TCPConn
	Cid       int
	SanDBFile sanface.SanDBFileFace
}

func (c *ConnModel) Start() {

}

func NewConnModel(conn *net.TCPConn, cid int) *ConnModel {
	return &ConnModel{
		Conn:      conn,
		Cid:       cid,
		SanDBFile: NewSanDBFileModel(conf.ConfigObj.SanDBFilePath),
	}
}
