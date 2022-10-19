package main

import (
	"SanDB/sanmodel"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
)

func main() {
	server := sanmodel.NewServerModel("Server v1.2", "127.0.0.1:6666")
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//main
	server.Server()
}
