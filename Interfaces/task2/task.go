package main

import (
    "fmt" 
    "strings"
)


    
    type B struct {
        str string
    }
    
    func (b2 B) String() string {
        a := b2.str
        count := strings.Count(a, "1")
        a = strings.Repeat(" ", 10-count)
        a += strings.Repeat("X", count)
        return fmt.Sprintf("[%v]", a)
    }

func main() {

    var txt string
    fmt.Scan(&txt)
    
    batteryForTest := B{str:txt}

	fmt.Printf(batteryForTest.str)

}