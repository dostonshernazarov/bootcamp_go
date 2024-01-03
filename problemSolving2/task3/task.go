package main

import "fmt"

func main() {
    var num string
    var max int
    fmt.Scan(&num)
    
    for i:=0;i<len(num);i++ {
        if max<int(num[i]) {
            max = int(num[i])
        }
    }
    fmt.Println(max-48)
}