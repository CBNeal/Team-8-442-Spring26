package main

import (
	"fmt"
	"github.com/secsy/goftp"

)

var config = goftp.Config{
	User: "anonymous",
	Password: "",
}


func FtpConnect() {
	conn, err := goftp.DialConfig(config, "138.47.99.21:21")
	if err != nil {
		return
	}

	entries, err := conn.ReadDir("/")
	if err != nil {
		return
	}

	for i := 0; i < len(entries); i++{
		a := entries[i]
		fmt.Println(a.Name)
	}

}


func main() {
	FtpConnect()
}
