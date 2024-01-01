package main

import "fmt"

func main() {
    
    var h,m,s int
    fmt.Scan(&s)
    h = s/3600
    m = s%3600/60
    fmt.Println("It is",h, "hours", m, "minutes.")
    
}