package main

import (
	s "chat/server"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	s.Start(":8080")
}
