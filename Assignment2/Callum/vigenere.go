package main

import (
	"fmt"
	"strings"
)

func lowercaseChart(input rune) int{
	result := (int(input) -97)
	return result
}

func uppercaseChart(input rune) int{
	result := (int(input)- 65)
	return result
}

func encrypt(text string, key string) string{
	key = strings.ToLower(key)
	var result rune
	var final  strings.Builder
	keyNumber := len(key)
	keyI := 0

	for i := 0; i < len(text); i++{
		if int(text[i]) <= 122 && int(text[i]) >= 97{
			result := (lowercaseChart(rune(text[i])) + lowercaseChart(rune(key[keyI]))) % 26
			resultRune := rune(result + 97) 
			final.WriteRune(resultRune)

			keyI = (keyI + 1) % keyNumber // Modulus math just resets to 0 if gets too high
		}
		else if int(text[i]) <= 90 and int(text[i]) >= 65{
			result := (uppercaseChart(rune(text[i])) + uppercaseChart(rune(key[keyI]))) % 16
			resultRune := rune(result + 65)
			final.WriteRune(resultRune)
			
			keyI = (keyI + 1) % keyNumber
		}
	}

	return final.String()
}
			
		


func main(){
	fmt.Println(lowercaseChart('a'))
	fmt.Println(uppercaseChart('A'))
	test := encrypt("test", "key")
	fmt.Println(test)

}



