package main

import (
	"fmt"
	"strconv"
	"strings"
)

// read through the stin input as one big string. blocks of binary may be in 8 or 7 bit segments. idk
func DecodeIn(bitString string) string {
	var decodedString string
	var bitSegLen int = 0

	//figure out if the incoming string is 8 bit or not
	if (len(bitString) % 8) == 0 {
		bitSegLen = 8
	} else {
		bitSegLen = 7
	}

	for i := 0; i < len(bitString); i += bitSegLen {
		bitSeg := string(bitString[i : i+bitSegLen]) //splice segment

		//convert to int
		decimalVal, err := strconv.ParseUint(bitSeg, 2, bitSegLen)
		if err != nil {
			panic(err)
		}
		decodedString += string(decimalVal)
	}

	return decodedString
}

func main() {
	var test string
	fmt.Println("Enter an string encoded in binary (7 or 8 bit segments): ")
	fmt.Scanln(&test)
	test = strings.ReplaceAll(test, "\r\n", "") //remove carriage returns
	fmt.Println(test)
	test = DecodeIn(test)
	fmt.Println(test)
}
