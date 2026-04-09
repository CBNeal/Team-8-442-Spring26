package main

//libaries
import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

//variables for ftp site and method decoding
const (
	//ip address
	addr = "localhost"
	//port
	port = "21"
//--ONE-OR-THE-OTHER--------------------
	//7bit permissions decode
	METHOD = "7bit"
	//10bit full permissions decode
	//METHOD = "10bit" 
//--------------------------------------
	//ftp login username and password
	user = "anonymous"
	passw = "anon"
)

//main function
func main(){
	//variables for storing file name data
	var lines []string
	//connects to ftp site
	connect, err := net.Dial("tcp", addr + ":" + port)
	if(err != nil){
		fmt.Println("Error connecting to server: ", err)
		return
	}
	//connected to server
	read := bufio.NewReader(connect)
	//input anonymous username and password
	fmt.Fprintf(connect, "USER %s\r\n", user)
	read.ReadString('\n')
	fmt.Fprintf(connect, "PASS %s\r\n", passw) 
	read.ReadString('\n')
	//go to either method folder
	if(METHOD == "7bit"){
		fmt.Fprintf(connect, "CWD 7/folder\r\n")
		read.ReadString('\n')
	}else if(METHOD == "10bit"){
		fmt.Fprintf(connect, "CWD 10/folder\r\n")
		read.ReadString('\n')
	}
	//fetch file listing including permissions
	fmt.Fprintf(connect, "LIST\r\n")
	//inf loop till done reading files
	for{
		line, err := read.ReadString('\n')
		if(err != nil){
			break
		}
		//appends file name data to array
		lines = append(lines, line)
		if(strings.Contains(line, "256")){
			break
		}
	}
	//disconnects from ftp server
	fmt.Fprintf(connect, "QUIT\r\n") //quit command
	connect.Close()					 //closes tcp connection
	//uses decode method to decode permissions and print results to cmd line
	decode(lines)
}
//-------------------------------------------------------------------
//decode function that grabs either 7 or 10 bits and stores for conversion
func decode(lines []string){
	//variables
	bits := ""
	//7bit enabled
	if(METHOD == "7bit"){
		//for all data in data 
		for _, line := range lines{
			//slices the data into different components
			fields := strings.Fields(line)
			if(len(fields) == 0){
				continue
			}
			//permissions slice of data 
			perm := fields[0]
			//filter first 3bits (owner permissions)
			if(len(perm) < 10){
				continue
			}
			if(perm[1] != '-' || perm[2] != '-' || perm[3] != '-'){
				continue
			}
			//keep and append last 7 bits to be decoded to text
			bits += perm[3:]
		}
	//10bit enabled
	}else if(METHOD == "10bit"){
		//for all data in data 
		for _, line := range lines{
			//for all data in data 
			fields := strings.Fields(line)
			if(len(fields) == 0){
				continue
			}
			//permissions slice of data 
			perm := fields[0]
			//keep and append all permissions to be decoded to text
			bits += perm
		}
	}
	//uses below function to convert to binary then text and output message
	fmt.Println(binaryToText(bits))
}

//binary decoder function
func binaryToText(input string) string {
	//variables
	clean := ""
	var out string
	//for data in input 
	for _, c := range input {
		if(c == '0' || c == '1'){
			//moves bits to clean variable
			clean += string(c)
		}
	}
	//for len of clean, increment by seven
	for i := 0; i+7 <= len(clean); i += 7 {
		//takes and separates 7-bit data chunk
		byteData := clean[i : i+7]
		val := 0
		//for range of byteData
		for _, b := range byteData {
			//creates integer from byteData 
			val <<= 1
			if(b == '1'){
				val |= 1
			}
		}
		//converts value to ASCII string and returns it
		//byte makes val be an 8-bit value
		//string convert to ASCII
		out += string(byte(val))
	}
	//return converted password 
	return out
}