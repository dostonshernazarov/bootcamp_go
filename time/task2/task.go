package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
    timeX, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    timeParse, err := time.Parse("2006-01-02 15:04:05", timeX[:19])
    if err != nil {
        fmt.Print(err)
        return
    }

    if timeParse.Hour()>13 {
        timeParse = timeParse.AddDate(0, 0, 1)
    }
    
    fmt.Println(timeParse.Format("2006-01-02 15:04:05")) 
}