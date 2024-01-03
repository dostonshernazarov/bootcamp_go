package main

import (
	"fmt"
)

func main() {

	first := make(chan int)
	done := make(chan struct{})
	fmt.Println(calculator(first, done))

}



func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
    channel:=make(chan int)
    sum := 0
    
    go func() {
        defer close(channel)
        
        for {
            select{
            case num:=<-arguments:
                sum += num
            case <-done:
                channel<-sum
                return
            }
        }
            
    } ()
    
    return channel
}