package main
import (
    "fmt"
    "math"
)

func gip(a, b int) float64 {
    return math.Sqrt(float64(a*a + b*b))
}
func main() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(gip(x ,y ))
}
