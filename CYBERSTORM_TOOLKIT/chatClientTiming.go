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

	var invBitBuffer string
		//stores decoded message 
	var covertMessage string

	var invCovertMessage string
	
	//connect to server 
	conn, err := net.Dial("tcp", "138.47.99.31:12321") 
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	//reads incoming bits 
	reader := bufio.NewReader(conn)

	fmt.Println("[Connected]")
	
	//added
	_, err = conn.Write([]byte("Deimos\n"))
	if err != nil {
		panic(err)
	}

	

	//done

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
			if ms <= 55 {
				bitBuffer += "1"
				invBitBuffer += "0"
			} else if ms > 55 {
				bitBuffer += "0"
				invBitBuffer += "1"
			} //else if ms < 18 || ms > 115 {
			//	continue 
			//} 
			
			//converts every 8 bits to char
			if len(bitBuffer) == 8 && len(invBitBuffer) == 8 {
				//converts binary to ints 
				val, _ := strconv.ParseInt(bitBuffer, 2, 64)
				valinv, _ := strconv.ParseInt(invBitBuffer, 2, 64)
				//converts int to ASCII
				c := string(rune(val))
				d := string(rune(valinv))
				//append to finished message 
				covertMessage += c
				invCovertMessage += d
				//reset buffer for next byte
				bitBuffer = ""
				invBitBuffer = ""

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
	fmt.Println(" _________________________________________ ")
	fmt.Println("inverse message: ", invCovertMessage)
}
