package main

import (
	"fmt"
	"github.com/secsy/goftp"
	"strconv"
)

const (
	Address = "138.47.99.21"

	decodenum = 7

	Dir = "/7"
)

var config = goftp.Config{
	User:     "anonymous",
	Password: "",
}

func FtpConnect() {
	conn, err := goftp.DialConfig(config, Address)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	entries, err := conn.ReadDir(Dir)
	if err != nil {
		fmt.Println("ReadDir error:", err)
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
		for k := 3; k < len(a); k++ {
			if a[k] != '-' {
				workingword += "1"
			} else {
				workingword += "0"
			}
		}

		fullBinary += workingword
	}

	message := ""

	for i := 0; i+decodenum <= len(fullBinary); i += decodenum {
		chunk := fullBinary[i : i+decodenum]

		decimal, err := strconv.ParseUint(chunk, 2, 32)
		if err != nil {
			continue
		}

		message += string(rune(decimal))
	}

	fmt.Println(message)
}

func main() {
	FtpConnect()
}
