package main

import "fmt"

func main() {

    var num uint
    fmt.Scan(&num)
    var tmp uint = 1
  
    for tmp<=num {
        fmt.Print(tmp, " ")
        tmp *= 2
    
    }
    
} 