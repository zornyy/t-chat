package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MessageInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Me: ")
	scanner.Scan()
	msg := scanner.Text()

	return msg
}

func ReadFromConnnection(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func WriteToConnection(message string, conn net.Conn) (int, error) {
	n, err := conn.Write([]byte(message))

	if err != nil {
		return n, fmt.Errorf("faild to write bytes. error: %w", err)
	}

	return n, nil
}
