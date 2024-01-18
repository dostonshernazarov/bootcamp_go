package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()
	// Открываем порт
	conn, err := ln.Accept()
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write([]byte("message"))
	if err != nil {
		log.Println(err)
	}
}
