package sanmodel

import (
	"SanDB/conf"
	"SanDB/sanface"
	"fmt"
	"net"
	"time"
)

type SanDBClientModel struct {
}

func (s *SanDBClientModel) Start() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", conf.ConfigObj.Ip+conf.ConfigObj.Port)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("client net.DialTCP error", err)
		return
	}
	var nums = 1
	for {
		//var command string
		//_, err := fmt.Scanln(command)
		//if err != nil {
		//	fmt.Println("check your command syntax")
		//	continue
		//}
		trandata := NewTranDataModel([]byte(fmt.Sprintf("keyno%d valno%d", nums, nums)), Put)
		nums++
		buf, _ := trandata.Encode()
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("[Error] write []byte To Server error: ", err)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func (s *SanDBClientModel) Server() {
	defer s.Stop()
	s.Start()
	select {}
}

func (s *SanDBClientModel) Stop() {

}

func NewSanDBClientModel() sanface.SanDBClientFace {
	return &SanDBClientModel{}
}
