package main

import "fmt"

func main() {
    var num, a1, a2,f, idx int
    check := -1
    a2 = 1
    idx = 1
    fmt.Scan(&num)
    
    for f<num {
        f = a1 + a2
        idx++
        if num == f {
            check = idx
            break
        }
        a1 , a2 = a2, f
        }
    fmt.Println(check)
    }

