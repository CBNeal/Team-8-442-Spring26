package main

import (
	"fmt"
	"strconv"
)

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
				result = result[:len(result)-1]
			}
		}else{
			result = append(result, rune(val))
		}
	}
	return string(result)
}



func main() {
	fmt.Println(decode("100100011001011101100110110011011110100000101011111011111110010110110011001000100001", 7))
	//Just the first test file I didn't want to deal with input yet
	fmt.Println("________________________________________________________________________________________________")
	fmt.Println(decode("100100011001011101100110110011011110100000101011111011111110010110110011001000100001", 8))
	
}
