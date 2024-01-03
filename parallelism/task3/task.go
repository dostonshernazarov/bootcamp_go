package main

import (
	"fmt"
)

func main() {


inputStream := make(chan string)
outputStream := make(chan string)
go removeDuplicates(inputStream, outputStream)

go func() {
   defer close(inputStream)

   for _, r := range "112334456" {
      inputStream <- string(r)
   }
}()

for x := range outputStream {
   fmt.Print(x)
}
}





func removeDuplicates(inputStream chan string, outputStream chan string) {
    var x string
    for i := range inputStream {
        if i != x {
            outputStream <- i
            x = i
        }
    }
    close(outputStream)
}
