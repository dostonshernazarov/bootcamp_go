package main
import "fmt"

func main() {
    // put your code here
    var n float64 
    fmt.Scan(&n)
    
    if n>0 && n<10000 {
        fmt.Printf("%.4f", (n*n))
    }
    if n<1 {
        fmt.Printf("число %.2f не подходит", n)
    }
    if n>10000 {
        fmt.Printf("%e",n)
    }
}