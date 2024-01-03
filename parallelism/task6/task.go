package main

import (
	"fmt"
)

func main() {

	first := make(chan int)
	second := make(chan int)
	stop := make(chan struct{})
 fmt.Println(calculator(first, second, stop))

}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
    channel := make(chan int)
    
    go func() {
        defer close(channel)
        select{
        case msg1:=<-firstChan:
            channel <- (msg1*msg1)
        case msg2:=<-secondChan:
            channel <- (msg2*3)
        case <-stopChan:
        }
    } ()
    return channel
}