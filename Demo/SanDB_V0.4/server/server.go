package main

import (
	"SanDB/conf"
	"SanDB/sanmodel"
)

func main() {
	SanDB := sanmodel.NewServerModel("SanDB V1.0 Server", conf.ConfigObj.Ip+conf.ConfigObj.Port)
	SanDB.Server()
}
