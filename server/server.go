package server

import (
	"bufio"
	"fmt"
	"net"
)

type Chatter struct {
	username       string
	userconnection net.Conn
}
type ChatRoom struct {
	chatters map[string]Chatter
}

func (cr *ChatRoom) Broadcast(message string, username string) {
	for _, chatter := range cr.chatters {
		broadcastWritter := bufio.NewWriter(chatter.userconnection)
		_, err := broadcastWritter.WriteString(username + ":" + message + "\n")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = broadcastWritter.Flush()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func handleConnection(conn net.Conn, chatroom *ChatRoom) {
	defer conn.Close()
	serverWritter := bufio.NewWriter(conn)
	_, err := serverWritter.Write([]byte("hello from the server, what is your name?\n"))
	if err != nil {
		fmt.Println(err.Error())
	}

	err = serverWritter.Flush()
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	chatter := Chatter{username: scanner.Text(), userconnection: conn}
	chatroom.chatters[chatter.username] = chatter
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(chatter.username, ": ", message)
		chatroom.Broadcast(message, chatter.username)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}
}

// this function starts the server
func StartServer(port string) {

	listener, err := net.Listen("tcp", port)
	chatroom := ChatRoom{chatters: make(map[string]Chatter)}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("chat server started on port 8080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go handleConnection(conn, &chatroom)
	}

}
