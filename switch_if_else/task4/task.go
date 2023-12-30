package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)
    if a%10 + a/10%10 + a/100%10 == a/1000%10 + a/10000%10 + a/100000 {
        fmt.Println("YES")
    }else {
        fmt.Println("NO")
    }
}