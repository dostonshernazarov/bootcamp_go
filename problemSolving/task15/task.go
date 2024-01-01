package main
import "fmt"

func main() {
  var n, m, l, rep int
    fmt.Scan(&n, &m)
  
  tmp := n
  for tmp != 0 {
    l++
    tmp /= 10
  }
  
  td := 1
  for i := 1; i < l; i++ {
    td *= 10
  }
  for i := 0; i < l; i++ {
    if n/td != m {
      rep = rep*10 + (n / td)
    }
    n = n % td
    td = td / 10
  }
  fmt.Println(rep)
}