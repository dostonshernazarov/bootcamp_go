package main

import "fmt"
 func main() {

	a, b := 4, 5
	test(&a, &b)
 }

func test(x1 *int, x2 *int) {
    fmt.Println(*x1 * *x2)
}