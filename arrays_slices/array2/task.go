package main
import "fmt"

func main() {
    // здесь должен быть ваш код
    var n,num int
    fmt.Scan(&n)
    arr := make([]int,n,n)
    for i:=0;i<n;i++ {
   
        fmt.Scan(&num)
        arr[i] = num
        
    }
    fmt.Println(arr[3])
}