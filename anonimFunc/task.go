package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(num uint) uint {
    
		t := strconv.FormatUint(uint64(num), 10)
		
		str := ""
		for _, digitChar := range t {
			digit, _ := strconv.Atoi(string(digitChar))
			if digit != 0 && digit%2 == 0 {
				str += string(digitChar)
			}
		}
	
		result, _ := strconv.Atoi(str)
		if result == 0 {
			return 100
		}
	
		return uint(result)
	}

	fmt.Println(fn)
}