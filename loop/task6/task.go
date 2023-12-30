package main

import "fmt"

func main() {

    var num int
    
    
    for {
        fmt.Scan(&num)
        if num>100 {
            break
        }else  if num<10 {
            continue
        }
        fmt.Println(num)
    }
}