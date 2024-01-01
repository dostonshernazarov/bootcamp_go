package main

import "fmt"

func main() {

    var n,a int
    fmt.Scan(&n)
    arr := make([]int,n,n)
    
    for i:=0;i<n;i++ {
    
        fmt.Scan(&a)
        arr[i] = a
        
    }
    for i:=0;i<n;i++ {
        if i%2 == 0 {
            fmt.Print(arr[i], " ")
        }
    }
}