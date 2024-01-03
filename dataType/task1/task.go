package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert(10))

}

func convert(x int64) uint16 {
    return uint16(x)
}