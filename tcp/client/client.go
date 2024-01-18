package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	message := make([]byte, 1024)
	n, err := conn.Read(message)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(message[:n]))
}
