package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type config struct {
	SanDBFilePath string
}

var ConfigObj *config

func init() {
	ConfigObj = &config{
		SanDBFilePath: "/SanDBFile/Conn/",
	}
	reload()
}

func reload() {
	conf, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Println("[Warning] open conf.ini fail", err)
	}
	connconf := conf.Section("Conn")
	ConfigObj.SanDBFilePath = connconf.Key("FilePath").String()

	fmt.Println("SanDB Read Config file success...")

}
