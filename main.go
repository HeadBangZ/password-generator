package main

import (
	"fmt"
	"strconv"
)

// Make these into slices?
// TODO: Make these into slices
const LOWERCASE = "abcdefghijklmnopqrstunwxyz"
const UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMBERS = "0123456789"
const SYMBOLS = "!\"#$%&'()*+,-.:;<>"

func main() {
	fmt.Println("How long you want your password?")
	fmt.Println("-------------------")

	length := getUserInput()
	fmt.Printf("%d\n", length)

	generatePassword(length)
}

func getUserInput() int {
	var input string
	fmt.Scanf("%s", &input)
	length, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input has to be a number!")
		getUserInput()
	}

	return length
}

func generatePassword(n int) string {
	fmt.Printf("%s\n", fmt.Sprintf("Generate password: %d is the chosen length", n))

	return ""
}
