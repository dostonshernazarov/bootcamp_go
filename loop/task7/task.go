package main

import "fmt"

func main() {

    var x,p,y, year int
    fmt.Scan(&x, &p, &y)
    
    for  {
        if x>=y {
            break
        }
        year ++
        x += (x * p)/100
        
    }
    fmt.Println(year)
}