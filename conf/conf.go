package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type config struct {
	SanDBFilePath   string
	EntryHeaderSize int
}

var ConfigObj *config

func init() {
	ConfigObj = &config{
		SanDBFilePath:   "/SanDBFile/Conn/",
		EntryHeaderSize: 10,
	}
	reload()
}

func reload() {
	flag := 0
	conf, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Println("[Warning] open conf.ini fail", err)
		flag = 1
	}
	connconf := conf.Section("Conn")
	ConfigObj.SanDBFilePath = connconf.Key("FilePath").String()

	entryconf := conf.Section("Entry")
	ConfigObj.EntryHeaderSize, err = entryconf.Key("entryHeaderSize").Int()
	if err != nil {
		fmt.Println("[Warning] read conf.ini: Entry-entryHeaderSize appear error", err)
		flag = 1
	}

	if flag == 0 {
		fmt.Println("SanDB Read Config file success...")
	}
}
