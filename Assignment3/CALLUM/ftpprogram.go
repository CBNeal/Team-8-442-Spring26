package main

import (
	"fmt"
	"github.com/goftp/client"

)

var config = client.Config{
	User: "anonymous",
	Password: "",
}


func FtpConnect() {
	conn, err := client.DialConfig(config, "138.47.99.21")
	if err != nil {
		return
	}

	defer conn.Quit()

	entries, err := conn.List("/")
	if err != nil {
		return
	}

	for i := 0; i < len(entries); i++{
		a := entries[i]
		fmt.Println(a.name)
	}

}


func main() {
	FtpConnect()
}
