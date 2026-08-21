package net

import (
	"bufio"
	"fmt"
	"net"
)

func Connect() error {
	fmt.Println("Запускаем сервер")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return err
	}
	defer listener.Close()

	fmt.Println("Ждем соединение")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Ошибка подключения клиента:", err)
		return err
	}

	fmt.Println("Соединение получили - слушаем данные")

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return err
		}

		fmt.Println(message)
	}

	return nil
}
