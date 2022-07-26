package main

import "SanDB/sanmodel"

func main() {
	server := sanmodel.NewSanDBServerModel("SANLIACK", "127.0.0.1:3665")
	server.Server()
}
