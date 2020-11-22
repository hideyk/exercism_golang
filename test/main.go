package main

import (
	"fmt"
)

// Reverse helps to return a reversed string
func Reverse(input string) string {
	length := len([]rune(input))
	reversed := ""
	for i := length - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}


func main() {
	input := "I'm hungry!"
	for _, char := range input {
		fmt.Println(string(char))
	}
	length := len([]rune(input))
	fmt.Println(length)
	reversed := Reverse(input)
	fmt.Println(reversed)
}
