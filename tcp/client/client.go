package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	for {
		resp, err := http.Get("http://127.0.0.1:8081")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		fmt.Println(string(body))

	}
}
