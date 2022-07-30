package main

import "SanDB/sanmodel"

func main() {
	server := sanmodel.NewServerModel("SanDB Server V1.0", "127.0.0.1:3665")
	server.Server()
}
