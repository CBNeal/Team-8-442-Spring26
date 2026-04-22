package main

import (
	"fmt"
	"net"
	
	
)

func main() {

	fmt.Println("Connecting to the chat server ...")

	conn, err := net.Dial("tcp", "1338.47.99.21:31337")
	if err != nil {
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected")
	fmt.Println()
}
