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
	// Used for testing, this allows you to choose the directory
	Dir = "/7"
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

	entries, err := conn.ReadDir(Dir)
	if err != nil {
		return
	}

	for i := 0; i < len(entries); i++{
		entry  := entries[i].Mode()
		
		a := entry.String()
		workingword := ""
		for k := 3; k < len(a); k++{

			if a[k] != '-'{
				workingword += "1"
			}else{
				workingword += "0"
			}
		}

		decimal, err := strconv.ParseUint(workingword, 2,8)
		fmt.Println(decimal)
		if err != nil {
			return
		}

		//char := string(rune(decimal))

		//fmt.Print(char)



	}

}

func main() {
	FtpConnect()
}
