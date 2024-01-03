package main

import (
	"fmt"
)

func main() {
    var a,b int
    _, errA := fmt.Scan(&a)
    _, errB := fmt.Scan(&b)
    
    if errA!=nil || errB!=nil || a == 0 || b==0 {
        fmt.Println("ошибка")
    } else {
        result, _ := divide(a, b)
        fmt.Println(result)
    }
    
}

func divide(a int, b int) (int, error) {

    return a / b, nil

}