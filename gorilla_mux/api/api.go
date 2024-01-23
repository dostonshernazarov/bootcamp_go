package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"nt_bootcamp/bootcamp_go/gorilla_mux/modelMux"
	storageMux "nt_bootcamp/bootcamp_go/gorilla_mux/storage"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/create", CreateUser).Methods("POST")
	err := http.ListenAndServe("localhost:8080", router)

	if err != nil {
		log.Println("Error while running server", err)
		return
	}
	fmt.Println("Server running")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body_byte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	var user *modelMux.UserMux

	err = json.Unmarshal(body_byte, &user)
	if err != nil {
		log.Println("error while unmarshaling create user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uuid.New().String()
	user.ID = id

	respUser, err := storageMux.CreateUser(user)
	if err != nil {
		log.Println("error while creating user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshaling user in create user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resBody)
}

func GetAllUSer(w http.ResponseWriter, r *http.Request) {

}
