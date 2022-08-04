package sanmodel

import (
	"SanDB/sanface"
	"net"
)

type ClientControlModel struct {
	Conn       *net.TCPConn
	SetControl sanface.SetControlFace
	StrControl sanface.StrControlFace
}

func (c *ClientControlModel) Set() sanface.SetControlFace {
	return c.SetControl
}

func (c *ClientControlModel) Str() sanface.StrControlFace {
	return c.StrControl
}

func NewClientContolModel(conn *net.TCPConn) sanface.ClientControlFace {
	return &ClientControlModel{
		Conn:       conn,
		StrControl: NewStrControl(conn),
		SetControl: NewSetControl(conn),
	}
}
