package main

import "SanDB/sanmodel"

func main() {
	server := sanmodel.NewServerModel("Server v1.2", "127.0.0.1:6666")
	server.Server()
}
