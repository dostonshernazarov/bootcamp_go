package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	fmt.Println(task2(channel, "Go"))
}


func task2(ch chan string, s string) chan string{
    for i:=0; i<5; i++ {
        ch <- (s + " ")
    }
    return ch
}