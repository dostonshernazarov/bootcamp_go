package main

import "fmt"

func main(){

    var a uint16
    fmt.Scan(&a)
    fmt.Println("It is", a/30, "hours", (a*2)%60, "minutes.")
}