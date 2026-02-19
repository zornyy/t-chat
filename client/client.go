package client

import (
	"fmt"
	"go-chat/utils"
	"net"

	"github.com/fatih/color"
)

func handleConnection(conn net.Conn) {
	color.Set(color.FgYellow)
	fmt.Println("Attempting to read data from connection")
	color.Unset()

	go utils.ReadMessages(conn)
}

func Start() {
	fmt.Println("Attempting to connect to server...")

	conn, err := net.Dial("tcp", "127.0.0.1:1080")

	if err != nil {
		return
	}

	defer conn.Close() // Close the connection when it is no longer needed

	handleConnection(conn)
}
