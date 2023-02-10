package server

import (
	"bufio"
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	username := conn.RemoteAddr().String()
	fmt.Println("new client connection from", username)
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Closing connection from", username)
			conn.Close()
			return
		}

		fmt.Println(username, ">", message)
	}
}

func Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started and listening on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}
		go HandleConnection(conn)
	}
}
