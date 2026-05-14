package main

//libraries 
import (
	"io"
	"os"
)

//buffer for inputs 
const BUF_SIZE = 4096

//main function
func main() {
	//Open key file
	keyFile, err := os.Open("key")
	if err != nil {
		os.Stderr.WriteString("Cannot open key\n")
		os.Exit(1)
	}
	//closes keyfile after program runs 
	defer keyFile.Close()

	//Read key into memory
	key, err := io.ReadAll(keyFile) 
	//checks for empty key or invalid key 
	if err != nil || len(key) == 0 {
		os.Stderr.WriteString("Invalid key\n")
		os.Exit(1)
	}
	//stores the length of the key 
	keyLen := len(key)
	//erro check for empty key 
	if keyLen == 0 {
		os.Stderr.WriteString("Empty key\n")
		os.Exit(1)
	}
	//allocate space for buffer reading input  
	buffer := make([]byte, BUF_SIZE)
	//variable for keeping key index value
	var keyIndex int
	
	for {
		//read input from cmd line
		n, err := os.Stdin.Read(buffer)
		//processes the data if the bytes were read 
		if n > 0 {
			//loops through each byte 
			for i := 0; i < n; i++ {
				//XORs each byte with key (can cycle through key if needed to)
				buffer[i] ^= key[keyIndex]
				//moves to next byte 
				keyIndex = (keyIndex + 1) % keyLen
			}

			//write result to stdout in terminal 
			_, writeErr := os.Stdout.Write(buffer[:n])
			if writeErr != nil {
				os.Stderr.WriteString("Error writing output\n")
				os.Exit(1)
			}
		}
		//checks to see if the end of input is reached 
		if err == io.EOF {
			break
		}
		//error message for any reading errors
		if err != nil {
			os.Stderr.WriteString("Error reading input\n")
			os.Exit(1)
		}
	}
}