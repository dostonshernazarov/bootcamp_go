package main

import (
	in "nt_bootcamp/bootcamp_go/postgres/gorm/insert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)



func main() {
	dsn := "user=doston password=doston dbname=najot_talim sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&in.User{}, &in.Message{}, &in.UserMessage{})

	user := in.User{
		Username: "Doston",
	}
	message := in.Message{
		Content: "How is it going?",
	}

	in.InsertData(db, &user, &message)

	var usr []in.User

	in.SelectAll(db, usr)

}
