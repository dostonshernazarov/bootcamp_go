package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
	"log"
	"net/http"
	model "nt_bootcamp/bootcamp_go/handler/models"
	"nt_bootcamp/bootcamp_go/handler/storage"
	"strconv"
)

func StringToUUID(s string) pgtype.UUID {
	data, err := parseUUID(s)
	if err != nil {
		return pgtype.UUID{
			Bytes: [16]byte{},
			Valid: false,
		}
	}
	return pgtype.UUID{
		Bytes: data,
		Valid: true,
	}
}
func parseUUID(src string) (dst [16]byte, err error) {
	switch len(src) {
	case 36:
		src = src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:]
	case 32:
		// dashes already stripped, assume valid
	default:
		// assume invalid.
		return dst, fmt.Errorf("cannot parse UUID %v", src)
	}

	buf, err := hex.DecodeString(src)
	if err != nil {
		return dst, err
	}

	copy(dst[:], buf)
	return dst, err
}

func main() {

	http.HandleFunc("/user/create", CreateUser)
	http.HandleFunc("/user/getAllUsers", GettAllUsers)
	http.HandleFunc("/user/getUser", GetUserById)
	http.HandleFunc("/user/update", UpdateUserByID)
	http.HandleFunc("/user/delete", DeleteUserByID)
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
	idUser := r.URL.Query().Get("id")

	respuser, err := storage.GetUser(idUser)
	if err != nil {
		log.Println("Error while getting user by id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	respBody, err := json.Marshal(&respuser)
	if err != nil {
		log.Println("Error while marshaling user by id", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("error while getting body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user *model.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		log.Println("error while unmasheling body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	respUser, err := storage.UpdateUserById(user)
	if err != nil {
		log.Println("error while updating user by id  ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(respUser)
	if err != nil {
		log.Println("error while marshalling to response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := storage.DeleteUserByID(id)
	if err != nil {
		log.Println("error while deleting user by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted "))
}
