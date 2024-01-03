package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		work()
	}()
		time.Sleep(1*time.Second)
}


func work() {
	time.Sleep(3 * time.Second)
	fmt.Println("nil")
 }