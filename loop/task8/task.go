
package main

import "fmt"
func main() {

    var a, b, copyB int
    fmt.Scan(&a)
    fmt.Scan(&b)
    a1 := a
    a=0
    
    
    for a1>0 {
        a = a*10 + a1%10
        a1 /= 10
    }
    
    
    for a>0 {
        copyB = b
        for copyB>0 {
            if a%10 == copyB%10 {
                fmt.Print(a%10, " ")
            }
            copyB = copyB / 10
            
            
        }
        a = a/10
    }
}