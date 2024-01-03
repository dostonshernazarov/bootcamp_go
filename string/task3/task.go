package main

import (
    "fmt"
    "bufio"
    "os"
    "unicode/utf8"
)

func main() {
    
    str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    
    txt := []rune(str)
    length := utf8.RuneCountInString(str) - 1
    x := string("")
    for i := 1; i<length+1;i=i+2{
        x += string(txt[i])
    }
    fmt.Println(x)
}