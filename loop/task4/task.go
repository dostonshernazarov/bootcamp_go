package main
import "fmt"

func main() {
    
    var a, count, max uint
    
    fmt.Scan(&a)
    max = a
    count = 1
    
    for a!=0 {
        fmt.Scan(&a)
        if a == max {
            count ++ 
        } else if a>max {
            max = a
            count = 1
            
        }
    }
    fmt.Println(count)
}