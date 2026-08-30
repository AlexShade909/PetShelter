package network

import (
	"bufio"
	"fmt"
	"net"
)

func Connect() error {
	fmt.Println("Start server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connection", err)
		return err
	}
	defer listener.Close()
	fmt.Println("Wait connection client...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error connection client", err)
		return err
	}
	defer conn.Close()
	fmt.Println("Connection success. Listen more...")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error of reading", err)
			return err
		}
		fmt.Print(message)
	}
	return nil
}
