package main

import (
	"fmt"
	"github.com/secsy/goftp"

)

const ( 
	Address = "138.47.99.21"

	// The Decode number chooses whether it is decoding 7 or 10 bits
	decodenum = '7'
)

var config = goftp.Config{
	User: "anonymous",
	Password: "",
}


func FtpConnect() {
	conn, err := goftp.DialConfig(config, Address)
	if err != nil {
		return
	}

	entries, err := conn.ReadDir("/7")
	if err != nil {
		return
	}

	for i := 0; i < len(entries); i++{
		entry  := entries[i].Mode()
		
		fmt.Printf("%s\n", entry)
		a := entry.String()
		workingword := ""
		for k := 3; k < len(a); k++{
			if a[k] != '-'{
				workingword += "1"
			}else{
				workingword += "0"
			}
		}
		fmt.Printf("%s\n", workingword)


	}

}

func main() {
	FtpConnect()
}
