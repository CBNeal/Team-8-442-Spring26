package main

//libraries
import(
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//function that processes ASCII backspace characters (8)
func applyBackspaces(input string) string{
	// builds the final string
	var result []rune
	//convert string to rune slice
	runes := []rune(input)
	//loop
	for i := 0; i < len(runes); i++{
		r := runes[i]
		//if ASCII backspace character
		if(r == 8){
			//remove last character 
			if(len(result) > 0){
				result = result[:len(result)-1]
			}
		}else{
			//append character
			result = append(result, r)
		}
	}
	//string result
	return string(result)
}

//function converts a binary string into text using either 7-bit or 8-bit ASCII (depending on the "bits" argument)
func decodeBinary(data string, bits int) string{
	//decoded characters
	var result []rune
	//convert string into rune for indexing
	runes := []rune(data)
	//process the binary string in chunks of size "bits" (7 or 8)
	for i := 0; i+bits <= len(runes); i += bits{
		//one binary chunk of data
		chunk := runes[i : i+bits]
		var value int
		//convert chunk to integer
		for j := 0; j < len(chunk); j++{
			b := chunk[j]
			//shift left (mult by 2)
			value <<= 1 
			if(b == '1'){
				//sets lsb if current bit is 1
				value |= 1 
			}
		}
		//convert integer to rune and stores in result
		result = append(result, rune(value))
	}
	//applies backspace before returning final string
	return applyBackspaces(string(result))
}

//main function
func main(){
	//creates a buffer to read input from stdin
	reader := bufio.NewReader(os.Stdin)
	//reads entire input
	inputBytes, err := io.ReadAll(reader)
	//error checking
	if(err != nil){
		//print error to stderr if reading fails
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	//convert input to string and trims leading and trailing whitespace
	input := strings.TrimSpace(string(inputBytes)) 
	//remove any newline or carriage return characters
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	//decodes using both 7-bit ASCII and 8-bit ASCII
	decoded7 := decodeBinary(input, 7)
	decoded8 := decodeBinary(input, 8)

	//prints both results
	fmt.Print("Bit-7 result: ")
	fmt.Print(decoded7 + "\n")
	fmt.Print("Bit-8 result: ")
	fmt.Print(decoded8)
}