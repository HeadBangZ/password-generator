package main

import (
	//"crypto/rand"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	LOWERCASE = "abcdefghijklmnopqrstunwxyz"
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS   = "0123456789"
	SYMBOLS   = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func main() {
	fmt.Println("How long you want your password?")
	fmt.Println("-------------------")
	length := getUserInput()
	fmt.Printf("%s\n", fmt.Sprintf("Generate password: %d is the chosen length", *length))

	password := generatePassword(length)
	fmt.Println(password)
}

func getUserInput() *int {
	var input string
	fmt.Scanln(&input)
	length, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input has to be a number!")
		getUserInput()
	}

	if length < 8 {
		fmt.Println("Length of password needs to be at least 8")
		getUserInput()
	}

	return &length
}

// TODO: change to crypto/rand instead of math/rand, math/rand ain't secure to use.
func generatePassword(length *int) string {
	chars := UPPERCASE + LOWERCASE + NUMBERS + SYMBOLS
	pw := make([]byte, *length)
	pw[0] = UPPERCASE[rand.Intn(len(UPPERCASE))]
	pw[1] = LOWERCASE[rand.Intn(len(LOWERCASE))]
	pw[2] = NUMBERS[rand.Intn(len(NUMBERS))]
	pw[3] = SYMBOLS[rand.Intn(len(SYMBOLS))]
	for i := 4; i <= *length-1; i++ {
		pw[i] = chars[rand.Intn(len(chars))]
	}

	rand.Shuffle(len(pw), func(i, j int) {
		pw[i], pw[j] = pw[j], pw[i]
	})

	return string(pw)
}
