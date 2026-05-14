package main

//libraries
import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

//main function
func main(){
	//main variables:
		//stores time of previous char gotten
	var prevTime time.Time
		//stores the bits unti a byte is gotten 
	var bitBuffer string
		//stores decoded message 
	var covertMessage string
	
	//connect to server 
	conn, err := net.Dial("tcp", "138.47.99.21:31337") 
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//reads incoming bits 
	reader := bufio.NewReader(conn)

	fmt.Println("[Connected]")
	
	//reads data from server 
	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}
		//prints overt message
		fmt.Printf("%c", char)
		//gets current time for delay calculation 
		now := time.Now()
		
		if !prevTime.IsZero() {
			//gets time difference between characters
			delay := now.Sub(prevTime)
			//converts to milliseconds
			ms := float64(delay.Microseconds()) / 1000.0

			//prints timing (for debugging)
			fmt.Printf(" (%.3f ms)", ms)

			//determines bit (switch if output looks weird)
			if ms < 35 {
				bitBuffer += "0"
			} else if ms > 90 {
				bitBuffer += "1"
			}

			//converts every 8 bits to char
			if len(bitBuffer) == 8 {
				//converts binary to ints 
				val, _ := strconv.ParseInt(bitBuffer, 2, 64)
				//converts int to ASCII
				c := string(rune(val))
				//append to finished message 
				covertMessage += c
				//reset buffer for next byte
				bitBuffer = ""

				//stops when 'EOF' is detected
				if len(covertMessage) >= 3 &&
					covertMessage[len(covertMessage)-3:] == "EOF" {
					break
				}
			}
		}
		//updates time 
		prevTime = now
	}
	//prints disconnection and covert message 
	fmt.Println("\n[Disconnected]")
	fmt.Println("Covert message:", covertMessage)
}
