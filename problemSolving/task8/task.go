package main

import "fmt"

func main() {

    var n, num, min int
    count := 0
    fmt.Scan(&n)
    fmt.Scan(&num)
    min = num
    count++
    for i:=1;i<n;i++ {
        fmt.Scan(&num)
        if num<min {
            min = num
            count = 1
        }else if num == min {
            count++
        }
        
    }
    fmt.Println(count)
}