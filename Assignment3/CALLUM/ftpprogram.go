package main

import (
	"fmt"
	"github.com/secsy/goftp"
	"strconv"
)

const ( 
	Address = "138.47.99.21"

	// The Decode number chooses whether it is decoding 7 or 10 bits
	decodenum = '7'
	// Used for testing, this allows us to choose the directory
	// Set to /7 or /10 in order to test off the schools ftp Server
	Dir = "/7"
)

var config = goftp.Config{
	User: "anonymous",
	Password: "",
}

func FtpConnect() {
	conn, err := goftp.DialConfig(config, Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	entries, err := conn.ReadDir(Dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	fullBinary := ""

	for i := 0; i < len(entries); i++ {
		entry := entries[i].Mode()
		a := entry.String()

		if len(a) < 10 {
			continue
		}

		workingword := ""
		for k := len(a) - 9; k < len(a); k++ {
			if a[k] != '-' {
				workingword += "1"
			} else {
				workingword += "0"
			}
		}

		fullBinary += workingword
	}

	chunkSize := 7
	if decodenum == "10" { 
		chunkSize = 10
	}

	for i := 0; i+chunkSize <= len(fullBinary); i += chunkSize {
		chunk := fullBinary[i : i+chunkSize]

		decimal, err := strconv.ParseUint(chunk, 2, 8)
		if err != nil {
			continue
		}

		char := string(rune(decimal))
		fmt.Print(char)
	}
}

func main() {
	FtpConnect()
}
