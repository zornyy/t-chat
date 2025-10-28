package main

import (
	"fmt"
	"go-chat/client"
	"go-chat/server"
)

func main() {
	var i int

	fmt.Print("Start go-chat Server (0) or Client (1) 0/1: ")
	fmt.Scan(&i)

	if i == 0 {
		server.Start()
	} else {
		client.Start()
	}
}
