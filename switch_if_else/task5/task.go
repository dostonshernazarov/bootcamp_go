package main

import "fmt"

func main(){

    var year uint
    
    fmt.Scan(&year)
    
    if year%100 == 0 && year%400 != 0{
        fmt.Println("NO")
    } else if year%4==0 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}