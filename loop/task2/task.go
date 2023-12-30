package main

import "fmt"

func main() {

    var a,b uint
    var sum uint
    
    fmt.Scan(&a)
    fmt.Scan(&b)
    
    for i:=a; i<=b;i++ {
        sum += i
        
    }
    fmt.Println(sum)
    
}