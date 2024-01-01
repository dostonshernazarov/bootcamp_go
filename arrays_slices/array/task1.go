package main

import (
	"fmt"
)

func main(){
var (a, b, i int
	workArray = [10]uint8{}
	w = &workArray)
for i=range(w) {
   fmt.Scan(&w[i])
}
for i=0;i<3;i++ {
   fmt.Scan(&a, &b)
   w[a], w[b] = w[b], w[a]
}
for _,value := range(w) {
   fmt.Printf("%v ", value)
}
}