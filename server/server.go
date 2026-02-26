package server

import (
	"fmt"
	"go-chat/utils"
	"net"
)

type Server struct {
	listener   net.Listener
	connection net.Conn
}

func New(port int) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to start server. Error: %w", err)
	}
	return &Server{listener: listener}, nil
}

func (s *Server) Broadcast(message string) error {
	_, err := utils.WriteToConnection(message, s.connection)

	return err
}

func (s *Server) Start() error {
	fmt.Println("Server listening for new connections...")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept connection. Error: %w", err)
		}
		s.connection = conn
	}
}
