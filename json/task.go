package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

type Student struct {
  LastName   string `json:"LastName"`
  FirstName  string `json:"FirstName"`
  MiddleName string `json:"MiddleName"`
  Birthday   string `json:"Birthday"`
  Address    string `json:"Address"`
  Phone      string `json:"Phone"`
  Rating     []int  `json:"Rating"`
}

type Info struct {
  ID       int       `json:"ID"`
  Number   string    `json:"Number"`
  Year     int       `json:"Year"`
  Students []Student `json:"Students"`
}

func main() {

  data, err := ioutil.ReadAll(os.Stdin)

  if err != nil {
    fmt.Println(err)
    return
  }

  var info Info

  err = json.Unmarshal(data, &info)


  if err != nil {
    fmt.Println(err)
    return
  }

rating := 0
count := 0
var sum float64

for _, student := range info.Students {
  rating += len(student.Rating)
  count++
}

if count == 0 {
  sum =  0.0
} else {
	sum = float64(rating) / float64(count)
}

  result := map[string]float64{"Average": sum}
  resultJSON, err := json.MarshalIndent(result, "", "    ")


  if err != nil {
    fmt.Println("Error encoding result:", err)
    return
  }

  fmt.Println(string(resultJSON))
}