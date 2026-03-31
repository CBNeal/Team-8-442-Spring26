package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

var L_alphabet_array [26]string = InitAlphabetArray(true)  // a constant to refer to throughout program
var U_alphabet_array [26]string = InitAlphabetArray(false) // a constant to refer to throughout program
var KEY string

func InitAlphabetArray(lowercase bool) [26]string {
	var result [26]string
	var start_index int

	if lowercase {
		start_index = 97 //ascii value for 'a'
	} else {
		start_index = 65 //ascii value for 'A'
	}

	for i := 0; i < 26; i++ {
		result[i] = string(rune(start_index + i))
	}
	return result
}

func DecodeIn(plaintext string) string {
	//decoding formula P = (c - k + 26) % 26
	var result string
	var key_position int = 0 //tracks index as the key is traversed

	//iterate through string and apply above formula to each character
	for i := 0; i < len(plaintext); i++ {
		//check if character is a letter
		if !unicode.IsLetter(rune(plaintext[i])) {
			if unicode.IsSpace(rune(plaintext[i])) {
				result += " "
				continue
			}
			result += string(plaintext[i])
			continue
		}

		var alphabet_array [26]string
		var character string = string(plaintext[i])

		//check if character is uppercase or lowercase
		alphabet_array = CheckCase(character)

		//find index of character in alphabet array
		character_index := slices.Index(alphabet_array[:], character)

		//navigating the key and returning its index. maybe I'll have it wrap around? idk
		key_character := string(KEY[key_position%len(KEY)])
		key_index := slices.Index(L_alphabet_array[:], key_character)
		key_position++

		//apply da freaking formula. sunglasses emoji
		decoded_index := (26 + character_index - key_index) % 26 //26 is there to avoid negative numbers
		alphabet_array = CheckCase(character)
		result += alphabet_array[decoded_index]

	}

	return result
}

func EncodeIn(plaintext string) string {
	//same exact thing as before but the math at the end is different P = = 26 + c - k) % 26
	var result string
	var key_position int = 0 //tracks index as the key is traversed

	//iterate through string and apply above formula to each character
	for i := 0; i < len(plaintext); i++ {
		if !unicode.IsLetter(rune(plaintext[i])) {
			if unicode.IsSpace(rune(plaintext[i])) {
				result += " "
				continue
			}
			result += string(rune(plaintext[i]))
			continue
		}

		var alphabet_array [26]string //init alphabet array

		var character string = string(plaintext[i])

		//check case of character. this is a really stupid way of doing this idgaf
		alphabet_array = CheckCase(character)

		//find index of character in alphabet array
		character_index := slices.Index(alphabet_array[:], character)

		//navigating the key and returning its index. maybe I'll have it wrap around? idk
		key_character := string(KEY[key_position%len(KEY)]) //this is redundant now but. oh well
		key_index := slices.Index(L_alphabet_array[:], key_character)
		key_position++

		//apply da freaking formula. sunglasses emoji
		decoded_index := (character_index + key_index) % 26

		//ensure case of output matches case of input; this should probably be its own function atp
		alphabet_array = CheckCase(character)
		result += alphabet_array[decoded_index]

	}

	return result
}

// returns the alphabet array to use based on the case of the incoming character
func CheckCase(character string) [26]string {
	if character == strings.ToUpper(character) {
		return U_alphabet_array
	} else {
		return L_alphabet_array
	}
}

func main() {
	//processing input from command line
	args := os.Args[1:]

	if len(args) != 2 {
		panic("too many arguments")
	}

	flag := args[0]

	KEY = args[1]
	KEY = strings.ToLower(KEY)             //the case of the key does not matter
	KEY = strings.ReplaceAll(KEY, " ", "") //remove spaces from key
	fmt.Printf("Key: %s\n", KEY)

	scanner := bufio.NewScanner(os.Stdin)
	var plaintext string
	var encode bool
	var decode bool

	//if flag is neither of the accepted ones
	if flag != "-e" && flag != "-d" {
		panic("invalid flag")
	}

	if flag == "-e" {
		encode = true
	} else if flag == "-d" {
		decode = true
	} else {
		panic("invalid flag, use -e for encode and -d for decode")
	}

	for {
		fmt.Println("Enter text to encode/decode: ")
		scanner.Scan()
		plaintext = scanner.Text()

		if encode {
			fmt.Printf("Encoded text: %s\n", EncodeIn(plaintext))
		} else if decode {
			fmt.Printf("Decoded text: %s\n", DecodeIn(plaintext))
		}
	}

}
