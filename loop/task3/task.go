package main

import "fmt"

func main() {

    var n, sum, num int
    fmt.Scan(&n)
    for i:=0;i<n;i++{
    
        fmt.Scan(&num)
        if num%8==0 && num>9 && num<100{
            sum += num
        }
    }
    fmt.Println(sum)
}