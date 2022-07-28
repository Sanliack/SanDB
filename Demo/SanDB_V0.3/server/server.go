package main

import (
	"SanDB/conf"
	"SanDB/sanmodel"
)

func main() {
	server := sanmodel.NewSanDBServerModel("SanDB V0.4 Server ", conf.ConfigObj.Ip+conf.ConfigObj.Port)
	server.Server()
}
