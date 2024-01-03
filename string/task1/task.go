package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	
)

func main() {
	text, _:= bufio.NewReader(os.Stdin).ReadString('\n')
	word := []rune(text)
	if unicode.IsUpper(word[0]) && word[len(word)-1] == '.'{
		fmt.Println("Right") 
	} else {
		fmt.Println("Wrong") 
	}
}