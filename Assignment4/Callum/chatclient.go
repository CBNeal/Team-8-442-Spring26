package main

import (
	"fmt"
	"net"
	"bufio"
	
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
	/*
	var ActualMessage []rune
	var timdelays []float64
	var bits []int
	*/
	//prev := time.Now()

	for {
		bit, err := Scanner.ReadByte()
		if err != nil {
			return
		}
		// This is the actual calc for the time delay
		/*
		now := time.Now()
		delay := now.Sub(prev).Seconds()
		prev = now
		*/


		char := rune(bit)
		fmt.Print(char)
		/*
		ActualMessage += char

		timdelays = append(timdelay, delay)
	*/	
	}

		


}
