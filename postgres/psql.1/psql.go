package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"

	_ "github.com/lib/pq"
	trncsh "nt_bootcamp/bootcamp_go/postgres/transaction"
)



func main() {

	conStr := "user=doston password=doston dbname=najot_talim sslmode=disable"
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	respMsg := []byte(`
	{
		"username":"Rustamjon Rahmonov",
		"message":{
			"content":"Hello world"
		}
	}
	`)

	var user1 trncsh.User

	if err = json.Unmarshal(respMsg, &user1); err!= nil{
		panic(err)
	}

	trncsh.InsertUserMessage(db, &user1, &user1.Messages)




}
