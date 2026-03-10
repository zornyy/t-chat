package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MessageInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Me: ")
	scanner.Scan()
	msg := scanner.Text()

	return msg
}

func readFromConnnection(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func ReadMessages(conn net.Conn) error {
	for {
		message, err := readFromConnnection(conn)

		if err != nil {
			return fmt.Errorf("failed to read from socket. error: %w", err)
		}

		fmt.Printf("MSG: %s\n", message)
	}
}

func WriteToConnection(message string, conn net.Conn) (int, error) {
	n, err := conn.Write([]byte(message))

	if err != nil {
		return n, fmt.Errorf("faild to write bytes. error: %w", err)
	}

	return n, nil
}
