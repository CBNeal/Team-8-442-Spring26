package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"bufio"
)
// I dont think importing strings is strictly necessary but I am using it for the string builder type
// bufio is for the scanner stuff
func decode(input string, sizeChunk int) string{
	var result []rune

	for i := 0; i+sizeChunk <= len(input); i += sizeChunk{
		currChunk := input[i : i+sizeChunk]

		val, err := strconv.ParseInt(currChunk, 2, 64) //The 2 is for taking it from binary
							   //From what I read its like Atoi but it can be bin
		if err != nil{
			continue
		}

		if val == 8 { 
			if len(result) > 0 {
				result = result[:len(result)-1] // I just shrinked the array by 1 to handle backspaces
			}
		}else{
			result = append(result, rune(val))
		}
	}
	return string(result)
}



func main() {

	scanner := bufio.NewScanner(os.Stdin) 
	var input strings.Builder

	for scanner.Scan(){		//Using scanner.scan to just run until the end of the bin string no matter the length
		line := scanner.Text()
		input.WriteString(line)
	}
		
	input2 := input.String() 
	
	fmt.Println(decode(input2, 7))
	fmt.Println("________________________________________________________________________________________________")
	fmt.Println(decode(input2, 8))
	
}
