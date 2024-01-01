package main

import "fmt"

func main() {
    
    var num, sum int
    fmt.Scan(&num)
    for num>0 {
        sum += num%10
        num /= 10
    }
    fmt.Println(sum)
    
}