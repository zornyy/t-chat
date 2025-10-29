package server

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) error {
	fmt.Println("Connection Added")

	res, err := conn.Write([]byte(string("Hello dude !")))

	if err != nil {
		return fmt.Errorf("faild to write bytes. error: %w", err)
	}

	fmt.Printf("Wrote %d bytes to client\n", res)

	return nil
}

func Start() error {
	listener, err := net.Listen("tcp", ":1080")

	if err != nil {
		return fmt.Errorf("failed to start server. Error: %w", err)
	}

	fmt.Println("Server listening for new connections...")

	for {
		conn, err := listener.Accept()

		if err != nil {
			// This might mean killing the server (I do no know yet). Perhaps the return should be removed to avoid killing the server here
			return fmt.Errorf("failed to accept connection. Error: %w", err)
		}

		go handleConnection(conn)
	}
}
