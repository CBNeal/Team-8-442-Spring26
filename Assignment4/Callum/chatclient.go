package main

import (
	"fmt"
	"net"
	"bufio"
	"time"
	"strings"
	
)

const (
	add = "138.47.99.31:12321"
	
)

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
	

	_, err = conn.Write([]byte("Deimos\n"))
	if err != nil {
		panic(err)
	}

	var bits []int
	var timedelays []float64
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


		_ = timedelays

	}
	fmt.Println()
	fmt.Println("Disconnected")

	// Im just taking the min and max delays average as a midpoint
	for i:= 0; i < len(timedelays); i++{
		fmt.Println(timedelays[i])
	}
	
	minD := timedelays[0]
	maxD := timedelays[0]

	for i := 0; i < len(timedelays); i++{
		if timedelays[i] < minD{
			minD = timedelays[i]
		}
		if timedelays[i] > maxD{
			maxD = timedelays[i]
		}
	}

	//fmt.Println(minD)
	//fmt.Println(maxD)
	midpoint := (maxD + minD) / 2.0

	_= midpoint 
	fmt.Println()
	for i := 0; i < len(timedelays); i++ {
    	if timedelays[i] < .01{
        	bits = append(bits, 0)
    	} else {
        	bits = append(bits, 1)
    	}
	}

	fmt.Println(bits)

var finalstring []rune

for i := 0; i+8 <= len(bits); i += 8 {
    var b byte = 0
    for j := 0; j < 8; j++ {
        b = (b << 1) | byte(bits[i+j])  // shift left and OR in the next bit
    }
    if b != 0 {  // skip null bytes/padding
        finalstring = append(finalstring, rune(b))
    }
}

fmt.Println(string(finalstring))

}	
