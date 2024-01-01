package main

import (
	"fmt"
)

func main() {
	x, a := sumInt(1,2,3,4)

	fmt.Println(x, a)
}

func sumInt(a ...int) (int, int) {
    count := 0
    sum := 0
    for _, elem := range a {
        count++
        sum += elem
    }
    
    return count, sum
}