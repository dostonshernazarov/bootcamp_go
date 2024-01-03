package main

import (
    "fmt"
    "strconv"
)

func main () {
    var num, res string
    fmt.Scan(&num)
    
    for _, n := range num {
        dig,_ := strconv.Atoi(string(n))
        res += strconv.Itoa((dig*dig))
        
    }
    fmt.Println(res)
    
}
