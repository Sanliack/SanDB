package main

import "SanDB/sanmodel"

func main() {
	server := sanmodel.NewServerModel("SanDB V1.0 Test", "127.0.0.1:3355")
	server.Server()
}
