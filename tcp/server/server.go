package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var input string
		fmt.Scan(&input)
		fmt.Fprint(w, input)
	})

	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
