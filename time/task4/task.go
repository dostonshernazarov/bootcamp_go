package main

import (
    "fmt"
    "time"
)

// вам это понадобится
const now = 1589570165



func main() {

    
    nowCopy := now
	var min, sec int64
	fmt.Scanf("%d мин. %d", &min, &sec)

    timeU := (min * 60 + sec) + int64(nowCopy)
	totalDate := time.Unix(timeU, 0)
	fmt.Println(totalDate.Format(time.UnixDate))
}