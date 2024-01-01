package main
import "fmt"

func main() {
    
    var n, num int
    fmt.Scan(&n)
    count := 0
    arr := make([]int, n, n)
    
    for i:=0; i<n;i++ {
        fmt.Scan(&num)
        arr[i] = num
        
    }
    for i:=0;i<n;i++ {
        if arr[i]>0 {
            count++
        }
    }
    fmt.Println(count)
    

}