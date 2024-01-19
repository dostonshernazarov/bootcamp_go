package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// http запрос с методом GET
	resp, err := http.Get("https://www.instagram.com/doston_shernazarov_/")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close() // закрываем тело ответа после работы с ним

	data, err := io.ReadAll(resp.Body) // читаем ответ
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s", data) // печатаем ответ как строку
}
