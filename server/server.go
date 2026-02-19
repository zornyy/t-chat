package server

import (
	"fmt"
	"go-chat/utils"
	"net"
)

func handleConnection(conn net.Conn) error {
	fmt.Println("Connection Added")

	n, err := utils.WriteToConnection("My amazing message", conn)

	if err != nil {
		return fmt.Errorf("faild to write bytes. error: %w", err)
	}

	fmt.Printf("Wrote %d bytes to client\n", n)

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
