package server

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Handling connection !")
}

func Start() {
	fmt.Println("Server starting...")

	listener, err := net.Listen("tcp", ":1080")

	if err != nil {
		fmt.Printf("Could not start server. Error: %s", err)
		return
	}

	fmt.Println("Server listening for new connections...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("Could not open connection. Error: %s", err)
			return
		}

		go handleConnection(conn)
	}
}
