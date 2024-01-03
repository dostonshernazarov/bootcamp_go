package main

import "fmt"

func main() {
    var txt, res string
    fmt.Scan(&txt)
    for idx, item :=range txt {
        if idx == len(txt)-1 {
            break
        }
        res += string(item) + "*"
    }
    res += string(txt[len(txt)-1])
    fmt.Println(res)
}
