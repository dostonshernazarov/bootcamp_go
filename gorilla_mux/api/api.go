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
	"strconv"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user/create", CreateUser).Methods("POST")
	router.HandleFunc("/user/getAll", GetAllUSer).Methods("GET")
	router.HandleFunc("/user/update/{id}", UpdateUserByID).Methods("GET")
	router.HandleFunc("/user/delete/{id}", DeletUserByID).Methods("GET")

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

	respUser, err := storageMux.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting all users", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("Error while marshaling all users to body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resBody)

}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	body_byte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *modelMux.UserMux

	err = json.Unmarshal(body_byte, &user)
	if err != nil {
		log.Println("error unmarshaling update user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.ID = userId

	respUser, err := storageMux.UpdateUserById(user)
	if err != nil {
		log.Println("error while updating user by id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error marshaling update user", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	err := storageMux.DeleteUserByID(userId)
	if err != nil {
		log.Println("error while deleting user by id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User deleted"))
}
