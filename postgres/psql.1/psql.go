package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type Students struct {
	ID int
	Full_name string
	Birth_date int
	Group Groups
}

type Groups struct{
	ID int
	Name string
}

type Student_group struct {
	ID int
	Student_id int
	Group_id int
}

func main() {

	conStr := "user=doston password=doston dbname=schoole sslmode=disable"
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	respMsg := []byte(`
	{
		"full_name":"Rustamjon Rahmonov",
		"birth_date":1999,
		"group":{
			"name":"bootcamp go N8"
		}
	}
	`)

	var user1 Students

	if err = json.Unmarshal(respMsg, &user1); err!= nil{
		panic(err)
	}

	var reqUserId Students

	rowUser := db.QueryRow(`INSERT INTO students (full_name, birth_date) VALUES ($1,$2) returning id`, user1.Full_name, user1.Birth_date)	

	if err = rowUser.Scan(&reqUserId.ID); err!=nil {
		panic(err)
	}
	
	var reqMsgId Groups

	rowMsg := db.QueryRow(`INSERT INTO groups (name) VALUES ($1) returning id`,user1.Group.Name)
	if err = rowMsg.Scan(&reqMsgId.ID); err!=nil {
		panic(err)
	}

	var U_M Student_group
	rowUM := db.QueryRow(`INSERT INTO student_group (student_id, group_id) VALUES ($1,$2) returning id`,reqUserId.ID,reqMsgId.ID)
	if err := rowUM.Scan(&U_M.ID); err!=nil {
		panic(err)
	}

	fmt.Println(U_M.ID)


	var st_g []Students

	rowS, errS := db.Query(`SELECT s.id, s.full_name, g.name, g.id from students s join student_group sg on s.id=sg.student_id join groups g on sg.group_id=g.id WHERE s.id=$1`,1)
	if errS!=nil {
		panic(errS)
	}

	defer rowS.Close()

	for rowS.Next() {
		var student Students

		err1 := rowS.Scan(&student.ID, &student.Full_name, &student.Group.Name, &student.Group.ID)
		if err1!=nil {
			panic(err1)
		}

		st_g = append(st_g, student)
	}

	fmt.Println(st_g)



}
