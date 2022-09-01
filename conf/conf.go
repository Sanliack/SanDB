package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type config struct {
	SanDBFilePath      string
	EntryHeaderSize    int
	SanDBFileMergePath string
	Ip                 string
	Port               string
	WorkerPoolSize     int
	WorkQueueLength    int
}

var ConfigObj *config

func init() {
	ConfigObj = &config{
		SanDBFilePath:      "/SanDBFile/SanDBFile.data",
		EntryHeaderSize:    10,
		SanDBFileMergePath: "/SanDBFile/SanDBFile.merge.data",
		Ip:                 "127.0.0.1",
		Port:               "33366",
		WorkerPoolSize:     100,
		WorkQueueLength:    1000,
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
	ConfigObj.SanDBFileMergePath = connconf.Key("FileMergePath").String()

	entryconf := conf.Section("Entry")
	ConfigObj.EntryHeaderSize, err = entryconf.Key("entryHeaderSize").Int()
	if err != nil {
		fmt.Println("[Warning] read conf.ini: Entry-entryHeaderSize appear error", err)
		flag = 1
	}

	serverconf := conf.Section("Server")
	ConfigObj.Ip = serverconf.Key("ip").String()
	ConfigObj.Port = serverconf.Key("port").String()
	ConfigObj.WorkerPoolSize, err = serverconf.Key("WorkerPoolSize").Int()
	if err != nil {
		fmt.Println("[Warning] read conf.ini: WorkerPoolSize appear error (can't turn it to int)", err)
		flag = 1
	}
	ConfigObj.WorkQueueLength, err = serverconf.Key("WorkQueueLength").Int()

	if err != nil {
		fmt.Println("[Warning] read conf.ini: WorkQueueLength appear error (can't turn it to int)", err)
		flag = 1
	}

	if flag == 0 {
		fmt.Println("SanDB Read Config file success...")
	}
}
