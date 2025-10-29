package client

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Attempting to read data from connection")

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Printf("failed to read from socket. error: %s", err)
		return
	}

	fmt.Printf("Read %d bytes from server: %s", n, buffer)
	conn.Close()
}

func Start() {
	fmt.Println("Attempting to connect to server...")

	conn, err := net.Dial("tcp", "127.0.0.1:1080")

	if err != nil {
		return
	}

	fmt.Println(conn)
	defer conn.Close()

	handleConnection(conn)
}
