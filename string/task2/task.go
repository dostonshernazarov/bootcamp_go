package main
import "fmt"

func main() {
    var x, s string
    fmt.Scan(&x, &s)
    check := 0
    txt := []rune(x)
    for i:=0;i<len([]rune(x));i++ {
        if string(txt[i:len([]rune(s))+i]) == s {
            fmt.Println(i)
            check = 1
            break
        }
    }
    if check == 0{
        fmt.Println(-1)
    }
}