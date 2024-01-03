package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) < 5 {
		return false
	}

	for _, char := range password {
		if !unicode.IsDigit(char) && !unicode.Is(unicode.Latin,char) {
			return false
		}
	}

	return true
}

func main() {
    // put your code here
	var password string
	fmt.Scanln(&password)

	if isValidPassword(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}