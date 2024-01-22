package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	model "nt_bootcamp/bootcamp_go/handler/models"
	"nt_bootcamp/bootcamp_go/handler/storage"
	"strconv"
)

func main() {

	http.HandleFunc("/user/create", CreateUser)
	http.HandleFunc("/user/getUsers", GettAllUsers)
	err := http.ListenAndServe("localhost:8082", nil)
	log.Println("Server is running")
	if err != nil {
		fmt.Println("Error while running server")

	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body_byte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	var user *model.User

	err = json.Unmarshal(body_byte, &user)
	if err != nil {
		log.Println("error while unmarshaling user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		log.Println("error while getting uuid", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.ID = id.String()

	respUser, err := storage.CreateUser(user)
	if err != nil {
		log.Println("error while creating user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshaling user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func GettAllUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error while converting page, is not int", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	limit := r.URL.Query().Get("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error while converting limit, is not int", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users, err := storage.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting all users", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respBody, err := json.Marshal(&users)
	if err != nil {
		log.Println("Error while marshaling all users", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)

}

func GetUserById(w http.ResponseWriter, r *http.Request) {

}
