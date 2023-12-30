package main

import "fmt"

func main() {
    var s int
    fmt.Scan(&s)
    if s>0 {
        fmt.Println("Число положительное")
    } else if s<0 {
        fmt.Println("Число отрицательное")
    } else {
        fmt.Println("Ноль")
    }

}