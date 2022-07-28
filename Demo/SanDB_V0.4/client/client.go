package main

import "SanDB/sanmodel"

func main() {
	client := sanmodel.NewSanDBClientModel()
	client.Server()
}
