package main 
import "fmt"

func main() {

    var num, sum uint
    fmt.Scan(&num)
    
    for num>0 {
        sum += num%10
        num/=10
    }
    num = 0
    for sum>0 {
    
       num += sum%10
       sum /=10
    }
    fmt.Println(num)
}