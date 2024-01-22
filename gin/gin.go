package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Age      int    `json:"age"`
}

var users = []User{
	{
		ID:       1,
		Name:     "Davron",
		LastName: "Nuriddinov",
		Age:      21,
	},
}

func main() {

	router := gin.Default()

	router.POST("/user", CreateUser)
	router.GET("/users", GetUser)

	err := router.Run("localhost:8082")
	if err != nil {
		panic(err)
	}
}

func CreateUser(g *gin.Context) {
	var newUser User

	err := g.BindJSON(&newUser)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}
	newUser.ID = len(users) + 1

	users = append(users, newUser)

}

func GetUser(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, users)
}
