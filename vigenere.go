package main

//libraries
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

//function converts the key into numeric shifts (A=0 ... Z=25)
func normalizeKey(key string) []int{
	var shifts []int
	// loop through each char in the key
	for i := 0; i < len(key); i++{
		r := rune(key[i])
		//conditional for only processing letters
		if unicode.IsLetter(r){
			//convert to uppercase and map to 0–25
			r = unicode.ToUpper(r)
			shifts = append(shifts, int(r-'A'))
		}
	}
	return shifts
}

//function applies the Vigenère cipher to the input text
//true = encryption && false = decryption
func vigenere(text string, key []int, encrypt bool) string{
	var result []rune
	//position in key
	keyIndex := 0           
	//length of key
	keyLen := len(key)      
	//process each char in the input text
	for i := 0; i < len(text); i++{
		r := rune(text[i])
		//conditional that only shifts alphabet
		if unicode.IsLetter(r){
			//get corresponding shift value from key
			shift := key[keyIndex%keyLen]
			//reverse shift for decryption
			if !encrypt{
				shift = -shift
			}
			//conditional for handling uppercase letters
			if unicode.IsUpper(r) {
				newChar := (int(r-'A') + shift + 26) % 26
				result = append(result, rune(newChar+'A'))
			}else{
				//handles lowercase letters
				newChar := (int(r-'a') + shift + 26) % 26
				result = append(result, rune(newChar+'a'))
			}
			//move to next key char
			keyIndex++
		}else{
			//keeps non-letter chars the same 
			result = append(result, r)
		}
	}
	return string(result)
}

//main function
func main(){
	//ensure correct number of command-line arguments
	if len(os.Args) < 3{
		fmt.Fprintln(os.Stderr, "Format: ./vigenere-[e|d] key")
		return
	}
	// -e or -d for encrypt or decrypt respectively
	mode := os.Args[1]     
	//encryption/decryption key
	keyInput := os.Args[2] 
	//determine if encrypt or decrypt
	encrypt := true
	if mode == "-d"{
		encrypt = false
	}else if mode != "-e"{
		fmt.Fprintln(os.Stderr, "Error: can either be -e or -d")
		return
	}
	//converts key into numeric shifts
	key := normalizeKey(keyInput)
	//error checking for if key contains at least one valid letter
	if len(key) == 0{
		fmt.Fprintln(os.Stderr, "Error: key must contain letters")
		return
	}
	//creates scanner to read input from stdin 
	scanner := bufio.NewScanner(os.Stdin)
	//processes each line of input
	for scanner.Scan(){
		line := scanner.Text()
		//applies Vigenère cipher and prints result
		output := vigenere(line, key, encrypt)
		fmt.Println(output)
	}
}