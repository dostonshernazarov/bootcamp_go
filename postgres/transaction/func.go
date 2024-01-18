package transaction

import (
	"database/sql"
	"fmt"
)


type User struct {
	ID int
	Username string
	Messages Message 
}

type Message struct {
	ID int
	Content string
}

type UserMessage struct {
	ID int
	UserID    uint
	MessageID uint
}


func InsertUserMessage(db *sql.DB, user *User, message *Message) {
	tr, err := db.Begin()
	if err!=nil {
		tr.Rollback()
		panic(err)
	}

	var userId User

	rowU := tr.QueryRow(`INSERT INTO users (username) VALUES ($1) returning id`,user.Username)
	if err = rowU.Scan(&userId.ID); err!=nil {
		tr.Rollback()
		panic(err)
	}

	var messageId Message

	rowM := tr.QueryRow(`INSERT INTO messages (content) VALUES ($1) returning id`,message.Content)

	if err := rowM.Scan(&messageId.ID); err!=nil {
		tr.Rollback()
		panic(err)
	}

	_, errUM := tr.Exec(`INSERT INTO user_message (user_id, message_id) VALUES ($1, $2)`, userId.ID, messageId.ID)
	if errUM!=nil {
		tr.Rollback()
		panic(errUM)
	}

	err = tr.Commit()
	if err!=nil {
		tr.Rollback()
		panic(err)
	}
}

func SelectUserById(db *sql.DB, userID int) (User, error) {
	var errUser User

	tr,err := db.Begin()
	if err != nil {
		tr.Rollback()
		return errUser, err
	}
	

	var selectUser User

	rowU := tr.QueryRow(`SELECT * FROM users WHERE id = $1`,userID)
	if err = rowU.Scan(&selectUser);err!=nil {
		tr.Rollback()
		return errUser, err
	}

	errC := tr.Commit()
	if errC!=nil {
		tr.Rollback()
		return errUser, err
	}
	return selectUser, nil

}

func UpdateUsername(db *sql.DB, newUsername string, id int) (int, error) {

	tr, err := db.Begin()
	if err != nil {
		tr.Rollback()
		return -1, err
	}



	var user User
	rowUser := tr.QueryRow(`UPDATE users SET username=$1 WHERE id=$2 returning id`,newUsername, id )
	if err = rowUser.Scan(&user.ID); err!=nil{
		tr.Rollback()
		return -1, err
	}
	 
	errT := tr.Commit()
	if errT!=nil {
		return -1, errT
	}
	return user.ID, nil
}

func DeleteAllUsers(db *sql.DB) {

	tr, err := db.Begin()
	if err!=nil {
		tr.Rollback()
		panic(err)
	}

	rowU, err := tr.Exec(`DELETE FROM users`)
	if err!=nil {
		tr.Rollback()
		panic(err)
	}

	err = tr.Commit()
	if err!= nil {
		tr.Rollback()
		panic(err)
	}


	fmt.Println(rowU.RowsAffected())

}

func DeleteUserById(db *sql.DB, id int) (int, error) {
	tr, err := db.Begin()
	if err!= nil {
		tr.Rollback()
		return -1, err
	}

	var user User

	rowU := tr.QueryRow(`DELETE FROM users WHER id=$1 returning id`,id)
	if err:=rowU.Scan(&user.ID); err!= nil {
		tr.Rollback()
		return -1, err
	}

	err = tr.Commit()
	if err!= nil {
		tr.Rollback()
		return -1, err
	}

	return user.ID, nil
}

func GetUserMessageById(db *sql.DB, id int) (int, error) {
	tr, err := db.Begin()
	if err!= nil {
		tr.Rollback()
		return -1, err
	}

	var user User

	rowUM := tr.QueryRow(
	`SELECT u.id,
		u.username,
	m.content  FROM users u JOIN user_message um ON u.id=um.user_id JOIN messages m ON um.message_id=m.id WHERE u.id=$1`,id)
	
	if err = rowUM.Scan(&user.ID, &user.Username,&user.Messages.Content); err!= nil {
		tr.Rollback()
		return -1, err 
	}

	errC := tr.Commit()
	if errC!=nil {
		tr.Rollback()
		return -1, errC
	}
	return id, nil
}