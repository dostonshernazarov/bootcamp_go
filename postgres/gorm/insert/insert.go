package Insert

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Messages []Message `gorm:"many2many:user_messages;"`
}

type Message struct {
	gorm.Model
	Content string
	Users   []User `gorm:"many2many:user_messages;"`
}

type UserMessage struct {
	gorm.Model
	UserID    uint
	MessageID uint
}

func InsertData(db *gorm.DB, user *User, message *Message) {

	db.Create(&user)

	db.Create(&message)

	user_message := UserMessage{
		UserID:    user.Model.ID,
		MessageID: message.Model.ID,
	}

	db.Create(&user_message)
}

func SelectAll(db *gorm.DB, str []User) {
	db.Find(&str)
	fmt.Println("Selected data:")
	for _, u := range str {
		fmt.Printf("%+v\n", u)
	}

}

func updateUserData(db *gorm.DB, userID uint, newUsername string) {
	var user User

	db.First(&user, userID)

	user.Username = newUsername

	db.Save(&user)
}

func deleteUserData(db *gorm.DB, userID uint) {
	var user User

	db.First(&user, userID)

	db.Delete(&user)
}

func selectUserByID(db *gorm.DB, userID uint) (User, error) {
	var user User

	result := db.First(&user, userID)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func selectMany2ManyByID(db *gorm.DB, userID uint) (User, error) {
	var user User

	result := db.Preload("Messages").First(&user, userID)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
