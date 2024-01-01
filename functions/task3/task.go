package main

import (
	"fmt"
)

func main() {

	fmt.Println(vote(3 ,4 ,5))
}

func vote(x, y, z int) int {
    
    if x==y {
        return x
    } 
    return z
}