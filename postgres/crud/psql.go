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


	userId := InsertValueStudent(db, user1.Full_name, user1.Birth_date)

	groupId := InsertValueGroup(db, "bootcamp Go N8")

	result := InsertValueStGr(db, userId, groupId)

	fmt.Println(result)


}


func InsertValueStudent(db *sql.DB, full_name string, birth_date int) int {
	defer db.Close()

	var reqUserId Students

	rowUser := db.QueryRow(`INSERT INTO students (full_name, birth_date) VALUES ($1,$2) returning id`, full_name, birth_date)	

	if err := rowUser.Scan(&reqUserId.ID); err!=nil {
		panic(err)
	}

	return reqUserId.ID
}


func InsertValueGroup(db *sql.DB, name string) int {
	defer db.Close()
	var reqMsgId Groups

	rowMsg := db.QueryRow(`INSERT INTO groups (name) VALUES ($1) returning id`,name)
	if err := rowMsg.Scan(&reqMsgId.ID); err!=nil {
		panic(err)
	}

	return reqMsgId.ID
}

func InsertValueStGr(db *sql.DB, student_id int, group_id int) int {

	defer db.Close()
	var reqMsgId Groups

	rowMsg := db.QueryRow(`INSERT INTO student_group (student_id, group_id) VALUES ($1, $2) returning id`,student_id, group_id)
	if err := rowMsg.Scan(&reqMsgId.ID); err!=nil {
		panic(err)
	}

	return reqMsgId.ID
}

func SelectStValue(db *sql.DB, column_name string, value any) {
	rowSel := db.QueryRow(`SELECT * FROM students WHERE $2 = $3`,column_name, value)

	defer db.Close()

	var selValue Students
	err1 := rowSel.Scan(&selValue.ID, &selValue.Full_name, &selValue.Birth_date)
		if err1!=nil {
			panic(err1)
		}

	fmt.Println(selValue)
						
}

func SelectGrValue(db *sql.DB, column_name string, value any) {
	rowSel := db.QueryRow(`SELECT * FROM groups WHERE $2 = $3`,column_name, value)

	defer db.Close()

	var selValue Students
	err1 := rowSel.Scan(&selValue.Group.ID, &selValue.Group.Name)
		if err1!=nil {
			panic(err1)
		}

	fmt.Println(selValue)
}

func SelectManyToMany(db *sql.DB, ID_student int) {
	var st_g []Students

	rowS, errS := db.Query(`SELECT s.id AS student_id, s.full_name, g.name, g.id AS group_id from students s join student_group sg on s.id=sg.student_id join groups g on sg.group_id=g.id WHERE s.id=$2`,ID_student)
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

func UpdateValue(db *sql.DB, table_name string, column_name string, new_value any, id int) {
	rowUpdate, err := db.Exec(`UPDATE $1 SET $2 = $2 WHERE id=$3`,table_name, column_name, new_value, id)

	defer db.Close()

	if err != nil {
		panic(err)
	} 

	fmt.Println(rowUpdate.RowsAffected())
}

func DelRow(db *sql.DB, table_name string, id int) {
	rowDel, err := db.Exec(`DELETE FROM $1 WHERE id=$2`,table_name, id)

	if err != nil {
		panic(err)
	}

	fmt.Println(rowDel.RowsAffected())
}