package main

import "fmt"

func main() {
    
    var a, b, check int 
    fmt.Scan(&a,&b)
    for i:=b;i>=a;i-- {
        if i%7==0 {
            check++
            fmt.Println(i)
            break
        }
    }
    if check==0 {
        fmt.Println("NO")
    }
}