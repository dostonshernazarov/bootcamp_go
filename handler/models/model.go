package models

type User struct {
	ID        string `json:"id" :"id"`
	FirstName string `json:"first_name" :"first-name"`
	LastName  string `json:"last_name" :"last-name"`
}
