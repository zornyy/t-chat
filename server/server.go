package server

import (
	"fmt"
	"net"
)

func Start(port int) (net.Conn, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to start server. Error: %w", err)
	}

	fmt.Println("Server listening for new connections...")
	for {
		return listener.Accept()
	}
}
