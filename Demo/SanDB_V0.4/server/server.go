package main

import (
	"SanDB/conf"
	"SanDB/sanmodel"
)

func main() {
	SanDB := sanmodel.NewSanDBServerModel("SanDB V0.4 Server", conf.ConfigObj.Ip+conf.ConfigObj.Port)
	SanDB.Server()
}
