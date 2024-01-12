package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	ID int
	Full_name string
	Username string
	Message Messages
}

type Messages struct{
	ID int
	MSG_text string
}

type User_massage struct {
	ID int
	User_id int
	Message_id int
}

func main() {

	conStr := "user=doston password=doston dbname=najot_talim sslmode=disable"
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	respMsg := []byte(`
	{
		"full_name":"Oybek atamatov",
		"username":"oybekDM",
		"message":{
			"msg_text":"hello world"
		}
	}
	`)

	var user1 Users

	if err = json.Unmarshal(respMsg, &user1); err!= nil{
		panic(err)
	}

	var reqUserId Users

	rowUser := db.QueryRow(`INSERT INTO users (full_name, username) VALUES ($1,$2) returning id`, user1.Full_name, user1.Username)	

	if err = rowUser.Scan(&reqUserId.ID); err!=nil {
		panic(err)
	}
	
	var reqMsgId Messages

	rowMsg := db.QueryRow(`INSERT INTO messages (message_text) VALUES ($1) returning id`,user1.Message.MSG_text)
	if err = rowMsg.Scan(&reqMsgId.ID); err!=nil {
		panic(err)
	}

	var U_M User_massage
	rowUM := db.QueryRow(`INSERT INTO user_message (usera, message_id) VALUES ($1,$2) returning id`,reqUserId.ID,reqMsgId.ID)
	if err := rowUM.Scan(&U_M.ID); err!=nil {
		panic(err)
	}

	fmt.Println(U_M.ID)


}
