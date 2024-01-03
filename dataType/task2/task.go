package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(adding("Hello", "Go"))

}


func adding(str1, str2 string) int64 {
	r1 := []rune(str1)
	r2 := []rune(str2)
	  
	sum1, sum2 := 0, 0
	  
	  
	for i := 0; i < len(r1); i++ {
	  item, err := strconv.Atoi(string(r1[i]))
	  if err != nil {
		continue
	  }
	  sum1 = sum1*10 + item
	}
	  
	  
	for i := 0; i < len(r2); i++ {
	  item, err := strconv.Atoi(string(r2[i]))
	  if err != nil {
		continue
	  }
	  sum2 = sum2*10 + item
	}
	return int64(sum1 + sum2)
  }