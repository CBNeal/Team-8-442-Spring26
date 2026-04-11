package main

import (
	"fmt"

	"github.com/goftp/client"
)

const (
	ftp_server = "ftp.example.com"    //change
	ip_address = "0.0.0.0"            // chnage
	directory  = "/path/to/directory" // change
)

func main() {
	config := client.Config{
		User:     "username", // change
		Password: "password", // change
	}

	conn, err := client.DialConfig(config, ftp_server)

	if err != nil {
		panic("Failed to connect to FTP server: %v\n", err)
	}

	defer conn.Quit()

	//track entries
	entries, err := conn.List(directory)
	if err != nil {
		panic("failed to reach directory")
	}

	//print out dir contents
	for _, entry := range entries {
		fmt.Printf(entry.Name + "\n")
	}
}
