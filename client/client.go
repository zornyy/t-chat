package client

import (
	"fmt"
	"net"
)

func Start() {
	fmt.Println("Attempting to connect to server...")

	conn, err := net.Dial("tcp", "127.0.0.1:1080")

	if err != nil {
		fmt.Printf("Could not connect to server. Error: %s", err)
		return
	}

	print(conn)
	conn.Close()
}
