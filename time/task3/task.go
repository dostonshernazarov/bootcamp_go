package main

import (
	"fmt"
	"strings"
	"time" 
    "bufio"
    "os"
    "io"
)

func main() {
    str, err := bufio.NewReader(os.Stdin).ReadString('\n')
    if err != nil && err != io.EOF { panic(err) }
	res := strings.Split(str, ",")

	res[1] = strings.TrimRight(res[1], "\n")

	res[0] = strings.TrimRight(res[0], "\n")

	time1, err1 := time.Parse("02.01.2006 15:04:05", res[0])
    if err1 != nil {
        return
    }
   
	time2, err2 := time.Parse("02.01.2006 15:04:05", res[1])
    if err2 != nil {
        return
    }

	if time1.Before(time2){
		fmt.Println(time2.Sub(time1))
	} else {
		fmt.Println(time1.Sub(time2))
	}

}