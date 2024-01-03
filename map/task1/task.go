package main

import (
	"fmt"
	"time"

)

func main() {

	m := make(map[int]int)

	for i:=0;i<10;i++ {
    var a int
    fmt.Scan(&a)
    if it, ok := m[a]; ok {
        fmt.Print(it," ")
    } else {
        m[a] = work(a)
        
        fmt.Print(m[a]," ")
    }
}
}

func work(n int) int {
	if n > 3 {
	   time.Sleep(time.Millisecond * 500)
	   return n + 1
	} else {
	   time.Sleep(time.Millisecond * 500)
	   return n - 1
	}
 }