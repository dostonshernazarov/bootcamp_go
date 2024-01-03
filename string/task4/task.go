package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "unicode/utf8"
)

func main() {
    // put your code here
    str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    
    txt := []rune(str)
    length := utf8.RuneCountInString(str) 
    x := string("")
    for i := 0; i<length;i++{
        if strings.Count(str, string(txt[i])) < 2{
        
            x += string(txt[i])
        }
    }
    fmt.Println(x)
    
}