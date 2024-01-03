package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
    var txt string
    txt, _ = bufio.NewReader(os.Stdin).ReadString('\n') 
    res := strings.Split(txt, ";")
	num1 := strings.ReplaceAll(res[0], " ", "")
	num2 := strings.ReplaceAll(res[1], " ","")
    num1 = strings.ReplaceAll(num1, ",", ".")
	num2 = strings.ReplaceAll(num2, ",",".")
	n1, err1 := strconv.ParseFloat(num1, 64)
	n2, err2 := strconv.ParseFloat(num2, 64)
    if err1!=nil || err2!=nil {
        fmt.Println("")
    } else {
        fmt.Printf("%.4f", n1/n2)
    }
}