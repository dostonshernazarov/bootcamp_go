package main

import (
	"encoding/json"
	"fmt"
	"github.com/k0kubun/pp"
	"io"
	"net/http"
	"strings"
)

type Translation struct {
	TranslatedText string `json:"translatedText"`
}

type TranslationData struct {
	Translations []Translation `json:"translations"`
}

type ResponseData struct {
	Data TranslationData `json:"data"`
}

func main() {

	var codeFrom, codeTo, wordFrom, checkFrom string
	for {

		lanCheck := "Choole language: \n For English press -> 1 \n For Russian press -> 2 "

		fmt.Println(lanCheck)
		fmt.Scan(&checkFrom)

		if checkFrom == "1" {
			codeFrom = "en"
			codeTo = "ru"
			fmt.Print("Write word: ")
			fmt.Scan(&wordFrom)
			break

		} else if checkFrom == "2" {
			codeFrom = "ru"
			codeTo = "en"
			fmt.Print("Напиши слово: ")
			fmt.Scan(&wordFrom)
			break

		} else {
			fmt.Println("Error Language \n")
			continue
		}

	}

	//httpReq := ("source=en&target=ru&q=Hello%2C%20world!")

	result := fmt.Sprintf("%s%s%s%s%s", "source=", codeFrom, "&target=", codeTo, "&q=", wordFrom)
	//fmt.Println(result)

	url := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	payload := strings.NewReader(result)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "9545fea372msh7bb03bc0f210d67p1cc72ajsn1f6b00e62b90")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	//pp.Println(string(body))

	var response ResponseData
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pp.Println(response.Data.Translations[0].TranslatedText)

}
