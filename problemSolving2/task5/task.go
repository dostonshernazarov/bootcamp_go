package main

import (
	"fmt"
	"math"
)



func main(){

var a string = "Hello"
b := a
 fmt.Println(b)


}


func M() float64 {
	p:=5.5
	v:=10.5
    return (p * v)
}

func W() float64 {
	k:=2.5
    return math.Sqrt(k/M())
}

func T() float64 {
    return 6/W()
}


