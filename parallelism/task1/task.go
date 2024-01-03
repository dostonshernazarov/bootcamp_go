package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	fmt.Println(task(channel, 5))
}

func task(ch chan int, N int) chan int{
    ch <- (N+1)
    return ch
}
