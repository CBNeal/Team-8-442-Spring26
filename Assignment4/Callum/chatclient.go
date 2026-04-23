package main

import (
	"fmt"
	"net"
	"bufio"
	"time"
	"strings"
	
)

const (
	add = "138.47.99.21:31337")

func main() {

	fmt.Println("Connecting to the chat server ...")

	conn, err := net.Dial("tcp", add)
	if err != nil {
		fmt.Println("FAILED TO CONNECT TO SERVER")
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected")
	fmt.Println()

	Scanner := bufio.NewReader(conn)
	
	var timedelays []float64
	var bits []int
	var NormalMessage strings.Builder

	prev := time.Now()

	for {
		bit, err := Scanner.ReadByte()
		if err != nil {
			return
		}
		// This is the actual calc for the time delay
		
		now := time.Now()
		delay := now.Sub(prev).Seconds()
		prev = now
		


		char := string(bit)
		NormalMessage.WriteByte(bit)
		fmt.Print(char)
		

		timedelays = append(timedelays, delay)

		messageCheck := NormalMessage.String()
		if len(messageCheck) >= 3{
			if messageCheck[len(messageCheck)-3:] == "EOF"{
				break
			}
		}

		fmt.Println("Disconnected")

		_ = timedelays

		_ = bits
		
	}

		


}
