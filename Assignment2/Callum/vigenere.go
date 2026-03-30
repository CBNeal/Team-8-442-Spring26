package main

import (
	"fmt"
)

func lowercaseChart(input rune) int{
	result := (int(input) -96)
	return result
}

func uppercaseChart(input rune) int{
	result := (int(input)- 64)
	return result
}
/*
func encrypt(text string, key string) string{
	var result strings.Builder
	keyNumber := len(key)
	keyI := 0

	for i := 0; i <= len(text); i++{
	*/

func main(){
	fmt.Println(lowercaseChart('a'))
	fmt.Println(uppercaseChart('A'))
}



