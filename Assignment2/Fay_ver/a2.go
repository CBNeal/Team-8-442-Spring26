package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var alphabet_array [26]string = InitAlphabetArray() // a constant to refer to throughout program
var KEY string

func InitAlphabetArray() [26]string {
	var result [26]string
	var start_index int = 65 //ascii value for 'A'
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
		if string(plaintext[i]) == " " {
			result += " "
			continue
		}
		var character string = string(plaintext[i])
		character = strings.ToUpper(character) //convert to uppercase for easier indexing

		//find index of character in alphabet array
		character_index := slices.Index(alphabet_array[:], character)

		//navigating the key and returning its index. maybe I'll have it wrap around? idk
		key_character := string(KEY[key_position%len(KEY)])
		key_character = strings.ToUpper(key_character)
		key_index := slices.Index(alphabet_array[:], key_character)
		key_position++

		//apply da freaking formula. sunglasses emoji
		decoded_index := (26 + character_index - key_index) % 26 //26 is there to avoid negative numbers
		result += alphabet_array[decoded_index]
		fmt.Printf("character: %s, character_index: %d, key_character: %s, key_index: %d, decoded_index: %d\n", character, character_index, key_character, key_index, decoded_index)

	}

	return result
}

func EncodeIn(plaintext string) string {
	//same exact thing as before but the math at the end is different P = = 26 + c - k) % 26
	var result string
	var key_position int = 0 //tracks index as the key is traversed

	//iterate through string and apply above formula to each character
	for i := 0; i < len(plaintext); i++ {
		if string(plaintext[i]) == " " {
			result += " "
			continue
		}
		var character string = string(plaintext[i])
		character = strings.ToUpper(character) //convert to uppercase for easier indexing

		//find index of character in alphabet array
		character_index := slices.Index(alphabet_array[:], character)

		//navigating the key and returning its index. maybe I'll have it wrap around? idk
		key_character := string(KEY[key_position%len(KEY)])
		key_character = strings.ToUpper(key_character)
		key_index := slices.Index(alphabet_array[:], key_character)
		key_position++

		//apply da freaking formula. sunglasses emoji
		decoded_index := (character_index + key_index) % 26
		result += alphabet_array[decoded_index]
		fmt.Printf("character: %s, character_index: %d, key_character: %s, key_index: %d, decoded_index: %d, final_result: %s\n", character, character_index, key_character, key_index, decoded_index, result)

	}

	return result
}

func main() {
	key := os.Args[1:]
	var encode bool
	var decode bool
	encode_test := "cyberstorm is going to be bussin"
	decode_test := "nGmni gc ezvi ry hcmvwzx Tskcxipo ggzlcb agdlmex rri ioc!"

	if len(key) < 2 {
		panic("usage: program -e|-d <value>")
	}

	if key[0] == "-e" {
		encode = true
	} else if key[0] == "-d" {
		decode = true
	} else {
		panic("invalid flag, use -e for encode and -d for decode. key being read: " + key[0])
	}

	_ = encode
	_ = decode

	KEY = "cryptids"
	encodeIn := EncodeIn(encode_test)
	decodeIn := DecodeIn(decode_test)
	fmt.Println("encodeIn: ", encodeIn)
	fmt.Println("decodeIn: ", decodeIn)

}
