package main

import (
    "fmt"
    "time"
)

func main() {
    var timeX string
    fmt.Scan(&timeX)
    timeParse, err := time.Parse("2006-01-02T15:04:05Z07:00", timeX)
    if err != nil {
        fmt.Print(err)
        return
    }
    
    fmt.Println(timeParse.Format("Mon Jan _2 15:04:05 MST 2006")) 
}